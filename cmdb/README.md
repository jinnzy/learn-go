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
go ini 配置管理
```
govendor fetch github.com/go-ini/ini
```
2 目录结构
conf  配置文件
common 通用方法，全局变量       
controller 写router的控制方法
database 数据库连接           
model 数据相关
middleware 中间件
router 路由
pkg 包
pkg/utiles 工具类方法 时间加密之类还有分页
main  入口
