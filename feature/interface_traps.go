package feature

import (
	"encoding/json"
	"fmt"
	"time"
)

type Check struct {
	time.Time
}

type Inheriet struct {
	Check `json:"-"`
	Filed string `json:"field"`
}

func marshalmethed() {

	a := Inheriet{
		Check{time.Now()},
		"good",
	}

	marshal, err := json.Marshal(&a)
	if err != nil {
		return
	}

	fmt.Println(string(marshal))

}
