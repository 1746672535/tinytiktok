from faker import Faker

fake = Faker()
# 建议这里写你的测试账号的一些信息
username = "demo_01"
password = "demo_01"
user_token = r"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MzQsIk5hbWUiOiJkZW1vXzAxIiwiZXhwIjoxNzY1ODA3NDUyLCJpc3MiOiJ0aW55dGlrdG9rIiwibmJmIjoxNjkzODA3NDUyfQ.aPxR_KJPH9FxY_ewme7aBsnzMzzqXQ9jZUilvb0pBj0"
user_id = 34

to_username = "demo_01_friend"
to_password = "demo_01_friend"
to_user_token = r"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MzUsIk5hbWUiOiJkZW1vXzAxX2ZyaWVuZCIsImV4cCI6MTc2NTgwNzU5NywiaXNzIjoidGlueXRpa3RvayIsIm5iZiI6MTY5MzgwNzU5N30"
to_user_id = 35

# 视频ID
video_id = 21

# 测试用评论：
test_comment = fake.text(max_nb_chars=25)
test_content = fake.text(max_nb_chars=40)
comment_id = 25

# 定义url
base_url = "http://tiktok.jiudan.ltd:5051"
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
