# fetchall

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

## 内容列表

- [背景](#背景)
- [安装](#安装)
- [使用说明](#使用说明)
  - [生成器](#生成器)
- [徽章](#徽章)
- [示例](#示例)
- [相关仓库](#相关仓库)
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
#接口地址1
http:xxx.com/aaa/bbb
http:xxx.com/aaa/bbb
#接口地址2
http:xxx.com/aaa/ccc
```

然后使用

```shell
./fetchall
```

## 使用说明

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
```。

## 维护者

[@kuliantnt](https://github.com/kuliantnt/)。

## 如何贡献

非常欢迎你的加入！[提一个 Issue](https://github.com/kuliantnt/fetchall/issues/new) 或者提交一个 Pull Request。

## 使用许可

[MIT](LICENSE) © Richard Littauer
