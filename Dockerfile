FROM golang:1.16-alpine3.14 AS builder

WORKDIR /fizzbuzz
ADD / .

RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o /go/bin/fizzbuzzapi main.go


FROM scratch AS runner

COPY --from="builder" /go/bin/ /bin
ADD /fizzbuzz.localhost-key.pem .
ADD /fizzbuzz.localhost.pem .

ENTRYPOINT ["fizzbuzzapi"]
