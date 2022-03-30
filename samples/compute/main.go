package main

import (
	"strconv"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	helpers "gitlab.com/k33g_org/side-projects/galago/galago-executor/api/function"
)

func main() {}

func oh(body string) string {
	firstName := gjson.Get(body, "FirstName")
	lastName := gjson.Get(body, "LastName")

	//var sum int64 = 0
	sum := int64(0)
	for i := int64(0); i < 1000000; i++ {
		sum += i
	}

	result, _ := sjson.Set(`{"message":""}`, "message", "🖐️😃 "+firstName.Str+" "+lastName.Str+" "+strconv.FormatInt(sum, 10))

	return result
}

//export handle
func handle(parameters *int32) *byte {
	return helpers.Use(oh, parameters)
}
