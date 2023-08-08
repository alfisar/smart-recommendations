FROM golang:alpine


RUN apk update && apk add --no-cache git
WORKDIR /app

COPY . .
ENV GO111MODULE=on

RUN go mod tidy

RUN go build -o main .
EXPOSE 8080
CMD ["./main"]