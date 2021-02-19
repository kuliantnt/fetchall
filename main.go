package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"mycode/fetchall/service"
	"os"
)

//是否通过项目名称过滤
var flagProject = flag.String("P", "", "Project name")

//是否通过URL过滤
var flagURL = flag.String("U", "", "Url name")

func init() {
	//解析flag
	flag.Parse()
	//设置日志输出位置
	log.SetOutput(os.Stdout)
}

func main() {
	service.DoFetch(flagProject, flagURL)
}
