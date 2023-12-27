FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go app
RUN go mod download
RUN go build -o tbot-app ./main.go

CMD ["./tbot-app"]