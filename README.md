# wordtwist
something

# compiling
You need to have [Go](https://go.dev) installed
```sh
go build main.go
```
This outputs a binary named `main`

# running
`./main encode` - encodes something, drops it to `out.txt`
`./main decode` - decode content from `out.txt`

For the seed you can put any number. It is used for the random number generator.
For the step you shouldn't put too high numbers as it could overflow. It is used for scrambling the contents of the string.
It is best to use wordtwist on regular ASCII characters