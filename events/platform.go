package events

type Platform interface {
	Platform() string
	OsVersion() string
	Build() string
}
