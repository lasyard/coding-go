FROM golang:1.24.4-alpine3.22 AS base

FROM base AS builder

WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /usr/local/bin/create-cm /workspace/create/cm
RUN go build -o /usr/local/bin/create-dyn /workspace/create/dyn
RUN go build -o /usr/local/bin/uploader /workspace/upload

FROM base AS runner

COPY --from=builder /usr/local/bin/create-cm /usr/local/bin/create-cm
COPY --from=builder /usr/local/bin/create-dyn /usr/local/bin/create-dyn
COPY --from=builder /usr/local/bin/uploader /usr/local/bin/uploader

EXPOSE 8080

#CMD ["sh", "-c", "trap exit INT TERM; sleep infinity & wait"]
CMD ["/usr/local/bin/uploader"]
