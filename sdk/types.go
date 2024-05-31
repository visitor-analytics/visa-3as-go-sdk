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
	ID          string `json:"id"`
	ExtID       string `json:"intpWebsiteId"`
	PackageID   string `json:"packageId"`
	PackageName string `json:"packageName"`
	ResetAt     string `json:"stpResetAt"`
	CreatedAt   string `json:"createdAt"`
	ExpiresAt   string `json:"expiresAt"`
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

type DowngradeArgs struct {
	WebsiteExtID string `json:"intpWebsiteId"`
	PackageID    string `json:"packageId"`
}
