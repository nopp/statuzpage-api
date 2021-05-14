FROM golang:1.15.8-alpine as builder

WORKDIR /statuzpage-api

ADD . /statuzpage-api

RUN go build

FROM golang:1.15.8-alpine

LABEL maintainer "Carlos Augusto Malucelli <camalucelli@gmail.com>"

WORKDIR /statuzpage-api

COPY  config.json /etc/statuzpage-api/config.json
COPY --from=builder /statuzpage-api/statuzpage-api .

EXPOSE 8000

CMD ["./statuzpage-api"]