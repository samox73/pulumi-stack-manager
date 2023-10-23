package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./stack.json")
	if err != nil {
		panic(err)
	}

	s := Stack{}
	err = json.Unmarshal(data, &s)
	if err != nil {
		panic(err)
	}

	// keys := make(map[string]struct{})
	// for _, r := range s.Deployment.Resources {
	// 	for k := range r {
	// 		keys[k] = struct{}{}
	// 	}
	// }

	// for k := range keys {
	// 	fmt.Println(k)
	// }

	fmt.Println(s)
}
