/*
@File    : code.go
@Author  : GuguLH
@Date    : 2026/3/26 11:43
@Desc    : 定义系统的错误码
*/

package errs

// User 相关
const (
	// UserInvalidInput 统一的用户模块的输入错误
	UserInvalidInput = 401001
	// UserInvalidOrPassword 用户名错误或者密码不对
	UserInvalidOrPassword = 401002
	// UserDuplicateEmail 用户邮箱冲突
	UserDuplicateEmail = 401003
	// UserDuplicatePhone 用户手机号冲突
	UserDuplicatePhone = 401004
	// UserInternalServerError 统一的用户模块的系统错误
	UserInternalServerError = 501001
)
