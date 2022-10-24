# This is a multi-stage Dockerfile and requires >= Docker 17.05
# https://docs.docker.com/engine/userguide/eng-image/multistage-build/
FROM golang:1.14-alpine3.13 as builder

ENV GO111MODULE on
ENV GOPROXY http://proxy.golang.org

RUN mkdir -p /src/mqtt-auth
WORKDIR /src/mqtt-auth

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

ADD . .
RUN CGO_ENABLED=0 /usr/local/go/bin/go build -o ./main -a -ldflags '-extldflags "-static"' main.go

FROM scratch

WORKDIR /mqtt-auth

COPY --from=builder /src/mqtt-auth/main /mqtt-auth/.
ENTRYPOINT ["/mqtt-auth/main"]
