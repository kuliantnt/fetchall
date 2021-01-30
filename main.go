package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	conf "mycode/fetchall/conf"
	"net/http"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var flagProject = flag.String("P", "", "Project name")
var flagURL = flag.String("U", "", "Url name")

func init() {
	//解析flag
	flag.Parse()
	//设置日志输出位置
	log.SetOutput(os.Stdout)
}

type input struct {
}

func main() {
	content, err := ioutil.ReadFile("api.yaml")
	if err != nil {
		panic(err)
	}
	// print(string(content))
	env := conf.Projects{}
	err = yaml.Unmarshal(content, &env)
	fmt.Println(err, env)
	return
}

func domain() {

	//获取当前时间
	start := time.Now()
	//创建两个chan
	ch := make(chan result)
	//获取当前目录
	pwd, _ := os.Getwd()
	//读取api.list
	bytes, err := ioutil.ReadFile(pwd + "/api.list")
	if err != nil {
		log.Fatal(err)
	}
	index := 0
	//根据空白分割符获取切片
	urlSince := strings.Fields(string(bytes))
	//新建一个变量
	var prefix string

	for _, url := range urlSince {
		if strings.HasPrefix(url, "#") {
			//去掉前面的#
			prefix = strings.TrimPrefix(url, "#")
			//prefix = url
			continue
		}
		//根据 -P 入参判断是否执行
		runProject := false
		var runURL bool = false
		if *flagProject == "" {
			runProject = true
		} else if strings.Contains(prefix, *flagProject) {
			runProject = true
		}

		//根据 -U 入参判断是否执行
		if *flagURL == "" {
			runURL = true
		} else if strings.Contains(url, *flagURL) {
			runURL = true
		} else {
			runURL = false
		}
		//实际上执行的语句
		if runProject && runURL {
			index++
			go fetch(url, prefix, ch)
		}
	}

	//用于统计错误信息
	errCount := 0
	successfulCount := 0
	//根据channel输出
	for i := 1; i <= index; i++ {
		fmt.Printf("%2d. ", i)
		output := <-ch

		if output.error != nil {
			//log.Error(output.name, "\t", output.info, "\t",output.error)
			log.WithFields(log.Fields{
				"Project": output.name,
				"Url":     output.info,
				"err":     output.error}).Error("Sorry don't connect:")
			errCount++
		} else {
			log.WithFields(log.Fields{
				"Times":   output.time,
				"Project": output.name,
				"Url":     output.info}).Info("Successful !")
			successfulCount++
		}
	}
	//显示统计信息
	fmt.Printf("Over time:\t%.2fs\n"+
		"Success count\t%d time\n"+
		"Error count:\t%d times\n"+
		"By Lin\n",
		time.Since(start).Seconds(), successfulCount, errCount)
}

//根据url获取
func fetch(url, prefix string, ch chan<- result) {
	//开始时间
	start := time.Now()
	//使用get方法获取resp
	resp, err := http.Get(url)
	if err != nil {
		//如果错误，发送到信道
		ch <- result{
			prefix, url, err, "",
		}
		return
	}
	_, err = io.Copy(ioutil.Discard, resp.Body)
	//不需要获取resource
	if err != nil {
		//出现错误
		ch <- result{
			prefix, url, err, "",
		}
		return
	}
	err = resp.Body.Close()
	//获取时间
	secs := time.Since(start).Seconds()
	//导出数据
	ch <- result{
		prefix,
		url,
		nil,
		fmt.Sprintf("%.2f", secs),
	}
}

type result struct {
	name  string
	info  string
	error error
	time  string
}
