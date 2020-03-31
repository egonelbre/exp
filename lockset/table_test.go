package lockset_test

import (
	"testing"

	"github.com/egonelbre/exp/lockset"
)

func TestTable(t *testing.T) {
	table := lockset.NewTable()

	A := &lockset.Lock{}
	B := &lockset.Lock{}
	C := &lockset.Lock{}

	// A -> B -> C
	assertNil(t, table.Locked(A))
	assertNil(t, table.Locked(B))
	assertNil(t, table.Locked(C))
	assertTrue(t, table.Unlocking(C))
	assertTrue(t, table.Unlocking(B))
	assertTrue(t, table.Unlocking(A))

	// B -> C
	assertNil(t, table.Locked(B))
	assertNil(t, table.Locked(C))
	assertTrue(t, table.Unlocking(C))
	assertTrue(t, table.Unlocking(B))

	// B -> C -> A!
	assertNil(t, table.Locked(B))
	assertNil(t, table.Locked(C))
	assertNotNil(t, table.Locked(A))
	assertTrue(t, table.Unlocking(A))
	assertTrue(t, table.Unlocking(C))
	assertTrue(t, table.Unlocking(B))
}

func assertNil(t *testing.T, inv *lockset.Inversion) {
	t.Helper()
	if inv != nil {
		t.Fatal(inv.String())
	}
}

func assertNotNil(t *testing.T, inv *lockset.Inversion) {
	t.Helper()
	if inv == nil {
		t.Fatal("expected inversion")
	}
}

func assertTrue(t *testing.T, v bool) {
	t.Helper()
	if v != true {
		t.Fatal("expected true")
	}
}
