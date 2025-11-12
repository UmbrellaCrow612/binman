package shared

type Config struct {
	Binaries []Binary `yaml:"binaries"`
}

type Binary struct {
	// Name
	Name string `yaml:"name"`

	// Key is platform, value is URL
	URL map[string]string `yaml:"url"`

	// Key value pair of platform then sha key
	SHA256 map[string]string `yaml:"sha256"`

	// Another other Key value pair
	Extra map[string]any `yaml:",inline"`
}
