package types

import (
	"entdemo/ent"
	"fmt"
)

// give me a interface of db client
type DbConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	DB   string `json:"db"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

type Db struct {
	Config DbConfig
	Client interface{}
}

// implement the interface of service

func (d *Db) Init() error {
	port := fmt.Sprintf("%v", d.Config.Port)
	client, err := ent.Open("mysql", d.Config.User+":"+d.Config.Pass+"@tcp("+d.Config.Host+":"+port+")/"+d.Config.DB+"?parseTime=True")
	if err != nil {
		return err
	}
	d.Client = client
	return nil
}

func (d *Db) Close() (callback func(), err error) {
	return func() {
		err = d.Client.(*ent.Client).Close()
	}, err
}
