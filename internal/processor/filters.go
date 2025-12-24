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

var PatternsDate = []string{
	`\d{2}/\d{2}/\d{4}`,
	`\d{2}-\d{2}-\d{4}`,
	`\d{4}/\d{2}/\d{2}`,
	`\d{4}-\d{2}-\d{2}`,
	`\d{2}/\d{4}/\d{2}`,
	`\d{2}-\d{4}-\d{2}`,
	`\d{4}/\d{2}`,
	`\d{4}-\d{2}`,
	`\d{2}/\d{4}`,
	`\d{2}-\d{4}`,
	`\d{2}:\d{2}:\d{2}`,
}

var PatternsLogLevel = `(?i)\b(debug|info|notice|warn(?:ing)?|error|critical|alert|emergency)\b`
