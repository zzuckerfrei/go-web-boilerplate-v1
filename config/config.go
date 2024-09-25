package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetConfig(env string) error {
	viper.SetConfigName("config-" + env) // config 파일 이름 (확장자 제외)
	viper.SetConfigType("yaml")          // 설정 파일 타입 지정
	viper.AddConfigPath(".")             // 현재 디렉토리에서 env.yaml 파일 찾기

	// 설정 파일 읽기
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 설정 파일을 찾을 수 없음
			return fmt.Errorf("Config file not found: %v", err)
		} else {
			// 다른 에러 발생
			return fmt.Errorf("Error reading config file: %v", err)
		}
	}

	return nil
}
