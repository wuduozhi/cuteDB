package cutedb

import (
	"fmt"
	"testing"
)

func TestDbOpen(t *testing.T) {
	db, _ := Open("test")
	var size = 800000
	for i := 0; i < size; i++ {
		db.Put(fmt.Sprintf("a-%v", i), fmt.Sprintf("value-%v", i))
		db.Put(fmt.Sprintf("b-%v", i), fmt.Sprintf("value-%v", i))
		db.Put(fmt.Sprintf("c-%v", i), fmt.Sprintf("value-%v", i))
	}

	for i := 0; i < size; i++ {
		value, _, _ := db.Get(fmt.Sprintf("c-%v", i))
		if i%100 == 0 {
			t.Log(value)
		}
	}
}
