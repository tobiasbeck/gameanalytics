package gameanalytics

type Platform struct {
	platform  string
	osVersion string
	build     string
}

func (p Platform) Platform() string {
	return p.platform
}

func (p Platform) OsVersion() string {
	return p.osVersion
}

func (p Platform) Build() string {
	return p.build
}

// NewPlatform instanciates a new Platform configuration
func NewPlatform(platform string, osVersion string) Platform {
	p := Platform{
		platform:  platform,
		osVersion: osVersion,
	}

	return p
}
