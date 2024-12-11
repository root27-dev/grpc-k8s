
FROM golang:1.22.2-alpine as builder


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY pb/ pb/

COPY api/ api/

RUN go build -o api-client api/main.go


FROM gcr.io/distroless/base


COPY --from=builder /app/api-client /app/api-client

COPY --from=builder /app/pb /app/pb


ENTRYPOINT ["/app/api-client"]
