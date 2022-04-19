# Memory Source

memory source 提供内存中的数据作为源

## Memory Format

预期的数据格式是 json

```json
data := []byte(`{
    "hosts": {
        "database": {
            "address": "10.0.0.1",
            "port": 3306
        },
        "cache": {
            "address": "10.0.0.2",
            "port": 6379
        }
    }
}`)
```

## New Source

Specify source with data

```go
memorySource := memory.NewSource(
	memory.WithJSON(data),
)
```

## Load Source

Load the source into config

```go
// Create new config
conf := config.NewConfig()

// Load memory source
conf.Load(memorySource)
```
