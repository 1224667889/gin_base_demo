package util

import "fzuhelper_launch_screen/pkg/setting"

// Setup 初始化工具类
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}