package obj

import (
	"fmt"

	"github.com/takeru56/tcompiler/code"
)

type ObjectType string

const (
	INTEGER_OBJ  = "INTEGER"
	FUNCTION_OBJ = "FUNCTION"
	CLASS_OBJ    = "CLASS"
)

type Object interface {
	Type() ObjectType
	Inspect() string
	Size() int
}

type Integer struct {
	Value int
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

// TODO: のちほど32bitに対応する
// ひとまず2byte(16bit)で表現
func (i *Integer) Size() int { return 2 }

type Function struct {
	Id           int
	Instructions code.Instructions
	NumArg       int
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string  { return fmt.Sprintf("function%p", f) }

func (f *Function) Size() int { return len(f.Instructions) }

type Class struct {
	Index          int
	NumInstanceVal int
	NumMethod      int
	ConstantPool   []Object
}

func (c *Class) Type() ObjectType { return CLASS_OBJ }
func (c *Class) Inspect() string  { return fmt.Sprintf("class%p", c) }

// それぞれ1byte
func (c *Class) Size() int { return 2 }
