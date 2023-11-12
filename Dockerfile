FROM golang:1.18
ENV GOPATH=/go
WORKDIR /go/src/github.com/slaxor/e
COPY .  /go/src/github.com/slaxor/e
RUN apt-get update && apt-get install -y neovim tmux
RUN go build -o e .
CMD ./test-session.sh
