package model

// Logger ...
type Logger struct {
	AuditLog string `json:"audit_log"`
	Message  string `json:"message"`
	UUID     string `json:"uuid"`
}
