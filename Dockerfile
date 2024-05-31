FROM golang:alpine as builder
WORKDIR /src
COPY . .
RUN apk update && apk add --no-cache git
RUN go build  -o api

FROM alpine
WORKDIR /app
COPY --from=builder /src/api /app
CMD ["/app/api"]
