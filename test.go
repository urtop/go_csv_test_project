package main

import (
	"github.com/maxim2266/csvplus"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Data struct{
	Flag string
	System string
	Note string
}
var engine *xorm.Engine

func main() {
	raws := csvplus.CsvFileDataSource("test.csv").SelectColumns("system", "flag", "note")
	data, _ := csvplus.Take(raws).ToRows()
	//fmt.Println(data[1]["flag"])
	//json.Unmarshal(data[1],Data{})
	engine, _ = xorm.NewEngine("mysql","root:@/csv?charset=utf8")
	sql := "select * from data"
	results, err := engine.Query(sql)
	println(results)
	fmt.Println((err))
	for key,value:= range data {
		fmt.Println(key, value)
		var data  Data
		data.Flag = value["flag"]
		data.System = value["system"]
		data.Note = value["note"]
		affected,err := engine.Insert(&data)
		fmt.Println(affected)
		if(err !=nil) {
			fmt.Println(err)
		}
	}

}
