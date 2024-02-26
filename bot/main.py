#!/usr/bin/env python

import random
from captrivia import Captrivia

# Simple bot to play Captrivia.  Starts a new game, answers all of the questions, ends the game, and starts
# a new one.

# Assumes we're running Captrivia locally, or that it's running in Docker portmapped to the same 8080
# port.  The default `docker compose up` should do this correctly.
CAPTRIVIA_URL = "http://localhost:8080"

# Instantiate a Captrivia object to talk to the server
cap = Captrivia(CAPTRIVIA_URL)
game_count = 0

# Main infinite loop: Start a game, answer the questions, finish the game
while True:
    # Start a game
    cap.start_game()

    # Answer all of the questions
    questions = cap.get_questions()
    for question in questions:
        qid = question["id"]
        # Choose the correct answer 70% of the time
        if random.random() < 0.7:
            answer = question["correctIndex"]
        else:
            answer = question["correctIndex"] + 1 % len(question["options"])

        cap.answer_question(qid, answer)

    # abandon game 5% of the time
    if random.random() > 0.05:
        # End the game
        cap.end_game()

    # Do some logging so we can see how many answers we've given
    game_count += 1
    if game_count % 100 == 0:
        print("Played %d games" % (game_count))
