FROM golang:1.25

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /publisher

EXPOSE 8080

CMD [ "/publisher" ]
