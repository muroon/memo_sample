package infra

import (
	"database/sql"
)

// InfraInfo インフラ情報
type InfraInfo struct {
	DB *sql.DB
} 

// NewInfraInfo インフラ情報を渡す
func NewInfraInfo() *InfraInfo {
	return &InfraInfo {DB: db}
}