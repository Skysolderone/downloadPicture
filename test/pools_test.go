package test

import (
	"fmt"
	"testing"

	"github.com/panjf2000/ants/v2"
)

func Testpool() {
	fmt.Println(1)
}
func TestPool(t *testing.T) {
	defer ants.Release()
	p, _ := ants.NewPool(10000, ants.WithPreAlloc(true))
	defer p.Release()
	for i := 0; i < 100; i++ {
		p.Submit(Testpool)
	}
}
