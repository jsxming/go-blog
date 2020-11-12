package response

import (
	"blog/pkg/errorcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
		"msg":  "操作成功",
	})
}

func Error(c *gin.Context, err *errorcode.Error) {
	response := gin.H{
		"code":      400,
		"errorCode": err.Code(),
		"msg":       err.Msg(),
	}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	c.JSON(err.StatusCode(), response)
}

func ToTable(c *gin.Context, data interface{}, total int) {
	c.JSON(http.StatusOK, gin.H{
		"data":  data,
		"total": total,
	})
}
