package untils
//untils存放我们的一些公共方法
//判断字符串是否在列表中
func InList (str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}