package create

type Flags struct {
	name       *string
	domainName *string
	entrypoint *string
	expose     *[]string
	publish    *[]string
	env        *[]string
}

type FlagsType func(*Flags)

func UpdateFlags(flags *Flags, args ...FlagsType) *Flags {
	for _, arg := range args {
		arg(flags)
	}
	return flags
}

func WithName(name *string) FlagsType {
	return func(f *Flags) {
		f.name = name
	}
}

func WithDomainName(dname *string) FlagsType {
	return func(f *Flags) {
		f.domainName = dname
	}
}

func WithEntrypoint(en *string) FlagsType {
	return func(f *Flags) {
		f.entrypoint = en
	}
}

func WithExpose(expose *[]string) FlagsType {
	return func(f *Flags) {
		f.expose = expose
	}
}

func WithPublish(ports *[]string) FlagsType {
	return func(f *Flags) {
		f.publish = ports
	}
}

func WithEnv(envs *[]string) FlagsType {
	return func(f *Flags) {
		f.env = envs
	}
}
