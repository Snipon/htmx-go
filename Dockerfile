FROM golang:alpine as builder
WORKDIR /src
COPY . .
RUN apk update && apk add --no-cache git
RUN go build  -o api

FROM alpine
WORKDIR /app
COPY --from=builder /src/api /app
ENV MONGO_URL=mongodb://mongo:27017
EXPOSE 8080
CMD ["/app/api"]
