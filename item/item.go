package item

type Interface interface {
	CanonicalName() string
	AllNames() []string
	LatestVersion() string
}
