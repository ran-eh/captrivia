import requests

class Captrivia:

    # Constructor: save the URL.
    def __init__(self, url):
        self.url = url
        self.session_id = None

    # Start a new game.  Saves the session id for future use (this is a
    # single-threaded bot, for now).
    def start_game(self):
        result = requests.post(self.url + "/game/start")
        self.session_id = result.json()["sessionId"]

    # Get a set of questions.  This isn't really necessary as the endpoint is just
    # selecting a random and randomly ordered set of questions without linking to a
    # specific session ID, but it mimics what the UI would do so it makes our bot
    # look more like a human.
    # Returns a list of questions with 'id', 'questionText', and 'options' properties.
    def get_questions(self):
        result = requests.get(self.url + "/questions")
        return result.json()

    # Answer a question.  Returns whether we were correct
    def answer_question(self, question_id, answer):
        self.ensure_in_game()

        # Do the request to answer the question
        body = {
            "sessionId": self.session_id,
            "questionId": question_id,
            "answer": answer
        }
        result = requests.post(self.url + "/answer", json=body)

        # Return whether we were correct
        return result.json()["correct"]

    # End a game.  Returns the final score
    def end_game(self):
        self.ensure_in_game()

        # Do the request to end the game, and save our score
        body = {
            "sessionId": self.session_id
        }
        result = requests.post(self.url + "/game/end", json=body)
        score = result.json()["finalScore"]

        # Clear the sesssion id and return the score
        self.session_id = None
        return score

    # Make sure we're currently in a game
    def ensure_in_game(self):
        if not self.session_id:
            raise Exception("Not currently in a game!")
