FROM golang:1.22-alpine AS builder

WORKDIR /usr/local/src

COPY ["go.mod", "go.sum", "./"]
RUN --mount=type=cache,target=/go/pkg/mod/ \
    go mod download -x
COPY . .

RUN go build -o bin/migrator main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/migrator /

CMD [ "/migrator" ]