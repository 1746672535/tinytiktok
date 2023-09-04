import time
from queue import Queue

import requests

from config import *

send_message_ti_queue = Queue()


def send_message(friend_id, token):
    params = {
        "token": token,
        "to_user_id": friend_id,
        "action_type": "1",
        "content": fake.text(max_nb_chars=40),
    }
    start_ti = time.time()
    response = requests.post(message_action_url, params=params)
    end_ti = time.time()
    if response.json()["status_code"] == 0:
        send_message_ti_queue.put(end_ti - start_ti)
        print(f"用户向好友发送信息成功")


def test_message_action():
    params = {
        "token": user_token,
        "to_user_id": to_user_id,
        "action_type": "1",
        "content": test_content,
    }
    print(test_content)

    # 发送 POST 请求
    start_time = time.time()
    response = requests.post(message_action_url, params=params)
    end_time = time.time()
    print(response.status_code)
    # 检查响应状态码
    if response.status_code == 200:
        # 打印返回信息
        print(response.json())
        # 解析响应数据
        data = response.json()
        # 如果 status_code = 0, 则表示该接口没有问题
        if data["status_code"] == 0:
            print("true")
        return end_time - start_time
    else:
        print("请求失败")
        return -10086


if __name__ == '__main__':
    test_message_action()
