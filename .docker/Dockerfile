#@version 3.7.5
FROM golang:1.14.6-alpine3.12 as base

WORKDIR /go/src/app
COPY ./go/src .

RUN go get -v ./...  \
    && go build -ldflags="-s -w"
RUN ls -la

CMD ["app"]

FROM scratch as prod

COPY --from=base /go/src/app/app .

CMD ["/app"]
