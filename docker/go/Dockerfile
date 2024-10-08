FROM golang:1.23-alpine

RUN go install github.com/air-verse/air@latest

WORKDIR /app

#COPY go.mod go.sum ./

COPY . .

RUN go mod download

#RUN go build -o main .

#CMD ["./main"]

CMD ["air", "-c", ".air.toml"]