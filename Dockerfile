FROM golang

RUN mkdir /opt/app

ADD ./server /opt/app

WORKDIR /opt/app

RUN go build -o main .

CMD ["/opt/app/main"]