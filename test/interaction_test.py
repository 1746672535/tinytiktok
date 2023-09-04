from interaction.comment_action import *
from interaction.comment_list import *
from interaction.favor_action import *
from interaction.favor_list import *

if __name__ == '__main__':
    # 点赞操作
    like_video_ti = test_favor_action(is_favor=1)
    print(f"点赞耗时: {like_video_ti}")

    like_video_ti = test_favor_action(is_favor=0)
    print(f"取消点赞耗时: {like_video_ti}")

    # 获取点赞列表
    get_like_video_list_ti = test_favor_list()
    print(f"获取点赞列表耗时: {get_like_video_list_ti}")

    # 评论
    comment_action_ti = test_comment_action(is_com=1)
    print(f"评论耗时: {comment_action_ti}")

    comment_action_ti = test_comment_action(is_com=0)
    print(f"取消评论耗时: {comment_action_ti}")

    # 获取评论列表
    get_comment_list_ti = test_comment_list()
    print(f"获取评论列表耗时: {get_comment_list_ti}")
