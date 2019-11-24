package main

import "fmt"
import "encoding/json"

type IT struct {
	Company  string `json:"company"`
	Subjects []string
	Isok     bool `json:",string"`
	Prince   float64
}

func main() {
	s := IT{"itcast", []string{"Go", "c++", "python"}, true, 666.66}
	buf, err := json.MarshalIndent(s, "   ", "    ")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("buf = ", string(buf))
}
