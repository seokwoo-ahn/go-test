package types

import (
	"encoding/json"
	"testing"
)

func TestJsonMarshal(t *testing.T) {
	db := Database{
		Name:       "db_name_test",
		Collection: "collection_test",
	}
	cfg := Configs{
		Port:     443,
		Url:      "seok-study.tistory.com",
		Database: db,
	}

	b, err := json.Marshal(&cfg)
	if err != nil {
		panic(err)
	}
	t.Log(string(b))

	b, err = json.MarshalIndent(&cfg, "*", "    ")
	if err != nil {
		panic(err)
	}
	t.Log(string(b))
}
