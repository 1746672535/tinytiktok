package randImg

import (
	"fmt"
	"github.com/awnumar/fastrand"
)

func gen() string {
	return fmt.Sprintf(`https://q.qlogo.cn/headimg_dl?dst_uin=%d&spec=640&img_type=jpg`, 800000000+fastrand.Intn(99999999))
}
