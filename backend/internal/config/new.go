package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	StatusUnpublished = iota + 1
	StatusPublished
	StatusClosed
)

type Configs struct {
	User         Server       `yaml:"user"`
	RelationalDB RelationalDB `yaml:"relationalDB"`
	CacheRedis   RedisCache   `yaml:"cacheRedis"`
	Twitter      Twitter      `yaml:"twitter"`
	AccessToken  AuthToken    `yaml:"access_token"`
	RefreshToken AuthToken    `yaml:"refresh_token"`
}

func (c Configs) GetUserAddr() string {
	return c.User.GetAddr()
}

func (c Configs) MasterDB() DB {
	return c.RelationalDB.Master
}

func (c Configs) RepricaDB() []DB {
	return c.RelationalDB.Repricas
}

func (c Configs) GetCacheRedis() RedisCache {
	return c.CacheRedis
}

func New() Configs {
	var conf Configs

	b, err := os.ReadFile("./configs.yaml")
	if err != nil {
		log.Fatalf("failed read configs.yaml. err :%s", err.Error())
	}
	err = yaml.Unmarshal(b, &conf)
	if err != nil {
		log.Fatalf("failed Unmarshal configs.yaml. err :%s", err.Error())
	}

	return conf
}
