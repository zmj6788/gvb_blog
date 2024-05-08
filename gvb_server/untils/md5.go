package untils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5加密函数
func MD5(bytedata []byte) string {
	m := md5.New()
	m.Write(bytedata)
	//将计算得到的哈希值转换为16进制字符串
	str := hex.EncodeToString(m.Sum(nil))
	return str
}
