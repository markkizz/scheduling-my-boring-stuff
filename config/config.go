package config

import (
	"github.com/markkizz/time-tracker-automation/config/jibble"
)

type Configuration struct {
	Jibble jibble.JibbleConfig
}

func Config() Configuration {
	jibbleConfig := jibble.LoadJibbleConfig()

	config := Configuration{Jibble: jibbleConfig}
	return config
}
