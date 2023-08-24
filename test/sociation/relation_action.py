import time
import requests
from test_publish import *
from config import *


# 测试登录功能
def test_relation_action():
    # is_relation=1时，进行关注操作，is_relation=2时，进行取消关注操作
    is_relation = int(input())
    if is_relation == 1:
        params = {
            "token": user_token,
            "to_user_id": video_id,
            "action_type": "1",
        }
    if is_relation == 2:
        params = {
            "token": user_token,
            "to_user_id": video_id,
            "action_type": "2",
        }

    # 发送 POST 请求
    start_time = time.time()
    response = requests.post(relation_action_url, params=params)
    end_time = time.time()
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
    test_relation_action()
