package dao

import (
	"fmt"
	"studygo2/CloudRestaurant/model"
	"studygo2/CloudRestaurant/tool"
)

type MemberDao struct {
	*tool.Orm
}

func (m *MemberDao) QueryMemberById(id int64) (*model.Member, error) {
	var member model.Member
	_, err := m.Orm.Where("id = ?", id).Get(&member)
	return &member, err

}

//验证手机号和验证码是否存在

func (m *MemberDao) ValidataSmsCode(phone string, code string) (*model.SmsCode, error) {
	var sms model.SmsCode
	_, err := m.Where("phone=? and code =?", phone, code).Get(&sms)
	if err != nil {
		return nil, err
	}
	return &sms, nil
}
func (m *MemberDao) QueryByPhone(phone string) *model.Member {
	var member model.Member
	_, err := m.Where("mobile= ?", phone).Get(&member)
	if err != nil {
		fmt.Println("查询用户数据库失败：", err)
	}
	return &member
}

func (m *MemberDao) InsertCode(sms *model.SmsCode) int64 {
	one, err := m.InsertOne(sms)
	if err != nil {
		fmt.Println("插入数据失败，err:", err)
		//return false
	}
	return one
}
func (m *MemberDao) InsertMember(member *model.Member) int64 {
	one, err := m.InsertOne(member)
	if err != nil {
		fmt.Println("插入数据失败，err:", err)
	}
	return one
}
func (m *MemberDao) QueryMemberByPassword(name string, password string) (*model.Member, error) {
	var member model.Member

	_, err := m.Where("user_name= ? and password= ?", name, tool.EncoderSha256(password)).Get(&member)
	if err != nil {
		fmt.Println("插入数据失败，err:", err)
	}
	return &member, err
}
func (m *MemberDao) UpdateMemberAvatar(userid int64, filename string) int64 {
	member := model.Member{Avatar: filename}
	result, err := m.Where("id= ?", userid).Update(&member)
	if err != nil {
		fmt.Println("头像插入数据库失败", err)
		return 0
	}
	return result
}
