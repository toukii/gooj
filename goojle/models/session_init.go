package models

import (
	"github.com/astaxie/beego/session"

	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var GlobalSessions *session.Manager

func init() {
	// GlobalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	GlobalSessions, _ = session.NewManager("file", `{"cookieName":"gosessionid","sessionsavepath":"./sessionpath/", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	// GlobalSessions, _ = session.NewManager("mysql", `{"cookieName":"gosessionid","sessionsavepath":"./sessionpath/", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": "root:1234@tcp(localhost:3306)/session?charset=utf8"}`)
	// GlobalSessions, _ = session.NewManager("mysql", `{"cookieName":"gosessionid","sessionsavepath":"./sessionpath/", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": "cdb_outerroot:root1234@tcp(55c354e17de4e.sh.cdb.myqcloud.com:7276)/session?charset=utf8"}`)
	defer GlobalSessions.GC()
}
