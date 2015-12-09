import os
import smtplib
from email.mime.text import MIMEText
import boto
from boto.redshift import *

sender = "noreply@gmail.com"
to = ['mail1@gmail.com', 'mail2@gmail.com', 'mail3@gmail.com', 'mail4@gmail.com', 'mail5@gmail.com']
smtp_server = 'email-smtp.us-east-1.amazonaws.com'
smtp_username = 'KEY' #AMAZON SES ACCESS Key
smtp_password = 'SECRET' #AMAZON SES SECRET Key
smtp_port = '587'
smtp_do_tls = True


connection1 = boto.redshift.connect_to_region("us-east-1")
cluster1 = connection1.describe_resize('redshift001')
status1 = str(cluster1['DescribeResizeResponse']['DescribeResizeResult']['Status'])
body = "resize status %s" %(status1)
msg = MIMEText(body)
msg['Subject'] = 'Redshift cluster resize status'
msg['From'] = sender
msg['To'] = ", ".join(to)
s = smtplib.SMTP(
        host = smtp_server,
        port = smtp_port,
        timeout = 10
        )
s.starttls()
s.ehlo()
s.login(smtp_username, smtp_password)
s.sendmail(sender, to, msg.as_string())
s.quit()
