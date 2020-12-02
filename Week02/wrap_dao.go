package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)
var ErrNoRows=errors.New("ErrNoRows")
func main() {
	if err:=biz();err!=nil{
		fmt.Printf("err:%+v",err)
	}
}
func biz() error{
	return dao()
}
func dao() error{
	sqlStr:="select name from users"
	err:=db.Select(sqlStr)//伪代码，表示执行查询
	if err==sql.ErrNoRows{
		//如果不用mysql换了MongoDB，也可以返回自定义的ErrNoRows
		return ErrNoRows
	}else{
		errors.Wrap(err,sqlStr)
	}
}