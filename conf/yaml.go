package conf

//context
type Context struct {
	URL  string `yaml:"url"`
	Name string `yaml:"name"`
}

type Projects struct {
	Project1 []Project `yaml:"projects"`
}

type Project struct {
	Projectname string    `yaml:project_name`
	Method      string    `yaml:"method"`
	Context     []Context `yaml:"context"`
}
