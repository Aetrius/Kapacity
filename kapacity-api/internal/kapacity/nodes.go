package kapacity

import (
	"time"

	"github.com/google/uuid"
)

// Node structure for output
// Use snake_case for gorm (Go Object Relational Management)
type Node struct {
	NodeID       *uuid.UUID `csv:"ClusterID" gorm:"type:uuid;column:cluster_id;primary_key"`
	ClusterID    *uuid.UUID `csv:"ClusterID" gorm:"type:uuid;column:cluster_id;foreign_key"`
	Name         string     `csv:"Name" gorm:"type:varchar(150);column:name;default:null"`
	MemoryValue  *time.Time `csv:"MemoryValue" gorm:"type:varchar(12);precision:6;column:memory_value;default:null"`
	MemoryLimit  *time.Time `csv:"MemoryLimit" gorm:"type:varchar(12);precision:6;column:memory_limit;default:null"`
	DiskValue    string     `csv:"DiskValue" gorm:"type:varchar(12);column:disk_value"`
	DiskLimit    string     `csv:"DiskLimit" gorm:"type:varchar(12);column:disk_limit"`
	CPUValue     string     `csv:"CPUValue" gorm:"type:varchar(12);column:cpu_value"`
	CPULimit     string     `csv:"CPULimit" gorm:"type:varchar(12);column:cpu_limit"`
	Location     string     `csv:"Location" gorm:"type:varchar(12);column:location"`
	NetworkSpeed string     `csv:"NetworkSpeed" gorm:"type:varchar(12);column:network_speed"`
	UpdatedAt    *time.Time `csv:"UpdatedAt" gorm:"type:timestamp;precision:6;column:updated_at;default:null"`
	ResolvedAt   *time.Time `csv:"ResolvedAt" gorm:"type:timestamp;precision:6;column:resolved_at;default:null"`
}
