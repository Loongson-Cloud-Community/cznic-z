// Copyright 2023 The Zlib-Go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package z // import "modernc.org/z/v2"

import (
	"modernc.org/libc/v2"
)

func Xcompress(tls *libc.TLS, dest uintptr, destLen uintptr, source uintptr, sourceLen uint64) (r int32) {
	return x_compress(tls, dest, destLen, source, sourceLen)
}

func Xuncompress(tls *libc.TLS, dest uintptr, destLen uintptr, source uintptr, _sourceLen uint64) (r int32) {
	return x_uncompress(tls, dest, destLen, source, _sourceLen)
}

func Xgzopen(tls *libc.TLS, path uintptr, mode uintptr) (r uintptr) {
	return x_gzopen(tls, path, mode)
}

func Xgzputc(tls *libc.TLS, file uintptr, c int32) (r int32) {
	return x_gzputc(tls, file, c)
}

func Xgzputs(tls *libc.TLS, file uintptr, s uintptr) (r int32) {
	return x_gzputs(tls, file, s)
}

func Xgzerror(tls *libc.TLS, file uintptr, errnum uintptr) (r uintptr) {
	return x_gzerror(tls, file, errnum)
}

func Xgzprintf(tls *libc.TLS, file uintptr, format uintptr, va uintptr) (r int32) {
	return x_gzprintf(tls, file, format, va)
}

func Xgzseek(tls *libc.TLS, file uintptr, offset int64, whence int32) (r int64) {
	return x_gzseek(tls, file, offset, whence)
}

func Xgzclose(tls *libc.TLS, file uintptr) (r int32) {
	return x_gzclose(tls, file)
}

func Xgzread(tls *libc.TLS, file uintptr, buf uintptr, len uint32) (r int32) {
	return x_gzread(tls, file, buf, len)
}

func Xgztell(tls *libc.TLS, file uintptr) (r int64) {
	return x_gztell(tls, file)
}

func Xgzgetc(tls *libc.TLS, file uintptr) (r int32) {
	return x_gzgetc(tls, file)
}

func Xgzungetc(tls *libc.TLS, c int32, file uintptr) (r int32) {
	return x_gzungetc(tls, c, file)
}

func Xgzgets(tls *libc.TLS, file uintptr, buf uintptr, len int32) (r uintptr) {
	return x_gzgets(tls, file, buf, len)
}

func XdeflateInit_(tls *libc.TLS, strm uintptr, level int32, version uintptr, stream_size int32) (r int32) {
	return x_deflateInit_(tls, strm, level, version, stream_size)
}

func Xdeflate(tls *libc.TLS, strm uintptr, flush int32) (r int32) {
	return x_deflate(tls, strm, flush)
}

func XdeflateEnd(tls *libc.TLS, strm uintptr) (r int32) {
	return x_deflateEnd(tls, strm)
}

func XinflateInit_(tls *libc.TLS, strm uintptr, version uintptr, stream_size int32) (r int32) {
	return x_inflateInit_(tls, strm, version, stream_size)
}

func Xinflate(tls *libc.TLS, strm uintptr, flush int32) (r int32) {
	return x_inflate(tls, strm, flush)
}

func XinflateEnd(tls *libc.TLS, strm uintptr) (r int32) {
	return x_inflateEnd(tls, strm)
}

func XdeflateParams(tls *libc.TLS, strm uintptr, level int32, strategy int32) (r int32) {
	return x_deflateParams(tls, strm, level, strategy)
}

func XinflateSync(tls *libc.TLS, strm uintptr) (r int32) {
	return x_inflateSync(tls, strm)
}

func XdeflateSetDictionary(tls *libc.TLS, strm uintptr, dictionary uintptr, dictLength uint32) (r int32) {
	return x_deflateSetDictionary(tls, strm, dictionary, dictLength)
}

func XinflateSetDictionary(tls *libc.TLS, strm uintptr, dictionary uintptr, dictLength uint32) (r int32) {
	return x_inflateSetDictionary(tls, strm, dictionary, dictLength)
}

func XzlibVersion(tls *libc.TLS) (r uintptr) {
	return x_zlibVersion(tls)
}

func XzlibCompileFlags(tls *libc.TLS) (r uint64) {
	return x_zlibCompileFlags(tls)
}

func Xgzwrite(tls *libc.TLS, file uintptr, buf uintptr, len uint32) (r int32) {
	return x_gzwrite(tls, file, buf, len)
}

func Xgzdopen(tls *libc.TLS, fd int32, mode uintptr) (r uintptr) {
	return x_gzdopen(tls, fd, mode)
}
