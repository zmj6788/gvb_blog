package res

type ErrorCode int

// 错误码
// 通常是从JSON文件中读取
const (
	SettingsError ErrorCode = 1001 //系统错误
<<<<<<< HEAD
	ArgumentError ErrorCode = 1002 //参数错误
=======
>>>>>>> 2f9e4d1a6a0ab0002a002517dace0301441cd6ca
)

// 错误码对应错误信息
var ErrorMap = map[ErrorCode]string{
	SettingsError: "系统错误",
<<<<<<< HEAD
	ArgumentError: "参数错误",
=======
>>>>>>> 2f9e4d1a6a0ab0002a002517dace0301441cd6ca
}
