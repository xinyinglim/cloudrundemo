FROM golang:1.14 as builder
WORKDIR /app
COPY go.* ./
RUN go mod download

COPY server.go server_test.go ./

#Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server

FROM alpine:3
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /server

# ENTRYPOINT ["bash"] # causes the following error
# Error response from daemon: OCI runtime create failed: container_linux.go:349: starting container process caused "exec: \"bash\": executable file not found in $PATH": unknown.

#port env variable is needed by cloud run deployment, otherwise will fail
ENV PORT=8081
CMD ["/server"]