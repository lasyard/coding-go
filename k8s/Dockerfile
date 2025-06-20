FROM golang:1.24.4-alpine3.22 AS base

FROM base AS builder

WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /usr/local/bin/create-cm /workspace/create/cm

FROM base AS runner

COPY --from=builder /usr/local/bin/create-cm /usr/local/bin/create-cm

CMD ["tail", "-f", "/dev/null"]
