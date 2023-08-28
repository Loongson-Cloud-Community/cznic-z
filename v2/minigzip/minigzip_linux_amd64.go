// Code generated for linux/amd64 by 'gcc -o minigzip/minigzip_linux_amd64.go --package-name=main minigzip_linux_amd64.c.go -lmodernc.org/z/v2', DO NOT EDIT.

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

const BUFLEN = 16384
const FD_SETSIZE = 1024
const GZ_SUFFIX = ".gz"
const MAX_NAME_LEN = 1024
const Z_OK = 0

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

var prog uintptr

// C documentation
//
//	/* ===========================================================================
//	 * Display error message and exit
//	 */
func error1(tls *libc.TLS, msg uintptr) {
	bp := tls.Alloc(24) /* tlsAllocs 0 maxVaListSize 16 */
	defer tls.Free(24)
	libc.Xfprintf(tls, libc.Xstderr, ts, libc.VaList(bp+8, prog, msg))
	libc.Xexit(tls, int32(1))
}

func gz_compress(tls *libc.TLS, in uintptr, out uintptr) {
	bp := tls.Alloc(16392) /* tlsAllocs 16392 maxVaListSize 0 */
	defer tls.Free(16392)
	var len1 int32
	var _ /* buf at bp+0 */ [16384]int8
	var _ /* err at bp+16384 */ int32
	for {
		len1 = int32(libc.Xfread(tls, bp, uint64(1), uint64(16384), in))
		if libc.Xferror(tls, in) != 0 {
			libc.Xperror(tls, ts+8)
			libc.Xexit(tls, int32(1))
		}
		if len1 == 0 {
			break
		}
		if z.Xgzwrite(tls, out, bp, uint32(len1)) != len1 {
			error1(tls, z.Xgzerror(tls, out, bp+16384))
		}
	}
	libc.Xfclose(tls, in)
	if z.Xgzclose(tls, out) != Z_OK {
		error1(tls, ts+14)
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Uncompress input to output then close both files.
//	 */
func gz_uncompress(tls *libc.TLS, in uintptr, out uintptr) {
	bp := tls.Alloc(16392) /* tlsAllocs 16392 maxVaListSize 0 */
	defer tls.Free(16392)
	var len1 int32
	var _ /* buf at bp+0 */ [16384]int8
	var _ /* err at bp+16384 */ int32
	for {
		len1 = z.Xgzread(tls, in, bp, uint32(16384))
		if len1 < 0 {
			error1(tls, z.Xgzerror(tls, in, bp+16384))
		}
		if len1 == 0 {
			break
		}
		if int32(libc.Xfwrite(tls, bp, uint64(1), uint64(uint32(len1)), out)) != len1 {
			error1(tls, ts+29)
		}
	}
	if libc.Xfclose(tls, out) != 0 {
		error1(tls, ts+43)
	}
	if z.Xgzclose(tls, in) != Z_OK {
		error1(tls, ts+14)
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Compress the given file: create a corresponding .gz file and remove the
//	 * original.
//	 */
func file_compress(tls *libc.TLS, file uintptr, mode uintptr) {
	bp := tls.Alloc(1048) /* tlsAllocs 1024 maxVaListSize 16 */
	defer tls.Free(1048)
	var in, out uintptr
	var _ /* outfile at bp+0 */ [1024]int8
	if libc.Xstrlen(tls, file)+libc.Xstrlen(tls, ts+57) >= uint64(1024) {
		libc.Xfprintf(tls, libc.Xstderr, ts+61, libc.VaList(bp+1032, prog))
		libc.Xexit(tls, int32(1))
	}
	libc.X__builtin_snprintf(tls, bp, uint64(1024), ts+84, libc.VaList(bp+1032, file, ts+57))
	in = libc.Xfopen(tls, file, ts+89)
	if in == libc.UintptrFromInt32(0) {
		libc.Xperror(tls, file)
		libc.Xexit(tls, int32(1))
	}
	out = z.Xgzopen(tls, bp, mode)
	if out == libc.UintptrFromInt32(0) {
		libc.Xfprintf(tls, libc.Xstderr, ts+92, libc.VaList(bp+1032, prog, bp))
		libc.Xexit(tls, int32(1))
	}
	gz_compress(tls, in, out)
	libc.Xunlink(tls, file)
}

// C documentation
//
//	/* ===========================================================================
//	 * Uncompress the given file and remove the original.
//	 */
func file_uncompress(tls *libc.TLS, file uintptr) {
	bp := tls.Alloc(1048) /* tlsAllocs 1024 maxVaListSize 16 */
	defer tls.Free(1048)
	var in, infile, out, outfile uintptr
	var len1 uint64
	var _ /* buf at bp+0 */ [1024]int8
	len1 = libc.Xstrlen(tls, file)
	if len1+libc.Xstrlen(tls, ts+57) >= uint64(1024) {
		libc.Xfprintf(tls, libc.Xstderr, ts+61, libc.VaList(bp+1032, prog))
		libc.Xexit(tls, int32(1))
	}
	libc.X__builtin_snprintf(tls, bp, uint64(1024), ts+113, libc.VaList(bp+1032, file))
	if len1 > libc.Uint64FromInt64(4)-libc.Uint64FromInt32(1) && libc.Xstrcmp(tls, file+uintptr(len1)-uintptr(libc.Uint64FromInt64(4)-libc.Uint64FromInt32(1)), ts+57) == 0 {
		infile = file
		outfile = bp
		*(*int8)(unsafe.Pointer(outfile + uintptr(len1-uint64(3)))) = int8('\000')
	} else {
		outfile = file
		infile = bp
		libc.X__builtin_snprintf(tls, bp+uintptr(len1), uint64(1024)-len1, ts+113, libc.VaList(bp+1032, ts+57))
	}
	in = z.Xgzopen(tls, infile, ts+89)
	if in == libc.UintptrFromInt32(0) {
		libc.Xfprintf(tls, libc.Xstderr, ts+92, libc.VaList(bp+1032, prog, infile))
		libc.Xexit(tls, int32(1))
	}
	out = libc.Xfopen(tls, outfile, ts+116)
	if out == libc.UintptrFromInt32(0) {
		libc.Xperror(tls, file)
		libc.Xexit(tls, int32(1))
	}
	gz_uncompress(tls, in, out)
	libc.Xunlink(tls, infile)
}

func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(48) /* tlsAllocs 24 maxVaListSize 16 */
	defer tls.Free(48)
	var bname, file, in uintptr
	var copyout, uncompr, v1, v2 int32
	var _ /* outmode at bp+0 */ [20]int8
	copyout = 0
	uncompr = 0
	libc.X__builtin_snprintf(tls, bp, uint64(20), ts+113, libc.VaList(bp+32, ts+119))
	prog = *(*uintptr)(unsafe.Pointer(argv))
	bname = libc.Xstrrchr(tls, *(*uintptr)(unsafe.Pointer(argv)), int32('/'))
	if bname != 0 {
		bname++
	} else {
		bname = *(*uintptr)(unsafe.Pointer(argv))
	}
	argc--
	argv += 8
	if !(libc.Xstrcmp(tls, bname, ts+124) != 0) {
		uncompr = int32(1)
	} else if !(libc.Xstrcmp(tls, bname, ts+131) != 0) {
		v1 = libc.Int32FromInt32(1)
		uncompr = v1
		copyout = v1
	}
	for argc > 0 {
		if libc.Xstrcmp(tls, *(*uintptr)(unsafe.Pointer(argv)), ts+136) == 0 {
			copyout = int32(1)
		} else {
			if libc.Xstrcmp(tls, *(*uintptr)(unsafe.Pointer(argv)), ts+139) == 0 {
				uncompr = int32(1)
			} else {
				if libc.Xstrcmp(tls, *(*uintptr)(unsafe.Pointer(argv)), ts+142) == 0 {
					(*(*[20]int8)(unsafe.Pointer(bp)))[int32(3)] = int8('f')
				} else {
					if libc.Xstrcmp(tls, *(*uintptr)(unsafe.Pointer(argv)), ts+145) == 0 {
						(*(*[20]int8)(unsafe.Pointer(bp)))[int32(3)] = int8('h')
					} else {
						if libc.Xstrcmp(tls, *(*uintptr)(unsafe.Pointer(argv)), ts+148) == 0 {
							(*(*[20]int8)(unsafe.Pointer(bp)))[int32(3)] = int8('R')
						} else {
							if int32(*(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(argv))))) == int32('-') && int32(*(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(argv)) + 1))) >= int32('1') && int32(*(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(argv)) + 1))) <= int32('9') && int32(*(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(argv)) + 2))) == 0 {
								(*(*[20]int8)(unsafe.Pointer(bp)))[int32(2)] = *(*int8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(argv)) + 1))
							} else {
								break
							}
						}
					}
				}
			}
		}
		argc--
		argv += 8
	}
	if int32((*(*[20]int8)(unsafe.Pointer(bp)))[int32(3)]) == int32(' ') {
		(*(*[20]int8)(unsafe.Pointer(bp)))[int32(3)] = 0
	}
	if argc == 0 {
		if uncompr != 0 {
			file = z.Xgzdopen(tls, libc.Xfileno(tls, libc.Xstdin), ts+89)
			if file == libc.UintptrFromInt32(0) {
				error1(tls, ts+151)
			}
			gz_uncompress(tls, file, libc.Xstdout)
		} else {
			file = z.Xgzdopen(tls, libc.Xfileno(tls, libc.Xstdout), bp)
			if file == libc.UintptrFromInt32(0) {
				error1(tls, ts+171)
			}
			gz_compress(tls, libc.Xstdin, file)
		}
	} else {
		if copyout != 0 {
		}
		for first := true; ; first = false {
			if !first {
				argv += 8
				argc--
				v2 = argc
				if !(v2 != 0) {
					break
				}
			}
			if uncompr != 0 {
				if copyout != 0 {
					file = z.Xgzopen(tls, *(*uintptr)(unsafe.Pointer(argv)), ts+89)
					if file == libc.UintptrFromInt32(0) {
						libc.Xfprintf(tls, libc.Xstderr, ts+92, libc.VaList(bp+32, prog, *(*uintptr)(unsafe.Pointer(argv))))
					} else {
						gz_uncompress(tls, file, libc.Xstdout)
					}
				} else {
					file_uncompress(tls, *(*uintptr)(unsafe.Pointer(argv)))
				}
			} else {
				if copyout != 0 {
					in = libc.Xfopen(tls, *(*uintptr)(unsafe.Pointer(argv)), ts+89)
					if in == libc.UintptrFromInt32(0) {
						libc.Xperror(tls, *(*uintptr)(unsafe.Pointer(argv)))
					} else {
						file = z.Xgzdopen(tls, libc.Xfileno(tls, libc.Xstdout), bp)
						if file == libc.UintptrFromInt32(0) {
							error1(tls, ts+171)
						}
						gz_compress(tls, in, file)
					}
				} else {
					file_compress(tls, *(*uintptr)(unsafe.Pointer(argv)), bp)
				}
			}
		}
	}
	return 0
}

func main() {
	libc.Start(main1)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var ts = (*reflect.StringHeader)(unsafe.Pointer(&ts1)).Data

var ts1 = "%s: %s\n\x00fread\x00failed gzclose\x00failed fwrite\x00failed fclose\x00.gz\x00%s: filename too long\n\x00%s%s\x00rb\x00%s: can't gzopen %s\n\x00%s\x00wb\x00wb6 \x00gunzip\x00zcat\x00-c\x00-d\x00-f\x00-h\x00-r\x00can't gzdopen stdin\x00can't gzdopen stdout\x00"
