package msg

var (
	ParameterError = "参数错误"
	ServerError    = "服务器错误"
	JwtError       = "无法获取jwt"
	AuthError      = "鉴权失败"
	RepeatError    = "请勿重复操作"

	Ok            = "ok"
	Success int32 = 0
	Fail    int32 = 1
)
