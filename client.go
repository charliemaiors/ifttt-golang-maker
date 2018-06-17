package maker

import (
	"encoding/json"
	"fmt"
)

type Client interface {
	Do(event string, values Values)
}

type Values struct {
	FirstValue  string `json:"value1"`
	SecondValue string `json:"value2"`
	ThirdValue  string `json:"value3"`
}

func NewValues(values [3]string) Values {
	return Values{
		FirstValue:  values[0],
		SecondValue: values[1],
		ThirdValue:  values[2],
	}
}

func (val Values) String() string {
	bts, err := json.Marshal(&val)

	if err != nil {
		fmt.Printf("Got error marshaling %v", err)
		return ""
	}
	return string(bts)
}
