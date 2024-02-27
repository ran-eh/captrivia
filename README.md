# CapTrivia: Cap Table Trivia Game - Ran Ever-Hadani

## Summary
- **Philosophy**: To control scope, and in the spirit of agile developmemt, it is assumed that event througput is small, and the design leaves room for improving performance down the line should it be needed.
- **Event stream**: And EventSender was added to the backend and it is being called by the various hamdlers.  Currently, the sender  writes events into the analytics database syncronously.  When performance needs arise, it can be made asyncronous: the sender will post the event to a queue (e.g. GCP Pubsub).  A separate service will subscribe to to the event stream, and it will write them to the analytics database in an efficient manner (e.g. by batching).
- **Analytics Database Platform**: Prostgres is not ideal for an Analytics, but it works as a staring point.  The provided docker setup is left as is.  
- **Analytics Pipeline Platform:** A dbt pipeline is used to trasform events into usable models.
- **Analytics Data Modeling**: Models are organized in the following layers:
  - **Staging**: Models reflecting source data with minor transformations.  There is one model per event type.  
  - **Main**: Business objects.  In this project it contains a single `games` model.
  - **Reports**: visualization models.
- **Visualization/Dashboard**: A Jupyter notebook is used as a starting point.  A more flexible platform may be used down the line.
- **Testing**: the bot script was changed to generate 10 days worth of events.  To make it easy to fake success rates, the correct answer id was added to the `/questions` end point.  Also, a `/debug/shiftdays` end point was added to allow time travel, e.g. creating events in the past.  A mechanism will be added to disable these test features in production.
- **Deployment**: the existing docker composed was kept.  A container was added for notebook visualization.

## Code Review
The captrivia code was forked into a private repo, which was then shared with A private fork was created for the captrivia repo.  It is shared with [mboorstin](https://github.com/mboorstin) and [aarondl0](https://github.com/aarondl0).  The new code may be reviewed in 

https://github.com/ProlificLabs/captrivia/compare/main...ran-eh:captrivia:main

## Running the code
- Clone the forked repo https://github.com/ran-eh/captrivia
- In it'a root directory, run `docker compose up`
- Create events table
```
$ export PGPASSWORD=postgres; psql -f db/up.pgsql -h localhost -U postgres -d captrivia
```
- 
## Quickstart

1. Make sure you have Docker installed. https://docs.docker.com/engine/install/
2. In the captrivia root directory, run `docker compose up`.
3. Open http://localhost:3000 in your browser to see the game.

## Tasks
Please complete the task appropriate to the position you are applying for.

### Software Engineers
1. **Refactor and Enhance**: The current code is intentionally basic and contained in single files for both front and back ends. Your first task is to clean up and structure the codebase.
2. **Implement Multiplayer Functionality**: 
    - Allow users to start a new game and generate a shareable link.
    - Enable real-time joining of players in a waiting room.
    - Implement a countdown and start the game simultaneously for all players.
    - Introduce a scoring system with a configurable number of questions (default: 10).
    - Players will compete to answer each question first. For each question, the player who answers first gets the points.
    - Display the winner at the end and allow users to start new games.

### Data Engineers
Your task is to build a data pipeline to record and visualize analytics information about a player's performance.
  - Stand up an analytics database: you may use the existing PostgreSQL database in the docker-compose.yml, or any other database of your choice.
  - Instrument the backend application to report data on each answered question to the database.
  - Start up the bot provided in [bot](bot) to start playing games and feeding data into your database.  Right now the bot uses a very simple strategy (always choosing the first answer); you will likely want to change its strategy to create more interesting data.  If you like, your bot can "cheat" by using [questions.json](backend/questions.json) to sometimes figure out the right answer.
  - Build and display some simple analytics on top of your analytics database.


## Time Allocation
There is no time restrictions on this challenge, you can take as much time as you need. This is to allow folks who might not know the tech stack to spend time learning or if you wanted to take the challenge to the next level. The most important part is achieving the task above.

As a baseline for someone proficient in the stack we expect this challenge to take around 3-5 hours. This figure is only to help set your own expectations about how much time you may need to allocate to complete it.

## Submission Guidelines
- **Zip File**: Zip up the code directory and send it to us. Please make sure you delete `node_modules` and any Go binaries that would bloat the zip. Leaving the `.git` directory is preferred so we can see the history of how you implemented your solution. To automatically clean files that are not in git, you can use `git clean -fdx` (with appropriate caution as this deletes files).
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
