# 1 依赖
依赖管理工具
```
go get -u github.com/kardianos/govendor
使用govendor 管理服务依赖
govendor init 初始化
项目完成后可以使用govendor init    govendor add +external 添加相关依赖 
```
web 框架
```
govendor fetch github.com/gin-gonic/gin
```
日志库
```
govendor fetch go.uber.org/zap/zapcore
govendor fetch go.uber.org/zap
```
beego的表单验证
```
govendor fetch github.com/astaxie/beego/validation
```
mysql/gorm
```
github.com/go-sql-driver/mysql

```
