# https://www.acmicpc.net/ problem set solutions

I am learning Golang for fun :)

## Golang switch statement is unlike other languages that I've tried.
if I have switch statement like this.

```Go
switch 0 {
	case 0:
	case 1:
		fmt.Println("case 0 and case 1")
		return
	case 2:
		fmt.Println("case 2")
		return
}
fmt.Println("default")
```

It will print out "deafult" instead of "case 0 and case 1".
So if I want to combine two cases I should write like this.

```Go
switch 0 {
	case 0, 1:
		fmt.Println("case 0 and case 1")
		return
	case 2:
		fmt.Println("case 2")
		return
}
fmt.Println("default)
```

This will give me desirable result.

# My go cheat sheets

build go binary file: go build

format go source file: go fmt fileName.go
