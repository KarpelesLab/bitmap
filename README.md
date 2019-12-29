[![GoDoc](https://godoc.org/github.com/KarpelesLab/bitmap?status.svg)](https://godoc.org/github.com/KarpelesLab/bitmap)

# Bitmap

Simple bitmaps in Go.

Provides the following methods on bitmap objects:

* Get(bit int)
* Set(bit int, value bool)
* GetAtomic(bit int)
* SetAtomic(bit int, value bool)

## Example

```Go
	map := bitmap.New(127)

	map.Set(42, true)

	if map.Get(42) {
		// OK
	}
```

