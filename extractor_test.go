package _markdown

import (
	"testing"
)

func TestExtractYamlMarkdown(t *testing.T) {
	t.Log(ExtractYamlMarkdown(`
	---
	title:  "Nginx 配置 MySQK 代理"
	summary: "通过 Nginx 代理 MySQL或者其他 TCP 服务，在网络不同的时候，可以帮助解决很多麻烦 "
	date: "2022-03-24T19:18:07+08:00"
	draft: false
	tags:
	   - "Nginx"
	   - "MySQL"
	   - "代理"
	categories:
	   - "技术文章"
	# description:
	# slug:
	---
	
	## 安装 Nginx 及 Stream 模块
	`))
}
