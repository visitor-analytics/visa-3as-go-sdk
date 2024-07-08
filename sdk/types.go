package sdk

type NewWebsiteArgs struct {
	ExtID     string `json:"intpWebsiteId"`
	Domain    string `json:"domain"`
	PackageID string `json:"packageId"`
}

type NewCustomerArgs struct {
	ExtID   string         `json:"intpCustomerId"`
	Email   string         `json:"email"`
	Website NewWebsiteArgs `json:"website"`
}

type Customer struct {
	ID        string `json:"id"`
	ExtID     string `json:"intpCustomerId"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}

type Website struct {
	ID                     string  `json:"id"`
	ExtID                  string  `json:"intpWebsiteId"`
	Status                 string  `json:"status"`
	PackageID              string  `json:"packageId"`
	PackageName            string  `json:"packageName"`
	ResetAt                string  `json:"stpResetAt"`
	CreatedAt              string  `json:"createdAt"`
	ExpiresAt              string  `json:"expiresAt"`
	TrackingCode           *string `json:"visaTrackingCode,omitempty"`
	TrackingCodeMaxPrivacy *string `json:"visaMaxPrivacyModeTrackingCode,omitempty"`
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
