package brew

import "testing"

func TestUsesFromMacos_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		args   []byte
		wanted UsesFromMacos
	}{
		{"object", []byte(`{"foo": "bar"}`), UsesFromMacos{obj: struct{ name, phase string }{name: "foo", phase: "bar"}}},
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
			if u.obj.phase != tt.wanted.obj.phase {
				t.Errorf("error: obj.phase got=%s, wanted=%s", u.obj.phase, tt.wanted.obj.phase)
			}
		})
	}
}
