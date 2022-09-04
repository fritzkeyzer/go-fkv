module tests

go 1.19

replace github.com/fritzkeyzer/go-fkv => ../

replace github.com/fritzkeyzer/go-fkv/disk => ../disk

replace github.com/fritzkeyzer/go-fkv/mem => ../mem

replace github.com/fritzkeyzer/go-fkv/storj => ../storj

require (
	github.com/fritzkeyzer/go-fkv v0.0.0-00010101000000-000000000000
	github.com/fritzkeyzer/go-fkv/disk v0.0.0-00010101000000-000000000000
	github.com/fritzkeyzer/go-fkv/mem v0.0.0-00010101000000-000000000000
	github.com/fritzkeyzer/go-fkv/storj v0.0.0-00010101000000-000000000000
)

require (
	github.com/calebcase/tmpfile v1.0.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/spacemonkeygo/monkit/v3 v3.0.17 // indirect
	github.com/vivint/infectious v0.0.0-20200605153912-25a574ae18a3 // indirect
	github.com/zeebo/errs v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20220131195533-30dcbda58838 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220128215802-99c3d69c2c27 // indirect
	storj.io/common v0.0.0-20220414110316-a5cb7172d6bf // indirect
	storj.io/drpc v0.0.30 // indirect
	storj.io/uplink v1.9.0 // indirect
)
