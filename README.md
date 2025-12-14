# go-basics

format file:

```
go fmt <path_to_file>
```

execute file:

```
go run <path_to_file>
```

or:

```
go build 1chapter/hello_world/hello_world.go
./hello_world
```

```
go mod init <project-name>
// file go.mod appears
go get golang.org/x/text
```

Styling:

Upper case - variable/function CAN be exported

Lower case - local usage only.

`fmt.Println` - upper case -> can be exported

`quantity` := 4 -> no


to get in terminal information about method:
```
go doc fmt.Println
```