package tests

import (
	"github.com/fritzkeyzer/go-fkv"
	"github.com/fritzkeyzer/go-fkv/disk"
	"github.com/fritzkeyzer/go-fkv/mem"
	"github.com/fritzkeyzer/go-fkv/storj"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestDiskKV(t *testing.T) {
	kv := disk.NewKV("test_dir")
	test(kv, t)
}

func TestMemKV(t *testing.T) {
	kv := mem.NewKV()
	test(kv, t)
}

func TestStorjKV(t *testing.T) {
	kv, err := storj.NewKV("test-bucket", os.Getenv("STORJ_ACCESS"))
	if err != nil {
		t.Fatal(err)
	}
	test(kv, t)
}

func test(kv fkv.FKV, t *testing.T) {
	type Object struct {
		Field string
	}

	o1 := Object{Field: "hello world"}
	if err := kv.Set("file0", o1); err != nil {
		t.Fatal(err)
	}

	var o2 Object
	if f, err := kv.Get("file0", &o2); err != nil || !f {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(o1, o2) {
		t.Fatalf("got != want:\n\tgot: %v\n\twant: %v", o2, o1)
	}

	//if err := kv.Delete("file0"); err != nil {
	//	t.Fatal(err)
	//}

	var o3 Object
	found, err := kv.Get("file1", &o3)
	if found {
		t.Fatalf("file1 should not exist!")
	}
	if err != nil {
		t.Fatalf("expect no error, got: %v", err)
	}
}

func BenchmarkDiskKV(b *testing.B) {
	kv := disk.NewKV("test_dir")
	bench(kv, b.N)
}

func BenchmarkMemKV(b *testing.B) {
	kv := mem.NewKV()
	bench(kv, b.N)
}

func BenchmarkStorjKV(b *testing.B) {
	kv, err := storj.NewKV("test-bucket", os.Getenv("STORJ_ACCESS"))
	if err != nil {
		log.Fatal(err)
	}
	bench(kv, b.N)
}

func bench(kv fkv.FKV, n int) {
	type Object struct {
		Field string
	}

	o1 := Object{Field: "hello world"}
	for i := 0; i < n; i++ {
		if err := kv.Set("file0", o1); err != nil {
			log.Fatal(err)
		}

		var o2 Object
		if f, err := kv.Get("file0", &o2); err != nil || !f {
			log.Fatal(err)
		}

		if err := kv.Delete("file0"); err != nil {
			log.Fatal(err)
		}
	}
}
