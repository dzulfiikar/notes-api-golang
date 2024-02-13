FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN ls -la

RUN CGO_ENABLED=0 GOOS=linux go build -o ./tmp/main ./cmd/main.go

EXPOSE 3000

CMD [ "./tmp/main" ]
