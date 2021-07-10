package Util

import (
	"fmt"
	"go/build"
	"time"
)

func FilePath(fileName string) string {
	//这边存在在go目录下，用时间标签区分文件命名，（如果同时上传同名会不会出现覆盖的情况）
	return fmt.Sprintf("%s/%d_%s", build.Default.GOPATH, time.Now().UnixNano(), fileName)
}