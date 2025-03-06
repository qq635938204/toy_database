package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	_ "toy_database/routers"
	"toy_database/tool"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	if logFile, err := os.OpenFile("./toy_database.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644); err != nil {
		fmt.Printf("open log file error:%v\r\n", err)
	} else {
		multiWriter := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(multiWriter)
	}
}

func getRunPath() (string, error) {
	var ret string
	var retErr error
	if runpath, err := os.Executable(); err != nil {
		retErr = err
	} else {
		if runtime.GOOS == "windows" {
			ret = runpath[:strings.LastIndex(runpath, "\\")]
		} else {
			ret = runpath[:strings.LastIndex(runpath, "/")]
		}
	}
	return ret, retErr
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	if runPath, err := getRunPath(); err != nil {
		log.Println("get run path error:", err)
	} else {
		tool.SetRootPath(runPath)
		logs.Info("run path:", runPath)
	}
	beego.SetStaticPath("/", "static")
	beego.SetStaticPath("/image", "data/image")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers",
			"Content-Type", "Access-Token", "Accept", "x-requested-with", "Domain"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowOrigins:  []string{"*"},
	}))
	go beego.Run()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	if err := beego.BeeApp.Server.Close(); err != nil {
		log.Printf("close server error:%v\n", err)
	}
	log.Println("End...")
}
