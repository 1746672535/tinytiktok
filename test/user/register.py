import time

import requests

from config import *


# 测试注册功能
def test_register():
    # 请求参数
    params = {
        # 第一次生成一组自己用的测试账号
        # "username": username,
        # "password": password,
        # 之后的测试请自动生成
        "username": fake.user_name()[:32],
        "password": fake.password()[:32],
    }
    # 发送 POST 请求
    start_time = time.time()
    response = requests.post(register_url, params=params)
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
        # 用户令牌在某些服务中是必须的, 请在config.py下手动修改token的值
        print(data["token"])
        return end_time - start_time
    else:
        print("请求失败")
        return -10086


if __name__ == '__main__':
    test_register()
