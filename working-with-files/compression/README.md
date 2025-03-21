# Compression in Go

Go provides support for both archiving and compression, which are often used 
together to package multiple files into a single, smaller file. A common format 
is `.tar.gz` (gzipped tar ball). Remember, `zip` and `gzip` are different formats.

Supported compression algorithms in Go:

- **bzip2**: bzip2 format
- **flate**: DEFLATE (RFC 1951)
- **gzip**: gzip format (RFC 1952)
- **lzw**: Lempel-Ziv-Welch format (from "A Technique for High-Performance Data 
  Compression," Computer, 17(6) (June 1984))
- **zlib**: zlib format (RFC 1950)

For more details, see the [Go Compression Packages](https://golang.org/pkg/compress/). 
This guide uses gzip compression, but switching to any other format is simple.
