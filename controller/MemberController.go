package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"studygo2/CloudRestaurant/Service"
	"studygo2/CloudRestaurant/model"
	"studygo2/CloudRestaurant/param"
	"studygo2/CloudRestaurant/tool"
	"time"
)

type MemberController struct {
}

func (m *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", m.sendSmsCode)
	//engine.OPTIONS("/api/login_sms",m.smsLogin)
	engine.POST("/api/login_sms", m.smsLogin)

	engine.GET("api/captcha", m.captcha)
	engine.POST("/api/verifycha", m.verifycha)
	//login_pwd 用户以账号密码登录
	engine.POST("/api/login_pwd", m.namelogin)
	engine.POST("/api/upload/avator", m.uploadAvator)
	//用户信息查询
	engine.GET("api/userinfo", m.userInfo)
}
func (m *MemberController) userInfo(ctx *gin.Context) {
	cookie, err := tool.CookieAuth(ctx)
	if err != nil {
		ctx.Abort()
		tool.Failed(ctx, "还未登录，请先登录")
	}
	memberService := Service.MemberService{}
	member, err := memberService.GetUserInfo(cookie.Value)
	if err != nil {
		//返回成功信息
		tool.Sucess(ctx, map[string]interface{}{
			"id":            member.Id,
			"user_name":     member.UserName,
			"mobile":        member.Mobile,
			"register_time": member.RegisterTime,
			"avatar":        member.Avatar,
		})
		return
	}
	tool.Failed(ctx, "获取用户信息失败")
}

//http://localhost:8090/api/sendcode?phone=18805053512
func (m *MemberController) sendSmsCode(ctx *gin.Context) {
	phone, exist := ctx.GetQuery("phone")
	if exist == false {
		tool.Failed(ctx, "参数解析失败")
		return
	}

	ms := &Service.MemberService{}
	isSend := ms.SendCode(phone)
	if isSend {
		tool.Sucess(ctx, "发送成功")
	} else {
		tool.Failed(ctx, "发送失败")
	}
}

func (m *MemberController) smsLogin(ctx *gin.Context) {
	var smsLoginParam *param.SmsLoginParam
	err := tool.Decode(ctx.Request.Body, &smsLoginParam)
	if err != nil {
		tool.Failed(ctx, "参数解析失败")
		return
	}
	// fmt.Println(smsLoginParam)
	//完成手机+验证码登录
	us := Service.MemberService{}
	member := us.SmsLogin(smsLoginParam)
	if member != nil {
		sess, _ := json.Marshal(member)
		tool.SetSession(ctx, 13, 13)
		err := tool.SetSession(ctx, "userID_"+string(member.Id), sess)
		if err != nil {
			tool.Failed(ctx, "登录失败")
			return
		}
		ctx.SetCookie("cookie_id", strconv.Itoa(int(member.Id)), 10*60, "/", "localohost", true, true)
		tool.Sucess(ctx, member)
		return
	} else {
		tool.Failed(ctx, "登录失败")
	}

}
func (m *MemberController) captcha(ctx *gin.Context) {
	//todo 生成二维码返回客户端
	tool.GenerateCaptcha(ctx)
}

//验证验证码是否正确
func (m *MemberController) verifycha(ctx *gin.Context) {
	/*fmt.Println("验证码验证redis：",tool.GetSession(ctx, 1))
	fmt.Println("这个为啥不执行？因为TMD前端根本没有定义这个接口")
	var captcha tool.CaptchaResult
	err := tool.Decode(ctx.Request.Body, &captcha)
	if err != nil {
		tool.Failed(ctx,"参数解析失败")
		return
	}
	result := tool.VerifyCha(captcha.Id, captcha.VerifyValue)
	if result{
		fmt.Println("验证通过")
	}else{
		fmt.Println("验证失败")
	}*/
}

//用户名+密码+验证码登录
func (m *MemberController) namelogin(ctx *gin.Context) {

	//解析用户登录参数
	var loginParam param.LoginParam
	err := tool.Decode(ctx.Request.Body, &loginParam)
	if err != nil {
		tool.Failed(ctx, "参数解析失败")
	}
	//验证验证码
	validata := tool.VerifyCha(loginParam.Id, loginParam.Value)

	if !validata {
		tool.Failed(ctx, "验证码不正确")
		return
	}
	//登录
	ms := &Service.MemberService{}
	member, err := ms.Login(loginParam.Name, loginParam.Password)
	if err == nil {
		//保存用户信息到session
		sess, _ := json.Marshal(member)
		tool.SetSession(ctx, 13, 13)
		err := tool.SetSession(ctx, "userID_"+string(member.Id), sess)
		if err != nil {
			fmt.Println("用户名密码登录设置session失败！")
		}
		//err := tool.SetSession(ctx, "userID_5", sess)
		//	fmt.Println("user_"+string(member.Id),member.Id)
		if err != nil {
			tool.Failed(ctx, "登录失败")
			return
		}
		ctx.SetCookie("cookie_id", strconv.Itoa(int(member.Id)), 10*60, "/", "localohost", true, true)
		tool.Sucess(ctx, member)
		return
	}
	tool.Failed(ctx, "登录失败，查无此人")
}
func (m *MemberController) uploadAvator(ctx *gin.Context) {
	//解析上传的参数 file、user_id
	userid := ctx.PostForm("user_id")
	fmt.Println(userid)
	formFile, err := ctx.FormFile("avatar")
	/*if err != nil {
		fmt.Println("读取头像文件失败")
		tool.Failed(ctx,"参数解析失败")
		return
	}*/
	//判断user_id对应的用户是否已经登录 sessions
	tool.SetSession(ctx, 15, 1)
	//time.Sleep(10*time.Second)
	fmt.Println(tool.GetSession(ctx, 15))
	session := tool.GetSession(ctx, "userID_"+string(userid))
	if session == nil {
		fmt.Println("用户未登录")
		tool.Failed(ctx, "参数不合法")
		return
	}
	var member model.Member
	//fmt.Println("session:",session)
	err = json.Unmarshal(session.([]byte), &member)
	if err != nil {
		fmt.Println("json解析失败")
	}
	//将file保存到本地
	filename := "./uploadfile/" + strconv.FormatInt(time.Now().Unix(), 10) + formFile.Filename
	//3.将保存的文件本地路径，保存到用户表的头像字段
	err = ctx.SaveUploadedFile(formFile, filename)
	if err != nil {
		fmt.Println("保存头像文件失败")
		tool.Failed(ctx, "图像更新失败")
		return
	}
	memberService := Service.MemberService{}
	path := memberService.UploadAvatar(member.Id, filename[1:])
	if path != "" {
		tool.Sucess(ctx, "http://localhost:8090"+path)
		return
	}
	fmt.Println("保存头像文件到数据库失败")
	tool.Failed(ctx, "上传失败")
	//4.返回结果

}
