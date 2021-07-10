FROM arm64v8/golang:1.16 AS GOLANG

RUN mkdir /home/app

COPY . /home/app

WORKDIR /home/app

RUN go build .

FROM arm64v8/debian:buster-slim

WORKDIR  /usr/local/bin

COPY --from=GOLANG /home/app/static /usr/local/bin/static
COPY --from=GOLANG /home/app/auth-service /usr/local/bin/

EXPOSE 8010/tcp
EXPOSE 8010/udp

ENTRYPOINT ["auth-service"]