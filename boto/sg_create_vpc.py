import boto
import boto.vpc
from sys import argv
connection = boto.vpc.connect_to_region('us-east-1')
groups = connection.get_all_security_groups()
for group in groups:
	if group.vpc_id == argv[1]:
		print group.id, "\t", group.name
#create = connection.create_security_group('testing1', 'testing1', vpc_id='vpc-123456')
#print create, create.id, create.name
#create.authorize('tcp', 80, 80, '0.0.0.0/0')
#create.authorize('tcp', 443, 443, '0.0.0.0/0')


#connection.authorize_security_group(group_id=create.id, ip_protocol='tcp', from_port=22, 
#	to_port=22, src_security_group_name='testing')

