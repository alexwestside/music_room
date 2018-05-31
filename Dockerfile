FROM golang:1.10.1

ADD bin/server /usr/bin/

CMD /usr/bin/server