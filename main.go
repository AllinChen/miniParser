package main

import (
	"fmt"
	"os"

	"github.com/AllinChen/miniParser/common"
	"github.com/AllinChen/miniParser/miniparser"
)

func main() {
	f := common.MyFlag{}
	f.Init()
	p := miniparser.NewParser()
	f.SQL = `select id from view1 where id >ANY(select id from view1)`
	// f.SQL = "select count(*) from tablex a,tablex b where a.colname = b.precolname;"
	// f.SQL = `select t02.*, id id_1 from db01.t01 inner join db02.t02 on t01.id=t02.id;`
	// f.SQL = `alter table t01 change column col1 col2 int(10) comment 'ddd';`
	// f.SQL = `alter table t01 comment 'ddd';`
	// f.SQL = `insert into t01(col1, col2, col3) values(1, 1, 1, 1);`
	// f.SQL = `select projectfile.id,projectfile2.filename from projectfile cross join projectfile2 where projectfile.id =1 and projectfile2.id =2;`

	sql := f.SQL
	result, warns, err := p.Parse(sql)
	if err != nil {
		fmt.Printf("parse error: %s", err.Error())
		os.Exit(1)
	}
	if warns != nil {
		for _, warn := range warns {
			fmt.Printf("parse warn: %s", warn.Error())
		}
		os.Exit(1)
	}
	fmt.Println(result.TableNames, result.ColumnNames)
	fmt.Println(miniparser.SQLInfo.SelectTabs)
	fmt.Println(miniparser.SQLInfo.JoinTables)
}

// 	jsonBytes, err := result.Marshal()
// 	if err != nil {
// 		fmt.Printf("marshal error: \n%s", err.Error())
// 		os.Exit(1)
// 	}

// 	fmt.Println(string(jsonBytes))
// }
