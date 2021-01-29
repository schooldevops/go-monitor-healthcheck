FROM golang:alpine as builder

WORKDIR /src/

RUN apk --update add --no-cache ca-certificates openssl git tzdata && \
update-ca-certificates

COPY *.go /src/
COPY go.mod go.sum /src/
RUN go mod download
RUN CGO_ENABLED=0 go build -o /bin/demo

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /bin/demo /bin/demo
COPY ./data/targets.json /bin/data/targets.json
WORKDIR /bin/
ENTRYPOINT ["/bin/demo"]