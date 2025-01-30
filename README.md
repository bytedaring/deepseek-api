# deepseek-api #
[**中文**](./README.md) | [**English**](./README_EN.md)

本项目实现了DeepSeek API的Go客户端，方便大家使用。

## 内容导引 ##
* [介绍](#介绍)
* [里程碑](#里程碑)
* [版本实现](#版本实现)
* [部署](#部署)
* [快速上手](#快速上手)
* [问题反馈](#问题反馈)

## 介绍 ##
#### 如何使用DeepSeek API服务
* [DeepSeek](https://www.deepseek.com/)
* [DeepSeek API 文档](https://api-docs.deepseek.com/zh-cn/)
* [DeepSeek API Docs](https://api-docs.deepseek.com/)
#### 目前为止，deepseek-api可以提供以下支持：
* 对话（Chat）
>> 对话补全
* 补全（Completions）
>> FIM补全（Beta）
* 模型（Model）
>> 列出模型
* 其他
>> 查询余额

## 里程碑 ##
* 实现DeepSeek API服务的基础客户端 < latest

## 版本实现 ##
#### 测试版
* master
#### 正式版
* 无

## 部署 ##
deepseek-api的部署依赖Go modules，如果你还没有go mod，你需要首先初始化:
```sh
go mod init myproject
```
安装 deepseek-api
```sh
go get -u github.com/ZSLTChenXiYin/deepseek-api
```

## 快速上手 ##
* 请参考 [deepseek_cli_test.go](./deepseek_api_test/deepseek_cli_test.go) 中的示例

## 问题反馈 ##
* 陈汐胤会在每周五至周日查看 [Issues](https://github.com/ZSLTChenXiYin/deepseek-api/issues)，还会不定期地在bilibili直播。
>> 陈汐胤的e-mail: imjfoy@163.com
>> 
>> 陈汐胤的bilibili UID: 352456302
