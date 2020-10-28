FROM golang:1.15.0-alpine3.12 as builder

WORKDIR /build
COPY . /build/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o webdriver-watcher cmd/main.go

FROM scratch
COPY --from=builder /build/webdriver-watcher /webdriver-watcher
ENTRYPOINT [ "/webdriver-watcher" ]
