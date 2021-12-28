package config

type mysqlConfig struct {
	DbHost  string `json:"db_host"`
	DbName  string `json:"db_name"`
	DbPwd   string `json:"db_pwd"`
	DbUser  string `json:"db_user"`
	DbPort  string `json:"db_port"`
	Maxidle int    `json:"maxidle" yaml:"maxidle"`
	Maxopen int    `json:"maxopen" yaml:"maxopen"`
}

func (t *mysqlConfig) Reload() error {
	println("reload mysql")
	return nil
}
func (t *mysqlConfig) IsLoad() bool {
	if len(t.DbHost) > 0 {
		return true
	}
	return false
}
