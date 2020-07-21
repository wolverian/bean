package brew

import "encoding/json"

type UsesFromMacos struct {
	str string
	obj struct{ name, phase string }
}

func (u UsesFromMacos) Name() string {
	if u.str != "" {
		return u.str
	} else {
		return u.obj.name
	}
}

func (u *UsesFromMacos) UnmarshalJSON(bytes []byte) error {
	if bytes[0] == '{' {
		return json.Unmarshal(bytes, &u.obj)
	} else {
		u.str = string(bytes)
		return nil
	}
}

type Revision struct {
	rev string
}

func (r *Revision) UnmarshalJSON(bytes []byte) error {
	if bytes[0] == '"' {
		r.rev = string(bytes)
	} else {
		r.rev = json.Number(bytes).String()
	}
	return nil
}

type Formula struct {
	Name                    string          `json:"name"`
	FullName                string          `json:"full_name"`
	OldName                 *string         `json:"oldname"`
	Aliases                 []string        `json:"aliases"`
	VersionedFormulae       []string        `json:"versioned_formulae"`
	Desc                    string          `json:"desc"`
	License                 string          `json:"license"`
	Homepage                string          `json:"homepage"`
	Versions                Versions        `json:"versions"`
	Urls                    Urls            `json:"urls"`
	Revision                Revision        `json:"revision"`
	VersionScheme           int             `json:"version_scheme"`
	Bottle                  Bottle          `json:"bottle"`
	KegOnly                 bool            `json:"keg_only"`
	BottleDisabled          bool            `json:"bottle_disabled"`
	Options                 []interface{}   `json:"options"`
	BuildDependencies       []string        `json:"build_dependencies"`
	Dependencies            []string        `json:"dependencies"`
	RecommendedDependencies []string        `json:"recommended_dependencies"`
	OptionalDependencies    []string        `json:"optional_dependencies"`
	UsesFromMacos           []UsesFromMacos `json:"uses_from_macos"`
	Requirements            []Requirements  `json:"requirements"`
	ConflictsWith           []string        `json:"conflicts_with"`
	Caveats                 string          `json:"caveats"`
	Installed               []Installed     `json:"installed"`
	LinkedKeg               *string         `json:"linked_keg"`
	Pinned                  bool            `json:"pinned"`
	Outdated                bool            `json:"outdated"`
}
type Versions struct {
	Stable string `json:"stable"`
	Devel  string `json:"devel"`
	Head   string `json:"head"`
	Bottle bool   `json:"bottle"`
}
type Stable struct {
	URL      string   `json:"url"`
	Tag      string   `json:"tag"`
	Revision Revision `json:"revision"`
}
type Urls struct {
	Stable Stable `json:"stable"`
}
type File struct {
	URL    string `json:"url"`
	Sha256 string `json:"sha256"`
}
type Files struct {
	Catalina   File `json:"catalina"`
	Mojave     File `json:"mojave"`
	HighSierra File `json:"high_sierra"`
	Sierra     File `json:"sierra"`
	ElCapitan  File `json:"el_capitan"`
	Yosemite   File `json:"yosemite"`
}
type BottleStable struct {
	Rebuild int    `json:"rebuild"`
	Cellar  string `json:"cellar"`
	Prefix  string `json:"prefix"`
	RootURL string `json:"root_url"`
	Files   Files  `json:"files"`
}
type Bottle struct {
	Stable BottleStable `json:"stable"`
}
type Requirements struct {
	Name     string   `json:"name"`
	Cask     string   `json:"cask"`
	Download string   `json:"download"`
	Version  string   `json:"version"`
	Contexts []string `json:"contexts"`
}
type Installed struct {
	Version               string                `json:"version"`
	UsedOptions           []interface{}         `json:"used_options"`
	BuiltAsBottle         bool                  `json:"built_as_bottle"`
	PouredFromBottle      bool                  `json:"poured_from_bottle"`
	RuntimeDependencies   []RuntimeDependencies `json:"runtime_dependencies"`
	InstalledAsDependency bool                  `json:"installed_as_dependency"`
	InstalledOnRequest    bool                  `json:"installed_on_request"`
}
type RuntimeDependencies struct {
	FullName string `json:"full_name"`
	Version  string `json:"version"`
}
