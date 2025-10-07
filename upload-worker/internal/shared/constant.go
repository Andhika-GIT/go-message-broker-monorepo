package shared

// Exchange Names
const (
	ExchangeGoApp = "go-app-exchange"
)

// Routing Keys
const (
	RoutingKeyUserDirectImport  = "user.import.direct"
	RoutingKeyUserSftpImport    = "user.import.sftp"
	RoutingKeyUserExport        = "user.export"
	RoutingKeyOrderDirectImport = "order.import.direct"
	RoutingKeyOrderImport       = "order.import"
	RoutingKeyOrderExport       = "order.export"
)

// Queue Names
const (
	QueueUserDirectImport  = "user.import.direct.q"
	QueueUserSftpImport    = "user.import.sftp.q"
	QueueUserExport        = "user.export.q"
	QueueOrderDirectImport = "order.import.direct.q"
	QueueOrderImport       = "order.import.q"
	QueueOrderExport       = "order.export.q"
)
