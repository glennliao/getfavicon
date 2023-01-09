## INSTALL
`go get github.com/glennliao/getfavicon`

## USE
```go
func TestGet(t *testing.T) {
	favicon, err := getfavicon.Get("https://www.github.com/")
	if err != nil {
		panic(err)
	}

	fmt.Println(favicon)
}

```