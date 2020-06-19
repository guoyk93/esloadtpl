FROM golang:1.13 AS builder
ENV CGO_ENABLED 0
WORKDIR /go/src/app
ADD . .
RUN go build -o /estpload

FROM scratch
COPY --from=builder /estpload /estpload
WORKDIR /workspace
CMD ["/estpload"]
