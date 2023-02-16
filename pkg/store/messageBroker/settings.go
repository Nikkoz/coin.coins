package messageBroker

type (
	Settings struct {
		Connection     string
		Broker         string
		SessionTimeout uint16

		Sasl           Sasl
		SchemaRegistry SchemaRegistry
	}

	Sasl struct {
		IsUse            bool
		Broker           string
		User             string
		Password         string
		Mechanisms       string
		SecurityProtocol string
	}

	SchemaRegistry struct {
		Type string
		Url  string
	}
)

func NewSettings(connection, broker string, timeout uint16, useSasl bool, saslBroker, user, password, mechanisms, protocol, srType, url string) Settings {
	return Settings{
		Connection:     connection,
		Broker:         broker,
		SessionTimeout: timeout,
		Sasl: Sasl{
			IsUse:            useSasl,
			Broker:           saslBroker,
			User:             user,
			Password:         password,
			Mechanisms:       mechanisms,
			SecurityProtocol: protocol,
		},
		SchemaRegistry: SchemaRegistry{
			Type: srType,
			Url:  url,
		},
	}
}
