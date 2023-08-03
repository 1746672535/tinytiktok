package srv

import (
	"tinytiktok/user/proto/server"
)

type Handle struct {
	server.UnimplementedUserServiceServer
}
