package main

// Разработайте пакет для чтения конфигурации типичного веб-приложения через флаги или переменные окружения.
// Пример конфигурации можно посмотреть здесь. По желанию вы можете задать другие имена полям, сгруппировать их или добавить собственные поля.
// Помимо чтения конфигурации приложение также должно валидировать её - например, все URLы должны соответствовать ожидаемым форматам.
// Работу с конфигурацией необходимо вынести в отдельный пакет (не в пакет main).

import (
	"fmt"
	"lesson9/config"
)

func main() {
	jsonConfig := config.JsonConfig{}
	fmt.Println("Json parsing...")
	if err := jsonConfig.Read(); err != nil {
		fmt.Printf("Unable to read the config: %v\n", err.Error())
	} else {
		jsonConfig.Print()
	}

	yamlConfig := config.YamlConfig{}
	fmt.Println("YAML parsing...")
	if err := yamlConfig.Read(); err != nil {
		fmt.Printf("Unable to read the config: %v\n", err.Error())
	} else {
		yamlConfig.Print()
	}

	flagGonfig := config.FlagConfig{}
	fmt.Println("Flags parsing...")
	if err := flagGonfig.Read(); err != nil {
		fmt.Printf("Unable to read the config: %v\n", err.Error())
	} else {
		flagGonfig.Print()
	}

	envConfig := config.EnvConfig{}
	fmt.Println("ENV parsing...")
	if err := envConfig.Read(); err != nil {
		fmt.Printf("Unable to read the config: %v\n", err.Error())
	} else {
		envConfig.Print()
	}

}
