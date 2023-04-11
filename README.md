# Wordify

Wordify returns the Spanish word for a given integer number ranging from -999_999_999_999 up to 999_999_999_999.

## Sample

Invoking `wordify` would involve the following piece of code:
```go
func main() {
    number := 21_324
    inWords := wordify.Int(number) // returns `veintiún mil trescientos veinticuatro`
}
```
