FROM golang:1-alpine as builder
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN mkdir /build
WORKDIR /build
COPY . .
RUN go get github.com/gorilla/mux
RUN go build -o main .

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]