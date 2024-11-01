// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameMenu = "colab_menu"

// Menu 菜单表
type Menu struct {
	ID            int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:datetime(3);not null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:datetime(3);not null" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
	Pid           int64          `gorm:"column:pid;type:bigint;not null;comment:父id" json:"pid"`                                  // 父id
	Title         string         `gorm:"column:title;type:varchar(100);not null;comment:名称" json:"title"`                         // 名称
	Icon          string         `gorm:"column:icon;type:varchar(100);not null;comment:菜单图标" json:"icon"`                         // 菜单图标
	URL           string         `gorm:"column:url;type:varchar(400);not null;comment:链接" json:"url"`                             // 链接
	Filepath      string         `gorm:"column:filepath;type:varchar(200);not null;comment:文件路径" json:"filepath"`                 // 文件路径
	Params        string         `gorm:"column:params;type:varchar(500);not null;comment:链接参数" json:"params"`                     // 链接参数
	Node          *string        `gorm:"column:node;type:varchar(255);not null;default:#;comment:权限节点" json:"node"`               // 权限节点
	Sort          int32          `gorm:"column:sort;type:int;not null;comment:菜单排序" json:"sort"`                                  // 菜单排序
	Status        *int32         `gorm:"column:status;type:tinyint;not null;default:1;comment:状态(0:禁用,1:启用)" json:"status"`       // 状态(0:禁用,1:启用)
	IsInner       bool           `gorm:"column:is_inner;type:tinyint(1);not null;comment:是否内页" json:"is_inner"`                   // 是否内页
	DefaultValues string         `gorm:"column:default_values;type:varchar(255);not null;comment:参数默认值" json:"default_values"`    // 参数默认值
	ShowSlider    *bool          `gorm:"column:show_slider;type:tinyint(1);not null;default:1;comment:是否显示侧栏" json:"show_slider"` // 是否显示侧栏
}

// TableName Menu's table name
func (*Menu) TableName() string {
	return TableNameMenu
}