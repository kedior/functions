# functions

 Provides some higher-order functions

### examples

These two functions are identical

```go
func SelectCodeFiles(files []string) []string {
  return Filter(func(s string) bool {
    return Some(Curry21(strings.HasSuffix)(s), ".go", ".cpp", ".c", ".java", ".js")
  }, files...)
}
```
```go
func SelectCodeFiles(files []string) []string {
  var codes []string
  for _, file := range files {
    for _, suffix := range []string{".go", ".cpp", ".c", ".java", ".js"} {
      if strings.HasSuffix(file, suffix) {
        codes = append(codes, file)
        break
      }
    }
  }
  return codes
}

```
