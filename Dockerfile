FROM golang:buster

WORKDIR /app
COPY . ./
RUN go mod download
RUN go install -v ./...

CMD ["go", "run", "main.go"]