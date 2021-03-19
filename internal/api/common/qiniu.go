package common

import (
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

/**
七牛云 token返回
*/

const accessKey = "b6qHuI07mVNVp1xe_Qwhd_rV90Kie9YebgWaVQAM"
const secretKey = "jka8-Gdx7GQohKqqcyZ3h9xPtEjtpy4FgtvYuPYA"
const bucket = "mtutlib"

func CreateQinuToken(c *gin.Context) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 // 2小时有效期
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	response.Success(c, upToken)
}
