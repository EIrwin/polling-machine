
FROM golang

WORKDIR /go/src/polling-machine
ADD . /go/src/polling-machine

RUN go get gopkg.in/validator.v2
RUN go get github.com/gorilla/mux
RUN go get github.com/thoas/stats
RUN go get github.com/codegangsta/negroni
RUN go get github.com/jinzhu/gorm
RUN go get github.com/eirwin/polling-machine/models
RUN go get github.com/eirwin/polling-machine/users
RUN go get github.com/eirwin/polling-machine/polls
RUN go get github.com/eirwin/polling-machine/data
RUN go get github.com/pborman/uuid

RUN go install polling-machine
EXPOSE 8181
ENTRYPOINT /go/bin/polling-machine