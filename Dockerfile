FROM golang:alpine as builder
RUN apk add --no-cache dep
ADD . /go/src/build/
WORKDIR /go/src/build
RUN dep ensure
RUN go build -o main .
FROM alpine
RUN adduser -S -D -H -h /app appuser
RUN apk add --no-cache imagemagick
COPY --from=builder /go/src/build/main /app/
COPY --from=builder /go/src/build/assets /app/assets
RUN mkdir /app/tmp && \
    mkdir /app/static && \
    chown -R appuser /app && \
    chmod 755 /app
WORKDIR /app
ENV PORT 5000
EXPOSE 5000
USER appuser
CMD ["./main"]
