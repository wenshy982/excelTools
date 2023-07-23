package randx

import (
	"testing"
)

func TestOrderNo(t *testing.T) {
	orderNo := OrderNo()
	if len(orderNo) != 20 {
		t.Errorf("OrderNo() = %v, want %v", len(orderNo), 20)
	}
	t.Logf("OrderNo() = %v", orderNo)
}

func TestGenOrderNo(t *testing.T) {
	orderNo := GenOrderNo()
	if len(orderNo) != 16 {
		t.Errorf("GenOrderNo() = %v, want %v", len(orderNo), 16)
	}
	t.Logf("GenOrderNo() = %v", orderNo)
}
