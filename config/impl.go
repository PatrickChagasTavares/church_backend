package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config interface de configuração
type Config interface {
	GetInt(key string) int
	GetBool(key string) bool
	GetString(key string) string
	GetFloat64(key string) float64
	GetDuration(key string) time.Duration
	GetStringSlice(key string) []string
	Close()
}

type configImpl struct {
	vmain *viper.Viper

	kill chan bool
}

func (c *configImpl) GetBool(key string) bool {
	return c.vmain.GetBool(key)
}

func (c *configImpl) GetString(key string) string {
	return c.vmain.GetString(key)
}

func (c *configImpl) GetDuration(key string) time.Duration {
	return c.vmain.GetDuration(key)
}

func (c *configImpl) GetInt(key string) int {
	return c.vmain.GetInt(key)
}

func (c *configImpl) GetFloat64(key string) float64 {
	return c.vmain.GetFloat64(key)
}

func (c *configImpl) GetStringSlice(key string) []string {
	return c.vmain.GetStringSlice(key)
}

// Close encerra a função de watch
func (c *configImpl) Close() {
	c.kill <- true
}
