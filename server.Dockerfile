FROM golang:1.22.2-alpine as builder


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY pb/ pb/

COPY server/ server/

RUN go build -o grpc-server server/main.go


FROM gcr.io/distroless/base


COPY --from=builder /app/grpc-server /app/grpc-server

COPY --from=builder /app/pb /app/pb


ENTRYPOINT ["/app/grpc-server"]
