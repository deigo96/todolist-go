FROM golang:alpine

WORKDIR /app 

COPY . /app

CMD ["go", "run", "cmd/main.go"]%