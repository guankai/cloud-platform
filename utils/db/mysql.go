package db

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"net"
	"time"
)

type mysql struct {
	host     string
	port     string
	user     string
	pwd      string
	database string
}

func NewMySQL(host, port, user, pwd, database string) Database {
	return &mysql{
		host:     host,
		port:     port,
		user:     user,
		pwd:      pwd,
		database: database,
	}
}

func (m *mysql) Register(alias ...string) error {
	if err := m.testConn(m.host, m.port); err != nil {
		return err
	}

	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		return err
	}

	an := "default"
	if len(alias) != 0 {
		an = alias[0]
	}

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", m.user, m.pwd, m.host, m.port, m.database)
	logs.Debug("conn:%s", conn)
	return orm.RegisterDataBase(an, "mysql", conn)
}

func (m *mysql) testConn(host, port string) error {
	ch := make(chan int, 1)
	go func() {
		var err error
		var c net.Conn
		for {
			c, err = net.DialTimeout("tcp", host + ":" + port, 20 * time.Second)
			if err == nil {
				c.Close()
				ch <- 1
			} else {
				logs.Error("failed to connect to db, retry after 2 second: %v", err)
				time.Sleep(2 * time.Second)
			}
		}
	}()

	select {
	case <-ch:
		return nil
	case <-time.After(60 * time.Second):
		return errors.New("failed to connect to database after 60 seconds")
	}
}

// Implement Database's Name method
func (m *mysql) Name() string {
	return "MySQL"
}

// Implement Database's String method
func (m *mysql) String() string {
	return fmt.Sprintf("type-%s host-%s port-%s user-%s database-%s", m.Name(), m.host, m.port, m.user, m.database)
}
