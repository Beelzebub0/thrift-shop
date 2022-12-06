package database

import (
	sqlx "database/sql"

	c "github.com/Beelzebub0/thrift-shop/src/conf"
)

type Interface interface {
	Connect() (*sqlx.DB, error)
}

type sql struct {
	conf c.SQLConfig
}

func Init(conf c.SQLConfig) Interface {
	s := &sql{
		conf: conf,
	}

	return s
}

func (s *sql) ping() {
	driver := s.conf.Driver
	host := s.conf.Host
	port := s.conf.Port
	user := s.conf.User
	pass := s.conf.Password
	name := s.conf.Database
	protocol := s.conf.Protocol
	connConfStr := user + ":" + pass + "@" + protocol + "(" + host + ":" + port + ")/" + name + "?parseTime=true"

	db, err := sqlx.Open(driver, connConfStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.Close()
}

func (s *sql) Connect() (*sqlx.DB, error) {
	driver := s.conf.Driver
	host := s.conf.Host
	port := s.conf.Port
	user := s.conf.User
	pass := s.conf.Password
	name := s.conf.Database
	protocol := s.conf.Protocol
	connConfStr := user + ":" + pass + "@" + protocol + "(" + host + ":" + port + ")/" + name + "?parseTime=true"

	db, err := sqlx.Open(driver, connConfStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
