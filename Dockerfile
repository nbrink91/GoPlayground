# Build app
FROM golang:1-alpine as builder
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh dep
RUN mkdir -p /go/src/github.com/nbrink91/GoPlayground
WORKDIR /go/src/github.com/nbrink91/GoPlayground

# Install deps
COPY ./Gopkg.lock .
COPY ./Gopkg.toml .
ENV GOHOME=/go
RUN dep ensure -vendor-only

# Copy and build app
COPY . .
RUN go build -o main .

# Copy app to final image
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
WORKDIR /app
COPY --from=builder /go/src/github.com/nbrink91/GoPlayground .

ENV GIN_MODE=release
ENV PORT=8080

EXPOSE 8080

CMD ["./main"]