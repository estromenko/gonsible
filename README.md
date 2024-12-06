# gonsible

Simple Ansible-like tool written in Go that uses toml instead of yaml

## Example usage:

```bash
go build .

gonsible run ./example/pipeline.toml -i ./example/inventory.toml
```
