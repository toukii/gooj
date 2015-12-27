FROM golang

WORKDIR /gopath/goojle
ENV GOPATH /gopath
ADD . /gopath/

RUN go get github.com/shaalx/gooj && go get github.com/everfore/exc && go get github.com/qiniu/log && go get github.com/astaxie/beego && go get github.com/everfore/rpcsv && go build -o gooe

EXPOSE 80

CMD ["/gopath/goojle/gooe"]
