# functions

 Provides some higher-order functions

### examples
Here are some higher-order functions

```
// make f(x)->y to f(t)->y by f(t)->x
func Substitute[X, Y, T any](fn func(X) Y, substitution func(T) X) func(T) Y

// currying, make f(a,b)->r to f(a)->(f(b)->r)
func Curry21[A, B, R any](fn func(A, B) R) func(A) func(B) R

// make f(a,b)->x to f(b,a)->x
func ArgRev1BA[A, B, X any](fn func(A, B) X) func(B, A) X

// select all x from x[] for which f(x)->true
func Filter[T any](fn func(T) bool, args []T) []T

// if all x in x[] make f(x)->true, this function returns true
func Some[T any](fn func(T) bool, args []T) bool
```
Using the above functions, we can make the following method identical

```
func SelectCodeFiles1(files []string) []string {
  suffix := ToArr(".go", ".cpp", ".c", ".java", ".js")

  some := Some[string]             // => (str->bool,str[]) -> bool
  revSome := ArgRev1BA(some)       // => (str[],str->bool) -> bool
  curryingSome := Curry21(revSome) // => (str[]) -> ((str->bool)->bool)
  a := curryingSome(suffix)        // => (str->bool) -> bool
  b := Curry21(strings.HasSuffix)  // => (str) -> ((str)->bool)

  isCodeFile := Substitute(a, b)   // => (str) -> bool
  return Filter(isCodeFile, files) // => str[]
}

func SelectCodeFiles2(files []string) []string {
  return Filter(
    Substitute(
      Curry21(ArgRev1BA(Some[string]))(ToArr(".go", ".cpp", ".c", ".java", ".js")),
      Curry21(strings.HasSuffix),
    ), files)
}

func SelectCodeFiles3(files []string) []string {
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
