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
// 反转切片
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}