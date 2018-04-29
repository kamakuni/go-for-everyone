package reflect

import (
	"reflect"
	"testing"
)

type StructAccess struct {
	Int int
}

func BenchmarkDetectTypeAssert(b *testing.B) {
	var s interface{} = StructAccess{Int: 100}
	for i := 0; i < b.N; i++ {
		if sa, ok := s.(StructAccess); ok {
			_ = sa.Int
		}
	}
}

func BenchmarkDetectTypeReflect(b *testing.B) {
	var s interface{} = StructAccess{Int: 100}
	for i := 0; i < b.N; i++ {
		rv := reflect.ValueOf(s)
		if rv.Type().Name() == "StructAccess" {
			_ = s.(StructAccess).Int
		}
	}
}

func BenchmarkDetectNone(b *testing.B) {
	s := StructAccess{Int: 100}
	for i := 0; i < b.N; i++ {
		_ = s.Int
	}
}
