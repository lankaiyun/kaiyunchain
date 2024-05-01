package ast

type Node interface {
	Node()
}

type Expr interface {
	Node
	ExprNode()
}

type Stmt interface {
	Node
	StmtNode()
}

type Contract struct {
	Name  string
	Stmts []Stmt
}

func (c *Contract) Node() {}

type Func struct {
	Name        string           // 函数名
	Parameters  []*ExprParameter // 函数参数
	ReturnValue []*ExprS         // 返回值
	Block       *ExprBlock
}

func (f *Func) Node() {}

func (f *Func) StmtNode() {}

type StmtIdent struct {
	Name string
	Type string
}

func (s *StmtIdent) Node() {}

func (s *StmtIdent) StmtNode() {}

type ExprIdent struct {
	Name string
	Type string
}

func (e *ExprIdent) Node() {}

func (e *ExprIdent) ExprNode() {}

type ExprParameter struct {
	Name string
	Type string
}

func (e *ExprParameter) Node() {}

func (e *ExprParameter) ExprNode() {}

type ExprS struct {
	Expr Expr
}

func (e *ExprS) Node() {}

func (e *ExprS) ExprNode() {}

type ExprBlock struct {
	Exprs []Expr
}

func (b *ExprBlock) Node() {}

func (b *ExprBlock) ExprNode() {}

type Assign struct {
	Obj   *ExprIdent
	Value Expr
}

func (b *Assign) Node() {}

func (b *Assign) ExprNode() {}
