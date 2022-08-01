package main

// You can edit this code!
// Click here and start typing.

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type Customer struct {
	Loyalty interface{} `json:"loyalty"`
}

func main() {

	// Receber 2 campo um int e uma string, descobrir e utilizar assert. Compo de interface Utilizar marshal

	var c Customer
	var str string = `{"loyalty":"1"}`
	err := json.Unmarshal([]byte(str), &c)

	fmt.Println(err)
	fmt.Println(c)
	fmt.Println(c.Loyalty)

	x := c.Loyalty

	// fmt.Printf("\nx:%s", x)

	/* if x.(type) == "String" {

		c.Loyalty, _ = strconv.Atoi(x.(string))
		fmt.Println(c.Loyalty, err)
	} */

	switch x.(type) {
	case int:
		c.Loyalty = x.(int)
	case string:
		c.Loyalty, _ = strconv.Atoi(x.(string))
	}
	fmt.Println(c.Loyalty)
	fmt.Println(reflect.TypeOf(c.Loyalty))
}
