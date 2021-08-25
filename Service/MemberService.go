package Service

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"math/rand"
	"strconv"
	dao2 "studygo2/CloudRestaurant/dao"
	"studygo2/CloudRestaurant/model"
	"studygo2/CloudRestaurant/param"
	"studygo2/CloudRestaurant/tool"
	"time"
)

type MemberService struct {
}

//根据用户ID查询
func (ms *MemberService) GetUserInfo(userId string) (*model.Member, error) {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}
	memberDao := dao2.MemberDao{tool.DbEngine}
	return memberDao.QueryMemberById(int64(id))
}

//登录
func (ms *MemberService) Login(name string, password string) (*model.Member, error) {
	//1.查询用户是否存在
	md := &dao2.MemberDao{tool.DbEngine}
	member, err := md.QueryMemberByPassword(name, password)
	if err == nil {
		if member.Id != 0 {
			return member, nil
		}
	}

	//return member,nil
	//不存在即注册一个新用户
	user := &model.Member{
		UserName:     name,
		Password:     tool.EncoderSha256(password),
		RegisterTime: time.Now().Unix(),
	}
	result := md.InsertMember(user)
	user.Id = result

	return user, nil

}
func (ms *MemberService) SmsLogin(loginParam *param.SmsLoginParam) *model.Member {
	//1.获取手机号和验证码

	//2.验证手机号加验证码是否正确
	md := &dao2.MemberDao{tool.DbEngine}
	smsCode, err := md.ValidataSmsCode(loginParam.Phone, loginParam.Code)
	if err != nil {
		return nil
	}
	if smsCode.Id == 0 {
		return nil
	}
	//3.根据手机号member表查询记录
	member := md.QueryByPhone(loginParam.Phone)
	if member.Id != 0 {
		return member
	}

	//4新创建一个member记录
	user := &model.Member{}
	user.UserName = loginParam.Phone
	user.Mobile = loginParam.Phone
	user.RegisterTime = time.Now().Unix()
	md.InsertMember(user)
	//fmt.Println(user)
	return user
}
func (ms *MemberService) SendCode(phone string) bool {
	//产生验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).
		Int31n(10000))

	//测试数据库插入
	smscode := &model.SmsCode{
		Phone:      phone,
		Code:       code,
		BizId:      "123",
		CreateTime: time.Now().Unix(),
	}
	dao := &dao2.MemberDao{tool.DbEngine}
	_, err := dao.InsertOne(smscode)
	if err != nil {
		fmt.Println("insert sms request fail,err", err, "result:")
	}
	return true
	//调用阿里云sdk完成发送

	/*  config:=tool.GetCofig()
		client,err:=CreateClient(&config.Sms.AppKey,&config.Sms.AppSecret)
		if err!=nil{
	        fmt.Println("create client failed")
		}
		request:=&dysmsapi20170525.SendSmsRequest{
			PhoneNumbers: tea.String("18805053512"),
	       SignName: tea.String(config.Sms.SignName),
	       TemplateCode: tea.String(config.Sms.TemplateCode),
	       TemplateParam: tea.String(code),
		}
		response, err := client.SendSms(request)
		fmt.Println(response)
	    if err!=nil{
	    	fmt.Println("send sms request fail,err",err)
	    	return false
		}

		if *response.Body.Code=="OK"{
			smscode:=&model.SmsCode{
				Phone: phone,
				Code: code,
				BizId: *response.Body.BizId,
				CreateTime: time.Now().Unix(),
			}
	       dao:=&dao2.MemberDao{tool.DbEngine}
			result, err := dao.InsertOne(smscode)
			if err!=nil{
				fmt.Println("insert sms request fail,err",err,"result:",request)
				return false
			}
			return result>0
			//return true
		}*/
	return false
}
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}
func (ms *MemberService) UploadAvatar(userId int64, filename string) string {
	dao := dao2.MemberDao{tool.DbEngine}
	result := dao.UpdateMemberAvatar(userId, filename)
	if result == 0 {
		return ""
	}
	return filename
}
