FROM golang:1.19-alpine as builder

WORKDIR /app

ARG tag

RUN apk update && apk add git

RUN git clone https://github.com/Qumber-ali/akv-secret-loader.git . &&\
               git checkout $tag

RUN go get -d ./... && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
        -ldflags="-w -s" -o /seclo

FROM scratch

WORKDIR /

COPY --from=builder /seclo /usr/bin/

CMD ["--help"] 

ENTRYPOINT ["seclo"]

