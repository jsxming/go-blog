package common

import (
	"blog/global"
	"blog/pkg/errorcode"
	"blog/pkg/response"
	"blog/pkg/upload"
	"blog/pkg/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func UploadFile(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := util.Int(c.PostForm("type"))
	fmt.Println(file, "File")
	if err != nil {
		response.Error(c, errorcode.InvalidParam.WithDetails(err.Error()))
		return
	}

	if fileHeader == nil || fileType <= 0 {
		response.Error(c, errorcode.InvalidParam)
		return
	}

	fileInfo, err := CheckFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("common.UploadFile err:%v", err)
		response.Error(c, errorcode.ErrorUploadFileFall.WithDetails(err.Error()))
		return
	}

	response.Success(c, gin.H{
		"url": fileInfo.AccessUrl,
	})

}

func CheckFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}

	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit .")
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
