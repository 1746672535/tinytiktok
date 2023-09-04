import time
from queue import Queue

import requests

from config import *

get_feed_ti_queue = Queue()


def get_feed():
    params = {
        "latest_time": int(time.time()) * 1000,
        "token": user_token,
    }
    response = requests.get(feed_url, params=params)
    start_ti = time.time()
    video_list = response.json()["video_list"]
    end_ti = time.time()
    if response.json()["status_code"] == 0:
        get_feed_ti_queue.put(end_ti - start_ti)
        return video_list


def test_feed():
    # 发送 GET 请求
    start_time = time.time()
    params = {
        "latest_time": int(start_time) * 1000,
        "token": user_token,
    }
    response = requests.get(feed_url, params=params)
    print(response.json())
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
        # print(data["token"])
        return end_time - start_time
    else:
        print("请求失败")
        return -10086


if __name__ == '__main__':
    # test_feed()
    get_feed()
