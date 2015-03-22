import boto
import boto.ec2

connection = boto.ec2.connect_to_region('us-east-1')
create_snap = connection.create_snapshot('vol-12345', 'test_vol')
print create_snap.id