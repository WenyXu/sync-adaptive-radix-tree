package art

const (
	maxPrefixLen int = 10
)

type Key []byte
type Kind uint8

const (
	Leaf Kind = iota
	Node4
	Node16
	Node48
	Node256
)

// At return a char at post
func (key Key) At(pos int) byte {
	if pos < 0 || pos >= len(key) {
		// imitate the C-like string termination character
		return 0
	}
	return key[pos]
}

type node[T any] interface {
	insert(*leaf[T], int, *olock, uint64) (n node[T], restart bool, updated bool)
	del(Key, int, *olock, uint64, func(node[T])) (deleted, restart bool, deletedNode node[T])
	get(Key, int, *olock, uint64) (value T, found bool, restart bool)
	walk(walkFn[T], int) bool
	addPrefixBefore(node *inner[T], key byte)
	Kind() Kind
	isLeaf() bool
	String() string
	leftmost() node[T]
}

type inner[T any] struct {
	lock      olock
	prefix    [maxPrefixLen]byte
	prefixLen int
	node      inode[T]
}

// walkFn should return false if iteration should be terminated.
type walkFn[T any] func(node[T], int) bool

type inode[T any] interface {
	leftmost() node[T]
	Kind() Kind
	// next returns child after the requested byte
	// if byte is nil - returns leftmost child
	next(*byte) (byte, node[T])
	prev(*byte) (byte, node[T])

	// child return index of the child together with the child
	child(byte) (int, node[T])
	// addChild inserts child at the specified byte
	addChild(byte, node[T])
	// replace updates node at specified index
	// if node is nil - delete the node and adjust metadata.
	// return replaced node
	replace(int, node[T]) node[T]

	// full is true if node reached max size
	full() bool
	// grow the node to next size
	// node256 can't grow and will return nil
	grow() inode[T]

	// min is true if node reached min size
	min() bool
	// shrink is the opposite to grow
	// if node is of the smallest type (node4) nil will be returned
	shrink() inode[T]

	// walk is internal helper to iterate in depth first order over all nodes, including inner nodes
	walk(walkFn[T], int) bool

	String() string
}
