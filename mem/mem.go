package mem

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"
)

type KV struct {
	sync.Mutex
	objects map[string][]byte
}

func NewKV() *KV {
	kv := KV{
		objects: make(map[string][]byte),
	}
	return &kv
}

func (kv *KV) Get(key string, ptr any) (found bool, err error) {
	kv.Lock()
	defer kv.Unlock()

	obytes, f := kv.objects[key]
	if !f {
		return false, nil
	}

	buf := bytes.NewBuffer(obytes)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(ptr); err != nil {
		return true, fmt.Errorf("decode: %v", err)
	}

	return true, nil
}

func (kv *KV) Set(key string, object any) error {
	kv.Lock()
	defer kv.Unlock()

	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(object)
	if err != nil {
		return fmt.Errorf("encode: %v", err)
	}
	kv.objects[key] = buffer.Bytes()

	return nil
}

func (kv *KV) Delete(key string) error {
	kv.Lock()
	defer kv.Unlock()

	delete(kv.objects, key)

	return nil
}

func (kv *KV) List(ptr *[]string) error {
	for k := range kv.objects {
		*ptr = append(*ptr, k)
	}

	return nil
}
