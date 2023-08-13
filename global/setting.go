package global

import (
	"blog-sevice/pkg/logger"
	"blog-sevice/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings

	Logger *logger.Logger
)
