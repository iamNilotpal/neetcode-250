package main

type HashSet struct {
	data map[int]int
}

func NewHashSet() *HashSet {
	return &HashSet{data: make(map[int]int)}
}

func (hs *HashSet) Add(key int) {
	hs.data[key] = key
}

func (hs *HashSet) Remove(key int) {
	delete(hs.data, key)
}

func (hs *HashSet) Contains(key int) bool {
	_, ok := hs.data[key]
	return ok
}

func main() {
	hs := NewHashSet()

	hs.Add(1)
	hs.Add(2)
	hs.Add(3)
	println("Contains 3", hs.Contains(3))
	println("Contains 4", hs.Contains(4))
}
