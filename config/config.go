package config

import (
	"flag"
	"time"
)

type Config struct {
	Ip, Proxy, Format, Listen string
	Debug, Version            bool
	Timeout                   time.Duration
	Httpd                     bool
}

var cfg Config

func init() {
	flag.StringVar(&cfg.Ip, "ip", "", "the IP to be search, the default is the IP of the machine currently executing the cmdline")
	flag.StringVar(&cfg.Proxy, "proxy", "", "http proxy using for debugging, no proxy by default, eg http://127.0.0.1:8888")
	flag.StringVar(&cfg.Format, "format", "text", "response message format, default is json (json|text)")
	flag.BoolVar(&cfg.Debug, "debug", false, "debug for request response content ")
	flag.DurationVar(&cfg.Timeout, "timeout", 10*time.Second, "set http request timeout seconds")
	flag.StringVar(&cfg.Listen, "listen", "127.0.0.1:6100", "the listen address for ip search http server, eg 127.0.0.1:6100")
	flag.BoolVar(&cfg.Httpd, "httpd", false, "the http server for ip search")
	flag.BoolVar(&cfg.Version, "version", false, "ipsearch version")
}

//Get 获取Config配置
func Get() *Config {
	return &cfg
}
