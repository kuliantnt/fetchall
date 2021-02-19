package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mycode/fetchall/conf"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

//Result 出信道的参数
type result struct {
	name  string
	info  string
	error error
	time  string
}

//InputStruct 输入信道的参数
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

//用于debug的打印方法
func (is inputStruct) PrintStr() string {
	return fmt.Sprintf("Project: %s Name: %s Method: %s url: %s", is.project, is.name, is.method, is.url)
}

func (res result) PrintStr() string {
	return fmt.Sprintf("name: %s info: %s error: %s time: %s", res.name, res.info, res.error.Error(), res.time)
}

//fetch 根据url获取
func fetch(ctx context.Context, input inputStruct, ch chan<- result) {
	//设定prefix
	prefix := input.project + " " + input.name
	//开始时间
	start := time.Now()
	var resp *http.Response
	var err error
	if input.method == "GET" {
		//使用get方法获取resp
		select {
		case <-ctx.Done():
			err = errors.New(prefix + "time out")
			ch <- result{
				prefix, input.url, err, "",
			}
			return
		default:
			resp, err = http.Get(input.url)
		}
	} else if input.method == "POST" {
		select {
		case <-ctx.Done():
			err = errors.New(prefix + "time out")
			ch <- result{
				prefix, input.url, err, "",
			}
			return
		default:
			params := url.Values{}
			resp, err = http.PostForm(input.url, params)
		}
	} else {
		err = errors.New("Wrong request method")
		//获取时间
		ch <- result{
			prefix, input.url, err, "",
		}
		return
	}
	if err != nil {
		//用来显示错误码
		if resp != nil {
			err = errors.New(prefix + "error code: " + strconv.Itoa(resp.StatusCode))
		}
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

//DoFetch 实际上执行的参数
func DoFetch(flagProject *string, flagURL *string) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	//创建一个waitGroup
	//wg := sync.WaitGroup{}
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
				go fetch(ctx, input, ch)
				index++
			}
		}
	}
	//wg.Add(index)
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
	//直接结束调所有开启的子进程
	cancelFunc()
	//显示统计信息
	fmt.Printf("Over time:\t%.2fs\n"+
		"Success count\t%d time\n"+
		"Error count:\t%d times\n"+
		"By Lin\n",
		time.Since(start).Seconds(), successfulCount, errCount)
}
