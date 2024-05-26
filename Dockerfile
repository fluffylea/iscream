FROM docker.io/library/golang

WORKDIR /usr/src/app

COPY go.mod resources ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]
