FROM arm64v8/golang:1.16 AS GOLANG

RUN mkdir /home/app

COPY . /home/app

WORKDIR /home/app

RUN go build .

FROM arm64v8/debian:buster-slim

WORKDIR  /usr/local/bin

COPY --from=GOLANG /home/app/user-service /usr/local/bin/
COPY --from=GOLANG /home/app/static /usr/local/bin/static

EXPOSE 8020/tcp
EXPOSE 8020/udp

ENTRYPOINT ["user-service"]