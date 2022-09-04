module example

go 1.19

replace github.com/fritzkeyzer/go-fkv/disk => ../disk

replace github.com/fritzkeyzer/go-fkv => ../../go-fkv

require (
	github.com/fritzkeyzer/go-fkv v0.0.0-00010101000000-000000000000
	github.com/fritzkeyzer/go-fkv/disk v0.0.0-00010101000000-000000000000
)

require (
	github.com/fritzkeyzer/go-utils/pretty v0.0.0-20220823233912-ca82e21eaee4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
)
