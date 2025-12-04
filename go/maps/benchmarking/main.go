package main

import (
	"encoding/binary"
	"fmt"

	"time"
)

func main() {

	n := uint64(3000)

	now := time.Now()
	m := mapMapMake(n)
	fmt.Printf("mapmap took %f to make\n", time.Since(now).Seconds())

	now = time.Now()
	mapMapRead(m, n)
	fmt.Printf("mapmap took %f to make\n", time.Since(now).Seconds())

	now = time.Now()
	mm := mapStrMake(n)
	fmt.Printf("mapStr took %f to make\n", time.Since(now).Seconds())

	now = time.Now()
	mapStrRead(mm, n)
	fmt.Printf("mapStr took %f to read\n", time.Since(now).Seconds())

	now = time.Now()
	mm = mapStr2Make(n)
	fmt.Printf("mapStr2 took %f to make\n", time.Since(now).Seconds())

	now = time.Now()
	mapStr2Read(mm, n)
	fmt.Printf("mapStr2 took %f to read\n", time.Since(now).Seconds())

	now = time.Now()
	mb := mapBytesMake(n)
	fmt.Printf("mapBytes took %f to make\n", time.Since(now).Seconds())

	now = time.Now()
	mapBytesRead(mb, n)
	fmt.Printf("mapBytes took %f to read\n", time.Since(now).Seconds())
}

type ObjectCache struct {
	BodyMetadataIndex uint64
	Body              []byte
}

func mapBytesMake(sz uint64) map[[16]byte]ObjectCache {
	m := make(map[[16]byte]ObjectCache)
	var key [16]byte
	var o ObjectCache
	for i := range sz {
		binary.LittleEndian.PutUint64(key[:8], i)
		for j := range sz {
			binary.LittleEndian.PutUint64(key[8:], j)

			o.BodyMetadataIndex = i*sz + j
			m[key] = o
		}
	}
	return m
}

func mapBytesRead(m map[[16]byte]ObjectCache, sz uint64) {
	var ok bool
	var key [16]byte
	for i := range sz {
		binary.LittleEndian.PutUint64(key[:8], i)
		for j := range sz {
			binary.LittleEndian.PutUint64(key[8:], j)
			_, ok = m[key]
			if !ok {
				panic(fmt.Errorf("missing entry i, j = %d, %d", i, j))
			}
		}
	}
}

func mapMapMake(sz uint64) map[uint64]map[uint64]ObjectCache {
	m := make(map[uint64]map[uint64]ObjectCache)
	var n map[uint64]ObjectCache
	var ok bool
	var o ObjectCache
	for i := range sz {
		for j := range sz {
			o.BodyMetadataIndex = i*sz + j
			n, ok = m[i]
			if !ok {
				m[i] = map[uint64]ObjectCache{
					j: o,
				}
			} else {
				n[j] = o
			}
		}
		if len(n) != int(sz) {
			panic(fmt.Errorf("n has only len %d instead of %d", len(n), sz))
		}
	}
	return m
}

func mapMapRead(m map[uint64]map[uint64]ObjectCache, sz uint64) {
	var n map[uint64]ObjectCache
	var ok bool
	for i := range sz {
		for j := range sz {
			n, ok = m[i]
			if ok {
				_, ok = n[j]
				if !ok {
					panic(fmt.Errorf("missing entry i, j, %d, %d", i, j))
				}
			} else {
				panic(fmt.Errorf("missing entry i = %d", i))
			}
		}
		if len(n) != int(sz) {
			panic(fmt.Errorf("n has only len %d instead of %d", len(n), sz))
		}
	}
}

func mapStrMake(sz uint64) map[string]ObjectCache {
	m := make(map[string]ObjectCache)
	var o ObjectCache
	for i := range sz {
		for j := range sz {
			o.BodyMetadataIndex = i*sz + j
			m[fmt.Sprintf("%d.%d", i, j)] = o
		}
	}

	return m
}

func mapStrRead(m map[string]ObjectCache, sz uint64) {
	var ok bool
	for i := range sz {
		for j := range sz {
			_, ok = m[fmt.Sprintf("%d.%d", i, j)]
			if !ok {
				panic(fmt.Errorf("missing entry i, j = %d, %d", i, j))
			}
		}
	}
}

func mapStr2Make(sz uint64) map[string]ObjectCache {
	m := make(map[string]ObjectCache)
	var o ObjectCache
	var key string
	for i := range sz {
		key = fmt.Sprintf("%d", i)
		for j := range sz {
			o.BodyMetadataIndex = i*sz + j
			m[fmt.Sprintf("%s.%d", key, j)] = o
		}
	}

	return m
}

func mapStr2Read(m map[string]ObjectCache, sz uint64) {
	var ok bool
	var key string
	for i := range sz {
		key = fmt.Sprintf("%d", i)
		for j := range sz {
			_, ok = m[fmt.Sprintf("%s.%d", key, j)]
			if !ok {
				panic(fmt.Errorf("missing entry i, j = %d, %d", i, j))
			}
		}
	}
}
