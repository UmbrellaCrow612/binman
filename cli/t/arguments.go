package t

// Holds the arguments / args passed to the CLI mapped to fields
type ArgOptions struct {
	// Holds the abs path to the base path the CLI is running in for example c:/dev/some/project
	BasePath string

	// Holds the abs path to the binman config file i.e the binman.yml file for example c:/dev/some/project/binman.yml
	ConfigPath string
}