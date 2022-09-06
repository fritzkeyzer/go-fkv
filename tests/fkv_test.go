package tests

import (
	"github.com/fritzkeyzer/go-fkv/disk"
	"github.com/fritzkeyzer/go-fkv/mem"
	"testing"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func TestDiskKV(t *testing.T) {
	kv := disk.NewKV("test_dir")
	if err := Test(kv); err != nil {
		t.Fatal(err)
	}
}

func TestMemKV(t *testing.T) {
	kv := mem.NewKV()
	if err := Test(kv); err != nil {
		t.Fatal(err)
	}
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func BenchmarkDiskKV(b *testing.B) {
	kv := disk.NewKV("test_dir")
	Bench(kv, b.N)
}

func BenchmarkMemKV(b *testing.B) {
	kv := mem.NewKV()
	Bench(kv, b.N)
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
