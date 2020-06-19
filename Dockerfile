FROM golang:1.13 AS builder
ENV CGO_ENABLED 0
WORKDIR /go/src/app
ADD . .
RUN go build -o /esloadtpl

FROM scratch
COPY --from=builder /esloadtpl /esloadtpl
WORKDIR /workspace
CMD ["/esloadtpl"]
