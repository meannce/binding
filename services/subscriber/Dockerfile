FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
# Copy go.sum only if it exists
COPY go.sum* ./

RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o subscriber .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/subscriber .
CMD ["./subscriber"]
