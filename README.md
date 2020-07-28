## nginx 灰度发布应用的后端

###
- 1 项目使用 beego web框架和go mod依赖管理
- 2 新建一个数据库，名字跟配置文件的一样，把 data.sql 脚本执行
- 3 进入项目根目录，执行 go run main.go 会自动下载相关依赖并启动
- 4 浏览器打开 https://localhost:8090/node/page 检验项目是否能正常运行
