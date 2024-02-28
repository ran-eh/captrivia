# CapTrivia: Cap Table Trivia Game - Ran Ever-Hadani

## Summary
- **Philosophy**: To control scope, and in the spirit of agile developmemt, it is assumed that event throughput is small, and the design leaves room for improving performance down the line should it be needed.
- **Event stream**: An EventSender was added to the backend, and it is being called by the various handlers.  Currently, the sender  writes events into the analytics database syncronously.  This is not ideal long term, and as needs evolve, it can be made asyncronous: the sender will post the event to a queue (e.g. GCP Pubsub).  A separate service will subscribe to to the event stream, and it will write them to the analytics database in an efficient manner (e.g. by batching).
- **Analytics Database Platform**: Prostgres is not ideal for analytics, but it works as a starting point, so it is kept in place to minimize scope.
- **Analytics Pipeline Platform:** A dbt pipeline is used to trasform events into usable models.
- **Analytics Data Modeling**: Models are organized in the following layers:
  - **Staging**: Models reflecting source data with minor transformations.  There is one model per event type.  
  - **Main**: Business objects.  In this project it contains a single `games` model.
  - **Reports**: visualization models.
- **Visualization/Dashboard**: A Jupyter notebook is used as a starting point.  In hindsight this was not the best choice, as it proved to be finicky with docker, but here we are.
- **Testing**: the bot script was changed to generate 10 days worth of events.  To make it easy to fake success rates, the correct answer id was added to the `/questions` end point resoponse.  Also, a `/debug/shiftdays` end point was added, to allow time travel, e.g. creating events in the past.  A mechanism will be added to disable these test features in production.
- **Deployment**: the existing docker composed was kept.  A container was added for notebook visualization.

## Code Review
The captrivia code was forked into a private repo, which was then shared with A private fork was created for the captrivia repo.  It is shared with [mboorstin](https://github.com/mboorstin) and [aarondl0](https://github.com/aarondl0).  The new code may be reviewed in 

https://github.com/ProlificLabs/captrivia/compare/main...ran-eh:captrivia:main

## Running the code
- Clone the forked repo https://github.com/ran-eh/captrivia
- In it'a root directory, run `docker compose up`
- Run the bot script to create ten days worth of events
  - Follow the instructions in ./bot/README.md .
- View games report
  - Open http://localhost:8888/lab/tree/dash.ipynb
