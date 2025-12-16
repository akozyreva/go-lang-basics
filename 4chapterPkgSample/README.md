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