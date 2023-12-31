import requests
import time
from basis.test_publish import *
from queue import Queue

publish_video_ti_queue = Queue()


def publish_video(token):
    files = {
        'data': ('test_video', test_video, 'application/octet-stream')
    }
    payloads = {
        "token": token,
        "title": fake.text(max_nb_chars=16),
    }
    start_ti = time.time()
    response = requests.post(pub_action_url, data=payloads, files=files)
    end_ti = time.time()
    if response.json()["status_code"] == 0:
        publish_video_ti_queue.put(end_ti - start_ti)
        print(f"用户发布视频成功")


def test_pub_action():
    # 请求参数
    files = {
        'data': ('test_video', test_video, 'application/octet-stream')
    }
    payloads = {
        "token": user_token,
        "title": test_title,
    }

    # 发送 POST 请求
    start_time = time.time()
    response = requests.post(pub_action_url, data=payloads, files=files)
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
    test_pub_action()
