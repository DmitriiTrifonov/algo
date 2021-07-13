package linked

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestSingleLinkedList_Reverse tests a Reverse func
func TestSingleLinkedList_Reverse(t *testing.T) {
	max := 100
	list := NewSingleLinkedList()
	tmp := list
	for i := 1; i < max; i++ {
		tmp.Add()
		tmp = tmp.next
	}

	listAddresses := list.GetAddresses()

	reversed := list.Reverse()
	reversedAddresses := reversed.GetAddresses()

	length := len(reversedAddresses)

	require.Equal(t, len(listAddresses), length)

	for i, addr := range listAddresses {
		rev := reversedAddresses[(length-1)-i]
		if addr != rev {
			t.Fatalf("%s is not equal to %s", addr, rev)
		}
	}

}
