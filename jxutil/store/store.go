package store

type Receptacle[T any] struct {
	box map[string]T
}

func NewReceptacle[T any]() *Receptacle[T] {
	return &Receptacle[T]{box: make(map[string]T)}
}

func (r *Receptacle[T]) Set(key string, toy T) {
	r.box[key] = toy
}

func (r *Receptacle[T]) Get(key string) (toy T, exist bool) {
	toy, exist = r.box[key]
	return
}
