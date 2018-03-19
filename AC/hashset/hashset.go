package hashset

type HashSet struct {
	Set map[int]bool
}

func NewHashSet() *HashSet {
	return &HashSet{make(map[int]bool)}
}

func (set *HashSet) Add(i int) bool {
	_, found := set.Set[i]
	set.Set[i] = true
	return !found //False if it existed already
}

func (set *HashSet) Get(i int) bool {
	_, found := set.Set[i]
	return found //true if it existed already
}

func (set *HashSet) Remove(i int) {
	delete(set.Set, i)
}
