FROM golang

WORKDIR /gopath/app
ENV GOPATH /gopath
ADD . /gopath/

RUN go get github.com/shaalx/gooj
RUN go get github.com/everfore/exc
RUN go get github.com/qiniu/log
RUN go get github.com/shaalx/gooj
RUN go get github.com/shaalx/goutils
RUN go get github.com/astaxie/beego
RUN go build -o Goojle ./goojle/ 

EXPOSE 80
CMD ["/gopath/app/Goojle"]