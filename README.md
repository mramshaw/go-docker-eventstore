# go-docker-eventstore

Example writing events to EventStore with Go.

The application and the EventStore instance are running on Docker.

The application is provisioned with Docker Compose.

The application is compiled and run with Gulp on the application container.

Clone the repository to your preferred location.

Install Gulp into the eventposter directory as follows

`$ npm install --save-dev gulp`

`$ npm install --save-dev gulp-util`

`$ npm install --save-dev node-notifier`

Navigate back to the application root directory and execute the command

`$ docker-compose up`

Eventstore will start and log to the console.

Gulp will compile and run the application. The application will insert 10,000 events into EventStore.
