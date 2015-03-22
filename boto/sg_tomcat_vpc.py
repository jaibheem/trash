import boto
import boto.vpc
from sys import argv
connection = boto.vpc.connect_to_region('ap-southeast-2')
create = connection.create_security_group('tomcat', 'tomcat', vpc_id='vpc-123456')
print create, create.id, create.name

connection.authorize_security_group(group_id=create.id, ip_protocol='tcp', from_port=22, 
	to_port=22, src_security_group_group_id='sg-12345', src_security_group_owner_id='123456')
connection.authorize_security_group(group_id=create.id, ip_protocol='tcp', from_port=8080, 
	to_port=8080, src_security_group_group_id='sg-12345', src_security_group_owner_id='123456')
	to_port=8050, src_security_group_group_id='sg-12345', src_security_group_owner_id='123456')
	to_port=7000, src_security_group_group_id='sg-12345', src_security_group_owner_id='123456')

create.authorize('tcp', 22, 22, '168.198.0.21/32')
create.authorize('tcp', 8080, 8080, '168.198.0.21/32')

create.authorize('tcp', 22, 22, '168.198.2.0/24')
create.authorize('tcp', 8080, 8080, '168.198.2.0/24')

create.authorize('tcp', 22, 22, '168.198.3.241/32')
create.authorize('tcp', 8080, 8080, '168.198.3.241/32')

connection.authorize_security_group(group_id=create.id, ip_protocol='tcp', from_port=8080, 
	to_port=8080, src_security_group_group_id='sg-1234', src_security_group_owner_id='123456')
	to_port=8080, src_security_group_group_id='sg-123', src_security_group_owner_id='123456')

