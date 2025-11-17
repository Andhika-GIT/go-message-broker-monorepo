package shared

import "github.com/spf13/viper"

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type RabbitMQRoutingKey struct {
	UserDirectImport  string
	UserSftpImport    string
	UserExport        string
	OrderDirectImport string
	OrderImport       string
	OrderExport       string
}

type RabbitMQQueue struct {
	UserDirectImport  string
	UserSftpImport    string
	UserExport        string
	OrderDirectImport string
	OrderImport       string
	OrderExport       string
}

type RabbitMQConfig struct {
	ConnectURL string
	Exchange   string
	RoutingKey RabbitMQRoutingKey
	Queue      RabbitMQQueue
}

type Config struct {
	// RabbitMQ
	RabbitMQConnectURL string
	RabbitMQExchange   string

	RabbitMQRoutingKey RabbitMQRoutingKey

	RabbitMQQueue RabbitMQQueue

	// Database
	Database DatabaseConfig
}

func InitConfig(v *viper.Viper) *Config {
	cfg := &Config{}

	// RabbitMQ
	cfg.RabbitMQConnectURL = getOrDefault(v, "RABBITMQ_CONNECT_URL", "amqp://guest:guest@localhost:5672/")
	cfg.RabbitMQExchange = getOrDefault(v, "MQ_EXCHANGE_GO_APP", "go-app-exchange")

	// Routing Keys
	cfg.RabbitMQRoutingKey.UserDirectImport = getOrDefault(v, "MQ_RK_USER_DIRECT_IMPORT", "user.import.direct")
	cfg.RabbitMQRoutingKey.UserSftpImport = getOrDefault(v, "MQ_RK_USER_SFTP_IMPORT", "user.import.sftp")
	cfg.RabbitMQRoutingKey.UserExport = getOrDefault(v, "MQ_RK_USER_EXPORT", "user.export")
	cfg.RabbitMQRoutingKey.OrderDirectImport = getOrDefault(v, "MQ_RK_ORDER_DIRECT_IMPORT", "order.import.direct")
	cfg.RabbitMQRoutingKey.OrderImport = getOrDefault(v, "MQ_RK_ORDER_IMPORT", "order.import")
	cfg.RabbitMQRoutingKey.OrderExport = getOrDefault(v, "MQ_RK_ORDER_EXPORT", "order.export")

	// Queues
	cfg.RabbitMQQueue.UserDirectImport = getOrDefault(v, "MQ_Q_USER_DIRECT_IMPORT", "user.import.direct.q")
	cfg.RabbitMQQueue.UserSftpImport = getOrDefault(v, "MQ_Q_USER_SFTP_IMPORT", "user.import.sftp.q")
	cfg.RabbitMQQueue.UserExport = getOrDefault(v, "MQ_Q_USER_EXPORT", "user.export.q")
	cfg.RabbitMQQueue.OrderDirectImport = getOrDefault(v, "MQ_Q_ORDER_DIRECT_IMPORT", "order.import.direct.q")
	cfg.RabbitMQQueue.OrderImport = getOrDefault(v, "MQ_Q_ORDER_IMPORT", "order.import.q")
	cfg.RabbitMQQueue.OrderExport = getOrDefault(v, "MQ_Q_ORDER_EXPORT", "order.export.q")

	// Database
	cfg.Database.Host = getOrDefault(v, "DB_HOST", "127.0.0.1")
	cfg.Database.Port = v.GetInt("DB_PORT")
	if cfg.Database.Port == 0 {
		cfg.Database.Port = 3306
	}
	cfg.Database.User = getOrDefault(v, "DB_USERNAME", "root")
	cfg.Database.Password = getOrDefault(v, "DB_PASSWORD", "")
	cfg.Database.Name = getOrDefault(v, "DB_NAME", "mydb")

	return cfg
}

// helper
func getOrDefault(v *viper.Viper, key string, def string) string {
	val := v.GetString(key)
	if val == "" {
		return def
	}
	return val
}
