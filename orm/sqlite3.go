package main
import (
"database/sql"
"log"
_"github.com/mattn/go-sqlite3"

)

func main(){

db,_:=sql.Open("sqlite3","gee.db")

defer func(){

_=db.Close()

}()
_,_=db.Exec("drop table if exists user;")

_,_=db.Exec("create table user(name text);")
result,err:=db.Exec("insert into user (`name`)values(?),(?)","tom","jack")
if err!=nil{
affected,_:=result.RowsAffected()
log.Println(affected)

} 

row:=db.QueryRow("select name from user limit 1;")
var name string

if err:=row.Scan(&name);err==nil{
log.Println(name)
}




}
