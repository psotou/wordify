# Wordify

Wordify returns a given positive number with a range up to `999_999_999_999` with the words that actually represent this number in Spanish.

## Sample

Invoking `wordify` would involve the following piece of code:
```go
func main() {
    number := 21_324
    inWords, err := wordify.Int(number) // would return `veintiún mil trescientos veinticuatro`
}
```
