package config

import (
	"github.com/markkizz/scheduling-my-boring-stuff/config/jibble"
)

type Configuration struct {
	Jibble jibble.JibbleConfig
}

func Config() Configuration {
	jibbleConfig := jibble.LoadJibbleConfig()

	config := Configuration{Jibble: jibbleConfig}
	return config
}
