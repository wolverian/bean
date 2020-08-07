package brew

import (
	"reflect"
	"testing"
)

//nolint:unused
type ex struct {
	name  string
	phase []string
}

func TestUsesFromMacos_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		args   []byte
		wanted UsesFromMacos
	}{
		{"object", []byte(`{"foo": "bar"}`), UsesFromMacos{obj: ex{"foo", []string{"bar"}}}},
		{"object_array", []byte(`{"foo": ["bar", "baz"]}`), UsesFromMacos{"", ex{name: "foo", phase: []string{"bar", "baz"}}}},
		{"string", []byte(`"foobar"`), UsesFromMacos{str: "foobar"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsesFromMacos{}
			if err := u.UnmarshalJSON(tt.args); err != nil {
				t.Errorf("UnmarshalJSON() error = %v", err)
			}
			if u.str != tt.wanted.str {
				t.Errorf("error: str got=%s, wanted=%s", u.str, tt.wanted.str)
			}
			if u.obj.name != tt.wanted.obj.name {
				t.Errorf("error: obj.name got=%s, wanted=%s", u.obj.name, tt.wanted.obj.name)
			}
			if !reflect.DeepEqual(u.obj.phase, tt.wanted.obj.phase) {
				t.Errorf("error: obj.phase got=%+v, wanted=%+v", u.obj.phase, tt.wanted.obj.phase)
			}
		})
	}
}
