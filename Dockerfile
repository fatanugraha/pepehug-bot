FROM golang:alpine as builder
RUN apk add --no-cache dep
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN dep ensure
RUN go build -o main .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]
