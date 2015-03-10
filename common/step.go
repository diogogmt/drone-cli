package common

// Step represents a step in the build process, including
// the execution environment and parameters.
type Step struct {
	Image       string
	Pull        bool
	Privileged  bool
	Environment []string
	Entrypoint  []string
	Command     []string
	Volumes     []string
	WorkingDir  string `yaml:"working_dir"`
	NetworkMode string `yaml:"net"`

	// Config represents the unique configuration details
	// for each plugin.
	Config map[string]interface{} `yaml:"config,inline"`
}