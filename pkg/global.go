package pkg

type AppName string

const (
	APIAppName            = AppName("ElectionAPI")
	MySQLConnectorAppName = AppName("MySQLConnector")
)

func (name AppName) ToString() string {
	return string(name)
}
