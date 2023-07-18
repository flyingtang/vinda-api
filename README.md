## 本项目：主要是文章服务

## 总项目概况

 * 本项目目的是为了熟悉Golang开发的一个博客，界面简陋🤦‍

 * 前端采用umi+antd  [项目地址](https://github.com/txg5214/vinda-web)

 * 后端采用Golang(gin) + mysql(sqlx) [项目地址](https://github.com/txg5214/vinda-api)

 * ~~可能还有移动端, flutter试开发 [项目地址](https://github.com/txg5214/sunshine)~~


## 部署情况

 * 目前采用最笨的部署方式：手动docker部署🤦‍♀️ （dockerfile + docker-compose + shell)

 * 前端采用nginx容器+静态文件+路由拦截转达到后端API

 * 后端采用两个小项目 [文章服务](https://github.com/txg5214/vinda-api)、[视频服务](https://github.com/txg5214/vinda-video)

 * 数据来源于用node.js写得[两个简单爬虫服务](https://github.com/txg5214/fetch-data)

 * [线上演示地址：](https://flyingtang.com) 交互还有问题，会改进😁

