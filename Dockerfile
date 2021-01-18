FROM golang:1.14-alpine AS build
RUN apk add git

WORKDIR /src/
COPY ./Golang-receiver/main.go /src/receiver/
COPY ./health-check/health.go /src/health/

RUN go get github.com/michaelbironneau/asbclient

WORKDIR /src/receiver
RUN CGO_ENABLED=0 go build  -o /bin/demo

WORKDIR /src/health
RUN CGO_ENABLED=0 go build  -o /bin/health

FROM alpine
COPY --from=build /bin/demo /bin/demo
COPY --from=build /bin/health /bin/health

ENTRYPOINT ["/bin/demo"]
