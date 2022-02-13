FROM golang:1.17

LABEL maintainer="Nolan Clark <tweakdeveloper@gmail.com>"

EXPOSE 8080
ENV GIN_MODE=release

WORKDIR /usr/src/app

RUN apt-get -y update && apt-get install -y pandoc latexmk

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app .

CMD ["app"]
