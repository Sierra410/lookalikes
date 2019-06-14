## Lookalikes
[![GoDoc](https://godoc.org/github.com/Sierra410/lookalikes?status.svg)](https://godoc.org/github.com/Sierra410/lookalikes)

It's a tiny library that converts 𝐬𝐭𝐮𝐟𝐟 𝒍𝒊𝒌𝒆 тЂіѕ (popular unicode confusables/lookalikes and other symbols that might be used as letters) into proper english text. Well, most of the times.

It doesn't use any fancy algorithms or anything like that, just compares stuff a lot.

It has exactly one function that does exactly one thing.

```golang
s := lookalikes.Decode("𝐬𝐭𝐮𝐟𝐟 𝒍𝒊𝒌𝒆 тЂіѕ")
fmt.Println(s)
```
prints
```
stuff like this
```
that's it.