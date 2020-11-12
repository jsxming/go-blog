package errorcode

var (
	Success             = NewError(1, "成功")
	ServerError         = NewError(1000, "服务器内部错误")
	InvalidParam        = NewError(1001, "参数错误")
	NotFound            = NewError(1002, "找不到")
	InvalidToken        = NewError(1003, "无效的Token")
	TooManyRequests     = NewError(1004, "请求过多")
	ErrorUploadFileFall = NewError(1005, "上传文件失败")
	UpdateError         = NewError(1006, "更新失败")
	CreateError         = NewError(1007, "新增失败")
	DeleteError         = NewError(1008, "删除失败")
	ExistError          = NewError(1009, "已存在")
)
