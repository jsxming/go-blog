package global

import (
	"blog/pkg"
	"github.com/sirupsen/logrus"
)

var (
	ServerSetting   *pkg.ServerSetting
	AppSetting      *pkg.AppSetting
	DatabaseSetting *pkg.DatabaseSettings
	Logger          *logrus.Logger
	JWTSetting      *pkg.JWTSetting
)
