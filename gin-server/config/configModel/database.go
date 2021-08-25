package configModel

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}
