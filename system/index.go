package system

type index struct {
	Name     string         `yaml:"name"`
	Services []indexService `yaml:"services"`
}

type indexType func(*index)

func withIndexName(name string) indexType {
	return func(i *index) {
		i.Name = name
	}
}

func withIndexServices(services []indexService) indexType {
	return func(i *index) {
		i.Services = services
	}
}

type indexService struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

type indexServiceType func(*indexService)

func withIndexServiceName(name string) indexServiceType {
	return func(is *indexService) {
		is.Name = name
	}
}

func withIndexServicePath(path string) indexServiceType {
	return func(is *indexService) {
		is.Path = path
	}
}
