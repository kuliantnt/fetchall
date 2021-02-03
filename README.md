# fetchall

<!-- ![GitHub issues](https://img.shields.io/github/issues/kuliantnt/fetchall)
![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/kuliantnt/fetchall?include_prereleases)
![GitHub last commit](https://img.shields.io/github/last-commit/kuliantnt/fetchall)
![GitHub issues](https://img.shields.io/github/issues-raw/kuliantnt/fetchall)
![GitHub pull requests](https://img.shields.io/github/issues-pr/kuliantnt/fetchall)
![GitHub](https://img.shields.io/github/license/kuliantnt/fetchall) -->

## 内容列表

- [背景](#背景)
- [安装](#安装)
- [使用说明](#使用说明)
  - [命令行参数](#命令行参数)
  - [结果参数](#结果参数)
- [示例](#示例)
- [维护者](#维护者)
- [如何贡献](#如何贡献)
- [使用许可](#使用许可)

## 背景

工作之余写的一个接口检测小程序，用于拥有大量接口的业务进行批量检测。

第一种情况是每次部署服务后需要进行测试，才能形成闭环。测试方法是通过`curl -I` 服务地址进行检测。在进行大规模服务部署或更新之后，手动`curl`比较麻烦，还容易输入错误。

第二种情况是客户运营人员，通知运维人员业务中断，运维人员可以通过该程序进行测试，精准定位故障原因。

至于使用GO语言而不是用Java或者Shell脚本进行编写，因为GO语言相对于臃肿的Java来说，可以生成可执行二进制文件，方便维护，可执行文件体积小，无需安装运行环境。相对于shell，Go编写的程序可扩展能力强，可以完成成更复杂的功能。

## 安装

这个项目无需安装其他包，只需要编写`api.yaml` 参考如下

```yaml
projects:
  - projectname: mail
    method: GET
    context:
      - { url: https://aaa/mail/test1, name: side:aaa }
      - { url: https://bbb/mail/test1, name: side:bbb }
  - projectname: api
    method: GET
    context:
      - { url: https://aaa/api/test1, name: side:aaa }
      - { url: https://bbb/api/test1, name: side:bbb }
      - { url: https://ccc/api/test1, name: size:ccc }
  - projectname: socket
    method: POST
    context:
      - { url: https://ccc/api/socket, name: side:socket }
```

然后使用

```shell
./fetchall
```

就可以启动

## 使用说明

### 命令行参数

#### -P [服务名称]

只测试哪些服务，如只测试Google服务。其他服务不检测

#### -U [url地址]

只测试哪些URL地址，如只测试www.google.com上面的服务

### 结果参数

|名称|说明|
|---|---|
|INFO|接口测试通过|
|ERRO|接口测试失败|
|Project|服务名称，由于一个服务分布式部署，可以有多个接口|
|Times|接口响应时间|
|URL|接口测试地址|
|err|接口测试地址|

## 示例

```yaml
projects:
  - projectname: google
    method: GET
    context:
      - { url: https://mail.google.com/mail/u/0/#inbox, name: google_mail }
      - { url: https://www.google.com/, name: google_search }
  - projectname: Golang
    method: GET
    context:
      - { url: https://docs.google.com/spreadsheets/u/0/, name: google_doc }
      - { url: https://studygolang.com/pkgdoc, name: pkgdoc }
      - { url: https://golang.org/pkg, name: golang_pkg }
  - projectname: cpepc
    method: GET
    context:
      - { url: http://cpepc.com.cn, name: 001 }
```

输出结果

```log
 1. INFO[0000] Successful !                                  Times=0.72 name="Golang pkgdoc"
 2. INFO[0005] Successful !                                  Times=5.26 name="google google_mail"
 3. INFO[0006] Successful !                                  Times=6.63 name="Golang golang_pkg"
 4. INFO[0012] Successful !                                  Times=12.03 name="google google_search"
 5. INFO[0012] Successful !                                  Times=12.10 name="Golang google_doc"
 6. ERRO[0027] Sorry don't connect:                          Url="http://cpepc.com.cn" err="Get \"http://cpepc.com.cn\": EOF" name="cpepc 001"
Over time:      27.32s
Success count   5 time
Error count:    1 times
```

## 维护者

[@kuliantnt](https://github.com/kuliantnt/)。

## 如何贡献

非常欢迎你的加入！[提一个 Issue](https://github.com/kuliantnt/fetchall/issues/new) 或者提交一个 Pull Request。

## 使用许可

