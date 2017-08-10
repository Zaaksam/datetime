package datetime

import (
	"encoding/json"
	"testing"
)

func TestDatetime(t *testing.T) {
	type tempStruct struct {
		Datetime *Datetime
	}

	temp := &tempStruct{
		Datetime: New(),
	}

	strBytes := []byte(`{"Datetime":"2000-01-01 00:00:00"}`)

	err := json.Unmarshal(strBytes, temp)
	if err != nil {
		t.Fatal(err)
	}

	if temp.Datetime.String() != "2000-01-01 00:00:00" {
		t.Fatal("Expect 2000-01-01 00:00:00, but got: ", temp.Datetime.String())
	}

	strBytes, err = json.Marshal(temp)
	if err != nil {
		t.Fatal(err)
	}

	if string(strBytes) != `{"Datetime":"2000-01-01 00:00:00"}` {
		t.Fatal(`Expect {"Datetime":"2000-01-01 00:00:00"}, but got: `, string(strBytes))
	}
}
