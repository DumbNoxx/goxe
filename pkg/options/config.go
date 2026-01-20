package options

type Config struct {
	Port                int                 `json:"port"`
	IdLog               string              `json:"idLog"`
	PatternsWords       []string            `json:"pattenersWords"`
	GenerateLogsOptions OptionsGenerateLogs `json:"generateLogsOptions"`
}

type OptionsGenerateLogs struct {
	GenerateLogs bool   `json:"generateLogs"`
	Hour         string `json:"hour"`
}
