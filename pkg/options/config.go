package options

type Config struct {
	Port                  int                   `json:"port"`
	IdLog                 string                `json:"idLog"`
	PatternsWords         []string              `json:"pattenersWords"`
	GenerateLogsOptions   GenerateLogsOptions   `json:"generateLogsOptions"`
	WebHookUrls           []string              `json:"webhookUrls"`
	BurstDetectionOptions BurstDetectionOptions `json:"bursDetectionOptions"`
}

type GenerateLogsOptions struct {
	GenerateLogsFile bool   `json:"generateLogsFile"`
	Hour             string `json:"hour"`
}

type BurstDetectionOptions struct {
	LimitBreak int `json:"limitBreak"`
}
