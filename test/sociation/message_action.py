import time
import requests
from config import *


# 测试登录功能
def test_message_action():
    # is_con=1时，进行评论操作，is_con=2时...
    is_con = int(input())
    params = {}
    if is_con == 1:
        params = {
            "token": user_token,
            "to_user_id": to_user_id,
            "action_type": "1",
            "content": test_content,
        }
        print(test_content)

    #if is_con == 2:


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
