FROM golang:latest as builder

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY ./manager/ ./manager/
COPY ./network/ ./network/
COPY ./main.go ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/main \
    -ldflags '-s -w'

EXPOSE 3000

FROM scratch as runner

COPY --from=builder /go/bin/main /app/main

# install ca
COPY cafile.pem /usr/local/share/ca-certificates/
RUN apk add --no-cache ca-certificates && \
    update-ca-certificates

ENTRYPOINT ["/app/main"]
