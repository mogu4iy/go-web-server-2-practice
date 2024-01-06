FROM golang:alpine AS builder
ENV CGO_ENABLED 0
ENV GOOS linux
RUN apk update --no-cache
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download && go mod verify
COPY . .
RUN go build -o /app/executable

FROM alpine
RUN apk update --no-cache
WORKDIR /app
COPY --from=builder /app/executable /app/executable
CMD ["./executable"]