// Package fkv provides a simple key-value storage interface.
package fkv

// FKV provides a simple key-value storage interface.
// Example implementations of FKV:
//	github.com/fritzkeyzer/go-fkv/disk
//	github.com/fritzkeyzer/go-fkv/mem
//	github.com/fritzkeyzer/go-fkv/storj
type FKV interface {
	Set(key string, object any) error
	Get(key string, ptr any) (found bool, err error)
	Delete(key string) error
}
