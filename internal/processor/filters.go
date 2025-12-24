package processor

// List of ignored words
var Ignored = []string{
	"healthcheck",
	"heartbeat",
	"ping",
	"pong",
	"keepalive",
	"metrics",
	"debug",
	"trace",
	"verbose",
	"request received",
	"response sent",
	"connection established",
	"connection closed",
}
