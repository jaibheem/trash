import boto
import boto.ec2
from sys import argv
connection = boto.ec2.connect_to_region('ap-southeast-2')
create = connection.create_security_group('elb-api', 'elb-api')
print create
create.authorize('tcp', 80, 80, '0.0.0.0/0', group_id='vpc-123456')
create.authorize('tcp', 443, 443, '0.0.0.0/0', group_id='vpc-123456')
