# fetchall

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

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

这个项目无需安装其他包，只需要编写api.list

```shell
#接口名称1
http://接口1地址1
http://接口1地址2
#接口名称2
http://接口2地址1
http://接口2地址2
```

然后使用

```shell
./fetchall
```

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

```shell
#Google
https://mail.google.com/mail/u/0/#inbox
https://www.google.com/
https://docs.google.com/spreadsheets/u/0/
https://translate.google.com/
https://docs.google.com/document/u/0/
https://drive.google.com/drive/my-drive
https://plus.google.com/
https://www.youtube.com/#test
https://www.baidu.com
https://www.google.com
#Golang
https://studygolang.com/pkgdoc
https://gorm.io/zh_CN/
https://golang.org/pkg/
```

输出结果

```log
 1. INFO[0000] Successful !                                  Project=Golang Times=0.50 Url="https://gorm.io/zh_CN/"
 2. INFO[0000] Successful !                                  Project=Golang Times=0.57 Url="https://studygolang.com/pkgdoc"
 3. INFO[0000] Successful !                                  Project=Google Times=0.92 Url="https://www.google.com"
 4. INFO[0000] Successful !                                  Project=Google Times=0.92 Url="https://www.google.com/"
 5. INFO[0001] Successful !                                  Project=Google Times=1.34 Url="https://translate.google.com/"
 6. INFO[0001] Successful !                                  Project=Google Times=1.37 Url="https://www.youtube.com/#test"
 7. INFO[0001] Successful !                                  Project=Golang Times=1.41 Url="https://golang.org/pkg/"
 8. INFO[0001] Successful !                                  Project=Google Times=1.57 Url="https://docs.google.com/spreadsheets/u/0/"
 9. INFO[0001] Successful !                                  Project=Google Times=1.69 Url="https://docs.google.com/document/u/0/"
10. INFO[0002] Successful !                                  Project=Google Times=2.54 Url="https://mail.google.com/mail/u/0/#inbox"
11. INFO[0002] Successful !                                  Project=Google Times=2.54 Url="https://plus.google.com/"
12. INFO[0002] Successful !                                  Project=Google Times=2.62 Url="https://drive.google.com/drive/my-drive"
Over time:      2.62s
Success count   12 time
Error count:    0 times
```

## 维护者

[@kuliantnt](https://github.com/kuliantnt/)。

## 如何贡献

非常欢迎你的加入！[提一个 Issue](https://github.com/kuliantnt/fetchall/issues/new) 或者提交一个 Pull Request。

## 使用许可

[MIT](LICENSE) © kuliantnt
