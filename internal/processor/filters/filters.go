package filters

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
	`[a-z]+\s\d{1,2},\s\d{4}`,
}

var PatternIpLogs = `\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`

var PatternsIdLogs = `\d{1,9}`

var PatternsLogLevel = `(?i)\b(debug|info|notice|warn(?:ing)?|error|critical|alert|emergency)\b`
