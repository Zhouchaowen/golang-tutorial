package main

import (
	"fmt"
	"unsafe"
)

/*
	1.map深浅拷贝
*/

// Steps4 浅拷贝
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
	fmt.Printf("\ttmpIntString:%+v len:%d\n",
		tmpIntString,
		len(tmpIntString))
	fmt.Printf("\ttmpIntString       addr:%p\n", &tmpIntString)
	fmt.Printf("\ttmpIntString value addr:%p\n", tmpIntString)

	tmpIntString = mpIntString // 将指向底层映射的指针赋值给了 tmpIntString

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

// Steps5 深拷贝
func Steps5() {
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
	fmt.Printf("\ttmpIntString:%+v len:%d\n",
		tmpIntString,
		len(tmpIntString))
	fmt.Printf("\ttmpIntString       addr:%p\n", &tmpIntString)
	fmt.Printf("\ttmpIntString value addr:%p\n", tmpIntString)

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

// Steps6 证明 map 的底层结构
func Steps6() {
	mpIntString := map[int]string{
		1: "golang",
		2: "tutorial",
		3: "World",
	}
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))
	fmt.Printf("\tmpIntString       size:%d\n", unsafe.Sizeof(mpIntString))
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

	fmt.Printf("\t-------------------------------------------\n")
	// 等同于上面的结果，参考 go/src/runtime/export_test.go 548 行
	h := *(**hmap)(unsafe.Pointer(&mpIntString))
	fmt.Printf("\thmap %+v\n", h)
	fmt.Printf("\thmap %p\n", h)
	fmt.Printf("\thmap     count %p\n", &h.count)
	fmt.Printf("\thmap     flags %p\n", &h.flags)
	fmt.Printf("\thmap         B %p\n", &h.B)
	fmt.Printf("\thmap noverflow %p\n", &h.noverflow)
	fmt.Printf("\thmap     hash0 %p\n", &h.hash0)
	fmt.Printf("\thmap   buckets %p\n", &h.buckets)

	fmt.Printf("\t-------------------------------------------\n")
	tmpMpIntString := mpIntString
	tmpMpIntString[4] = "World"
	fmt.Printf("\ttmpMpIntString:%+v len:%d\n",
		tmpMpIntString,
		len(tmpMpIntString))
	fmt.Printf("\ttmpMpIntString       addr:%p\n", &tmpMpIntString)
	fmt.Printf("\ttmpMpIntString value addr:%p\n", tmpMpIntString)

	tmp := *(**hmap)(unsafe.Pointer(&tmpMpIntString))
	fmt.Printf("\thmap %+v\n", tmp)
	fmt.Printf("\thmap %p\n", tmp)
	fmt.Printf("\thmap     count %p\n", &tmp.count)
	fmt.Printf("\thmap     flags %p\n", &tmp.flags)
	fmt.Printf("\thmap         B %p\n", &tmp.B)
	fmt.Printf("\thmap noverflow %p\n", &tmp.noverflow)
	fmt.Printf("\thmap     hash0 %p\n", &tmp.hash0)
	fmt.Printf("\thmap   buckets %p\n", &tmp.buckets)
}

func main() {
	fmt.Println("Steps4():")
	Steps4()
	fmt.Println("Steps5():")
	Steps5()
	fmt.Println("Steps6():")
	Steps6()
}

type bmap struct {
	tophash [8]uint8
}
type mapextra struct {
	overflow     *[]*bmap
	oldoverflow  *[]*bmap
	nextOverflow *bmap
}
type hmap struct {
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
