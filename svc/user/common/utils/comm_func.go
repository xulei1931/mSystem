package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

const pass_salt = "xulei$xulei"

// 密码加密
func Md5Password(pass string) string {
	w := md5.New()

	io.WriteString(w, pass+pass_salt)    //将str写入到w中
	return fmt.Sprintf("%x", w.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
}
