package goex

type MapTree[K comparable, V any] map[K]any

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
			r, ok = r[name].(MapTree[K, V])
			if !ok {
				return
			}
		}
	}

	return true // empty path is ture
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
			r, ok = r[name].(MapTree[K, V])
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
			r, ok = r[name].(MapTree[K, V])
			if !ok {
				return
			}
		}
	}

	ok = false
	return
}

func (m MapTree[K, V]) Set(v any, path ...K) (ok bool) {
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
