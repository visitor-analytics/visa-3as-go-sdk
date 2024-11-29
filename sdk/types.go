package sdk

type NewWebsiteArgs struct {
	ExtID       string  `json:"intpWebsiteId"`
	Domain      string  `json:"domain"`
	PackageID   string  `json:"packageId"`
	BillingDate *string `json:"billingDate,omitempty"`
}

type NewSSRWebsiteArgs struct {
	ExtID string `json:"externalWebsiteId"`
	Name  string `json:"name"`
}

type SSRWebsiteSettings struct {
	Paused          bool `json:"paused"`
	AnyPage         bool `json:"anyPage"`
	ClickAndScroll  bool `json:"clickAndScroll"`
	TextObfuscation bool `json:"textObfuscation"`

	Pages        []string        `json:"pages"`
	URLPatterns  []SSRUrlPattern `json:"UrlPatterns"`
	MinDuration  int64           `json:"minDuration"`
	DynamicPages []string        `json:"dynamicPages"`
}

type SSRUrlPattern struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}

type NewIntpcArgs struct {
	ExtID   string         `json:"intpCustomerId"`
	Email   string         `json:"email"`
	Website NewWebsiteArgs `json:"website"`
}

type Intpc struct {
	ID        string `json:"id"`
	ExtID     string `json:"intpCustomerId"`
	VisaID    string `json:"visaId"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}

type Website struct {
	ID                     string  `json:"id"`
	ExtID                  string  `json:"intpWebsiteId"`
	IntpID                 string  `json:"intpId"`
	IntpCustomerID         string  `json:"intpCustomerId"`
	VisaCustomerID         string  `json:"visaCustomerId"`
	Status                 string  `json:"status"`
	Domain                 string  `json:"domain"`
	PackageID              string  `json:"packageId"`
	PackageName            string  `json:"packageName"`
	ResetAt                string  `json:"stpResetAt"`
	CreatedAt              string  `json:"createdAt"`
	ExpiresAt              string  `json:"expiresAt"`
	TrackingCode           *string `json:"visaTrackingCode"`
	TrackingCodeMaxPrivacy *string `json:"visaMaxPrivacyModeTrackingCode"`
	LastPackageChangeAt    *string `json:"lastPackageChangeAt"`
	PDPackageID            *string `json:"plannedDowngradePackageId"`
	PDPackageName          *string `json:"plannedDowngradePackageName"`
	PDPackageCycle         *string `json:"plannedDowngradeBillingInterval"`
}

type Subscription struct {
	WebsiteExtID string `json:"intpWebsiteId"`
	PackageID    string `json:"packageId"`
}

type UpgradeArgs struct {
	WebsiteExtID string `json:"intpWebsiteId"`
	PackageID    string `json:"packageId"`
	Trial        bool   `json:"trial"`
	ProRate      bool   `json:"proRate"`
}

type ReactivateArgs struct {
	WebsiteExtID string `json:"intpWebsiteId"`
	PackageID    string `json:"packageId"`
	Trial        bool   `json:"trial"`
}

type DowngradeArgs struct {
	WebsiteExtID string `json:"intpWebsiteId"`
	PackageID    string `json:"packageId"`
}

type APIErr struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Title   string `json:"title,omitempty"`
	Time    string `json:"time,omitempty"`
	Code    int    `json:"code,omitempty"`
}
