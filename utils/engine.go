package utils

// 每个service维护一个双向链表
type Engine struct {
	Id			uint8
	Addr		string
	Service		string
	Prev		*Engine
	Next		*Engine
}

func NewEngineHead()  {
	
}
