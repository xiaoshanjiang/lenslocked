FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o ./server ./cmd/server/

FROM alpine
WORKDIR /app
COPY ./assets ./assets
COPY .env .env
# copy over the go binary built from the previous stage
COPY --from=builder /app/server ./server
CMD ["./server"]