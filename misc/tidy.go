package misc

import (
	"fmt"
	"net/url"
	"strings"
)

func TidyName(name string) string {
	// replace name space to '-'
	return strings.TrimSpace(name)
}

func TidyUrl(target string) string {
	u, _ := url.Parse(target)
	hostname := u.Hostname()
	if "www.youtube.com" == hostname {
		return tidyYtb(target)
	}

	if "www.bilibili.com" == hostname {
		return tidyBl(target)
	}
	return target
}

func tidyBl(target string) string {
	res := strings.Split(target, "?")
	return res[0]
}

func tidyYtb(target string) string {
	u, _ := url.Parse(target)
	v := u.Query().Get("v")
	return fmt.Sprintf("https://www.youtube.com/watch?v=%s", v)
}
