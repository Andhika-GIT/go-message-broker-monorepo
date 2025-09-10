package shared

// Exchange Names
const (
	ExchangeGoApp = "go-app-exchange"
)

// Routing Keys
const (
	RoutingKeyUserDirectImport = "user.import.direct"
	RoutingKeyUserSftpImport   = "user.import.sftp"
	RoutingKeyUserExport       = "user.export"
	RoutingKeyOrderImport      = "order.import"
	RoutingKeyOrderExport      = "order.export"
)

// Queue Names
const (
	QueueUserDirectImport = "user.import.direct.q"
	QueueUserSftpImport   = "user.import.sftp.q"
	QueueUserExport       = "user.export.q"
	QueueOrderImport      = "order.import.q"
	QueueOrderExport      = "order.export.q"
)
