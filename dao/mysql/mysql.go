package mysql

import (
	"fmt"

	"github.com/miaogu-go/bluebell/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(dbConf *settings.DbConf) {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?loc=Local", dbConf.User, dbConf.Password, dbConf.Host,
		dbConf.Port, dbConf.DbName)
	//dsn := "root:123456@tcp(127.0.0.1:3306)/blog?parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(dbConf.MaxOpenConn)
	db.SetMaxIdleConns(dbConf.MaxIdleConn)

	return
}

func Close() {
	_ = db.Close()
}
