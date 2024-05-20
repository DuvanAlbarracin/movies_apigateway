package config

import (
	"os"

	"github.com/DuvanAlbarracin/movies_apigateway/pkg/utils"
)

type Config struct {
	Port       string
	AuthUrl    string
	MovieUrl   string
	ProfileUrl string
	CastUrl    string
}

func LoadConfig() (c Config, err error) {
	apiPortData, err := os.ReadFile("/run/secrets/api_port")
	if err != nil {
		return
	}

	// authUrlData, err := os.ReadFile("/run/secrets/auth_url")
	// if err != nil {
	// 	return
	// }

	profileUrlData, err := os.ReadFile("/run/secrets/profiles_url")
	if err != nil {
		return
	}

	c.Port = utils.TrimString(string(apiPortData))
	// c.AuthUrl = utils.TrimString(string(authUrlData))
	c.ProfileUrl = utils.TrimString(string(profileUrlData))

	return
}
