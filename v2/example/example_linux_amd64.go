// Code generated for linux/amd64 by 'gcc -o example/example_linux_amd64.go --package-name=main example_linux_amd64.go -lmodernc.org/z/v2', DO NOT EDIT.

//go:build linux && amd64
// +build linux,amd64

package main

import (
	"reflect"
	"unsafe"

	"modernc.org/libc/v2"
	"modernc.org/z/v2"
)

var (
	_ reflect.Type
	_ unsafe.Pointer
)

const FD_SETSIZE = 1024
const SEEK_CUR = 1
const TESTFILE = "foo.gz"
const ZLIB_VERNUM = 4864
const ZLIB_VERSION = "1.3"
const Z_BEST_COMPRESSION = 9
const Z_DEFAULT_STRATEGY = 0
const Z_FILTERED = 1
const Z_FINISH = 4
const Z_FULL_FLUSH = 3
const Z_NEED_DICT = 2
const Z_NO_COMPRESSION = 0
const Z_NO_FLUSH = 0
const Z_NULL = 0
const Z_OK = 0
const Z_STREAM_END = 1

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int64

type wchar_t = int32

type max_align_t = struct {
	__ll int64
	__ld float64
}

type size_t = uint64

type ptrdiff_t = int64

type z_size_t = uint64

type Byte = uint8

type uInt = uint32

type uLong = uint64

type Bytef = uint8

type charf = int8

type intf = int32

type uIntf = uint32

type uLongf = uint64

type voidpc = uintptr

type voidpf = uintptr

type voidp = uintptr

type z_crc_t = uint32

type ssize_t = int64

type register_t = int64

type time_t = int64

type suseconds_t = int64

type int8_t = int8

type int16_t = int16

type int32_t = int32

type int64_t = int64

type u_int64_t = uint64

type mode_t = uint32

type nlink_t = uint64

type off_t = int64

type ino_t = uint64

type dev_t = uint64

type blksize_t = int64

type blkcnt_t = int64

type fsblkcnt_t = uint64

type fsfilcnt_t = uint64

type timer_t = uintptr

type clockid_t = int32

type clock_t = int64

type pid_t = int32

type id_t = uint32

type uid_t = uint32

type gid_t = uint32

type key_t = int32

type useconds_t = uint32

type pthread_t = uintptr

type pthread_once_t = int32

type pthread_key_t = uint32

type pthread_spinlock_t = int32

type pthread_mutexattr_t = struct {
	__attr uint32
}

type pthread_condattr_t = struct {
	__attr uint32
}

type pthread_barrierattr_t = struct {
	__attr uint32
}

type pthread_rwlockattr_t = struct {
	__attr [2]uint32
}

type pthread_attr_t = struct {
	__u struct {
		__vi [0][14]int32
		__s  [0][7]uint64
		__i  [14]int32
	}
}

type pthread_mutex_t = struct {
	__u struct {
		__vi [0][10]int32
		__p  [0][5]uintptr
		__i  [10]int32
	}
}

type pthread_cond_t = struct {
	__u struct {
		__vi [0][12]int32
		__p  [0][6]uintptr
		__i  [12]int32
	}
}

type pthread_rwlock_t = struct {
	__u struct {
		__vi [0][14]int32
		__p  [0][7]uintptr
		__i  [14]int32
	}
}

type pthread_barrier_t = struct {
	__u struct {
		__vi [0][8]int32
		__p  [0][4]uintptr
		__i  [8]int32
	}
}

type u_int8_t = uint8

type u_int16_t = uint16

type u_int32_t = uint32

type caddr_t = uintptr

type u_char = uint8

type u_short = uint16

type ushort = uint16

type u_int = uint32

type uint1 = uint32

type u_long = uint64

type ulong = uint64

type quad_t = int64

type u_quad_t = uint64

type uint16_t = uint16

type uint32_t = uint32

type uint64_t = uint64

type timeval = struct {
	tv_sec  int64
	tv_usec int64
}

type timespec = struct {
	tv_sec  int64
	tv_nsec int64
}

type sigset_t = struct {
	__bits [16]uint64
}

type __sigset_t = sigset_t

type fd_mask = uint64

type fd_set = struct {
	fds_bits [16]uint64
}

type va_list = uintptr

type intptr_t = int64

type alloc_func = uintptr

type free_func = uintptr

type z_stream = struct {
	next_in   uintptr
	avail_in  uint32
	total_in  uint64
	next_out  uintptr
	avail_out uint32
	total_out uint64
	msg       uintptr
	state     uintptr
	zalloc    uintptr
	zfree     uintptr
	opaque    uintptr
	data_type int32
	adler     uint64
	reserved  uint64
}

type z_stream_s = z_stream

type z_streamp = uintptr

type gz_header = struct {
	text         int32
	time         uint64
	xflags       int32
	os           int32
	extra        uintptr
	extra_len    uint32
	extra_max    uint32
	name         uintptr
	name_max     uint32
	comment      uintptr
	comm_max     uint32
	hcrc         int32
	done         int32
	__ccgo_pad13 [4]byte
}

type gz_header_s = gz_header

type gz_headerp = uintptr

type in_func = uintptr

type out_func = uintptr

type gzFile = uintptr

type gzFile_s = struct {
	have uint32
	next uintptr
	pos  int64
}

type __isoc_va_list = uintptr

type fpos_t = struct {
	__lldata [0]int64
	__align  [0]float64
	__opaque [16]int8
}

type _G_fpos64_t = fpos_t

type cookie_io_functions_t = struct {
	read   uintptr
	write  uintptr
	seek   uintptr
	close1 uintptr
}

type _IO_cookie_io_functions_t = cookie_io_functions_t

type locale_t = uintptr

type div_t = struct {
	quot int32
	rem  int32
}

type ldiv_t = struct {
	quot int64
	rem  int64
}

type lldiv_t = struct {
	quot int64
	rem  int64
}

var hello = [14]int8{'h', 'e', 'l', 'l', 'o', ',', ' ', 'h', 'e', 'l', 'l', 'o', '!'}

/* "hello world" would be more standard, but the repeated "hello"
 * stresses the compression code better, sorry...
 */

var dictionary = [6]int8{'h', 'e', 'l', 'l', 'o'}
var dictId uint64 /* Adler32 value of the dictionary */

var zalloc = uintptr(0)

func init() {
	p := unsafe.Pointer(&zalloc)
	*(*uintptr)(unsafe.Add(p, 0)) = libc.UintptrFromInt32(0)
}

var zfree = uintptr(0)

func init() {
	p := unsafe.Pointer(&zfree)
	*(*uintptr)(unsafe.Add(p, 0)) = libc.UintptrFromInt32(0)
}

// C documentation
//
//	/* ===========================================================================
//	 * Test compress() and uncompress()
//	 */
func test_compress(tls *libc.TLS, compr uintptr, _comprLen uint64, uncompr uintptr, _uncomprLen uint64) {
	bp := tls.Alloc(40) /* tlsAllocs 16 maxVaListSize 16 */
	defer tls.Free(40)
	*(*uint64)(unsafe.Pointer(bp)) = _comprLen
	*(*uint64)(unsafe.Pointer(bp + 8)) = _uncomprLen
	var err int32
	var len1 uint64
	len1 = libc.Xstrlen(tls, uintptr(unsafe.Pointer(&hello))) + uint64(1)
	err = z.Xcompress(tls, compr, bp, uintptr(unsafe.Pointer(&hello)), len1)
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+24, ts+14, err))
		libc.Xexit(tls, int32(1))
	}
	libc.Xstrcpy(tls, uncompr, ts+23)
	err = z.Xuncompress(tls, uncompr, bp+8, compr, *(*uint64)(unsafe.Pointer(bp)))
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+24, ts+31, err))
		libc.Xexit(tls, int32(1))
	}
	if libc.Xstrcmp(tls, uncompr, uintptr(unsafe.Pointer(&hello))) != 0 {
		libc.Xfprintf(tls, libc.Xstderr, ts+42, 0)
		libc.Xexit(tls, int32(1))
	} else {
		libc.Xprintf(tls, ts+58, libc.VaList(bp+24, uncompr))
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Test read/write of .gz files
//	 */
func test_gzio(tls *libc.TLS, fname uintptr, uncompr uintptr, uncomprLen uint64) {
	bp := tls.Alloc(32) /* tlsAllocs 8 maxVaListSize 16 */
	defer tls.Free(32)
	var file, v2, v3 uintptr
	var len1, v1 int32
	var pos int64
	var _ /* err at bp+0 */ int32
	len1 = int32(libc.Xstrlen(tls, uintptr(unsafe.Pointer(&hello)))) + int32(1)
	file = z.Xgzopen(tls, fname, ts+76)
	if file == libc.UintptrFromInt32(0) {
		libc.Xfprintf(tls, libc.Xstderr, ts+79, 0)
		libc.Xexit(tls, int32(1))
	}
	z.Xgzputc(tls, file, int32('h'))
	if z.Xgzputs(tls, file, ts+93) != int32(4) {
		libc.Xfprintf(tls, libc.Xstderr, ts+98, libc.VaList(bp+16, z.Xgzerror(tls, file, bp)))
		libc.Xexit(tls, int32(1))
	}
	if z.Xgzprintf(tls, file, ts+114, libc.VaList(bp+16, ts+120)) != int32(8) {
		libc.Xfprintf(tls, libc.Xstderr, ts+126, libc.VaList(bp+16, z.Xgzerror(tls, file, bp)))
		libc.Xexit(tls, int32(1))
	}
	z.Xgzseek(tls, file, int64(1), int32(SEEK_CUR)) /* add one zero byte */
	z.Xgzclose(tls, file)
	file = z.Xgzopen(tls, fname, ts+144)
	if file == libc.UintptrFromInt32(0) {
		libc.Xfprintf(tls, libc.Xstderr, ts+79, 0)
		libc.Xexit(tls, int32(1))
	}
	libc.Xstrcpy(tls, uncompr, ts+23)
	if z.Xgzread(tls, file, uncompr, uint32(uncomprLen)) != len1 {
		libc.Xfprintf(tls, libc.Xstderr, ts+147, libc.VaList(bp+16, z.Xgzerror(tls, file, bp)))
		libc.Xexit(tls, int32(1))
	}
	if libc.Xstrcmp(tls, uncompr, uintptr(unsafe.Pointer(&hello))) != 0 {
		libc.Xfprintf(tls, libc.Xstderr, ts+163, libc.VaList(bp+16, uncompr))
		libc.Xexit(tls, int32(1))
	} else {
		libc.Xprintf(tls, ts+179, libc.VaList(bp+16, uncompr))
	}
	pos = z.Xgzseek(tls, file, -int64(8), int32(SEEK_CUR))
	if pos != int64(6) || z.Xgztell(tls, file) != pos {
		libc.Xfprintf(tls, libc.Xstderr, ts+193, libc.VaList(bp+16, pos, z.Xgztell(tls, file)))
		libc.Xexit(tls, int32(1))
	}
	if (*gzFile_s)(unsafe.Pointer(file)).have != 0 {
		(*gzFile_s)(unsafe.Pointer(file)).have--
		(*gzFile_s)(unsafe.Pointer(file)).pos++
		v3 = file + 8
		v2 = *(*uintptr)(unsafe.Pointer(v3))
		*(*uintptr)(unsafe.Pointer(v3))++
		v1 = int32(*(*uint8)(unsafe.Pointer(v2)))
	} else {
		v1 = z.Xgzgetc(tls, file)
	}
	if v1 != int32(' ') {
		libc.Xfprintf(tls, libc.Xstderr, ts+228, 0)
		libc.Xexit(tls, int32(1))
	}
	if z.Xgzungetc(tls, int32(' '), file) != int32(' ') {
		libc.Xfprintf(tls, libc.Xstderr, ts+242, 0)
		libc.Xexit(tls, int32(1))
	}
	z.Xgzgets(tls, file, uncompr, int32(uncomprLen))
	if libc.Xstrlen(tls, uncompr) != uint64(7) { /* " hello!" */
		libc.Xfprintf(tls, libc.Xstderr, ts+258, libc.VaList(bp+16, z.Xgzerror(tls, file, bp)))
		libc.Xexit(tls, int32(1))
	}
	if libc.Xstrcmp(tls, uncompr, uintptr(unsafe.Pointer(&hello))+uintptr(6)) != 0 {
		libc.Xfprintf(tls, libc.Xstderr, ts+287, 0)
		libc.Xexit(tls, int32(1))
	} else {
		libc.Xprintf(tls, ts+312, libc.VaList(bp+16, uncompr))
	}
	z.Xgzclose(tls, file)
}

// C documentation
//
//	/* ===========================================================================
//	 * Test deflate() with small buffers
//	 */
func test_deflate(tls *libc.TLS, compr uintptr, comprLen uint64) {
	bp := tls.Alloc(136) /* tlsAllocs 112 maxVaListSize 16 */
	defer tls.Free(136)
	var err int32
	var len1 uint64
	var v1 uint32
	var _ /* c_stream at bp+0 */ z_stream
	/* compression stream */
	len1 = libc.Xstrlen(tls, uintptr(unsafe.Pointer(&hello))) + uint64(1)
	(*(*z_stream)(unsafe.Pointer(bp))).zalloc = zalloc
	(*(*z_stream)(unsafe.Pointer(bp))).zfree = zfree
	(*(*z_stream)(unsafe.Pointer(bp))).opaque = libc.UintptrFromInt32(0)
	err = z.XdeflateInit_(tls, bp, -int32(1), ts+339, libc.Int32FromInt64(112))
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+343, err))
		libc.Xexit(tls, int32(1))
	}
	(*(*z_stream)(unsafe.Pointer(bp))).next_in = uintptr(unsafe.Pointer(&hello))
	(*(*z_stream)(unsafe.Pointer(bp))).next_out = compr
	for (*(*z_stream)(unsafe.Pointer(bp))).total_in != len1 && (*(*z_stream)(unsafe.Pointer(bp))).total_out < comprLen {
		v1 = libc.Uint32FromInt32(1)
		(*(*z_stream)(unsafe.Pointer(bp))).avail_out = v1
		(*(*z_stream)(unsafe.Pointer(bp))).avail_in = v1 /* force small buffers */
		err = z.Xdeflate(tls, bp, Z_NO_FLUSH)
		if err != Z_OK {
			libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+355, err))
			libc.Xexit(tls, int32(1))
		}
	}
	/* Finish the stream, still forcing small buffers: */
	for {
		(*(*z_stream)(unsafe.Pointer(bp))).avail_out = uint32(1)
		err = z.Xdeflate(tls, bp, int32(Z_FINISH))
		if err == int32(Z_STREAM_END) {
			break
		}
		if err != Z_OK {
			libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+355, err))
			libc.Xexit(tls, int32(1))
		}
	}
	err = z.XdeflateEnd(tls, bp)
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+363, err))
		libc.Xexit(tls, int32(1))
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Test inflate() with small buffers
//	 */
func test_inflate(tls *libc.TLS, compr uintptr, comprLen uint64, uncompr uintptr, uncomprLen uint64) {
	bp := tls.Alloc(136) /* tlsAllocs 112 maxVaListSize 16 */
	defer tls.Free(136)
	var err int32
	var v1 uint32
	var _ /* d_stream at bp+0 */ z_stream
	/* decompression stream */
	libc.Xstrcpy(tls, uncompr, ts+23)
	(*(*z_stream)(unsafe.Pointer(bp))).zalloc = zalloc
	(*(*z_stream)(unsafe.Pointer(bp))).zfree = zfree
	(*(*z_stream)(unsafe.Pointer(bp))).opaque = libc.UintptrFromInt32(0)
	(*(*z_stream)(unsafe.Pointer(bp))).next_in = compr
	(*(*z_stream)(unsafe.Pointer(bp))).avail_in = uint32(0)
	(*(*z_stream)(unsafe.Pointer(bp))).next_out = uncompr
	err = z.XinflateInit_(tls, bp, ts+339, libc.Int32FromInt64(112))
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+374, err))
		libc.Xexit(tls, int32(1))
	}
	for (*(*z_stream)(unsafe.Pointer(bp))).total_out < uncomprLen && (*(*z_stream)(unsafe.Pointer(bp))).total_in < comprLen {
		v1 = libc.Uint32FromInt32(1)
		(*(*z_stream)(unsafe.Pointer(bp))).avail_out = v1
		(*(*z_stream)(unsafe.Pointer(bp))).avail_in = v1 /* force small buffers */
		err = z.Xinflate(tls, bp, Z_NO_FLUSH)
		if err == int32(Z_STREAM_END) {
			break
		}
		if err != Z_OK {
			libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+386, err))
			libc.Xexit(tls, int32(1))
		}
	}
	err = z.XinflateEnd(tls, bp)
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+394, err))
		libc.Xexit(tls, int32(1))
	}
	if libc.Xstrcmp(tls, uncompr, uintptr(unsafe.Pointer(&hello))) != 0 {
		libc.Xfprintf(tls, libc.Xstderr, ts+405, 0)
		libc.Xexit(tls, int32(1))
	} else {
		libc.Xprintf(tls, ts+418, libc.VaList(bp+120, uncompr))
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Test deflate() with large buffers and dynamic change of compression level
//	 */
func test_large_deflate(tls *libc.TLS, compr uintptr, comprLen uint64, uncompr uintptr, uncomprLen uint64) {
	bp := tls.Alloc(136) /* tlsAllocs 112 maxVaListSize 16 */
	defer tls.Free(136)
	var err int32
	var _ /* c_stream at bp+0 */ z_stream
	/* compression stream */
	(*(*z_stream)(unsafe.Pointer(bp))).zalloc = zalloc
	(*(*z_stream)(unsafe.Pointer(bp))).zfree = zfree
	(*(*z_stream)(unsafe.Pointer(bp))).opaque = libc.UintptrFromInt32(0)
	err = z.XdeflateInit_(tls, bp, int32(1), ts+339, libc.Int32FromInt64(112))
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+343, err))
		libc.Xexit(tls, int32(1))
	}
	(*(*z_stream)(unsafe.Pointer(bp))).next_out = compr
	(*(*z_stream)(unsafe.Pointer(bp))).avail_out = uint32(comprLen)
	/* At this point, uncompr is still mostly zeroes, so it should compress
	 * very well:
	 */
	(*(*z_stream)(unsafe.Pointer(bp))).next_in = uncompr
	(*(*z_stream)(unsafe.Pointer(bp))).avail_in = uint32(uncomprLen)
	err = z.Xdeflate(tls, bp, Z_NO_FLUSH)
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+355, err))
		libc.Xexit(tls, int32(1))
	}
	if (*(*z_stream)(unsafe.Pointer(bp))).avail_in != uint32(0) {
		libc.Xfprintf(tls, libc.Xstderr, ts+433, 0)
		libc.Xexit(tls, int32(1))
	}
	/* Feed in already compressed data and switch to no compression: */
	z.XdeflateParams(tls, bp, Z_NO_COMPRESSION, Z_DEFAULT_STRATEGY)
	(*(*z_stream)(unsafe.Pointer(bp))).next_in = compr
	(*(*z_stream)(unsafe.Pointer(bp))).avail_in = uint32(uncomprLen) / uint32(2)
	err = z.Xdeflate(tls, bp, Z_NO_FLUSH)
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+355, err))
		libc.Xexit(tls, int32(1))
	}
	/* Switch back to compressing mode: */
	z.XdeflateParams(tls, bp, int32(Z_BEST_COMPRESSION), int32(Z_FILTERED))
	(*(*z_stream)(unsafe.Pointer(bp))).next_in = uncompr
	(*(*z_stream)(unsafe.Pointer(bp))).avail_in = uint32(uncomprLen)
	err = z.Xdeflate(tls, bp, Z_NO_FLUSH)
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+355, err))
		libc.Xexit(tls, int32(1))
	}
	err = z.Xdeflate(tls, bp, int32(Z_FINISH))
	if err != int32(Z_STREAM_END) {
		libc.Xfprintf(tls, libc.Xstderr, ts+453, 0)
		libc.Xexit(tls, int32(1))
	}
	err = z.XdeflateEnd(tls, bp)
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+363, err))
		libc.Xexit(tls, int32(1))
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Test inflate() with large buffers
//	 */
func test_large_inflate(tls *libc.TLS, compr uintptr, comprLen uint64, uncompr uintptr, uncomprLen uint64) {
	bp := tls.Alloc(136) /* tlsAllocs 112 maxVaListSize 16 */
	defer tls.Free(136)
	var err int32
	var _ /* d_stream at bp+0 */ z_stream
	/* decompression stream */
	libc.Xstrcpy(tls, uncompr, ts+23)
	(*(*z_stream)(unsafe.Pointer(bp))).zalloc = zalloc
	(*(*z_stream)(unsafe.Pointer(bp))).zfree = zfree
	(*(*z_stream)(unsafe.Pointer(bp))).opaque = libc.UintptrFromInt32(0)
	(*(*z_stream)(unsafe.Pointer(bp))).next_in = compr
	(*(*z_stream)(unsafe.Pointer(bp))).avail_in = uint32(comprLen)
	err = z.XinflateInit_(tls, bp, ts+339, libc.Int32FromInt64(112))
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+374, err))
		libc.Xexit(tls, int32(1))
	}
	for {
		(*(*z_stream)(unsafe.Pointer(bp))).next_out = uncompr /* discard the output */
		(*(*z_stream)(unsafe.Pointer(bp))).avail_out = uint32(uncomprLen)
		err = z.Xinflate(tls, bp, Z_NO_FLUSH)
		if err == int32(Z_STREAM_END) {
			break
		}
		if err != Z_OK {
			libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+489, err))
			libc.Xexit(tls, int32(1))
		}
	}
	err = z.XinflateEnd(tls, bp)
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+394, err))
		libc.Xexit(tls, int32(1))
	}
	if (*(*z_stream)(unsafe.Pointer(bp))).total_out != uint64(2)*uncomprLen+uncomprLen/uint64(2) {
		libc.Xfprintf(tls, libc.Xstderr, ts+503, libc.VaList(bp+120, (*(*z_stream)(unsafe.Pointer(bp))).total_out))
		libc.Xexit(tls, int32(1))
	} else {
		libc.Xprintf(tls, ts+527, 0)
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Test deflate() with full flush
//	 */
func test_flush(tls *libc.TLS, compr uintptr, comprLen uintptr) {
	bp := tls.Alloc(136) /* tlsAllocs 112 maxVaListSize 16 */
	defer tls.Free(136)
	var err int32
	var len1 uint32
	var _ /* c_stream at bp+0 */ z_stream
	/* compression stream */
	len1 = uint32(libc.Xstrlen(tls, uintptr(unsafe.Pointer(&hello)))) + uint32(1)
	(*(*z_stream)(unsafe.Pointer(bp))).zalloc = zalloc
	(*(*z_stream)(unsafe.Pointer(bp))).zfree = zfree
	(*(*z_stream)(unsafe.Pointer(bp))).opaque = libc.UintptrFromInt32(0)
	err = z.XdeflateInit_(tls, bp, -int32(1), ts+339, libc.Int32FromInt64(112))
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+343, err))
		libc.Xexit(tls, int32(1))
	}
	(*(*z_stream)(unsafe.Pointer(bp))).next_in = uintptr(unsafe.Pointer(&hello))
	(*(*z_stream)(unsafe.Pointer(bp))).next_out = compr
	(*(*z_stream)(unsafe.Pointer(bp))).avail_in = uint32(3)
	(*(*z_stream)(unsafe.Pointer(bp))).avail_out = uint32(*(*uint64)(unsafe.Pointer(comprLen)))
	err = z.Xdeflate(tls, bp, int32(Z_FULL_FLUSH))
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+355, err))
		libc.Xexit(tls, int32(1))
	}
	*(*uint8)(unsafe.Pointer(compr + 3))++ /* force an error in first compressed block */
	(*(*z_stream)(unsafe.Pointer(bp))).avail_in = len1 - uint32(3)
	err = z.Xdeflate(tls, bp, int32(Z_FINISH))
	if err != int32(Z_STREAM_END) {
		if err != Z_OK {
			libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+355, err))
			libc.Xexit(tls, int32(1))
		}
	}
	err = z.XdeflateEnd(tls, bp)
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+363, err))
		libc.Xexit(tls, int32(1))
	}
	*(*uint64)(unsafe.Pointer(comprLen)) = (*(*z_stream)(unsafe.Pointer(bp))).total_out
}

// C documentation
//
//	/* ===========================================================================
//	 * Test inflateSync()
//	 */
func test_sync(tls *libc.TLS, compr uintptr, comprLen uint64, uncompr uintptr, uncomprLen uint64) {
	bp := tls.Alloc(136) /* tlsAllocs 112 maxVaListSize 16 */
	defer tls.Free(136)
	var err int32
	var _ /* d_stream at bp+0 */ z_stream
	/* decompression stream */
	libc.Xstrcpy(tls, uncompr, ts+23)
	(*(*z_stream)(unsafe.Pointer(bp))).zalloc = zalloc
	(*(*z_stream)(unsafe.Pointer(bp))).zfree = zfree
	(*(*z_stream)(unsafe.Pointer(bp))).opaque = libc.UintptrFromInt32(0)
	(*(*z_stream)(unsafe.Pointer(bp))).next_in = compr
	(*(*z_stream)(unsafe.Pointer(bp))).avail_in = uint32(2) /* just read the zlib header */
	err = z.XinflateInit_(tls, bp, ts+339, libc.Int32FromInt64(112))
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+374, err))
		libc.Xexit(tls, int32(1))
	}
	(*(*z_stream)(unsafe.Pointer(bp))).next_out = uncompr
	(*(*z_stream)(unsafe.Pointer(bp))).avail_out = uint32(uncomprLen)
	err = z.Xinflate(tls, bp, Z_NO_FLUSH)
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+386, err))
		libc.Xexit(tls, int32(1))
	}
	(*(*z_stream)(unsafe.Pointer(bp))).avail_in = uint32(comprLen) - uint32(2) /* read all compressed data */
	err = z.XinflateSync(tls, bp)                                              /* but skip the damaged part */
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+548, err))
		libc.Xexit(tls, int32(1))
	}
	err = z.Xinflate(tls, bp, int32(Z_FINISH))
	if err != int32(Z_STREAM_END) {
		libc.Xfprintf(tls, libc.Xstderr, ts+560, 0)
		libc.Xexit(tls, int32(1))
	}
	err = z.XinflateEnd(tls, bp)
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+394, err))
		libc.Xexit(tls, int32(1))
	}
	libc.Xprintf(tls, ts+596, libc.VaList(bp+120, uncompr))
}

// C documentation
//
//	/* ===========================================================================
//	 * Test deflate() with preset dictionary
//	 */
func test_dict_deflate(tls *libc.TLS, compr uintptr, comprLen uint64) {
	bp := tls.Alloc(136) /* tlsAllocs 112 maxVaListSize 16 */
	defer tls.Free(136)
	var err int32
	var _ /* c_stream at bp+0 */ z_stream
	/* compression stream */
	(*(*z_stream)(unsafe.Pointer(bp))).zalloc = zalloc
	(*(*z_stream)(unsafe.Pointer(bp))).zfree = zfree
	(*(*z_stream)(unsafe.Pointer(bp))).opaque = libc.UintptrFromInt32(0)
	err = z.XdeflateInit_(tls, bp, int32(9), ts+339, libc.Int32FromInt64(112))
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+343, err))
		libc.Xexit(tls, int32(1))
	}
	err = z.XdeflateSetDictionary(tls, bp, uintptr(unsafe.Pointer(&dictionary)), uint32(libc.Int32FromInt64(6)))
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+624, err))
		libc.Xexit(tls, int32(1))
	}
	dictId = (*(*z_stream)(unsafe.Pointer(bp))).adler
	(*(*z_stream)(unsafe.Pointer(bp))).next_out = compr
	(*(*z_stream)(unsafe.Pointer(bp))).avail_out = uint32(comprLen)
	(*(*z_stream)(unsafe.Pointer(bp))).next_in = uintptr(unsafe.Pointer(&hello))
	(*(*z_stream)(unsafe.Pointer(bp))).avail_in = uint32(libc.Xstrlen(tls, uintptr(unsafe.Pointer(&hello)))) + uint32(1)
	err = z.Xdeflate(tls, bp, int32(Z_FINISH))
	if err != int32(Z_STREAM_END) {
		libc.Xfprintf(tls, libc.Xstderr, ts+453, 0)
		libc.Xexit(tls, int32(1))
	}
	err = z.XdeflateEnd(tls, bp)
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+363, err))
		libc.Xexit(tls, int32(1))
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Test inflate() with a preset dictionary
//	 */
func test_dict_inflate(tls *libc.TLS, compr uintptr, comprLen uint64, uncompr uintptr, uncomprLen uint64) {
	bp := tls.Alloc(136) /* tlsAllocs 112 maxVaListSize 16 */
	defer tls.Free(136)
	var err int32
	var _ /* d_stream at bp+0 */ z_stream
	/* decompression stream */
	libc.Xstrcpy(tls, uncompr, ts+23)
	(*(*z_stream)(unsafe.Pointer(bp))).zalloc = zalloc
	(*(*z_stream)(unsafe.Pointer(bp))).zfree = zfree
	(*(*z_stream)(unsafe.Pointer(bp))).opaque = libc.UintptrFromInt32(0)
	(*(*z_stream)(unsafe.Pointer(bp))).next_in = compr
	(*(*z_stream)(unsafe.Pointer(bp))).avail_in = uint32(comprLen)
	err = z.XinflateInit_(tls, bp, ts+339, libc.Int32FromInt64(112))
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+374, err))
		libc.Xexit(tls, int32(1))
	}
	(*(*z_stream)(unsafe.Pointer(bp))).next_out = uncompr
	(*(*z_stream)(unsafe.Pointer(bp))).avail_out = uint32(uncomprLen)
	for {
		err = z.Xinflate(tls, bp, Z_NO_FLUSH)
		if err == int32(Z_STREAM_END) {
			break
		}
		if err == int32(Z_NEED_DICT) {
			if (*(*z_stream)(unsafe.Pointer(bp))).adler != dictId {
				libc.Xfprintf(tls, libc.Xstderr, ts+645, 0)
				libc.Xexit(tls, int32(1))
			}
			err = z.XinflateSetDictionary(tls, bp, uintptr(unsafe.Pointer(&dictionary)), uint32(libc.Int32FromInt64(6)))
		}
		if err != Z_OK {
			libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+667, err))
			libc.Xexit(tls, int32(1))
		}
	}
	err = z.XinflateEnd(tls, bp)
	if err != Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+120, ts+394, err))
		libc.Xexit(tls, int32(1))
	}
	if libc.Xstrcmp(tls, uncompr, uintptr(unsafe.Pointer(&hello))) != 0 {
		libc.Xfprintf(tls, libc.Xstderr, ts+685, 0)
		libc.Xexit(tls, int32(1))
	} else {
		libc.Xprintf(tls, ts+708, libc.VaList(bp+120, uncompr))
	}
}

func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(40) /* tlsAllocs 8 maxVaListSize 24 */
	defer tls.Free(40)
	var compr, uncompr, v1 uintptr
	var uncomprLen uint64
	var _ /* comprLen at bp+0 */ uint64
	uncomprLen = uint64(20000)
	*(*uint64)(unsafe.Pointer(bp)) = uint64(3) * uncomprLen
	if int32(*(*int8)(unsafe.Pointer(z.XzlibVersion(tls)))) != int32(*(*int8)(unsafe.Pointer(myVersion))) {
		libc.Xfprintf(tls, libc.Xstderr, ts+737, 0)
		libc.Xexit(tls, int32(1))
	} else if libc.Xstrcmp(tls, z.XzlibVersion(tls), ts+339) != 0 {
		libc.Xfprintf(tls, libc.Xstderr, ts+764, libc.VaList(bp+16, z.XzlibVersion(tls)))
	}
	libc.Xprintf(tls, ts+808, libc.VaList(bp+16, ts+339, int32(ZLIB_VERNUM), z.XzlibCompileFlags(tls)))
	compr = libc.Xcalloc(tls, uint64(uint32(*(*uint64)(unsafe.Pointer(bp)))), uint64(1))
	uncompr = libc.Xcalloc(tls, uint64(uint32(uncomprLen)), uint64(1))
	/* compr and uncompr are cleared to avoid reading uninitialized
	 * data and to ensure that uncompr compresses well.
	 */
	if compr == uintptr(Z_NULL) || uncompr == uintptr(Z_NULL) {
		libc.Xprintf(tls, ts+857, 0)
		libc.Xexit(tls, int32(1))
	}
	test_compress(tls, compr, *(*uint64)(unsafe.Pointer(bp)), uncompr, uncomprLen)
	if argc > int32(1) {
		v1 = *(*uintptr)(unsafe.Pointer(argv + 1*8))
	} else {
		v1 = ts + 872
	}
	test_gzio(tls, v1, uncompr, uncomprLen)
	test_deflate(tls, compr, *(*uint64)(unsafe.Pointer(bp)))
	test_inflate(tls, compr, *(*uint64)(unsafe.Pointer(bp)), uncompr, uncomprLen)
	test_large_deflate(tls, compr, *(*uint64)(unsafe.Pointer(bp)), uncompr, uncomprLen)
	test_large_inflate(tls, compr, *(*uint64)(unsafe.Pointer(bp)), uncompr, uncomprLen)
	test_flush(tls, compr, bp)
	test_sync(tls, compr, *(*uint64)(unsafe.Pointer(bp)), uncompr, uncomprLen)
	*(*uint64)(unsafe.Pointer(bp)) = uint64(3) * uncomprLen
	test_dict_deflate(tls, compr, *(*uint64)(unsafe.Pointer(bp)))
	test_dict_inflate(tls, compr, *(*uint64)(unsafe.Pointer(bp)), uncompr, uncomprLen)
	libc.Xfree(tls, compr)
	libc.Xfree(tls, uncompr)
	return 0
}

var myVersion = ts + 339

func main() {
	libc.Start(main1)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var ts = (*reflect.StringHeader)(unsafe.Pointer(&ts1)).Data

var ts1 = "%s error: %d\n\x00compress\x00garbage\x00uncompress\x00bad uncompress\n\x00uncompress(): %s\n\x00wb\x00gzopen error\n\x00ello\x00gzputs err: %s\n\x00, %s!\x00hello\x00gzprintf err: %s\n\x00rb\x00gzread err: %s\n\x00bad gzread: %s\n\x00gzread(): %s\n\x00gzseek error, pos=%ld, gztell=%ld\n\x00gzgetc error\n\x00gzungetc error\n\x00gzgets err after gzseek: %s\n\x00bad gzgets after gzseek\n\x00gzgets() after gzseek: %s\n\x001.3\x00deflateInit\x00deflate\x00deflateEnd\x00inflateInit\x00inflate\x00inflateEnd\x00bad inflate\n\x00inflate(): %s\n\x00deflate not greedy\n\x00deflate should report Z_STREAM_END\n\x00large inflate\x00bad large inflate: %ld\n\x00large_inflate(): OK\n\x00inflateSync\x00inflate should report Z_STREAM_END\n\x00after inflateSync(): hel%s\n\x00deflateSetDictionary\x00unexpected dictionary\x00inflate with dict\x00bad inflate with dict\n\x00inflate with dictionary: %s\n\x00incompatible zlib version\n\x00warning: different zlib version linked: %s\n\x00zlib version %s = 0x%04x, compile flags = 0x%lx\n\x00out of memory\n\x00foo.gz\x00"
