FROM golang:1.19-alpine as builder

WORKDIR /app

RUN apk update && apk add git

ARG tag

RUN git clone https://github.com/Qumber-ali/akv-secret-loader.git . &&\
         git checkout $tag && git describe --tags

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
        -ldflags="-w -s" -o /seclo

FROM scratch

WORKDIR /

COPY --from=builder /seclo /usr/bin/

CMD ["seclo", "--help"] 


