from faker import Faker

fake = Faker()
# 建议这里写你的测试账号的一些信息
username = "jiudan_01"
password = "jiudan_01"
user_token = (r"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9."
              r"eyJJRCI6MTAsIk5hbWUiOiJqaXVkYW5fMDEiLCJleHAiOjE3NjQ3MTg0MzIsImlzcyI6InRpbnl0aWt0b2siLCJuYmYiOjE2OTI3MTg0MzJ9."
              r"E2nq37RrZZkjFHUqfPbxkw-DmJioZE1IHlGh0xGwz6s")

# 定义url
base_url = "http://101.43.225.43:5051"
register_url = base_url + "/douyin/user/register/"
login_url = base_url + "/douyin/user/login/"
