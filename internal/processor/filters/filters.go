package filters

// PatternsDate contains a list of common regular expressions for detecting dates and times in text.
//
//   - Type: []string
//   - Purpose: Used to search for date/time patterns in logs for sanitization or information extraction.
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

// PatternIpLogs is a regular expression for detecting IPv4 addresses.
//
//   - Type: string
//   - Purpose: Identifiers IP addresses in logs for potential anonymization of filtering.
var PatternIpLogs = `\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`

// PatternsIdLogs is a regular expression for detecting numeric Identifiers up to Nine digits long.
//   - Type: string
//   - Purpose: Recognizes numeric IDs in logs, useful for correlation or replacement.
var PatternsIdLogs = `\d{1,9}`

// PatternsLogLevel is a case-insentitive regular expression for detecting common log levels.
//   - Type: string
//   - Purpose: Captures keywords such as debug, info, warning, error, etc, to classify the input.
var PatternsLogLevel = `(?i)\b(debug|info|notice|warn(?:ing)?|error|critical|alert|emergency)\b`
