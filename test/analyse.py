import os
import shelve
from queue import Queue

from matplotlib import pyplot as plt

plt.rcParams["font.sans-serif"] = ["SimHei"]
plt.rcParams["axes.unicode_minus"] = False


def save_queue_to_file(queue_object, filename):
    queue_list = list(queue_object.queue)  # 转换队列为列表
    with shelve.open(filename) as db:
        db['queue'] = queue_list


def get_queue_to_file(filename):
    queue_object = Queue()
    with shelve.open(filename) as db:
        queue_list = db['queue']
        for item in queue_list:
            queue_object.put(item)
        return queue_object


def get_figure_size(num_x_ticks):
    tick_width = 0.02  # 每个刻度的宽度
    min_width = 6.0  # 图表的最小宽度

    # 计算图表的宽度
    width = max(num_x_ticks * tick_width, min_width)

    # 计算图表的高度
    height = width * 0.75

    # 返回图表尺寸
    return width, height


def analyze_request_timings(request_timings_queue, title):
    # 转换队列为列表
    request_timings = list(request_timings_queue.queue)
    # 仅统计1.5秒以下的请求（除了"发布视频"）
    if title != "发布视频":
        request_timings = [i for i in request_timings if i < 1.5]
    # 输出信息
    total_time = sum(request_timings)  # 计算总耗时
    average_time = total_time / len(request_timings)  # 计算平均耗时
    # 分析其他指标 最大耗时、最小耗时
    max_time = max(request_timings)
    min_time = min(request_timings)
    # 绘制柱状图 - 将时间转换为毫秒
    plt.figure(figsize=get_figure_size(len(request_timings)))
    plt.bar(range(len(request_timings)), [i * 1000 for i in request_timings])
    plt.xlabel('索引')
    plt.ylabel('请求时间/ms')
    plt.title(title)

    # 保存图表
    save_path = os.path.join('image', f'{title}.png')
    plt.savefig(save_path)

    print(f'''
--------------------------------------------------------------
title: {title}
min_time: {min_time}
max_time: {max_time}
times: {len(request_timings)}
total_time: {total_time}
average_time: {average_time}
--------------------------------------------------------------''')


if __name__ == '__main__':
    login_ti_queue = get_queue_to_file('data/login_queue.db')
    register_ti_queue = get_queue_to_file('data/register_queue.db')
    get_user_info_ti_queue = get_queue_to_file('data/get_user_info_queue.db')
    publish_video_ti_queue = get_queue_to_file('data/publish_video_queue.db')
    like_video_ti_queue = get_queue_to_file('data/like_video_queue.db')
    comment_video_ti_queue = get_queue_to_file('data/comment_video_queue.db')
    concern_user_ti_queue = get_queue_to_file('data/concern_user_queue.db')
    send_message_ti_queue = get_queue_to_file('data/send_message_queue.db')
    get_message_list_ti_queue = get_queue_to_file('data/get_message_list_queue.db')

    # 具体分析 - 总耗时 - 平均耗时
    analyze_request_timings(login_ti_queue, "登录")
    analyze_request_timings(register_ti_queue, "注册")
    analyze_request_timings(get_user_info_ti_queue, "获取用户信息")
    analyze_request_timings(publish_video_ti_queue, "发布视频")
    analyze_request_timings(like_video_ti_queue, "点赞")
    analyze_request_timings(comment_video_ti_queue, "评论")
    analyze_request_timings(concern_user_ti_queue, "关注用户")
    analyze_request_timings(send_message_ti_queue, "发送信息")
    analyze_request_timings(get_message_list_ti_queue, "获取聊天记录")
