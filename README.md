## nginx 灰度发布项目

#### 项目大致介绍

- 背景：一种常见的应用灰度发布的实现
```
一个线上的应用部署有 4台服务器，nginx 做了负载均衡，策略可能为 ip_hash，轮询，权重，最少连接之类的。
每次的请求可能分发到不同的服务器

在应用发布的时候，先把 发布的文件部署到一台服务器上，并且把这台服务器设置为内部访问，nginx 这时把
这台服务器不加入到负载均衡里面，线上的流量只访问到其他3台节点。测试人员通过内部ip访问到这台新发布应用的
节点，在测试完成之后就把这台服务器重新加入到负载均衡里面。

以上的工作都是运维完成的，相当于重复的工作

```

- **[灰度发布的重点](#)**
    - **[在于在不重启 nginx 的情况下进行节点的上下线以及动态扩容](#)**
    - **[需要能够动态修改一个 服务名 对应的 upstream](#)**
    - **[这个动态的 upstream 需要提供几个接口，新增节点，移除节点，列出当前所有节点](#)**

- **[灰度发布的实现](#)**
    - **[nginx lua 部分](#)**
        - 1 利用lua中 "lua_shared_dict" 指令开辟一个共享内存空间，所有进程都可访问
        - 2 lua_shared_dict 这个数据结构相当于一个 Map或者 Dict，key 是服务名，value 是节点列表
        - 3 动态的 upstream 就是能够动态的修改 lua_shared_dict 里面一个服务名对应节点列表，增加或者删除
        - 4 upstream 的动态修改操作是通过接口，可以用一个web后台访问接口达到 增加或者删除的效果
    - **[管理后台部分](#)**
        - 后台提供页面，访问接口达到 增加或者删除的效果
        - 后台提供页面，访问接口查看当前某一个服务有几个负载的节点
        - 这个后台管理 是用 go + beego + vue 实现的
        - **[当前这个beego项目就是后台管理部分](#)**
        
    - **[管理后台前端页面部分](#)**
        - 前端部分的代码在 nginx-gray-web
- lua 动态修改 upstream 的接口
    - 添加一个节点
        - http://localhost:6000/upstream_add?upstream=localhost&server=127.0.0.1:8082
    - 移除一个节点
        - http://localhost:9090/upstream_remove?server=127.0.0.1:8082
    - 列出当前服务名对应有几个节点
        - http://localhost:6000/upstream_list?upstream=localhost
```
接口的 upstream 参数就是对应哪个服务， server 参数就是哪一个节点
```

- 方案需要的工具
    - 1 openresty 或者 nginx
    - 2 luajit 环境
    - 3 nginx 添加依赖的模块
    
    ```
    ./configure --add-module=D:\dev\app\nginx\module\ngx_devel_kit-0.3.1 --add-module=D:\dev\app\nginx\module\lua-nginx-module-0.10.16rc5
    ```
- 运行步骤
    - openresty 配置步骤
        - 把 conf 文件夹的配置文件放到 openresty 的 conf/vhost 目录下
        - 把 lue-file 文件夹下的 lua 脚本放到 openresty 的 lualib 目录下
        - 启动3个web服务，对应到 8081,8082,8083 接口当作负载均衡的3个节点
        - 为这3个web服务做一个动态负载均衡的配置，在上面提到的配置文件里，访问这个 服务的一个接口
    - beego 项目运行步骤
        - 1 项目使用 beego web框架和go mod依赖管理
        - 2 新建一个数据库，名字跟配置文件的一样，把 data.sql 脚本执行
        - 3 进入项目根目录，执行 go run main.go 会自动下载相关依赖并启动
        - 4 浏览器打开 https://localhost:8090/node/page 检验项目是否能正常运行
