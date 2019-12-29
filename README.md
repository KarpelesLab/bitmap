[![GoDoc](https://godoc.org/github.com/KarpelesLab/bitmap?status.svg)](https://godoc.org/github.com/KarpelesLab/bitmap)

# Bitmap

Simple bitmaps in Go.

Provides the following methods on bitmap objects:

* Get(bit int)
* Set(bit int, value bool)
* Toggle(bit)
* GetAtomic(bit int)
* SetAtomic(bit int, value bool)
* ToggleAtomic(bit int)

## Example

```Go
	m := bitmap.New(127)

	m.Set(42, true)

	if m.Get(42) {
		// OK
	}
```

# Why?

There are already a few bitmap implementations in Go available out there, however they have some shortfalls and/or are too complex for what a bitmap should do.

* [boljen's go-bitmap](https://godoc.org/github.com/boljen/go-bitmap) has separate implementations for Bitmap/Concurrent/Threadsafe, which feels a bit un-needed, and looks generally okay except for lack of support for big endian.
* [ShawnMilo's bitmap](https://www.godoc.org/github.com/ShawnMilo/bitmap) has a nice feel but lacks atomic methods and adds string methods using json/gzip/base64 which feels a bit overkill
* [Roaring Bitmaps](https://godoc.org/github.com/RoaringBitmap/roaring) are simply too complex for what I need bitmaps for. You should however definitely use that if you store more than 200k bits or so.

