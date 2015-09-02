FROM golang

WORKDIR /gopath/gooj
ENV GOPATH /gopath
ADD . /gopath/

RUN go get github.com/shaalx/gooj
RUN go build

EXPOSE 80
CMD ["/gopath/gooj/gooj"]