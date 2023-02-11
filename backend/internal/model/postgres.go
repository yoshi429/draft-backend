package model

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/go-pg/pg"
	"github.com/yoshihiro-shu/draft-backend/internal/config"
)

type DBContext struct {
	master   *pg.DB
	repricas []*pg.DB
}

func (c DBContext) Master() *pg.DB {
	return c.master
}

func (c DBContext) Reprica() *pg.DB {
	numOfDB := big.NewInt(int64(len(c.repricas)))
	n, err := rand.Int(rand.Reader, numOfDB)
	if err != nil {
		return c.repricas[0]
	}
	return c.repricas[n.Int64()]
}

func connectToMaster(conf config.DB) *pg.DB {
	return getDBConnection(conf)
}

func connectToRepricas(conf []config.DB) []*pg.DB {
	dbs := make([]*pg.DB, len(conf))
	for i, v := range conf {
		dbs[i] = getDBConnection(v)
	}
	return dbs
}

func getDBConnection(c config.DB) *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", c.Host, c.Port),
		User:     c.User,
		Password: c.Password,
		Database: c.Name,
	})
}

func New(conf config.Configs) *DBContext {
	return &DBContext{
		master:   connectToMaster(conf.MasterDB()),
		repricas: connectToRepricas(conf.RepricaDB()),
	}
}
