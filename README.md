# yanblue
这是一个基于我自己常用的 [Go_Web](https://github.com/xyb7910/web_go) 开发常用的脚手架训练项目。

项目较为简单，实现了用户登录，注册， 创建帖子，发布帖子，评论以及点赞等功能。

目前功能较为简单，日后会进行更新迭代。

欢迎大家 Star 🌟 ！！！

## 项目架构

前段使用 Vue 开发

后端使用 Go-Web 常用的 `CLD` 分层设计

## 涉及技术

- Vue
- Gin
- MySQL
- Redis
- Viper
- Zap
- JWT
- ...

## 部分实现原理
- 用户ID，帖子ID，采用雪花生成算法
- 用户密码加密以及验证机制使用 JWT Token
- 路由请求限制使用令牌桶限流策略
- ...

## 项目下载
```bash
git clone git@github.com:xyb7910/yanblue.git

cd yanblue

go mod tidy
```

## 感谢

最后，感谢[李文周](https://www.liwenzhou.com/)博客分享相关技术。