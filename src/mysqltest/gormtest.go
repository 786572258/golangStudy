package mysqltest

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Stu struct {
	Id int
	Name string
	Age int
}
func (Stu) TableName() string {
	return "stu"
}

func query() (u []Stu) {
	//查询所有记录

	//err := gormdb.Where("1").Raw("for update").Take(&u).Error
	err := gormdb.Exec("select * from2 stu where 1 for update").Take(&u).Error
	if err != nil {
		panic(err.Error())
	}
	return u
}

var gormdb *gorm.DB
func init() {
	var err error
	gormdb, err = gorm.Open("mysql", "root:root@unix(/Applications/MAMP/tmp/mysql/mysql.sock)/test?charset=utf8&parseTime=True&loc=Local")
	gormdb.DB().SetMaxIdleConns(10)
	gormdb.DB().SetMaxOpenConns(100)
	gormdb.LogMode(true)


	if err != nil {
		panic("gorm.open 错误：" + err.Error())
	}
}

func Gormtest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
	}()
	res := query()
	fmt.Println("来：", res)
	defer gormdb.Close()

}