```
go mod init 4chapterPkgSample
```

after that package with name `greeting` will work

import in `main.go` is:

```
import "4chapterPkgSample/greeting"
```

package naming convention:

- starts with small letter and short word if possible: `fmt`
- use 1 word, if not possible, write it all as one: `strconv`

to run multiple files with `func main()` run them directly:

```
go run tocelsius.go
```

greeting package - sample of nested package files

get documentation from package

```
❯ go doc keyboard
package keyboard // import "4chapterPkgSample/keyboard"

Package keyboard this documentation will be visible, when you run `go doc` cmd

func GetFloat() (float64, error)
```

get documentation from function:

```
go doc keyboard.GetFloat
```