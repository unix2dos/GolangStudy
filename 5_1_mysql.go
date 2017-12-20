/*
	mysql 驱动 https://github.com/go-sql-driver/mysql


	mysql -h localhost -u root -p
	show databases;
	CREATE DATABASE IF NOT EXISTS goweb DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
	use goweb;

	CREATE TABLE `userinfo` (
		`uid` INT(10) NOT NULL AUTO_INCREMENT,
		`username` VARCHAR(64) NULL DEFAULT NULL,
		`departname` VARCHAR(64) NULL DEFAULT NULL,
		`created` DATE NULL DEFAULT NULL,
		PRIMARY KEY (`uid`)
	);

	CREATE TABLE `userdetail` (
		`uid` INT(10) NOT NULL DEFAULT '0',
		`intro` TEXT NULL,
		`profile` TEXT NULL,
		PRIMARY KEY (`uid`)
	);
*/

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "root:123456@/goweb?charset=utf8")

	//插入数据
	stmt, _ := db.Prepare("Insert userinfo SET username=?, departname=?,created=?")
	res, _ := stmt.Exec("liuwei", "golang", "2017-12-20")
	id, _ := res.LastInsertId()
	fmt.Println(id)
	//修改数据
	stmt, _ = db.Prepare("Update userinfo set username=? where uid=?")
	res, _ = stmt.Exec("xuanyuan", id)
	affect, _ := res.RowsAffected()
	fmt.Println(affect)
	//查询数据
	rows, _ := db.Query("Select * FROM userinfo")
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		rows.Scan(&uid, &username, &department, &created)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	//删除数据
	stmt, _ = db.Prepare("DELETE from userinfo where uid=?")
	res, _ = stmt.Exec(id)
	affect, _ = res.RowsAffected()
	fmt.Println(affect)

	db.Close()
}
