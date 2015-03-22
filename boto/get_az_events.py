from boto import ec2
import sys


az = sys.argv[1]
region = az[:-1]
conn = ec2.connect_to_region(region)
counter = 0
filter = {'availability-zone': az}
instances = conn.get_all_instance_status(filters=filter)
for instance in instances:
    event = instance.events
    if event and "Completed" not in instance.events[0].description:
        counter += 1
        print "%s : %s : %s " % (instance.id, instance.state_name, event)
print "%s : %s" % (region, counter)
