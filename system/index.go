package system

type index struct {
	name     string         `yaml:"name"`
	services []indexService `yaml:"services"`
}

type indexType func(*index)

func withIndexName(name string) indexType {
	return func(i *index) {
		i.name = name
	}
}

func withIndexServices(services []indexService) indexType {
	return func(i *index) {
		i.services = services
	}
}

type indexService struct {
	name string `yaml:"name"`
	path string `yaml:"path"`
}

type indexServiceType func(*indexService)

func withIndexServiceName(name string) indexServiceType {
	return func(is *indexService) {
		is.name = name
	}
}

func withIndexServicePath(path string) indexServiceType {
	return func(is *indexService) {
		is.path = path
	}
}
