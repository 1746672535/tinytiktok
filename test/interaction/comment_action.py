import requests
import time
from queue import Queue

comment_video_ti_queue = Queue()


def comment_video(video_id, token):
    comment = fake.text(max_nb_chars=50)
    params = {
        "token": token,
        "video_id": video_id,
        "action_type": "1",
        "comment_text": comment,
    }
    start_ti = time.time()
    response = requests.post(comment_action_url, params=params)
    end_ti = time.time()
    if response.json()["status_code"] == 0:
        comment_video_ti_queue.put(end_ti - start_ti)


def test_comment_action(is_com=1):
    params = {}
    if is_com == 1:
        params = {
            "token": user_token,
            "video_id": video_id,
            "action_type": "1",
            "comment_text": test_comment,
        }
        print(test_comment)

    if is_com == 0:
        params = {
            "token": user_token,
            "video_id": video_id,
            "action_type": "2",
            "comment_id": comment_id,  # 根据评论ID可修改
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
