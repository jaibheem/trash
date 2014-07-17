import boto
import boto.ec2

connection = boto.ec2.connect_to_region('us-east-1')
#print connection
all_eips = connection.get_all_addresses()
print "EIP:\t\t\tAssociated Instance ID:"
for eip in all_eips:
	print eip.public_ip, "\t\t", eip.instance_id