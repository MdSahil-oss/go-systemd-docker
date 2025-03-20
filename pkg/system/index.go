package system

type Index struct {
	Name     string         `yaml:"name"`
	Services []IndexService `yaml:"services"`
}

type IndexType func(*Index)

func NewIndex(opts ...IndexType) *Index {
	var index Index
	for _, opt := range opts {
		opt(&index)
	}
	return &index
}

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
	Image  string `yaml:"image" header:"Image"`
	Status string `header:"Status"`
}

type indexServiceType func(*IndexService)

func NewIndexService(opts ...indexServiceType) *IndexService {
	var is IndexService
	for _, opt := range opts {
		opt(&is)
	}
	return &is
}

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

func withIndexServiceImage(img string) indexServiceType {
	return func(is *IndexService) {
		is.Image = img
	}
}
