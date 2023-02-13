package types

import (
	"encoding/json"
	"testing"
)

func TestJsonMarshal(t *testing.T) {
	t.Run("marshal json struct", func(t *testing.T) {
		db := database{
			Name:       "db_name_test",
			Collection: "collection_test",
		}
		cfg := configs{
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
	})

	t.Run("load configs", func(t *testing.T) {
		cfg, err := loadConfigs()
		if err != nil {
			panic(err)
		}

		b, err := json.Marshal(&cfg)
		if err != nil {
			panic(err)
		}
		t.Log(string(b))

		b, err = json.MarshalIndent(&cfg, "!", "  ")
		if err != nil {
			panic(err)
		}
		t.Log(string(b))
	})
}
