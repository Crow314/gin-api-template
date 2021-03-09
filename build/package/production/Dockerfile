FROM golang:1.15.8 as builder

WORKDIR /go/src/app

# download modules (usually using cache)
COPY go.mod .
COPY go.sum .

RUN go mod download

# copy source codes
COPY . .

# build
RUN CGO_ENABLED=0 go build -o /go/bin/app ./cmd/main.go


FROM gcr.io/distroless/base

# copy application binary
COPY --from=builder /go/bin/app .

# copy necessary files
COPY --from=builder /go/src/app/configs/ ./configs/

ENV APP_ENV=production TZ='Asia/Tokyo'

EXPOSE 8080
ENTRYPOINT [ "./app" ]
