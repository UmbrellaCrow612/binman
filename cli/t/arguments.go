package t

// Holds the arguments / args passed to the CLI mapped to fields
type ArgOptions struct {
	// Holds the abs path to the base path the CLI is running in for example c:/dev/some/project
	BasePath string

	// Holds the abs path to the binman config file i.e the binman.yml file for example c:/dev/some/project/binman.yml
	ConfigPath string

	// Contains a array of specific architectures to build for example [x64]
	Architectures []string

	// Contains a list of specific platforms to build for example [linux, darwin]
	Platforms []string

	// Contains a list of specific packages to download i.e if you want to just get ripgrep or some other package you defined
	Packages []string
}
