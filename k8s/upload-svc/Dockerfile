FROM golang:1.26-alpine AS base

FROM base AS builder

WORKDIR /

# download may fail, so do it outside and `go mod vendor`, then do vendor building here
ENV GOFLAGS="-mod=vendor"

COPY . .
RUN go build -o /usr/local/bin/upload-svc /cmd/upload-svc

FROM base AS runner

COPY --from=builder /usr/local/bin/upload-svc /usr/local/bin/upload-svc

EXPOSE 8080

CMD ["/usr/local/bin/upload-svc"]
