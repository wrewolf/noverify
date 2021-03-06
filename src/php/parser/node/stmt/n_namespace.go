package stmt

import (
	"github.com/VKCOM/noverify/src/php/parser/freefloating"
	"github.com/VKCOM/noverify/src/php/parser/node"
	"github.com/VKCOM/noverify/src/php/parser/position"
	"github.com/VKCOM/noverify/src/php/parser/walker"
)

// Namespace node
type Namespace struct {
	FreeFloating  freefloating.Collection
	Position      *position.Position
	NamespaceName node.Node
	Stmts         []node.Node
}

// NewNamespace node constructor
func NewNamespace(NamespaceName node.Node, Stmts []node.Node) *Namespace {
	return &Namespace{
		FreeFloating:  nil,
		NamespaceName: NamespaceName,
		Stmts:         Stmts,
	}
}

// SetPosition sets node position
func (n *Namespace) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Namespace) GetPosition() *position.Position {
	return n.Position
}

func (n *Namespace) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Namespace) Walk(v walker.Visitor) {
	if !v.EnterNode(n) {
		return
	}

	if n.NamespaceName != nil {
		n.NamespaceName.Walk(v)
	}

	if n.Stmts != nil {
		for _, nn := range n.Stmts {
			if nn != nil {
				nn.Walk(v)
			}
		}
	}

	v.LeaveNode(n)
}
