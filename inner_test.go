package art

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type LeftmostSet struct {
	nodeFactor func() inode[Value]
	expected   node[Value]
}

func TestLeftmost(t *testing.T) {
	l := &leaf[Value]{key: Key("a key"), value: Value("value")}
	assert.Equal(t, l, l.leftmost())

	sets := []LeftmostSet{
		{
			nodeFactor: func() inode[Value] {
				return &node4[Value]{}
			},
			expected: l,
		},
		{
			nodeFactor: func() inode[Value] {
				return &node16[Value]{}
			},
			expected: l,
		},
		{
			nodeFactor: func() inode[Value] {
				return &node48[Value]{}
			},
			expected: l,
		},
		{
			nodeFactor: func() inode[Value] {
				return &node256[Value]{}
			},
			expected: l,
		},
	}
	var child node[Value]
	for _, set := range sets {
		// leaf level
		child = l

		// level + 1
		n := set.nodeFactor()
		n.addChild('a', child)
		child = &inner[Value]{node: n}

		// level + 1
		nn := set.nodeFactor()
		nn.addChild('a', child)
		upper := &inner[Value]{node: n}

		leftmost := upper.leftmost()
		assert.NotNil(t, leftmost)
		assert.Equal(t, set.expected, leftmost)
	}

}
