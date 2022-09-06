# FKV
Package fkv provides a simple key-value storage interface.

Designed to abstract storage implementations away from usage.

Simplifies app testing, since it is trivial to swap between different storage implementations. 
Eg: local file system, cloud-storage or in-memory. 

Naive implementations of FKV are provided:
- github.com/fritzkeyzer/go-fkv/disk
- github.com/fritzkeyzer/go-fkv/mem
