package configuration

type SiteList struct {
	Site []Site `yaml:"Site"`
}

type Site struct {
	Url      string `yaml:"Url"`
	DataPath string `yaml:"DataPath"`
}

//type test1 struct {
//	name string
//	age  int
//}
//
//func NewTest() {
//	var a1 *test1
//	a2 := test1{}
//	List := SiteList{}
//	t := 5
//	t1 := &t
//
//}

//type configuration struct {
//	Mailer []Mailer yaml:"Mailer"
//}
//
//type Mailer struct {
//	Smtp Smtp yaml:"Smtp"
//}
//type Smtp struct {
//	Host string yaml:"Host"
//	Port int    yaml:"Port"
//	From string yaml:"From"
//	Pass string yaml:"Pass"
//}
