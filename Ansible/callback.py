import os
import smtplib
from email.mime.text import MIMEText

from ansible import utils
from prettytable import PrettyTable

try:
    import prettytable
    HAS_PRETTYTABLE = True
except ImportError:
    HAS_PRETTYTABLE = False


class CallbackModule(object):
    def __init__(self):
    self.sender = "jaibheemsen@gmail.com"
    self.to = "jaibheemsen@gmail.com"
    self.playbook_name = None
    self.body = " "
    self.table = PrettyTable([ " " ])

    def send_mail(self, table, notify=False):
    self.body = table
    msg = MIMEText(self.body)
    msg['Subject'] = 'List of hosts with chattr on /etc/sudoers file'
    msg['From'] = self.sender
    msg['To'] = self.to
    s = smtplib.SMTP('localhost')
        s.sendmail(self.sender, self.to, msg.as_string())
        s.quit()

    def playbook_on_stats(self, stats):
        hosts = sorted(stats.processed.keys())
        t = prettytable.PrettyTable(['Host', 'Unreachable', 'Failures'])
        failures = False
        unreachable = False
        for h in hosts:
            s = stats.summarize(h)
            if s['failures'] > 0:
                failures = True
        t.add_row([h] + [s[k] for k in ['unreachable', 'failures']])
            if s['unreachable'] > 0:
                unreachable = True
        t.add_row([h] + [s[k] for k in ['unreachable', 'failures']])
            #t.add_row([h] + [s[k] for k in ['unreachable', 'failures']])
    t = unicode(t)
        if failures or unreachable:
            self.send_mail(t, notify=True)
