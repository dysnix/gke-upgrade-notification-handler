FROM golang:1.17-alpine as builder
LABEL stage=intermediate

RUN apk add --update git openssh-client ca-certificates

# Change workdir
WORKDIR /go/src/github.com/dysnix/gke-upgrade-notification-handler

COPY go.mod .
COPY go.sum .

RUN go mod download
# Copy code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags=-w -o /go/bin/app

FROM alpine:latest as runtime-image

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/bin/app /app

ENTRYPOINT [ "/app" ]
