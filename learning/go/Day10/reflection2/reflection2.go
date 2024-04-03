package main

import (
	"fmt"
	"reflect"
)

func main() {
	p := person{
		firstName: "Bipin",
		lastName:  "Radhakrishnan",
		email:     "br@mail.com",
	}
	fmt.Println(createInsertQuery(p))

	o := order{
		orderNumber: 123,
		totalOrder:  350.59,
	}

	fmt.Println(createInsertQuery(o))
}

type person struct {
	firstName string
	lastName  string
	email     string
}

type order struct {
	orderNumber int
	totalOrder  float32
}

func createInsertQuery(q interface{}) string {
	typeOf := reflect.TypeOf(q)
	if typeOf.Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		numberOfFields := typeOf.NumField() - 1
		queryBase := fmt.Sprintf("insert into %s", typeOf.Name())
		queryFields := ""
		queryValues := ""
		for i := 0; i <= numberOfFields; i++ {
			if len(queryFields) == 0 {
				queryFields = fmt.Sprintf("%s", typeOf.Field(i).Name)
			} else {
				queryFields = fmt.Sprintf("%s, %s", queryFields, typeOf.Field(i).Name)
			}

			fieldValue := v.Field(i)
			switch fieldValue.Kind() {
			case reflect.Int:
				if len(queryValues) == 0 {
					queryValues = fmt.Sprintf("%d", fieldValue.Int())
				} else {
					queryValues = fmt.Sprintf("%s, %d", queryValues, fieldValue.Int())
				}
				break
			case reflect.Float32:
				if len(queryValues) == 0 {
					queryValues = fmt.Sprintf("%f", fieldValue.Float())
				} else {
					queryValues = fmt.Sprintf("%s, %f", queryValues, fieldValue.Float())
				}
				break
			case reflect.String:
				if len(queryValues) == 0 {
					queryValues = fmt.Sprintf("\"%s\"", fieldValue.String())
				} else {
					queryValues = fmt.Sprintf("%s, \"%s\"", queryValues, fieldValue.String())
				}
				break
			}
		}
		queryBase = fmt.Sprintf("%s (%s) values (%s)", queryBase, queryFields, queryValues)
		return queryBase
	}
	panic("Unsupported type")
}
