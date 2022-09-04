# FKV
Package fkv provides a simple key-value storage interface.

Designed to abstract storage implementations away from usage.

Simplifies app testing, since it is trivial to swap between different storage implementations. 
Eg: local file system, cloud-storage or in-memory. 

- Provided implementations of FKV:
  - github.com/fritzkeyzer/go-fkv/disk
  - github.com/fritzkeyzer/go-fkv/mem
  - github.com/fritzkeyzer/go-fkv/storj

- Roadmap implementations
  - postgres sql
  - S3 bucket
  - Google cloud storage