FROM golang:1.24-alpine

ENV CMD $cmd

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o simulation

ENTRYPOINT if [ "$cmd" = "test" ]; then \
    go test ./...; else \
    /app/simulation; \
    fi
