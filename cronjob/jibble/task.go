package jibble

import (
	"fmt"

	globalConfig "github.com/markkizz/scheduling-my-boring-stuff/config"
	services "github.com/markkizz/scheduling-my-boring-stuff/services/jibble"
)

func Task() {

	jibbleService := services.JibbleService{}
	jibbleConfig := globalConfig.Config().Jibble
	fmt.Println("------- Loging in -------")
	jibbleService.Login(jibbleConfig.Username, jibbleConfig.Password)
	jibbleService.Clock("In")
}
