package object

import (
	"bytes"
	"strings"
)

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

type Hashable interface {
	Hashkey() HashKey
}

func (h *Hash) Type() ObjectType {
	return HASH_OBJ
}

func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}

	for _, pair := range h.Pairs {
		pairs = append(pairs, pair.Key.Inspect()+":"+pair.Value.Inspect())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ","))
	out.WriteString("}")

	return out.String()
}
