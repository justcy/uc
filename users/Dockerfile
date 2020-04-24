FROM golang

RUN mkdir -p /go/src/github.com/justcy/uc

ADD . /go/src/github.com/justcy/uc

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/justcy/uc/users/cmd  -watch github.com/justcy/uc/users
