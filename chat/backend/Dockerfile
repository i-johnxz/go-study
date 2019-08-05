FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
RUN go build .
CMD ["/app/realtime-chat-go-react"]
