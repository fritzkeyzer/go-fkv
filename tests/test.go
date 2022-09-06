package tests

import (
	"encoding/json"
	"fmt"
	"github.com/fritzkeyzer/go-fkv"
)

// Test an implementation of FKV
func Test(kv fkv.FKV) error {
	if err := testSetGet(kv); err != nil {
		return fmt.Errorf("testSetGet: %v", err)
	}

	if err := testListAndDelete(kv); err != nil {
		return fmt.Errorf("testListAndDelete: %v", err)
	}

	return nil
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func testSetGet(kv fkv.FKV) error {
	{
		key := "0/0"
		value := 1337
		var got int

		if err := kv.Set(key, value); err != nil {
			return fmt.Errorf("case[%s]: error: %v", key, err)
		}
		if f, err := kv.Get(key, &got); err != nil {
			return fmt.Errorf("case[%s]: error: %v", key, err)
		} else if !f {
			return fmt.Errorf("case[%s]: not found", key)
		}
		gotBuf, err := json.MarshalIndent(got, "", "\t")
		if err != nil {
			return fmt.Errorf("case[%s]: error marshalling got to json: %v", key, err)
		}
		wantBuf, err := json.MarshalIndent(value, "", "\t")
		if err != nil {
			return fmt.Errorf("case[%s]: error marshalling got to json: %v", key, err)
		}
		if string(gotBuf) != string(wantBuf) {
			return fmt.Errorf("case[%s]: got != want:\ngot: %v\nwant: %v", key, string(gotBuf), string(wantBuf))
		}
	}
	{
		key := "0/1"
		value := 13.37
		var got float64

		if err := kv.Set(key, value); err != nil {
			return fmt.Errorf("case[%s]: error: %v", key, err)
		}
		if f, err := kv.Get(key, &got); err != nil {
			return fmt.Errorf("case[%s]: error: %v", key, err)
		} else if !f {
			return fmt.Errorf("case[%s]: not found", key)
		}
		gotBuf, err := json.MarshalIndent(got, "", "\t")
		if err != nil {
			return fmt.Errorf("case[%s]: error marshalling got to json: %v", key, err)
		}
		wantBuf, err := json.MarshalIndent(value, "", "\t")
		if err != nil {
			return fmt.Errorf("case[%s]: error marshalling got to json: %v", key, err)
		}
		if string(gotBuf) != string(wantBuf) {
			return fmt.Errorf("case[%s]: got != want:\ngot: %v\nwant: %v", key, string(gotBuf), string(wantBuf))
		}
	}
	{
		key := "0/2"
		value := "1337"
		var got string

		if err := kv.Set(key, value); err != nil {
			return fmt.Errorf("case[%s]: error: %v", key, err)
		}
		if f, err := kv.Get(key, &got); err != nil {
			return fmt.Errorf("case[%s]: error: %v", key, err)
		} else if !f {
			return fmt.Errorf("case[%s]: not found", key)
		}
		gotBuf, err := json.MarshalIndent(got, "", "\t")
		if err != nil {
			return fmt.Errorf("case[%s]: error marshalling got to json: %v", key, err)
		}
		wantBuf, err := json.MarshalIndent(value, "", "\t")
		if err != nil {
			return fmt.Errorf("case[%s]: error marshalling got to json: %v", key, err)
		}
		if string(gotBuf) != string(wantBuf) {
			return fmt.Errorf("case[%s]: got != want:\ngot: %v\nwant: %v", key, string(gotBuf), string(wantBuf))
		}
	}
	{
		key := "0/3"
		value := struct{ V string }{V: "1337"}
		var got struct{ V string }

		if err := kv.Set(key, value); err != nil {
			return fmt.Errorf("case[%s]: error: %v", key, err)
		}
		if f, err := kv.Get(key, &got); err != nil {
			return fmt.Errorf("case[%s]: error: %v", key, err)
		} else if !f {
			return fmt.Errorf("case[%s]: not found", key)
		}
		gotBuf, err := json.MarshalIndent(got, "", "\t")
		if err != nil {
			return fmt.Errorf("case[%s]: error marshalling got to json: %v", key, err)
		}
		wantBuf, err := json.MarshalIndent(value, "", "\t")
		if err != nil {
			return fmt.Errorf("case[%s]: error marshalling got to json: %v", key, err)
		}
		if string(gotBuf) != string(wantBuf) {
			return fmt.Errorf("case[%s]: got != want:\ngot: %v\nwant: %v", key, string(gotBuf), string(wantBuf))
		}
	}
	return nil
}

func testListAndDelete(kv fkv.FKV) error {
	// list keys
	var keysGot []string
	if err := kv.List(&keysGot); err != nil {
		return fmt.Errorf("list keys error: %v", err)
	}

	// delete all
	for i := range keysGot {
		if err := kv.Delete(keysGot[i]); err != nil {
			return fmt.Errorf("delete[%d]: error: %v", i, err)
		}
	}

	// list keys again
	keysGot = nil
	if err := kv.List(&keysGot); err != nil {
		return fmt.Errorf("list keys error: %v", err)
	}

	// list should be empty
	if len(keysGot) != 0 {
		return fmt.Errorf("after deleting listed keys should be empty: got: %v", keysGot)
	}

	return nil
}
