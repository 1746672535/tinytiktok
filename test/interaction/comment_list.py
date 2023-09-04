import requests
import time
from queue import Queue

get_comment_list_ti_queue = Queue()


def get_comment_list(video_id, token):
    params = {
        "token": token,
        "video_id": video_id,
    }
    start_ti = time.time()
    response = requests.get(comment_list_url, params=params)
    end_ti = time.time()
    if response.json()["status_code"] == 0:
        get_comment_list_ti_queue.put(end_ti - start_ti)


def test_comment_list():
    # 发送 GET 请求
    params = {
        "token": user_token,
        "video_id": video_id,
    }

    start_time = time.time()
    response = requests.get(comment_list_url, params=params)
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
    test_comment_list()
