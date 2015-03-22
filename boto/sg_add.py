import boto
from boto.ec2 import *
from sys import argv
connection = boto.ec2.connect_to_region('ap-southeast-2')
sg_list = connection.get_all_security_groups()
sg = sg_list[4]
print sg
sg.authorize('tcp', int(argv[1]), int(argv[1]), argv[2])
#print sg.name, sg.id
#sg.authorize('tcp', argv[1], argv[1], argv[2])
#for i in sg_list:
#	print i.name
#print argv[2], type(argv[2])
#connection.authorize_security_group(group_id='sg-ced01cab', ip_protocol='tcp', from_port=int(argv[1]), 
	#to_port=int(argv[1]), src_security_group_group_id=argv[2], src_security_group_owner_id='414481387506')
