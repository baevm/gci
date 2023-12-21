package object

// Used to keep track of variables and its values
type Environment struct {
	// Hashmap of variable names and its values
	store map[string]Object
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)

	return &Environment{store: s}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
