package brew

func (f Formula) CanonicalName() string {
	return f.Name
}

func (f Formula) AllNames() []string {
	return append([]string{f.Name}, f.Aliases...)
}

func (f Formula) LatestVersion() string {
	return f.Versions.Stable
}
