package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// TODO: instead of using a global, use a closure as shown in
// https://stackoverflow.com/questions/34046194/how-to-pass-arguments-to-router-handlers-in-golang-using-gin-web-framework
var eventSender *EventSender;

type Question struct {
	ID           string   `json:"id"`
	QuestionText string   `json:"questionText"`
	Options      []string `json:"options"`
	CorrectIndex int      `json:"correctIndex"`
}

type PlayerSession struct {
	Score int
}

type SessionStore struct {
	sync.Mutex
	Sessions map[string]*PlayerSession
}

func (store *SessionStore) CreateSession() string {
	store.Lock()
	defer store.Unlock()

	uniqueSessionID := generateSessionID()
	store.Sessions[uniqueSessionID] = &PlayerSession{Score: 0}

	return uniqueSessionID
}

func (store *SessionStore) GetSession(sessionID string) (*PlayerSession, bool) {
	store.Lock()
	defer store.Unlock()

	session, exists := store.Sessions[sessionID]
	return session, exists
}

func generateSessionID() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

type GameServer struct {
	Questions []Question
	Sessions  *SessionStore
}

func main() {
	eventSender = NewEventSender()

	defer eventSender.Close()

	// Setup the server
	router, err := setupServer()
	if err != nil {
		log.Fatalf("Server setup failed: %v", err)
	}

	// set port to PORT or 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Println("Server starting on port " + port)
	log.Fatal(router.Run(":" + port))
}

// setupServer configures and returns a new Gin instance with all routes.
// It also returns an error if there is a failure in setting up the server, e.g. loading questions.
func setupServer() (*gin.Engine, error) {
	questions, err := loadQuestions()
	if err != nil {
		return nil, err
	}

	sessions := &SessionStore{Sessions: make(map[string]*PlayerSession)}
	server := NewGameServer(questions, sessions)

	// Create Gin router and setup routes
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	config := cors.DefaultConfig()
	// allow all origins
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.POST("/game/start", server.StartGameHandler)
	router.GET("/questions", server.QuestionsHandler)
	router.POST("/answer", server.AnswerHandler)
	router.POST("/game/end", server.EndGameHandler)
	router.POST("/debug/shiftdays", server.ShiftDaysHandler)

	return router, nil
}

func NewGameServer(questions []Question, store *SessionStore) *GameServer {
	return &GameServer{
		Questions: questions,
		Sessions:  store,
	}
}

func (gs *GameServer) StartGameHandler(c *gin.Context) {

	sessionID := gs.Sessions.CreateSession()
	c.JSON(http.StatusOK, gin.H{"sessionId": sessionID})

	eventSender.Send(&Event{
		SessionID: sessionID,
		Type:      "start_game",
		Data:      fmt.Sprintf("{\"sessionId\": \"%s\"}", sessionID),
	})
}

func (gs *GameServer) QuestionsHandler(c *gin.Context) {
	shuffledQuestions := shuffleQuestions(gs.Questions)
	questions := shuffledQuestions[:10]
	json, _ := json.Marshal(questions)
	c.JSON(http.StatusOK, questions)
	
	// TODO: Session ID not available for this call.  Suggest including
	// it in the call to allow linking the call to a session.
	eventSender.Send(&Event{
		Type:      "get_questions",
		Data: string(json),
	})
}

func (gs *GameServer) AnswerHandler(c *gin.Context) {
	type SubmittedAnswer struct {
		SessionID  string `json:"sessionId"`
		QuestionID string `json:"questionId"`
		Answer     int    `json:"answer"`
	}

	var submittedAnswer SubmittedAnswer;
	var eventData struct {
		Answer SubmittedAnswer;
		Correct bool;
		CurrentScore int;
	}
	if err := c.ShouldBindJSON(&submittedAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	session, exists := gs.Sessions.GetSession(submittedAnswer.SessionID)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}

	correct, err := gs.checkAnswer(submittedAnswer.QuestionID, submittedAnswer.Answer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	if correct {
		session.Score += 10 // Increment score for correct answer
	}

	c.JSON(http.StatusOK, gin.H{
		"correct":      correct,
		"currentScore": session.Score, // Return the current score
	})
	eventData.Answer = submittedAnswer
	eventData.Correct = correct
	eventData.CurrentScore = session.Score

	json, _ := json.Marshal(eventData)

		eventSender.Send(&Event{
		SessionID: submittedAnswer.SessionID,
		Type:      "answer",
		Data:      string(json),
	})

}

func (gs *GameServer) EndGameHandler(c *gin.Context) {
	var request struct {
		SessionID string `json:"sessionId"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	session, exists := gs.Sessions.GetSession(request.SessionID)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"finalScore": session.Score})
	eventSender.Send(&Event{
		SessionID: request.SessionID,
		Type:      "end_game",
		Data:      fmt.Sprintf("{\"finalScore\": \"%d\"}", session.Score),
	})
}

func (gs *GameServer) ShiftDaysHandler(c *gin.Context) {
	var request struct {
		ShiftDateDays int `json:"ShiftDateDays"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	eventSender.ShiftDateDays = request.ShiftDateDays

	c.JSON(http.StatusOK, gin.H{})
}

func (gs *GameServer) checkAnswer(questionID string, submittedAnswer int) (bool, error) {
	for _, question := range gs.Questions {
		if question.ID == questionID {
			return question.CorrectIndex == submittedAnswer, nil
		}
	}
	return false, errors.New("question not found")
}

func shuffleQuestions(questions []Question) []Question {
	rand.Seed(time.Now().UnixNano())
	qs := make([]Question, len(questions))

	// Include correct answer with questionm to make debugging easier.
	// TODO: only include it in debug build
	for i, q := range questions {
		qs[i] = Question(q)
	}

	rand.Shuffle(len(qs), func(i, j int) {
		qs[i], qs[j] = qs[j], qs[i]
	})
	return qs
}

func loadQuestions() ([]Question, error) {
	fileBytes, err := ioutil.ReadFile("questions.json")
	if err != nil {
		return nil, err
	}

	var questions []Question
	if err := json.Unmarshal(fileBytes, &questions); err != nil {
		return nil, err
	}

	return questions, nil
}
