package view

import (
	"encoding/json"
	"time"
)

// TUserVo 用户表 视图对象
type TUserVo struct {
	Id         uint64    `json:"id,omitempty"`         // 主键ID
	Username   string    `json:"username,omitempty"`   // 用户名
	Email      string    `json:"email,omitempty"`      // 邮箱
	Password   string    `json:"password,omitempty"`   // 密码（哈希）
	Role       string    `json:"role,omitempty"`       // 角色: user | admin
	CreateTime time.Time `json:"createTime,omitempty"` // 创建时间
	UpdateTime time.Time `json:"updateTime,omitempty"` // 更新时间
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *TUserVo) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *TUserVo) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}
