package _markdown

import (
	"strings"
)

// ExtractYamlMarkdown 一个简单的 Yaml Markdown 混写的提取器
func ExtractYamlMarkdown(content string) (yaml, markdown string) {
	items := strings.Split(content, "---")
	l := len(items)
	switch l {
	case 0:
		return
	case 1:
		markdown = items[0]
		return
	default:
		yaml = items[1]
		markdown = strings.Join(items[2:], "")
		return
	}
}
