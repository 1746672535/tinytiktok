from faker import Faker

fake = Faker()
# 建议这里写你的测试账号的一些信息
username = "jiudan_01"
password = "jiudan_01"
user_token = r"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MTAsIk5hbWUiOiJqaXVkYW5fMDEiLCJleHAiOjE3NjQ3MTg0MzIsImlzcyI6InRpbnl0aWt0b2siLCJuYmYiOjE2OTI3MTg0MzJ9.E2nq37RrZZkjFHUqfPbxkw-DmJioZE1IHlGh0xGwz6s"
user_id = 6

# 定义url
base_url = "http://101.43.225.43:5051"
register_url = base_url + "/douyin/user/register/"
login_url = base_url + "/douyin/user/login/"
user_url = base_url + "/douyin/user/"
feed_url = base_url + "/douyin/feed/"
pub_list_url = base_url + "/douyin/publish/list/"
pub_action_url = base_url + "/douyin/publish/action/"
favor_action_url = base_url + "/douyin/favorite/action/"
favor_list_url = base_url + "/douyin/favorite/list/"
comment_action_url = base_url + "/douyin/comment/action/"
comment_list_url = base_url + "/douyin/comment/list/"

#记录评论ID：
comment_id = []