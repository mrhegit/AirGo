package initialize

import (
	"AirGo/global"
	"AirGo/model"
)

func InitializeAll() {
	InitLogrus()            //logrus
	global.VP = InitViper() //初始化Viper
	global.DB = Gorm()      //gorm连接数据库
	if global.DB != nil {
		if !global.DB.Migrator().HasTable(&model.User{}) {
			global.Logrus.Info("未找到sys_user库表,开始建表并初始化数据...")
			InitCasbin()          //加载casbin,生成casbin_rule 表
			RegisterTables()      //创建table
			InsertInto(global.DB) //导入数据
			InitCasbin()          //重新加载casbin
		} else {
			RegisterTables() //AutoMigrate 自动迁移 schema
			InitCasbin()     //加载casbin
		}
	}
	InitServer()        //全局系统配置
	InitTheme()         //全局主题
	InitLocalCache()    //local cache
	InitBase64Captcha() //Base64Captcha
	InitCrontab()       //定时任务
	InitAlipayClient()  //alipay
	InitEmailDialer()   //gomail Dialer
	InitWebsocket()     //websocket
	InitRatelimit()     //限流
	InitRouter()        //初始总路由
}
