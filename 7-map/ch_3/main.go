package main

import (
	"fmt"
	"unsafe"
)

/*
	1.map深浅拷贝
*/

// Steps3 浅拷贝
func Steps3() {
	mpIntString := map[int]string{
		1: "golang",
		2: "tutorial",
	}
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))
	fmt.Printf("\tmpIntString       addr:%p\n", &mpIntString)
	fmt.Printf("\tmpIntString value addr:%p\n", mpIntString)
	fmt.Println("\t-------------------------")
	tmpIntString := make(map[int]string, 2)
	tmpIntString = mpIntString
	fmt.Printf("\ttmpIntString:%+v len:%d\n",
		tmpIntString,
		len(tmpIntString))
	fmt.Printf("\ttmpIntString       addr:%p\n", &tmpIntString)
	fmt.Printf("\ttmpIntString value addr:%p\n", tmpIntString)

	tmpIntString[2] = "IMianBa"
	fmt.Println("\t-------------------------")
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))
	fmt.Printf("\ttmpIntString:%+v len:%d\n",
		tmpIntString,
		len(tmpIntString))
}

// Steps4 深拷贝
func Steps4() {
	mpIntString := map[int]string{
		1: "golang",
		2: "tutorial",
	}
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))
	fmt.Printf("\tmpIntString       addr:%p\n", &mpIntString)
	fmt.Printf("\tmpIntString value addr:%p\n", mpIntString)
	fmt.Println("\t-------------------------")

	tmpIntString := make(map[int]string, 2)
	for k, v := range mpIntString {
		tmpIntString[k] = v
	}

	fmt.Printf("\ttmpIntString:%+v len:%d\n",
		tmpIntString,
		len(tmpIntString))
	fmt.Printf("\ttmpIntString       addr:%p\n", &tmpIntString)
	fmt.Printf("\ttmpIntString value addr:%p\n", tmpIntString)

	tmpIntString[2] = "IMianBa"
	fmt.Println("\t-------------------------")
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))
	fmt.Printf("\ttmpIntString:%+v len:%d\n",
		tmpIntString,
		len(tmpIntString))
}

// Steps5 证明 map 的底层结构
func Steps5() {
	mpIntString := map[int]string{
		1: "golang",
		2: "tutorial",
		3: "World",
	}
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))
	fmt.Printf("\tmpIntString       addr:%p\n", &mpIntString)
	fmt.Printf("\tmpIntString value addr:%p\n", mpIntString)

	/*
		// A header for a Go map.
		type hmap struct {
			// Note: the format of the hmap is also encoded in cmd/compile/internal/reflectdata/reflect.go.
			// Make sure this stays in sync with the compiler's definition.
			count     int // # live cells == size of map.  Must be first (used by len() builtin)
			flags     uint8
			B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
			noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
			hash0     uint32 // hash seed

			buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
			oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
			nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

			extra *mapextra // optional fields
		}
	*/

	fmt.Printf("\tmpIntString value addr:0x%x\n", *(*uintptr)(unsafe.Pointer(&mpIntString)))
	fmt.Printf("\tmpIntString value struct.filed_1      count %d\n", *(*int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&mpIntString)))))
	fmt.Printf("\tmpIntString value struct.filed_2      flags %d\n", *(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&mpIntString)) + uintptr(8))))
	fmt.Printf("\tmpIntString value struct.filed_3          B %d\n", *(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&mpIntString)) + uintptr(9))))
	fmt.Printf("\tmpIntString value struct.filed_4  noverflow %d\n", *(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&mpIntString)) + uintptr(10))))
	fmt.Printf("\tmpIntString value struct.filed_5      hash0 %d\n", *(*uint32)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&mpIntString)) + uintptr(12))))
	fmt.Printf("\tmpIntString value struct.filed_6    buckets 0x%x\n", *(*uintptr)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&mpIntString)) + uintptr(16))))
	fmt.Printf("\tmpIntString value struct.filed_7 oldbuckets 0x%x\n", *(*uintptr)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&mpIntString)) + uintptr(32))))

	// 等同于上面的结果，参考 go/src/runtime/export_test.go 548 行
	mapToBMap(mpIntString)
}

func mapToBMap(mpIntString map[int]string) {
	type bmap struct {
		// tophash generally contains the top byte of the hash value
		// for each key in this bucket. If tophash[0] < minTopHash,
		// tophash[0] is a bucket evacuation state instead.
		tophash [8]uint8
		// Followed by bucketCnt keys and then bucketCnt elems.
		// NOTE: packing all the keys together and then all the elems together makes the
		// code a bit more complicated than alternating key/elem/key/elem/... but it allows
		// us to eliminate padding which would be needed for, e.g., map[int64]int8.
		// Followed by an overflow pointer.
	}
	type mapextra struct {
		// If both key and elem do not contain pointers and are inline, then we mark bucket
		// type as containing no pointers. This avoids scanning such maps.
		// However, bmap.overflow is a pointer. In order to keep overflow buckets
		// alive, we store pointers to all overflow buckets in hmap.extra.overflow and hmap.extra.oldoverflow.
		// overflow and oldoverflow are only used if key and elem do not contain pointers.
		// overflow contains overflow buckets for hmap.buckets.
		// oldoverflow contains overflow buckets for hmap.oldbuckets.
		// The indirection allows to store a pointer to the slice in hiter.
		overflow    *[]*bmap
		oldoverflow *[]*bmap

		// nextOverflow holds a pointer to a free overflow bucket.
		nextOverflow *bmap
	}
	type hmap struct {
		// Note: the format of the hmap is also encoded in cmd/compile/internal/reflectdata/reflect.go.
		// Make sure this stays in sync with the compiler's definition.
		count     int // # live cells == size of map.  Must be first (used by len() builtin)
		flags     uint8
		B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
		noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
		hash0     uint32 // hash seed

		buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
		oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
		nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

		extra *mapextra // optional fields
	}

	h := *(**hmap)(unsafe.Pointer(&mpIntString))
	fmt.Printf("\thmap %+v\n", h)
}

func main() {
	fmt.Println("Steps3():")
	Steps3()
	fmt.Println("Steps4():")
	Steps4()
	fmt.Println("Steps5():")
	Steps5()
}
