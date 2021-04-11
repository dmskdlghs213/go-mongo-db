FROM golang:1.15 as build-image

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

FROM alpine

COPY --from=build-image /go/bin/app /go/bin/app
COPY --from=build-image ./mongodb/migrations /migrations
RUN chmod -R 400 /migrations

EXPOSE 8888

CMD ["/go/bin/app"]