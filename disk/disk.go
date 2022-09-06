package disk

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// todo file locks
type KV struct {
	dir string
}

func NewKV(dir string) *KV {
	kv := KV{
		dir: dir,
	}
	return &kv
}

func (kv *KV) Get(key string, ptr any) (found bool, err error) {
	fullpath := filepath.Join(kv.dir, key)

	if _, err := os.Stat(fullpath); errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	buf, err := os.ReadFile(fullpath)
	if err != nil {
		return true, fmt.Errorf("reading file: %w", err)
	}

	if err := json.Unmarshal(buf, ptr); err != nil {
		return true, fmt.Errorf("decoding: %w", err)
	}

	return true, nil
}

func (kv *KV) Set(key string, object any) error {
	data, err := json.MarshalIndent(object, "", "\t")
	if err != nil {
		return fmt.Errorf("encoding: %v", err)
	}

	dir := filepath.Join(kv.dir, filepath.Dir(key))
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("creating path: %w", err)
	}

	fullpath := filepath.Join(kv.dir, key)
	err = os.WriteFile(fullpath, data, os.ModePerm)
	if err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	return nil
}

func (kv *KV) Delete(key string) error {
	fullpath := filepath.Join(kv.dir, key)
	err := os.Remove(fullpath)
	if err != nil {
		return fmt.Errorf("deleting file: %w", err)
	}

	return nil
}

func (kv *KV) List(ptr *[]string) error {
	err := filepath.Walk(kv.dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		path = strings.TrimPrefix(path, kv.dir)
		cleanPath := strings.TrimPrefix(path, "/")

		if !info.IsDir() {
			*ptr = append(*ptr, cleanPath)
		}
		return nil
	})

	return err
}
