# go table

# example

```go
package main

import (
	"os"
	"github.com/skaji/go-table"
)

func main() {
	t := table.New()
	t.Add([]string{"module_name", "owner", "co_maintainers"})
	t.Add([]string{"App::cpm", "SKAJI", "N/A"})
	t.Add([]string{"App::cpm::Distribution", "SKAJI", "N/A"})
	t.Add([]string{"App::cpm::Job", "SKAJI", "N/A"})
	t.Add([]string{"App::cpm::Logger", "SKAJI", "N/A"})
	t.Add([]string{"App::cpm::Logger::File", "SKAJI", "N/A"})
	t.Add([]string{"App::cpm::Master", "SKAJI", "N/A"})
	t.Render(os.Stdout)
}
```
```
❯ go run _example/main.go
module_name            owner co_maintainers
App::cpm               SKAJI N/A
App::cpm::Distribution SKAJI N/A
App::cpm::Job          SKAJI N/A
App::cpm::Logger       SKAJI N/A
App::cpm::Logger::File SKAJI N/A
App::cpm::Master       SKAJI N/A
```
