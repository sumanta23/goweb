from locust import HttpUser, TaskSet, HttpLocust, FastHttpUser, task
import json
import jwt
import time
# from gevent import pool
from geventhttpclient.client import HTTPClientPool

class UserBehavior(TaskSet):

    def __init__(self, parent):
        super(UserBehavior, self).__init__(parent)
        self.headers = {}
        self.apiprefix = ""

    def count(self):
        self.client.get(self.apiprefix+"/count")


    def on_start(self):
        # The on_start method is called
        # when a simulated user starts
        # executing that TaskSet class
        pass


    @task(1)
    def user_workflow_mymessage(self):
        self.count()


class WebsiteUser(FastHttpUser):
    """
    User class that does requests to the locust web server running on localhost,
    using the fast HTTP client
    """
    # some things you can configure on FastHttpUser
    # connection_timeout = 60.0
    # insecure = True
    # max_redirects = 5
    # max_retries = 1
    # network_timeout = 60.0

    shared_client_pool = HTTPClientPool(concurrency=1000)

    host = "http://127.0.0.1:8080"
    client_pool=shared_client_pool
    # host = "https://ipl-chat.api.engageapps.jio"


    tasks = [UserBehavior]
    min_wait = 100
    max_wait = 300
