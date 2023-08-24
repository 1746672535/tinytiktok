import time
import requests
from config import *
from test_publish import *


# 测试登录功能
def test_comment_action():
    # is_com=1时，进行评论操作，is_com=2时，进行取消评论操作
    is_com = int(input())
    params = {}
    if is_com == 1:
        params = {
            "token": user_token,
            "video_id": video_id,
            "action_type": "1",
            "comment_text": test_comment,
        }
        print(test_comment)

    if is_com == 2:
        params = {
            "token": user_token,
            "video_id": video_id,
            "action_type": "2",
            "comment_id": comment_id,    #根据评论ID可修改
        }

    # 发送 POST 请求
    start_time = time.time()
    response = requests.post(comment_action_url, params=params)
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
    test_comment_action()
