import time
import requests
from config import *

def test_favor_list():

    # 发送 GET 请求
    params = {
        "token": user_token,
        "user_id": user_id,
    }

    start_time = time.time()
    response = requests.get(favor_list_url, params=params)
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
    test_favor_list()