package controller

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_parseLine(t *testing.T) { //nolint:funlen
	t.Parallel()
	data := []struct {
		name     string
		txt      string
		state    *State
		isErr    bool
		expState *State
		exp      string
	}{
		{
			name: "simple",
			txt:  `[3.541828, "o", "Bye!"]`,
			state: &State{
				Speed: 1,
			},
			exp: `[3.541828, "o", "Bye!"]`,
			expState: &State{
				Speed:      1,
				Prev:       3.541828,
				ActualPrev: 3.541828,
			},
		},
		{
			name: "change speed",
			txt:  `*2`,
			state: &State{
				Speed: 1,
			},
			exp: "",
			expState: &State{
				Speed: 2,
			},
		},
		{
			name: "trim time",
			txt:  `1`,
			state: &State{
				Speed: 1,
			},
			exp: "",
			expState: &State{
				Speed: 1,
				Delta: 1,
			},
		},
		{
			name: "trim time 2",
			txt:  `[6.0, "o", "Bye!"]`,
			state: &State{
				Speed:      1,
				Prev:       2.0,
				ActualPrev: 3.0,
				Delta:      1,
			},
			exp: `[4, "o", "Bye!"]`,
			expState: &State{
				Speed:      1,
				Prev:       4.0,
				ActualPrev: 6.0,
			},
		},
		{
			name: "change speed 2",
			txt:  `[7.0, "o", "Bye!"]`,
			state: &State{
				Speed:      2,
				Prev:       2.0,
				ActualPrev: 3.0,
			},
			exp: `[4, "o", "Bye!"]`,
			expState: &State{
				Speed:      2,
				Prev:       4.0,
				ActualPrev: 7.0,
			},
		},
		{
			name: "change speed and trim",
			txt:  `[6.0, "o", "Bye!"]`,
			state: &State{
				Speed:      2,
				Prev:       2.0,
				ActualPrev: 3.0,
				Delta:      1,
			},
			exp: `[3, "o", "Bye!"]`,
			expState: &State{
				Speed:      2,
				Prev:       3.0,
				ActualPrev: 6.0,
			},
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			s, err := parseLine(d.txt, d.state)
			if err != nil {
				if d.isErr {
					return
				}
				t.Fatal(err)
			}
			if d.isErr {
				t.Fatal("error must be returned")
			}
			if s != d.exp {
				t.Fatalf("wanted %s, got %s", d.exp, s)
			}
			if diff := cmp.Diff(d.state, d.expState); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
