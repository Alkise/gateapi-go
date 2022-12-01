package gateapi

type APIV4Key struct {
	UserId      int64          `json:"user_id,omitempty"`
	Name        string         `json:"name,omitempty"`
	Perms       []APIV4KeyPerm `json:"perms,omitempty"`
	IpWhitelist []string       `json:"ip_whitelist,omitempty"`
	Key         string         `json:"key,omitempty"`
	Secret      string         `json:"secret,omitempty"`
	State       int32          `json:"state,omitempty"`
	CreatedAt   int64          `json:"created_at,omitempty"`
	UpdatedAt   int64          `json:"updated_at,omitempty"`
}

type APIV4KeyPerm struct {
	Name     string `json:"name,omitempty"`
	ReadOnly bool   `json:"read_only,omitempty"`
}
