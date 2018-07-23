Go ETL
=======

Go ETL using [Ratchet](https://github.com/dailyburn/ratchet) pipelines

## Start

`make start`

## Configure

To configure, edit ./config/config.yaml to load a new pipeline.

To add a custom ETL, create a new plugin on ./plugins and add is on config.yaml.

## Examples

### Add an ETL code on config.yaml

1) Add on config/config.yaml:
```yaml
workers:
  - schedule:
      hour: 20
      minute: 0
    job:
      name: http-requestor
      code: |
        package main

        import (
          "github.com/Henrod/go-etl/processors"
          "github.com/dailyburn/ratchet"
        )

        type etl string

        func (e etl) Extract() ratchet.DataProcessor {
          return processors.NewHTTPRequestor("GET", "http://localhost:8080")
        }

        func (e etl) Transform() ratchet.DataProcessor {
          return &processors.Logger{}
        }

        func (e etl) Load() ratchet.DataProcessor {
          return &processors.Null{}
        }

        // ETL is the exported symbol of this plugin
        var ETL etl
```

2) Start:
```bash
make start
```

### Create a new ETL plugin

1) Create a new plugin on ./plugins like this: 
```golang
// ./plugins/http-requestor/main.go

package main

import (
	"github.com/Henrod/go-etl/processors"
	"github.com/dailyburn/ratchet"
)

type etl string

func (e etl) Extract() ratchet.DataProcessor {
	return processors.NewHTTPRequestor("GET", "http://localhost:8080")
}

func (e etl) Transform() ratchet.DataProcessor {
	return &processors.Logger{}
}

func (e etl) Load() ratchet.DataProcessor {
	return &processors.Null{}
}

// ETL is the exported symbol of this plugin
var ETL etl
```

2) Build the plugin binary:

```bash
make plugins
```

3) Add on config/config.yaml:
```yaml
workers:
  - period: 1h
    job:
      name: http-requestor
```

4) Start:
```bash
make start
```
