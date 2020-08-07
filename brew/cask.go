package brew

type Cask struct {
	Token    string   `json:"token"`
	Names    []string `json:"name"`
	Homepage string   `json:"homepage"`
	Url      string   `json:"url"`
	Version  string   `json:"version"`
}

func (c Cask) CanonicalName() string {
	return c.Token
}

func (c Cask) AllNames() []string {
	return c.Names
}

func (c Cask) LatestVersion() string {
	return c.Version
}
