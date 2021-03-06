## ipsearch
> ip serach 查询是基于 http://ip.taobao.com/ipSearch.html 接口代理获取的数据（有请求频率限制）
>
> 以前使用ip.cn会出现频次限制，可以基于命令行走淘宝接口查询

### 安装
```
// 安装ipsearch命令工具，以及httpd服务
go get -u -v github.com/lupguo/ipsearch
```

### ipsearch 使用
```
// 命令行使用
Usage of ipsearch:
  -debug
    	debug for request response content
  -format string
    	response message format, default is json (json|text) (default "text")
  -httpd
    	the http server for ip search
  -ip string
    	the IP to be search, the default is the IP of the machine currently executing the cmdline
  -listen string
    	the listen address for ip search http server, eg 127.0.0.1:6100 (default "127.0.0.1:6100")
  -proxy string
    	http proxy using for debugging, no proxy by default, eg http://127.0.0.1:8888
  -timeout duration
    	set http request timeout seconds (default 10s)
  -version
    	ipsearch version

// 客户端查询
$ ipsearch
Ip: 210.21.233.226, Network: 联通, Address: 中国 广东 深圳

// http服务
$ ./ipsearch -httpd -listen '0.0.0.0:6100'
2020/04/12 01:08:03 ipshttpd listen on http://0.0.0.0:6100, ipshttd version 0.4.0'

// 请求查询
$ curl localhost:6100
Version 0.4.0
Usage:
	//search current client ip information
	curl localhost:6100/ips

	//search for target ip information
	curl localhost:6100/ips?ip=targetIp

// 通过curl查询
$ curl localhost:6100/ips
{"addr":"中国 广东 深圳","network":"鹏博士","ip":"175.191.11.165"}
$ curl 'localhost:6100/ips?ip=175.190.11.16'
{"addr":"中国 辽宁 大连","network":"鹏博士","ip":"175.190.11.16"}
```

### Centos Systemd安装

编辑保存 /etc/systemd/system/ipshttpd.service 文件，启动 `systemctl start ipshttpd.service`

```bash
[Unit]
Description=Ipsearch used for searching ip infomation
After=syslog.target network.target remote-fs.target nss-lookup.target

[Service]
Type=simple
User=www
PIDFile=/var/run/ipshttpd.pid
ExecStart=/data/go/bin/ipsearch -httpd -listen 127.0.0.1:6100
TimeoutStartSec=3s
ExecStop=/bin/kill -s QUIT $MAINPID
PrivateTmp=true

[Install]
WantedBy=multi-user.target
```

### 添加Nginx 代理
```shell script
server {
  listen 80;
  listen 443 ssl http2;
  server_name ioio.cool;
  root /data/www/ioio.cool/;

  #ssl
  include ssl_cert.conf;

  # ioio.cool
  location / {
      proxy_pass  http://127.0.0.1:6100;
      proxy_set_header X-Real-IP $remote_addr;
  }
}
```

### 执行请求
```shell script
$ curl ioio.cool/ips
{"addr":"中国 广东 深圳","network":"鹏博士","ip":"14.103.83.43"}
```

### 变更内容
- 2020-04-12：更新了目录结构，调整了重试机制
- 2019-08-07：代码目前版本还比较粗糙，会持续完善！
    - [ ] 代理问题，寻求更好用的代理
    - [x] 程序中一些已知的Bug修复 
- 2019-08-08
    - [x] 修复了客户端请求ipshttpd没有获取到正确IP的问题
    - [x] 修复了ipshttpd的handler处理
    - [x] 新增了版本展示
- 2019-08-13
    - [x] 新增Docker环境支持
