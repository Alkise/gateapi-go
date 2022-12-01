package gateapi

type SubAccount struct {
	Remark     string `json:"remark,omitempty"`
	LoginName  string `json:"login_name"`
	Password   string `json:"password,omitempty"`
	Email      string `json:"email,omitempty"`
	State      int32  `json:"state,omitempty"`
	UserId     int64  `json:"user_id,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
}
