FROM golang:1.17.1 as develop

WORKDIR /go/src/app

# download modules (usually using cache)
COPY go.mod .
COPY go.sum .

RUN go mod download

# copy source codes
COPY . .

ENV APP_ENV=development TZ='Asia/Tokyo'

EXPOSE 8080
CMD go run cmd/main.go
