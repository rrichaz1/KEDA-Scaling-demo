FROM golang:1.14-alpine AS build
RUN apk add git

WORKDIR /src/
COPY ./Golang-receiver/main.go /src/

RUN go get github.com/michaelbironneau/asbclient

RUN CGO_ENABLED=0 go build  -o /bin/demo

FROM alpine
COPY --from=build /bin/demo /bin/demo

ENTRYPOINT ["/bin/demo"]
