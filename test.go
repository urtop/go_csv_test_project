package main

import (
	"github.com/maxim2266/csvplus"
	"fmt"
)

type Data struct{
	flag string
	system string
	note string
}

func main() {
	raws := csvplus.CsvFileDataSource("test.csv").SelectColumns("system", "flag", "note")
	data, _ := csvplus.Take(raws).ToRows()
    for key,value:= range data {
		fmt.Println("Key: %s  Value: %s\n", key, value)
	}
	//err := csvplus.Take(people).
	//	Filter(csvplus.Like(csvplus.Row{"name": "Amelia"})).
	//	Map(func(row csvplus.Row) csvplus.Row { row["name"] = "Julia"; return row }).
	//	ToCsvFile("out.csv", "name", "surname")
   //
   //if err != nil {
	//   fmt.Println(err)
   //}
}
