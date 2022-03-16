package matcher

import "testing"

const (
	singles int = iota
	double
	triples
)

const (
	a20 int = iota
	a30
	a40
)

func TestCompatibility(t *testing.T) {
	possible_players := []int{singles, double, triples}
	possible_points := []int{a20, a30, a40}

	m := NewMatcher([][]int{possible_players, possible_points})

	cA := [][]int{{singles, triples}, {a20, a40}}
	cB := [][]int{{singles}, {a20}}

	if ok := m.Match(cA, cB); !ok {
		t.Error("deberian ser compatibles")
	}

	c1 := [][]int{{singles, triples}, {a20, a40}}
	c2 := [][]int{{singles}, {a20}}
	c3 := [][]int{{double}, {a20}}
	c4 := [][]int{{singles, triples}, {a30}}
	c5 := [][]int{{double, triples}, {a30, a40}}

	if ok := m.Match(c1, c2); !ok { //
		t.Log("c1 y c2 deberian ser compatibles")
	}

	if ok := !m.Match(c1, c3); !ok { //
		t.Log("c1 y c3 NO deberian ser compatibles")
	}

	if ok := !m.Match(c1, c4); !ok { //
		t.Log("c1 y c4 NO deberian ser compatibles")
	}

	if ok := m.Match(c1, c5); !ok { //
		t.Log("c1 y c5 deberian ser compatibles")
	}

	if ok := !m.Match(c2, c3); !ok { //
		t.Log("c2 y c3 NO deberian ser compatibles")
	}

	if ok := !m.Match(c2, c4); !ok { //
		t.Log("c2 y c4 NO deberian ser compatibles")
	}

	if ok := !m.Match(c3, c4); !ok { //
		t.Log("c3 y c4 NO deberian ser compatibles")
	}

}
