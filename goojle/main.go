package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	// "github.com/shaalx/gooj/goojle/models"
	_ "github.com/shaalx/gooj/goojle/routers"
)

func main() {
	// beego.EnableXSRF = true
	// go TaskSessionGC()
	beego.Run()
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

func TaskSessionGC() {
	tk := toolbox.NewTask("taska", "0/10 * * * * *", func() error {
		fmt.Println("hello world")
		// models.GlobalSessions.GC()
		return nil
	},
	)
	err := tk.Run()
	if err != nil {
		beego.Error(err)
	}
	toolbox.AddTask("taska", tk)
	toolbox.StartTask()
}
