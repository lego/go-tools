package pkg

import "encoding/json"

func fn1(i3 interface{}) {
	var v map[string]interface{}
	var i1 interface{} = v
	var i2 interface{} = &v
	p := &v
	_ = json.Unmarshal([]byte(`{}`), v) // MATCH /Unmarshal expects to unmarshal into a pointer/
	_ = json.Unmarshal([]byte(`{}`), &v)
	_ = json.Unmarshal([]byte(`{}`), i1) // MATCH /Unmarshal expects to unmarshal into a pointer/
	_ = json.Unmarshal([]byte(`{}`), i2)
	_ = json.Unmarshal([]byte(`{}`), i3)
	_ = json.Unmarshal([]byte(`{}`), p)

	_ = json.NewDecoder(nil).Decode(v) // MATCH /Decode expects to unmarshal into a pointer/
}
