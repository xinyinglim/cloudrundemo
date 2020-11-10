FROM golang:1.14 as builder
WORKDIR /app
COPY go.* ./
RUN go mod download

COPY main.go ./

#Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server

FROM alpine:3
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /server

ENV PORT=8080
CMD ["/server"]