FROM golang:1.19-alpine as builder

WORKDIR /app

RUN apk update && apk add git py3-pip gcc musl-dev python3-dev libffi-dev openssl-dev cargo make && pip install --upgrade pip && pip install azure-cli

RUN git clone https://github.com/Qumber-ali/akv-secret-loader.git . &&\
               git checkout $tag

RUN go get -d ./... && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
        -ldflags="-w -s" -o /seclo

FROM scratch

WORKDIR /

COPY --from=builder /seclo /usr/bin/

CMD ["--help"] 

ENTRYPOINT ["seclo"]

