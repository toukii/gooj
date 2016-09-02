FROM golang

WORKDIR /gopath/goojle
ENV GOPATH /gopath
ADD . /gopath/

RUN go get github.com/toukii/gooj/goojle && go get github.com/toukii/gooj/model_util && go get github.com/everfore/exc && go get github.com/qiniu/log && go get github.com/astaxie/beego && go get github.com/go-sql-driver/mysql && go get github.com/astaxie/beego/orm && go get github.com/astaxie/beego/toolbox  && go get github.com/astaxie/beego/session && go get github.com/astaxie/beego/session/mysql && go get github.com/astaxie/beego/validation && go get github.com/everfore/rpcsv && go get github.com/everfore/oauth/oauth2 && go get github.com/toukii/jsnm && go build -o gooe

EXPOSE 80

CMD ["/gopath/goojle/gooe"]
