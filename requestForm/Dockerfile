# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app

COPY go.mod ./
#COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /docker-alex

#EXPOSE 8081

CMD [ "/docker-alex" ]