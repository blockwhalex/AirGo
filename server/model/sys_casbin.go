package model

// Casbin info structure
type CasbinItem struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

// Casbin structure for input parameters
type CasbinInfo struct {
	RoleID      int          `json:"roleID"` // 权限id
	CasbinItems []CasbinItem `json:"casbinItems"`
}
type CasbinData struct {
	RoleID int      `json:"roleID"` // 权限id
	Data   []string `json:data`
}
