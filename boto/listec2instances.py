import boto
import boto.ec2
from boto.ec2.ec2object import TaggedEC2Object

region = boto.ec2.connect_to_region('us-east-1')
print region
reservation = region.get_all_instances()
print "\tInstance ID: Instance Status"
for res in reservation:
	for instance in res.instances:
		print "\t%s :\t%s" % (instance.id, instance.state)
#		print instance.get_all_tags()
#instance = reservation.instances[0]
#print instance.instance_i
#tags = region.get_all_tags()
#for tag in tags:
#    print tag.value
#i = 1
#for list in reservation:
#	print i, ":", list
#	i = i+1
