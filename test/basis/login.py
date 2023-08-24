import time

import requests

from config import *


# 测试登录功能
def test_login():
    # 请求参数
    params = {
        "username": username,
        "password": password,
    }
    # 发送 POST 请求
    start_time = time.time()
    response = requests.post(login_url, params=params)
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
    # 测试登录
    test_login()
