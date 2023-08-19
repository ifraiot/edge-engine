package commander

import "encoding/json"

type Command struct {
	Command    string      `json:"command"`
	Sender     string      `json:"sender"`
	Parameters interface{} `json:"parameters"`
}

func (c Command) Bind(params interface{}) error {

	data, err := json.Marshal(c.Parameters)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, params)
}
