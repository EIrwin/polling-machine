# Poll Machine

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
- Break out domains into separately executable services.
- Use document storage for aggregating quick snapshots for poll results.
- Take better advantage of cache layer for data access.
- API Model Validation (currently, there is none).
- Add supervisor process to monitor golang app to keep it running.
- Use a better abstraction for configuration.
- Pass in init script when initializing Postgres database.
- Hash passwords instead of storing them as plain text :)
- Use Consul or other mechanism for discovering infrastructure service endpoints.
### Known Problems
- Email notifications might be blocked by SMTP relay

### Screenshots




