# CapTrivia: Cap Table Trivia Game

## Introduction
Welcome to CapTrivia, a cap table trivia game developed by Pulley. This project is part of our engineering interview process, designed to assess your coding skills and problem-solving abilities. Currently, CapTrivia is a simple, single-player game. Your challenge is to develop it into a multiplayer experience.

## Project Overview
- **Current State**: CapTrivia is a functioning single-player game with basic features. We tried our best to make it very ugly so that you have lots of room to improve it!
- **Tech Stack**: The front end is built with React (in `App.js`), and the back end with Go (in `main.go`). We also included a docker-compose.yml so that you can easily run this locally (it includes a PostgreSQL instance).
- **Objective**: Transform CapTrivia into a multiplayer game where users can create games, invite others via a shareable link, and compete in real-time trivia challenges.

## Quickstart

1. Make sure you have Docker installed. https://docs.docker.com/engine/install/
2. In the captrivia root directory, run `docker compose up`.
3. Open http://localhost:3000 in your browser to see the game.

## Your Task
1. **Refactor and Enhance**: The current code is intentionally basic and contained in single files for both front and back ends. Your first task is to clean up and structure the codebase.
2. **Implement Multiplayer Functionality**: 
    - Allow users to start a new game and generate a shareable link.
    - Enable real-time joining of players in a waiting room.
    - Implement a countdown and start the game simultaneously for all players.
    - Introduce a scoring system with a configurable number of questions (default: 10).
    - Players will compete to answer each question first. For each question, the player who answers first gets the points.
    - Display the winner at the end and allow users to start new games.

## Time Allocation
You should aim to spend no more than 5 hours on this challenge. We respect your time and want to see what you can achieve within these constraints.

## Submission Guidelines
- **Git Patch**: Submit your solution as a Git patch file(s). This allows us to easily review your changes against our original code. You can do this with `git diff starting_commit..HEAD > captrivia.patch` or `git format-patch --root start_commit..HEAD` (start_commit should be the commit *before* your first commit, it will not be included in the diff/patch).
- **Running Locally**: Please ensure that we can run your version with `docker compose up`.

## Evaluation Criteria
- **Code Quality**: Clean, readable, and well-structured code.
- **Problem Solving**: Effective and efficient solutions to the challenges presented.
- **Creativity**: We appreciate innovative approaches and ideas.
- **Tool Utilization**: Feel free to use tools like GitHub Copilot and ChatGPT. We value efficiency and resourcefulness.

## After Submission
Once we receive your solution, we'll schedule a live code pairing session. Here, you'll walk us through your code, and we'll collaboratively work on an additional small feature.

## Conclusion
Thank you for participating in Pulley's engineering challenge. We're excited to see your approach to enhancing CapTrivia. Good luck, and have fun!
