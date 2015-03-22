import boto
import boto.ec2
from sys import argv
connection = boto.ec2.connect_to_region('ap-southeast-2')
sg_list = connection.get_all_security_groups()
for i in sg_list:
	print i.name, "\t\t\t", i.id