package conf

//Context 一条url以及名字的结构
type Context struct {
	URL  string `yaml:"url"`
	Name string `yaml:"name"`
}

//Projects 实际上需要用结构体总和
type Projects struct {
	Project []ProjectContext `yaml:"projects"`
}

//ProjectContext 每个项目的结构
type ProjectContext struct {
	Projectname string    `yaml:"projectname"`
	Method      string    `yaml:"method"`
	Context     []Context `yaml:"context"`
}
