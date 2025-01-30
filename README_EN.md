# deepseek-api #
[**中文**](./README.md) | [**English**](./README_EN.md)

This project has implemented a Go client for the DeepSeek API, making it convenient for everyone to use.

## Content Guide ##
* [Introduction](#Introduction)
* [Milestone](#Milestone)
* [Version](#Version)
* [Deployment](#Deployment)
* [Quick Start](#Quick-Start)
* [Feedback](#Feedback)

## Introduction ##
#### How to use the DeepSeek API service
* [DeepSeek](https://www.deepseek.com/)
* [DeepSeek API 文档](https://api-docs.deepseek.com/zh-cn/)
* [DeepSeek API Docs](https://api-docs.deepseek.com/)
#### So far, deepseek-api can provide the following support:
* Chat
>> Create Chat Completion
* Completions
>> Create FIM Completion (Beta)
* Models
>> Lists Models
* Others
>> Get User Balance

## Milestone ##
* Implement the basic client for the DeepSeek API service < latest

## Version ##
#### Beta Version
* master
#### Release Version
* null

## Deployment ##
The deployment of deepseek-api relies on Go modules. If you don't have go mod yet, you need to initialize it first.
```sh
go mod init myproject
```
Install deepseek-api
```sh
go get -u github.com/ZSLTChenXiYin/deepseek-api
```

## Quick Start ##
* Please refer to the examples in [deepseek_cli_test.go](./deepseek_api_test/deepseek_cli_test.go).

## Feedback ##
* Chen Xiyin will check the [Issues](https://github.com/ZSLTChenXiYin/deepseek-api/issues) from Friday to Sunday every week and will also conduct live broadcasts on Bilibili irregularly.
>> Chen Xiyin's e-mail: imjfoy@163.com
>> 
>> Chen Xiyin's Bilibili UID: 352456302
