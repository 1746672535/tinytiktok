import random
import threading
import time

from analyse import *
from basis import register, login, user, feed, pubaction, publishlist
from interaction import favor_action, favor_list, comment_action, comment_list
from sociation import relation_action, relation_follower_list, relation_friend_list, message_chat, message_action


def generate_password(length=6):
    letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"  # 包含所有英文字母的字符串
    password = ''.join(random.choice(letters) for _ in range(length))
    return password


class PartI(threading.Thread):
    def __init__(self, test_prefix, test_id):
        super().__init__()
        self.test_prefix = test_prefix
        self.test_id = test_id

    def run(self):
        # 注册用户
        data = register.register(self.test_prefix, self.test_id)
        if data is None:
            print("服务器错误")
            return
        username, password, user_id, token = data
        time.sleep(random.random() * 3)
        # 获取用户信息
        user.get_user_info(user_id, token)
        time.sleep(random.random() * 3)
        # 发布一条视频 - 用户有20%的概率发布视频 - 视频文件为 demo.mp4
        if random.random() < 0.2:
            # 发布视频
            pubaction.publish_video(token)
            # 获取发布列表
            publishlist.get_publish_list(user_id, token)
            time.sleep(random.random() * 3)


class PartII(threading.Thread):
    def __init__(self, test_prefix, test_id):
        super().__init__()
        self.test_prefix = test_prefix
        self.test_id = test_id

    def run(self):
        # 登录用户
        data = login.login(f"{self.test_prefix}-{self.test_id}", "123456")
        if data is None:
            print("服务器错误")
            return
        user_id, token = data
        # 获取视频列表
        video_list = feed.get_feed()
        if video_list is None:
            return
        # 用户会在视频列表中随机点赞视频或发表评论
        for video in video_list:
            video_id = video["id"]
            # 有100%的概率点开点赞列表
            favor_list.get_like_list(user_id, token)
            # 有100%的概率点开评论区
            comment_list.get_comment_list(video_id, token)
            # 有50%概率点赞视频
            if random.random() < 0.5:
                favor_action.like_video(video_id, token)
            # 有15%的概率发表评论
            if random.random() < 0.15:
                comment_action.comment_video(video_id, token)
            # 用户有15%的概率关注视频作者
            if random.random() < 0.15:
                author_id = video["author"]["id"]
                relation_action.concern_user(author_id, token)
            time.sleep(random.random() * 3)


class PartIII(threading.Thread):
    def __init__(self, test_prefix, test_id):
        super().__init__()
        self.test_prefix = test_prefix
        self.test_id = test_id

    def run(self):
        data = login.login(f"{self.test_prefix}-{self.test_id}", "123456")
        if data is None:
            print("服务器错误")
            return
        user_id, token = data
        # 等待其他线程完成关注操作, 回关所有粉丝, 建立好友关系, 完成接下来的测试操作
        # 获取粉丝列表
        follower_list = relation_follower_list.get_follower_list(user_id, token)
        # 如果实在没有人关注他则默认返回
        if follower_list is None:
            return
        for follower in follower_list:
            # 用户只有25%概率向朋友发送信息
            if random.random() > 0.25:
                continue
            # 随机等待
            time.sleep(random.random() * 3)
            # 回关粉丝
            follower_id = follower["id"]
            relation_action.concern_user(follower_id, token)
            # 获取好友列表
            friend_list = relation_friend_list.get_friend_list(user_id, token)
            if friend_list is None:
                return
            for friend in friend_list:
                friend_id = friend["id"]
                # 给好友发送信息1-5条
                for i in range(int(random.random() * 5) + 1):
                    message_action.send_message(friend_id, token)
                # 获取聊天记录
                message_chat.get_message_list(friend_id, token)
                time.sleep(random.random() * 3)


if __name__ == '__main__':
    test_prefix = generate_password(6)
    test_times = 100
    print(test_prefix)
    start_time = time.time()
    # 第一阶段测试
    wg = []
    for i in range(test_times):
        thread = PartI(test_prefix, i)
        wg.append(thread)
        thread.start()
    for w in wg:
        w.join()
    print("第一阶段测试结束")
    # 第二阶段测试
    wg = []
    for i in range(test_times):
        thread = PartII(test_prefix, i)
        wg.append(thread)
        thread.start()
    for w in wg:
        w.join()
    print("第二阶段测试结束")
    # 第三阶段测试
    wg = []
    for i in range(test_times):
        thread = PartIII(test_prefix, i)
        wg.append(thread)
        thread.start()
    for w in wg:
        w.join()
    print("第三阶段测试结束")
    # 测试结束
    print(f"总耗时: {time.time() - start_time}")
    # 保存队列到本地文件
    save_queue_to_file(login.login_ti_queue, 'data/login_queue.db')
    save_queue_to_file(register.register_ti_queue, 'data/register_queue.db')
    save_queue_to_file(user.get_user_info_ti_queue, 'data/get_user_info_queue.db')
    save_queue_to_file(pubaction.publish_video_ti_queue, 'data/publish_video_queue.db')
    save_queue_to_file(favor_action.like_video_ti_queue, 'data/like_video_queue.db')
    save_queue_to_file(comment_action.comment_video_ti_queue, 'data/comment_video_queue.db')
    save_queue_to_file(relation_action.concern_user_ti_queue, 'data/concern_user_queue.db')
    save_queue_to_file(message_action.send_message_ti_queue, 'data/send_message_queue.db')
    save_queue_to_file(message_chat.get_message_list_ti_queue, 'data/get_message_list_queue.db')
    # 具体分析 - 总耗时 - 平均耗时
    analyze_request_timings(login.login_ti_queue, "登录")
    analyze_request_timings(register.register_ti_queue, "注册")
    analyze_request_timings(user.get_user_info_ti_queue, "获取用户信息")
    analyze_request_timings(pubaction.publish_video_ti_queue, "发布视频")
    analyze_request_timings(favor_action.like_video_ti_queue, "点赞")
    analyze_request_timings(comment_action.comment_video_ti_queue, "评论")
    analyze_request_timings(relation_action.concern_user_ti_queue, "关注用户")
    analyze_request_timings(message_action.send_message_ti_queue, "发送信息")
    analyze_request_timings(message_chat.get_message_list_ti_queue, "获取聊天记录")
