package mysqltest

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //导入并且没有使用，所以默认调用init方法，在底层会注册一个叫mysql的驱动
)

var db *sql.DB
func init() {
	//fmt.Println("测试mysql使用")
}

//数据库三个字段 主键id,name,age
func initDB() (err error) {
	//连接数据库 账号root 密码123123 数据库名 godb
	dsn := "root:root@unix(/Applications/MAMP/tmp/mysql/mysql.sock)/test"
	//Open不会校验用户名密码是否正确，只能校验数据格式对不对
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Println("ping 不通")
		return
	}
	//设置数据库连接池的最大连接数
	db.SetMaxIdleConns(10)

	return
}

type user struct {
	id   int
	name string
	age  int
}

//删
func delete(id int) {
	sqlStr := `delete from user where id=?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("删除失败", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("获取删除记录错误", err)
		return
	}
	fmt.Println("更新了", n, "条数据")

}

//查询单条
func queryOne(id int) {
	var u1 user
	//sql语句编写
	sqlStr := `select id,name,age from stu where2 id=?`
	rowObj := db.QueryRow(sqlStr, id) //从连接池那里连接查询出一个单挑记录
	//拿到结果，
	err := rowObj.Scan(&u1.id, &u1.name, &u1.age) //必须对rowObj对象调用Scan方法，因为该方法会释放数据库连接

	fmt.Printf("u1:%#v\n err:%#v\n", u1, err)
}
func insert() {
	//sql
	sqlStr := `insert into user(name,age) values("名字",20)`
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println("执行sql错误", err)
		return
	}
	//如果是插入数据的操作,能够拿到插入数据的id值
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Println("获取id错误", err)
		return
	}
	fmt.Println("id:", id)
}

//查询多条
func queryMore(n int) (res string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
		res = "ttttt"
	}()

	//sql
	sqlStr := `select id,name,age from stu where id > ? for update;`
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		panic("语法错误" + err.Error())
		fmt.Println("查询错误", err)
		return "bb"

	}
	fmt.Println("继续", rows)

	//关闭rows
	//defer rows.Close()
	//循环取值

	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Println("解析错误", err)
		}
		fmt.Println(u1) //取一条打印一条
	}
	return res
}

//改
func updateRow(newAge int, id int) {
	sqlStr := `update user set age=? where id=?`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Println("更新错误", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("a", err)
		return
	}

	fmt.Println("更新了", n, "行数据")
}

func Main() {
	err := initDB()
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}
	fmt.Println("连接数据库成功")
	//queryOne(2)
	str := queryMore(0)
	fmt.Println("来：", str)
	// insert()
	// updateRow(99, 2)
	//delete(5)
}