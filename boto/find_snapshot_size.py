from boto import ec2
#conn = ec2.connect_to_region('ap-southeast-2')
#list = conn.get_all_snapshots()
#for snap in list:
#	print snap, "==>", snap.volume_size
snapshot = "snap-12345"
conn = ec2.connect_to_region("us-west-2")
print snapshot.volume_size
