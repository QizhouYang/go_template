package database

import (
	"fmt"
	"go_template/pkg/util/encrypt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

const phaseName = "db"

type InitDB struct {
	Host         string
	Port         int
	Name         string
	User         string
	Password     string
	MaxOpenConns int
	MaxIdleConns int
}

func CreateInitDB() *InitDB {

	return &InitDB{
		Host:         viper.GetString("db.host"),
		Port:         viper.GetInt("db.port"),
		Name:         viper.GetString("db.name"),
		User:         viper.GetString("db.user"),
		Password:     viper.GetString("db.password"),
		MaxOpenConns: viper.GetInt("db.max_open_conns"),
		MaxIdleConns: viper.GetInt("db.max_idle_conns"),
	}
}
func (db *InitDB) Exec(q string) error {
	return fmt.Errorf("mysql: not implemented <%s>", q)
}

func (i *InitDB) Init() error {
	p, err := encrypt.StringDecrypt(i.Password)

	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Asia%%2FShanghai",
		i.User,
		p,
		i.Host,
		i.Port,
		i.Name)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		return err
	}

	gorm.DefaultTableNameHandler = func(DB *gorm.DB, defaultTableName string) string {
		if !strings.HasPrefix(defaultTableName, "t_") {
			return "t_" + defaultTableName
		}
		return defaultTableName
	}
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(i.MaxOpenConns)
	db.DB().SetMaxIdleConns(i.MaxIdleConns)
	DB = db
	DB.LogMode(false)
	return nil
}

func (i *InitDB) PhaseName() string {
	return phaseName
}
