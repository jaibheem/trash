import os
import smtplib
from email.mime.text import MIMEText
import boto
from boto.redshift import *
from boto.ses import SESConnection


sender = "jaibheemsen@gmail.com"
to = "jaibheemsen@gmail.com"
smtp_server = 'email-smtp.us-east-1.amazonaws.com'
smtp_username = opts.user
smtp_password = opts.passw
smtp_port = '587'

connection1 = boto.redshift.connect_to_region("us-east-1")
cluster1 = connection1.describe_cluster('redshift001')
status1 = cluster1['DescribeClustersResponse']['DescribeClustersResult']['Clusters']
for item in status1:
	restore_status = item['RestoreStatus']['Status']

body = "%s resize Status %s \n %s resize status %s" %(restore_status)
msg = MIMEText(body)
msg['Subject'] = 'Redshift cluster resize status'
msg['From'] = sender
msg['To'] = to
s = = smtplib.SMTP(
		host = smtp_server,
		port = smtp_port,
		timeout = 10 
		)
s.login(smtp_username, smtp_password)
s.sendmail(sender, to, msg.as_string())
s.quit()
