# syntax=docker/dockerfile:1
FROM golang:trixie AS builder
WORKDIR /src

COPY go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/hello ./main.go

FROM scratch AS production
COPY --from=builder /bin/hello /bin/hello
EXPOSE 8080
ENTRYPOINT ["/bin/hello"]