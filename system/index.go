package system

type Index struct {
	Name     string         `yaml:"name"`
	Services []IndexService `yaml:"services"`
}

type IndexType func(*Index)

func withIndexName(name string) IndexType {
	return func(i *Index) {
		i.Name = name
	}
}

func withIndexServices(services []IndexService) IndexType {
	return func(i *Index) {
		i.Services = services
	}
}

type IndexService struct {
	Name   string `yaml:"name" header:"Name"`
	Path   string `yaml:"path" header:"Path"`
	Status string `header:"Status"`
}

type indexServiceType func(*IndexService)

func withIndexServiceName(name string) indexServiceType {
	return func(is *IndexService) {
		is.Name = name
	}
}

func withIndexServicePath(path string) indexServiceType {
	return func(is *IndexService) {
		is.Path = path
	}
}
