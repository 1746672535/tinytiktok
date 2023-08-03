package srv

import (
	"tinytiktok/video/proto/server"
)

type Handle struct {
	server.UnimplementedVideoServiceServer
}
