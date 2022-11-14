package hash

import (
	"crypto/md5"
	"encoding/hex"
)

/**
* @Title StringMd5
* @Description:   字符串md5加密
* @Param:
* @return:
* @Author: liwei
* @Date: 2020/4/24
**/
func StringMd5(s string) string {
	md5S := md5.New()
	md5S.Write([]byte(s))
	return hex.EncodeToString(md5S.Sum(nil))
}
