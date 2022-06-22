package set

import "fmt"

type HashSet struct {
	valMap map[interface{}]bool
}

func makeHashMap(valMap map[interface{}]bool) *HashSet {
	return &HashSet{valMap: valMap}
}

func (s *HashSet) add(key interface{}) {
	s.valMap[key] = true
}

func (s *HashSet) remove(key interface{}) {
	delete(s.valMap, key)
}

func (s *HashSet) contains(key interface{}) bool {
	return s.valMap[key]
}

func (s *HashSet) size() int {
	return len(s.valMap)
}

func TestHashSet() {
	set := makeHashMap(map[interface{}]bool{})
	set.add(1)
	set.add(2)
	set.add(3)

	set.remove(4)
	fmt.Printf(" Size of set is %d \n", set.size())
	set.remove(3)
	fmt.Printf(" Size of set is %d \n", set.size())
	set.add(5)
	fmt.Printf(" Size of set is %d \n", set.size())
	fmt.Println(set.contains(1))
	fmt.Println(set.contains(100))
}
