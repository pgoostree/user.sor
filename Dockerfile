FROM golang

RUN mkdir /opt/app

ADD . /opt/app

WORKDIR /opt/app

RUN go build -o main .

EXPOSE 9000

CMD ["/opt/app/main"]