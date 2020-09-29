package oksdk

type NoProxyResult struct {
	General     General     `json:"general"`
	Risks       Risks       `json:"risks"`
	Score       Score       `json:"score"`
	Suggestions Suggestions `json:"suggestions"`
}

type General struct {
	IP       string `json:"ip"`
	Asn      int64  `json:"asn"`
	Provider string `json:"provider"`
	Country  string `json:"country"`
}

type Risks struct {
	Total    int64 `json:"total"`
	Proxy    bool  `json:"proxy"`
	Country  bool  `json:"country"`
	Asn      bool  `json:"asn"`
	Provider bool  `json:"provider"`
}

type Score struct {
	Noproxy   int64 `json:"noproxy"`
	Abuseipdb int64 `json:"abuseipdb"`
}

type Suggestions struct {
	Verify bool `json:"verify"`
	Block  bool `json:"block"`
}
