FROM golang:1.17-alpine

LABEL maintainer="Nolan Clark <tweakdeveloper@gmail.com>"

EXPOSE 8080
ENV GIN_MODE=release

WORKDIR /usr/src/app

RUN apk add texlive

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app .

CMD ["app"]
