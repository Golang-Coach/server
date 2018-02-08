FROM golang:latest
EXPOSE 80

RUN mkdir /app

ADD server_* /app/

WORKDIR /app

CMD ["/app/server_linux_amd64"]