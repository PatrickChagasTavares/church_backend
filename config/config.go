package config

import (
	"strings"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Watch fica escutando modificações no config remoto
func Watch(fn func(c Config)) {
	kill := make(chan bool)
	vmain := viper.New()

	c := &configImpl{
		vmain: vmain,
		kill:  kill,
	}

	//Substitui o _ por .
	vmain.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// realiza o bind das variaveis de ambiente
	vmain.AutomaticEnv()

	// seta o arquivo de configuração local
	vmain.SetConfigFile("./config.json")

	// realiza a leitura das configurações locais
	if err := vmain.ReadInConfig(); err != nil {
		logger.Fatal("não consegui ler o arquivo de configuração local")
	}

	// inicia o server
	go fn(c)

	<-kill
}
