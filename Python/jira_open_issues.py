import os
import json
import requests
from requests.auth import HTTPBasicAuth
user_id = os.getenv('JIRA_USER')
user_password = os.getenv('JIRA_PASS')
jira_url = os.getenv('JIRA_URL')
r = requests.get(jira_url + "/rest/api/2/search?jql=assignee="+os.getenv('USERVARIABLE1')+"\\u0040apigee.com%20AND%20status%3DInvestigate", auth=HTTPBasicAuth(user_id, user_password)).json()
for i in r['issues']:
        print i['key']
