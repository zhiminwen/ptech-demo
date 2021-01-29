FROM golang as builder
WORKDIR /build
COPY *.go /build

RUN CGO_ENABLED=0 go build -o serving *.go

FROM alpine
WORKDIR /app
COPY --from=builder /build/serving /app

ENTRYPOINT ["./serving"]
  
