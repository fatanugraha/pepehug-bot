FROM golang:alpine as builder
RUN apk add --no-cache dep
ADD . /go/src/build/
WORKDIR /go/src/build
RUN dep ensure
RUN go build -o main .
FROM alpine
COPY --from=builder /go/src/build/main /app/
COPY --from=builder /go/src/build/assets /app/assets
RUN mkdir /app/tmp
RUN mkdir /app/static
WORKDIR /app
ENV PORT 5000
EXPOSE 5000
CMD ["./main"]
