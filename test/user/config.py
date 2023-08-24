from faker import Faker

fake = Faker()
# 建议这里写你的测试账号的一些信息
username = "jiudan_01"
password = "jiudan_01"
user_token = r"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MTAsIk5hbWUiOiJqaXVkYW5fMDEiLCJleHAiOjE3NjQ3MTg0MzIsImlzcyI6InRpbnl0aWt0b2siLCJuYmYiOjE2OTI3MTg0MzJ9.E2nq37RrZZkjFHUqfPbxkw-DmJioZE1IHlGh0xGwz6s"
user_id = 6

to_username = "hholmes"
#to_password = []
to_user_token = r"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MTEsIk5hbWUiOiJoaG9sbWVzIiwiZXhwIjoxNzY0ODY1MTA1LCJpc3MiOiJ0aW55dGlrdG9rIiwibmJmIjoxNjkyODY1MTA1fQ.rfFG084Z8qUBoFmhEGYhuDcHGi31lCpP9XwAhe0mE_4"
to_user_id = 11

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
relation_action_url = base_url + "/douyin/relation/action/"
relation_follow_list_url = base_url + "/douyin/relation/follow/list/"
relation_follower_list_url = base_url + "/douyin/relation/follower/list/"
relation_friend_list_url = base_url + "/douyin/relation/friend/list/"
message_chat_url = base_url + "/douyin/message/chat/"
message_action_url = base_url + "/douyin/message/action/"

# 测试用评论：
test_comment = fake.text(max_nb_chars=25)
test_content = fake.text(max_nb_chars=40)
comment_id = 10