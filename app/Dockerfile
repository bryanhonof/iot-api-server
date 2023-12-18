FROM golang:bookworm AS builder

ENV CGO_ENABLED=1
RUN apt update -y
RUN apt upgrade -y
RUN apt install -y gcc build-essential
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
RUN go get -d -v
RUN go mod download
RUN go mod verify
RUN go build -o /dist/bin/server


FROM golang:bookworm AS runner

RUN apt update -y
RUN apt upgrade -y
COPY --from=builder /dist/bin/server /usr/local/bin/server
EXPOSE 8080
WORKDIR /var/lib/iot-api-server
ENTRYPOINT ["/usr/local/bin/server"]