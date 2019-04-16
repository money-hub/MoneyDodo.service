package db

import (
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	yaml "gopkg.in/yaml.v2"
)

type config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Schema   string `yaml:"schema"`
}

type DBService struct {
	conf   config
	engine *xorm.Engine
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func (dbsvc *DBService) Bind(dbconf string) error {
	confFile, err := ioutil.ReadFile(dbconf)
	checkErr(err)

	// fmt.Println(string(confFile))
	err = yaml.Unmarshal(confFile, &dbsvc.conf)
	checkErr(err)
	dataSourceName := dbsvc.conf.Username + ":" + dbsvc.conf.Password + "@tcp(" + dbsvc.conf.Host + ":" + dbsvc.conf.Port + ")/" + dbsvc.conf.Schema + "?charset=utf8"
	// fmt.Println(dataSourceName)
	dbsvc.engine, err = xorm.NewEngine("mysql", dataSourceName)
	checkErr(err)
	return err
}

func (dbsvc *DBService) Engine() *xorm.Engine {
	return dbsvc.engine
}
