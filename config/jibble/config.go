package jibble

import "os"

type JibbleConfig struct {
	Username, Password, ApiIdentityUrl, ApiTimeTrackingUrl string
}

const (
	USERNAME             string = "JIBBLE_USERNAME"
	PASSWORD             string = "JIBBLE_PASSWORD"
	API_IDENTITY_URL     string = "JIBBLE_API_IDENTITY_URL"
	API_TIME_TRACKER_URL string = "JIBBLE_API_TIME_TRACKER_URL"
)

func LoadJibbleConfig() JibbleConfig {
	jibbleConfig := JibbleConfig{os.Getenv(USERNAME), os.Getenv(PASSWORD), os.Getenv(API_IDENTITY_URL), os.Getenv(API_TIME_TRACKER_URL)}
	return jibbleConfig
}
