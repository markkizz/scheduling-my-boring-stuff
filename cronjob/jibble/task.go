package jibble

import (
	"fmt"

	globalConfig "github.com/markkizz/time-tracker-automation/config"
	services "github.com/markkizz/time-tracker-automation/services/jibble"
)

func Task() {

	jibbleService := services.JibbleService{}
	jibbleConfig := globalConfig.Config().Jibble
	fmt.Println("------- Loging in -------")
	jibbleService.Login(jibbleConfig.Username, jibbleConfig.Password)
	jibbleService.Clock("In")
}
