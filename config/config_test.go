package config_test

import (
	"github.com/annakallo/travel-log-server/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfig(t *testing.T) {
	c := config.NewConfig()
	assert.NotEqual(t, "", c.MysqlIP)
}
