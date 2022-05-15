package easy_json

import (
	"fmt"
	"github.com/mailru/easyjson"
	"testing"
)

var jsonStr = "{\"name\":\"spier\",\"skills\":[\"c\",\"c++\",\"java\",\"go\"]}"

func TestEasyJson(t *testing.T) {
	e := BasicInfo{Name: "spier", Skills: []string{"c", "c++", "java", "go"}}
	rawBytes, _ := easyjson.Marshal(e)
	fmt.Println(string(rawBytes))
	tmp := new(BasicInfo)
	err := easyjson.Unmarshal([]byte(jsonStr), tmp)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(*tmp)
}
