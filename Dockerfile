FROM golang

WORKDIR /gopath/goojle
ENV GOPATH /gopath
ADD . /gopath/

RUN go get github.com/shaalx/gooj
RUN go get github.com/everfore/exc
RUN go get github.com/qiniu/log
RUN go get github.com/shaalx/goutils
RUN go get github.com/astaxie/beego
RUN go get github.com/everfore/rpcsv
RUN go build -o gooe

EXPOSE 80
CMD ["/gopath/goojle/gooe"]
