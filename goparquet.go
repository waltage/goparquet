/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.1
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./swig/goparquet.swigcxx

package goparquet

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef _gostring_ swig_type_1;
typedef _gostring_ swig_type_2;
typedef _gostring_ swig_type_3;
typedef _gostring_ swig_type_4;
typedef _gostring_ swig_type_5;
typedef long long swig_type_6;
typedef _gostring_ swig_type_7;
typedef _gostring_ swig_type_8;
typedef _gostring_ swig_type_9;
typedef _gostring_ swig_type_10;
typedef _gostring_ swig_type_11;
typedef _gostring_ swig_type_12;
typedef _gostring_ swig_type_13;
typedef _gostring_ swig_type_14;
typedef _gostring_ swig_type_15;
typedef long long swig_type_16;
extern void _wrap_Swig_free_goparquet_498a767b677f9faf(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_goparquet_498a767b677f9faf(swig_intgo arg1);
extern uintptr_t _wrap_pq_column_goparquet_498a767b677f9faf(swig_type_1 arg1);
extern uintptr_t _wrap_pq_encoding_goparquet_498a767b677f9faf(swig_type_2 arg1);
extern uintptr_t _wrap_pq_compression_goparquet_498a767b677f9faf(swig_type_3 arg1);
extern uintptr_t _wrap_arrow_file_to_table_goparquet_498a767b677f9faf(swig_type_4 arg1);
extern void _wrap_arrow_table_to_parquet_goparquet_498a767b677f9faf(uintptr_t arg1, swig_type_5 arg2, swig_type_6 arg3);
extern void _wrap_SetCompression_goparquet_498a767b677f9faf(swig_type_7 arg1);
extern void _wrap_SetColumnCompression_goparquet_498a767b677f9faf(swig_type_8 arg1, swig_type_9 arg2);
extern void _wrap_SetColumnEncoding_goparquet_498a767b677f9faf(swig_type_10 arg1, swig_type_11 arg2);
extern void _wrap_EnableColumnDictionary_goparquet_498a767b677f9faf(swig_type_12 arg1);
extern void _wrap_DisableColumnDictionary_goparquet_498a767b677f9faf(swig_type_13 arg1);
extern void _wrap_ConvertArrowFileToParquet_goparquet_498a767b677f9faf(swig_type_14 arg1, swig_type_15 arg2, swig_type_16 arg3);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_goparquet_498a767b677f9faf(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_goparquet_498a767b677f9faf(C.swig_intgo(_swig_i_0)))
	return swig_r
}

func Pq_column(arg1 string) (_swig_ret Shared_ptr_Sl_parquet_schema_ColumnPath_Sg_) {
	var swig_r Shared_ptr_Sl_parquet_schema_ColumnPath_Sg_
	_swig_i_0 := arg1
	swig_r = (Shared_ptr_Sl_parquet_schema_ColumnPath_Sg_)(SwigcptrShared_ptr_Sl_parquet_schema_ColumnPath_Sg_(C._wrap_pq_column_goparquet_498a767b677f9faf(*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_0)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func Pq_encoding(arg1 string) (_swig_ret Parquet_Encoding_type) {
	var swig_r Parquet_Encoding_type
	_swig_i_0 := arg1
	swig_r = (Parquet_Encoding_type)(SwigcptrParquet_Encoding_type(C._wrap_pq_encoding_goparquet_498a767b677f9faf(*(*C.swig_type_2)(unsafe.Pointer(&_swig_i_0)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func Pq_compression(arg1 string) (_swig_ret Parquet_Compression_type) {
	var swig_r Parquet_Compression_type
	_swig_i_0 := arg1
	swig_r = (Parquet_Compression_type)(SwigcptrParquet_Compression_type(C._wrap_pq_compression_goparquet_498a767b677f9faf(*(*C.swig_type_3)(unsafe.Pointer(&_swig_i_0)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func Arrow_file_to_table(arg1 string) (_swig_ret Shared_ptr_Sl_arrow_Table_Sg_) {
	var swig_r Shared_ptr_Sl_arrow_Table_Sg_
	_swig_i_0 := arg1
	swig_r = (Shared_ptr_Sl_arrow_Table_Sg_)(SwigcptrShared_ptr_Sl_arrow_Table_Sg_(C._wrap_arrow_file_to_table_goparquet_498a767b677f9faf(*(*C.swig_type_4)(unsafe.Pointer(&_swig_i_0)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func Arrow_table_to_parquet(arg1 Shared_ptr_Sl_arrow_Table_Sg_, arg2 string, arg3 int64) {
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_arrow_table_to_parquet_goparquet_498a767b677f9faf(C.uintptr_t(_swig_i_0), *(*C.swig_type_5)(unsafe.Pointer(&_swig_i_1)), C.swig_type_6(_swig_i_2))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func SetCompression(arg1 string) {
	_swig_i_0 := arg1
	C._wrap_SetCompression_goparquet_498a767b677f9faf(*(*C.swig_type_7)(unsafe.Pointer(&_swig_i_0)))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
}

func SetColumnCompression(arg1 string, arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_SetColumnCompression_goparquet_498a767b677f9faf(*(*C.swig_type_8)(unsafe.Pointer(&_swig_i_0)), *(*C.swig_type_9)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func SetColumnEncoding(arg1 string, arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_SetColumnEncoding_goparquet_498a767b677f9faf(*(*C.swig_type_10)(unsafe.Pointer(&_swig_i_0)), *(*C.swig_type_11)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func EnableColumnDictionary(arg1 string) {
	_swig_i_0 := arg1
	C._wrap_EnableColumnDictionary_goparquet_498a767b677f9faf(*(*C.swig_type_12)(unsafe.Pointer(&_swig_i_0)))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
}

func DisableColumnDictionary(arg1 string) {
	_swig_i_0 := arg1
	C._wrap_DisableColumnDictionary_goparquet_498a767b677f9faf(*(*C.swig_type_13)(unsafe.Pointer(&_swig_i_0)))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
}

func ConvertArrowFileToParquet(arg1 string, arg2 string, arg3 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_ConvertArrowFileToParquet_goparquet_498a767b677f9faf(*(*C.swig_type_14)(unsafe.Pointer(&_swig_i_0)), *(*C.swig_type_15)(unsafe.Pointer(&_swig_i_1)), C.swig_type_16(_swig_i_2))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}


type SwigcptrShared_ptr_Sl_arrow_Table_Sg_ uintptr
type Shared_ptr_Sl_arrow_Table_Sg_ interface {
	Swigcptr() uintptr;
}
func (p SwigcptrShared_ptr_Sl_arrow_Table_Sg_) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrParquet_Encoding_type uintptr
type Parquet_Encoding_type interface {
	Swigcptr() uintptr;
}
func (p SwigcptrParquet_Encoding_type) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrParquet_Compression_type uintptr
type Parquet_Compression_type interface {
	Swigcptr() uintptr;
}
func (p SwigcptrParquet_Compression_type) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrShared_ptr_Sl_parquet_schema_ColumnPath_Sg_ uintptr
type Shared_ptr_Sl_parquet_schema_ColumnPath_Sg_ interface {
	Swigcptr() uintptr;
}
func (p SwigcptrShared_ptr_Sl_parquet_schema_ColumnPath_Sg_) Swigcptr() uintptr {
	return uintptr(p)
}



var swigDirectorTrack struct {
	sync.Mutex
	m map[int]interface{}
	c int
}

func swigDirectorAdd(v interface{}) int {
	swigDirectorTrack.Lock()
	defer swigDirectorTrack.Unlock()
	if swigDirectorTrack.m == nil {
		swigDirectorTrack.m = make(map[int]interface{})
	}
	swigDirectorTrack.c++
	ret := swigDirectorTrack.c
	swigDirectorTrack.m[ret] = v
	return ret
}

func swigDirectorLookup(c int) interface{} {
	swigDirectorTrack.Lock()
	defer swigDirectorTrack.Unlock()
	ret := swigDirectorTrack.m[c]
	if ret == nil {
		panic("C++ director pointer not found (possible	use-after-free)")
	}
	return ret
}

func swigDirectorDelete(c int) {
	swigDirectorTrack.Lock()
	defer swigDirectorTrack.Unlock()
	if swigDirectorTrack.m[c] == nil {
		if c > swigDirectorTrack.c {
			panic("C++ director pointer invalid (possible memory corruption")
		} else {
			panic("C++ director pointer not found (possible use-after-free)")
		}
	}
	delete(swigDirectorTrack.m, c)
}


