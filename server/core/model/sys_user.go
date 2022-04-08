package model

type SysUser struct {
	GPaasModel
	Name     string
	NickName string
	Password string
	Phone    string
	Email    string
}

func (entity *SysUser) TableName() string {
	return "sys_user"
}
