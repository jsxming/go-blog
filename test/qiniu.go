package main

import (
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

const accessKey = "b6qHuI07mVNVp1xe_Qwhd_rV90Kie9YebgWaVQAM"
const secretKey = "jka8-Gdx7GQohKqqcyZ3h9xPtEjtpy4FgtvYuPYA"

func main() {
	bucket := "your bucket name"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	fmt.Println(upToken)
}
