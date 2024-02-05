package union_find

func NewUnionFind(obj []int) *unionFind {
	u := &unionFind{
		parent: make(map[int]int),
		rank:   make(map[int]int),
	}
	for _, val := range obj {
		// 一开始每个人都指向自己
		// 自己是自己的根节点
		u.parent[val] = val
		u.rank[val] = 1
	}
	return u
}

// UnionFind 并查集
type unionFind struct {
	parent map[int]int // index代表当前节点 value代表当前节点的父节点
	rank   map[int]int // 代表当前节点的优先级 （低的往高的合并）
}

func (u *unionFind) GetSize() int {
	return len(u.parent)
}

func (u *unionFind) Find(p int) int {
	for p != u.parent[p] {
		u.parent[p] = u.parent[u.parent[p]]
		p = u.parent[p]
	}
	return p
}

func (u *unionFind) IsConnected(p, q int) bool {
	pRoot := u.Find(p)
	qRoot := u.Find(q)
	return qRoot == pRoot
}

func (u *unionFind) UnionElements(p, q int) {
	pRoot := u.Find(p)
	qRoot := u.Find(q)
	if pRoot == qRoot {
		return
	}

	if u.rank[pRoot] < u.rank[qRoot] {
		u.parent[pRoot] = u.parent[qRoot]
	} else if u.rank[pRoot] > u.rank[qRoot] {
		u.parent[qRoot] = u.parent[pRoot]
	} else {
		u.parent[qRoot] = u.parent[pRoot]
		u.rank[pRoot] += 1
	}
}
