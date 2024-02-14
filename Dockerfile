FROM golang:1.22-alpine AS builder

RUN apk add alpine-sdk

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY Makefile .

RUN go mod download

COPY . .

RUN make build

FROM alpine AS runner

WORKDIR /app

COPY ./data/ ./data

COPY --from=builder /app/bin/importer .

CMD ["time","./importer"]