package global

import (
	"LiuYanXiBlog/pkg/logger"
	"LiuYanXiBlog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	APPSetting      *setting.APPSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
