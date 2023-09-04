from sociation.message_action import *
from sociation.message_chat import *
from sociation.relation_action import *
from sociation.relation_follower_list import *
from sociation.relation_friend_list import *

if __name__ == '__main__':
    # 关注
    concern_ti = test_relation_action(1)
    print(f"关注耗时: {concern_ti}")

    # 取消关注
    concern_cancel_ti = test_relation_action(0)
    print(f"取消关注耗时: {concern_cancel_ti}")

    # 粉丝列表
    get_follower_list_ti = test_relation_follower_list()
    print(f"获取关注列表: {get_follower_list_ti}")

    # 好友列表
    get_friend_list_ti = test_relation_friend_list()
    print(f"获取好友列表: {get_friend_list_ti}")

    # 发送信息
    send_message_ti = test_message_action()
    print(f"发送信息耗时: {send_message_ti}")

    # 获取聊天记录
    get_message_list_ti = test_message_chat()
    print(f"获取聊天记录耗时: {get_message_list_ti}")
