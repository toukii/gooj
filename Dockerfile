FROM golang

WORKDIR /gopath/gooj
ENV GOPATH /gopath
ADD . /gopath/

RUN go get github.com/shaalx/gooj
RUN go get github.com/everfore/exc
RUN go get github.com/qiniu/log
RUN go get github.com/shaalx/gooj
RUN go get github.com/shaalx/goutils
RUN go build

EXPOSE 80
CMD ["/gopath/gooj/gooj"]