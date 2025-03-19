package system

type System struct {
	Name        string   `yaml:"name"`
	DisplayName string   `yaml:"displayName"`
	Description string   `yaml:"description"`
	Executable  string   `yaml:"executable"`
	Arguments   []string `yaml:"arguments"`
}

type SystemOption func(*System)

func NewSystem(opts ...SystemOption) *System {
	sys := System{}

	for _, opt := range opts {
		opt(&sys)
	}

	return &sys
}

func WithName(name string) SystemOption {
	return func(sys *System) {
		sys.Name = name
	}
}

func WithDisplayName(dname string) SystemOption {
	return func(sys *System) {
		sys.DisplayName = dname
	}
}

func WithDescription(desc string) SystemOption {
	return func(sys *System) {
		sys.Description = desc
	}
}

func WithExecutable(exec string) SystemOption {
	return func(sys *System) {
		sys.Executable = exec
	}
}

func WithArguments(args []string) SystemOption {
	return func(sys *System) {
		sys.Arguments = args
	}
}
