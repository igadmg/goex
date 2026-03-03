package goex

import (
	"reflect"
)

type equal[T any] interface {
	Equal(other T) bool
}

type EqualedAny interface {
	EqualAny(other EqualedAny) bool
}

func Equal[T equal[T]](a, b T) bool {
	return a.Equal(b)
}

func EqualAny(a, b any) bool {
	aeq, aok := a.(EqualedAny)
	beq, bok := b.(EqualedAny)
	if aok && bok {
		return aeq.EqualAny(beq)
	} else {
		av := reflect.ValueOf(a)
		if av.Kind() == reflect.Ptr {
			av = av.Elem()
		}
		return av.Equal(reflect.ValueOf(b))
	}
}
