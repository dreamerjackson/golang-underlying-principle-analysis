package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Age      int
	Name     string
}

func (s*Student) CreateSQL() string{
	sql := fmt.Sprintf("insert into student values(%d, %s)", s.Age, s.Name)
	return sql
}

type Trade struct {
	tradeId      int
	Price        int
}

func main() {
	createQuery(Student{Age:  20, Name: "jonson",})
	createQuery(Trade{tradeId: 123, Price:   456,})
}

func createQuery(q interface{}) string{
	// 判断类型为结构体
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		// 获取结构体名字
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return ""
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return query
	}
	return ""
}