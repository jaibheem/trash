import boto
import boto.vpc
from sys import argv
connection = boto.vpc.connect_to_region('ap-southeast-2')
create = connection.create_security_group('rds-test', 'rds-test', vpc_id='vpc-123')
print create, create.id, create.name


create.authorize('tcp', 3306, 3306, '100.123.30.0/24')
create.authorize('tcp', 3306, 3306, '100.123.31.0/24')
