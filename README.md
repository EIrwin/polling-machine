# Poll Machine

### Prerequisites
* In order to to run Poll Machine, you need an accessible docker host to run `docker-compose`.
* If you are running on Mac or Windows, you can use Docker for Mac or Docker for Windows.

### Installation
1. Clone this project or download and extract ZIP file.
2. `cd /polling-machine`
3. Run `docker-compose up`
4. Navigate to `http://$DOCKER_HOST`

**Note** - *This project assumes development docker host will run on `192.168.99.100`. If you want to change this, you can modify `${WORKSPACE}/polling-machine/client/development.js` to use another IP address.*

**Targeting** - *The `docker-compose.yml` file currently targets the `development` configuration when running the `client` container by using `gulp development`. This can be changed by replacing `development` with one of the other possible environment targets `(local|staging|release)`*

### Poll Machine Stack
* Go-lang
* Redis
* PostgresSQL
* Javascript / AngularJS
* Gulp
* NGINX (Reverse proxy)
* Docker

### Stories

- [X] As a user I would like to create an account so that I can create and manage my polls.
- [X] As a user I would like to create a poll so that I can share it and collect feedback.
- [X] As a user I would like a way to share my polls on a variety of social networks so that I can gather my results.
- [X] As a user I would like to be notified whenever someone answers my poll.
- [X] As a user I would like to prevent multiple people from taking the same poll twice without requiring them to login.
- [X] As a user I would like to display the results of my post in a graphical manner so that I can visually understand the various poll answers.
- [X] As a user I would like to set an end time to my poll and prevent any additional responses when that date is met.

### Needed Improvements
- Pool database and cache connections more efficiently.
- Improve build process with Makefile to pass in static golang binary (Dockerfile currently is `go-getting` all dependencies causing the images to be very large.
- Refactor `GetResponseCounts(...)` because its EXTREMELY inefficient. Might be better just storing count in Redis.
- Break out domains into separately executable services and use gRPC or RabbitMQ for between services.
- Use document storage for aggregating quick snapshots for poll results.
- Take better advantage of cache layer for data access.
- API Model Validation (currently, there is none...oops).
- Add supervisor process to monitor golang app to keep it running.
- Use a better abstraction for centralized configuration.
- Pass in init script when initializing Postgres database.
- Hash passwords instead of storing them as plain text :)
- Use Consul or other mechanism for discovering infrastructure service endpoints.
- Relay error messages from backend to frontend.

### Known Problems
- Email notifications might be blocked by SMTP relay



