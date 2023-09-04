from basis.feed import *
from basis.login import *
from basis.pubaction import *
from basis.publishlist import *
from basis.register import *
from basis.user import *

if __name__ == '__main__':
    # 注册
    ti = test_register()
    print(f"注册耗时: {ti}")

    # 登录
    login_ti = test_login()
    print(f"登录耗时: {login_ti}")

    # 获取用户信息
    get_user_info_ti = test_user()
    print(f"获取用户信息耗时: {get_user_info_ti}")

    # 获取视频流
    get_feed_list_ti = test_feed()
    print(f"获取视频流列表耗时: {get_feed_list_ti}")

    # 获取发布列表
    get_pub_list_ti = test_pub_list()
    print(f"获取发布列表耗时: {get_pub_list_ti}")

    # 投稿视频
    post_video_ti = test_pub_action()
    print(f"发布视频: {get_pub_list_ti}")
