package errorcode

var (
	//ErrorGetTagListFail = NewError(2001, "获取标签列表失败")
	//ErrorCreateTagFail  = NewError(2002, "创建标签失败")
	//ErrorUpdateTagFail  = NewError(2003, "更新标签失败")
	//ErrorDeleteTagFail  = NewError(2004, "删除标签失败")
	NotFoundUser    = NewError(2001, "账号不存在")
	InvalidPassword = NewError(2002, "密码不正确")
	//NotFoundUser   = NewError(2002, "Token")

)
