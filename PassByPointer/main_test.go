package main 
import (
	"testing"
)

func BenchmarkPBP(b *testing.B) {
	var obj Bigstruct = Bigstruct{}
	for i := 0; i < b.N; i++ {
		fooPBP(&obj)
	}
}

func BenchmarkPBV(b*testing.B){
	var obj Bigstruct = Bigstruct{}
	for i := 0; i < b.N; i++ {
		fooPBV(obj)
	}
}

