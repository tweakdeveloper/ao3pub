FROM golang:1.17-alpine

LABEL maintainer="Nolan Clark <tweakdeveloper@gmail.com>"

EXPOSE 8080

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /usr/src/app

RUN apk add texlive texlive-xetex texmf-dist-latexextra texmf-dist-langchinese ttf-dejavu font-noto-cjk

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o app .

CMD ["./app"]
