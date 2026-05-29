FROM golang:1.26-alpine AS base

FROM base AS builder

WORKDIR /

# COPY go.mod go.sum ./
# RUN go mod download && go mod verify

# download may fail, so do it outside and `go mod vendor`, then do vendor building here
ENV GOFLAGS="-mod=vendor"

COPY . .
RUN go build -o /usr/local/bin/create-cm /cmd/create-cm
RUN go build -o /usr/local/bin/create-cm-dyn /cmd/create-cm-dyn

FROM base AS runner

COPY --from=builder /usr/local/bin/create-cm /usr/local/bin/create-cm
COPY --from=builder /usr/local/bin/create-cm-dyn /usr/local/bin/create-cm-dyn

CMD ["sh", "-c", "trap exit INT TERM; sleep infinity & wait"]
