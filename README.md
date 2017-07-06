# go-docker-eventstore

An example of writing events to EventStore with Go.

The application and the EventStore instance run on Docker.

The application is provisioned with Docker Compose.

The application is compiled and run with Gulp on the application container.

Clone the repository to your preferred location.

Install Gulp into the eventposter directory as follows

	$ cd eventposter
	$ npm install --save-dev gulp
	$ npm install --save-dev gulp-util
	$ npm install --save-dev node-notifier

Navigate back to the application root directory and build & run as follows:

	$ cd ..
	$ sudo docker-compose up

If everything goes correctly:

1. Eventstore will start up and log to the console.
2. Gulp will compile and run the application.
3. The application will insert 10,000 events into EventStore.

The messages logged to the console show as:

	eventposter_1  | 2017/07/06 21:05:22 201 Created

The 201 is an HTTP Status Code (Created) and not an event counter!

Things can be terminated with __Ctrl-C__ once the following message is displayed:

	eventposter_1  | 2017/07/06 21:30:10 Have now created 10000 events

The final dialogue:

	eventposter_1  | 2017/07/06 21:30:10 201 Created
	eventposter_1  | 2017/07/06 21:30:10 Have now created 10000 events
	^CGracefully stopping... (press Ctrl+C again to force)
	Stopping godockereventstore_eventposter_1 ... done
	Stopping godockereventstore_eventstore_1 ... done
	$

## To build (or rebuild if things don't go according to plan)

	$ sudo docker-compose build

The Docker images produced are:

Image | Size
----- | ----
godockereventstore_eventposter | 339MB
godockereventstore_eventstore | 293MB

## To run

	$ sudo docker-compose up
