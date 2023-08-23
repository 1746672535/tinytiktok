package msg

var (
	ServerRegisterError = "服务注册失败"
	ServerFindError     = "服务发现失败"
	ParameterError      = "参数错误"
	ServerError         = "服务器错误"
	JwtError            = "无法获取jwt"
	AuthError           = "鉴权失败"
	RepeatError         = "请勿重复操作"
	UnableReadConfig    = "读取配置失败"
	UnableConnectDB     = "无法连接数据库"
	UnableCreateTable   = "创建数据表失败"
)

var (
	Ok            = "ok"
	NotOk         = "not ok"
	Success int32 = 0
	Fail    int32 = 1
)
