package goex

import (
	"iter"
)

type MapTree[K comparable, V any] map[K]any

func (m MapTree[K, V]) Root() map[K]any {
	return m
}

func (m MapTree[K, V]) IsEmpty() bool {
	return len(m) == 0
}

//func (m MapTree[K, V]) AllAny() iter.Seq2[K, any] {
//
//}

func (m MapTree[K, V]) Clone() any {
	r := MapTree[K, V]{}

	for k, v := range m {
		switch v := v.(type) {
		case Cloner:
			r[k] = v.Clone()
		default:
			r[k] = v
		}
	}

	return r
}

func (m MapTree[K, V]) Contains(path ...K) (ok bool) {
	r := m
	li := len(path) - 1
	for i, name := range path {
		if i == li {
			_, ok = r[name]
			return
		} else {
			r, ok = r.getNode(name)
			if !ok {
				return ok
			}
		}
	}

	return true // empty path is ture
}

func (m MapTree[K, V]) EnumObjects(path ...K) iter.Seq2[K, MapTree[K, V]] {
	if o, ok := m.GetNode(path...); ok {
		return func(yield func(K, MapTree[K, V]) bool) {
			for name, v := range o.Root() {
				if v, ok := m.toNode(v); ok {
					if !yield(name, v) {
						return
					}
				}
			}
		}
	}

	return func(yield func(K, MapTree[K, V]) bool) {}
}

func (m MapTree[K, V]) Get(path ...K) (v V, ok bool) {
	r := m
	li := len(path) - 1
	for i, name := range path {
		if i == li {
			if av, ok := r[name]; ok {
				switch av := av.(type) {
				case V:
					return av, true
				}
			}

			ok = false
			return
		} else {
			r, ok = r.getNode(name)
			if !ok {
				return
			}
		}
	}

	ok = false
	return
}

func (m MapTree[K, V]) GetNode(path ...K) (v MapTree[K, V], ok bool) {
	r := m
	li := len(path) - 1
	for i, name := range path {
		if i == li {
			if v, ok = r.getNode(name); ok {
				return
			}

			ok = false
			return
		} else {
			r, ok = r.getNode(name)
			if !ok {
				return
			}
		}
	}

	ok = false
	return
}

func (m MapTree[K, V]) GetAny(path ...K) (v any, ok bool) {
	r := m
	li := len(path) - 1
	for i, name := range path {
		if i == li {
			if v, ok = r[name]; ok {
				return
			}

			ok = false
			return
		} else {
			r, ok = r.getNode(name)
			if !ok {
				return
			}
		}
	}

	ok = false
	return
}

func (m MapTree[K, V]) Set(v V, path ...K) (ok bool) {
	if len(path) == 0 {
		return
	}

	r := m
	li := len(path) - 1
	for i, name := range path {
		if i == li {
			r[name] = v

			return true
		} else {
			if ar, ok := r[name]; ok {
				switch ar := ar.(type) {
				case MapTree[K, V]:
					r = ar
				default:
					return false // not a node can not traverse deeper
				}
			} else { // no node found, create new
				n := MapTree[K, V]{}
				r[name] = n
				r = n
			}
		}
	}

	return false
}

func (m MapTree[K, V]) SetNode(v MapTree[K, V], path ...K) (ok bool) {
	if len(path) == 0 {
		return
	}

	r := m
	li := len(path) - 1
	for i, name := range path {
		if i == li {
			r[name] = v

			return true
		} else {
			if ar, ok := r[name]; ok {
				switch ar := ar.(type) {
				case MapTree[K, V]:
					r = ar
				default:
					return false // not a node can not traverse deeper
				}
			} else { // no node found, create new
				n := MapTree[K, V]{}
				r[name] = n
				r = n
			}
		}
	}

	return false
}

func (m MapTree[K, V]) Merge(data MapTree[K, V]) (ok bool) {
	for k, v := range data {
		if mv, ok := m[k]; ok {
			data_node, okv := v.(MapTree[K, V])
			m_node, okmv := mv.(MapTree[K, V])

			if okv && okmv {
				m_node.Merge(data_node)
			} else if !okv && !okmv {
				m[k] = v
			} else {
				return false
			}
		} else {
			m[k] = v
		}
	}

	return true
}

func (m MapTree[K, V]) getNode(key K) (r MapTree[K, V], ok bool) {
	ar, ok := m[key]
	if !ok {
		return nil, false
	}

	return m.toNode(ar)
}

func (m MapTree[K, V]) toNode(v any) (r MapTree[K, V], ok bool) {
	switch v := v.(type) {
	case MapTree[K, V]:
		return v, true
	case map[K]any:
		return MapTree[K, V](v), true
	}

	return nil, false
}
