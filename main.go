package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"mycode/fetchall/conf"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
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
	//获取当前时间
	start := time.Now()
	//创建一个chan
	ch := make(chan result)
	index := 0
	readFile, err := ioutil.ReadFile("api.yaml")
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(readFile))
	//projects 用户传入的参数
	projects := conf.Projects{}
	err = yaml.Unmarshal(readFile, &projects)
	//debug
	//fmt.Println(projects)
	//根据yaml读取输入参数
	for _, project := range projects.Project {
		input := inputStruct{}
		//传入的参数
		input.project = project.Projectname
		input.method = project.Method
		for _, content := range project.Context {
			input.url = content.URL
			input.name = content.Name

			runProject := false
			runURL := false

			//根据 -P 入参判断是否执行
			if *flagProject == "" {
				runProject = true
			} else if strings.Contains(input.project, *flagProject) {
				runProject = true
			}

			//根据 -U 入参判断是否执行
			if *flagURL == "" {
				runURL = true
			} else if strings.Contains(input.url, *flagURL) {
				runURL = true
			} else {
				runURL = false
			}
			//如果执行了
			if runProject && runURL {
				go fetch(input, ch)
				index++
			}
		}
	}
	errCount := 0
	successfulCount := 0
	//根据channel输出
	for i := 1; i <= index; i++ {
		fmt.Printf("%2d. ", i)
		output := <-ch

		if output.error != nil {
			//log.Error(output.name, "\t", output.info, "\t",output.error)
			log.WithFields(log.Fields{
				"name": output.name,
				"Url":  output.info,
				"err":  output.error}).Error("Sorry don't connect:")
			errCount++
		} else {
			log.WithFields(log.Fields{
				"Times": output.time,
				"name":  output.name,
			}).Info("Successful !")
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
func fetch(input inputStruct, ch chan<- result) {
	//设定prefix
	prefix := input.project + " " + input.name
	//开始时间
	start := time.Now()
	var resp *http.Response
	var err error
	if input.method == "GET" {
		//使用get方法获取resp
		resp, err = http.Get(input.url)
	} else if input.method == "POST" {
		params := url.Values{}
		resp, err = http.PostForm(input.url, params)
	} else {
		err = errors.New(" Wrong request method")
		ch <- result{
			prefix, input.url, err, "",
		}
	}
	if err != nil {
		//如果错误，发送到信道
		ch <- result{
			prefix, input.url, err, "",
		}
		return
	}
	_, err = io.Copy(ioutil.Discard, resp.Body)
	//不需要获取resource
	if err != nil {
		//出现错误
		ch <- result{
			prefix, input.url, err, "",
		}
		return
	}
	err = resp.Body.Close()
	//获取时间
	secs := time.Since(start).Seconds()
	//导出数据
	ch <- result{
		prefix,
		"",
		nil,
		fmt.Sprintf("%.2f", secs),
	}
}

//result 出信道的参数
type result struct {
	name  string
	info  string
	error error
	time  string
}

//inputStruct 输入信道的参数
type inputStruct struct {
	//project name 项目名称
	project string
	//name 部署的名称
	name string
	//method GET or POST方法
	method string
	//url URL地址
	url string
}
