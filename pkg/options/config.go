package options

type Config struct {
	Port                  int                   `json:"port"`
	IdLog                 string                `json:"idLog"`
	PatternsWords         []string              `json:"pattenersWords"`
	GenerateLogsOptions   GenerateLogsOptions   `json:"generateLogsOptions"`
	WebHookUrls           []string              `json:"webhookUrls"`
	BurstDetectionOptions BurstDetectionOptions `json:"bursDetectionOptions"`
	ShipperConfig         ShipperConfig         `json:"shipper"`
	BufferUdpSize         int                   `json:"bufferUdpSize"`
}

type GenerateLogsOptions struct {
	GenerateLogsFile bool   `json:"generateLogsFile"`
	Hour             string `json:"hour"`
}

type BurstDetectionOptions struct {
	LimitBreak float64 `json:"limitBreak"`
}

type ShipperConfig struct {
	Address       string `json:"address"`
	FlushInterval int    `json:"flushInterval"`
	Protocol      string `json:"protocol"`
}
