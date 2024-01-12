package trie

type node struct {
	isWord bool
	child  map[rune]*node
}

type Trie struct {
	root *node
}

func (t *Trie) Create(str string) {
	cur := t.root
	for _, value := range str {
		val, ok := cur.child[value]
		if !ok {
			val = &node{
				child: make(map[rune]*node),
			}
			cur.child[value] = val
		}
		cur = val
	}
	cur.isWord = true
}

func (t *Trie) Find(str string) bool {
	cur := t.root
	for _, value := range str {
		val, ok := cur.child[value]
		if !ok {
			return false
		}
		cur = val
	}
	return cur.isWord
}

func (t *Trie) Contain(prefix string) bool {
	cur := t.root
	for _, value := range prefix {
		val, ok := cur.child[value]
		if !ok {
			return false
		}
		cur = val
	}
	return true
}
