package urltool

import (
	"errors"
	"net/url"
	"path"
)

// GetBasePath 获取URL路径的最后一节  aszhc.top/hello -> aszhc.top/a/b/c -> c
func GetBasePath(targetUrl string) (string, error) {
	myUrl, err := url.Parse(targetUrl)
	if err != nil {
		return "", err
	}
	if len(myUrl.Host) == 0 {
		return "", errors.New("no host in targetUrl")
	}
	return path.Base(myUrl.Path), nil
}
