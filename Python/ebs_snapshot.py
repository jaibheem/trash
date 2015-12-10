#!/usr/bin/env python
# -*- coding: utf-8 -*-


import boto.ec2
from ansible.module_utils.basic import *

class Snapshotmodule:
    def __init__(self):
        self.module = AnsibleModule(
            argument_spec=dict(
                instance_id=dict(required=True),
                hostname=dict(required=True),
                region=dict(required=True),
                aws_access_key=dict(required=True),
                aws_secret_key=dict(required=True)
            )
        )
        self.instance_id = self.module.params.get("instance_id")
        self.hostname= self.module.params.get("hostname")
        self.region = self.module.params.get("region")
        self.aws_access_key = self.module.params.get("aws_access_key")
        self.aws_secret_key = self.module.params.get("aws_secret_key")
        self.success = True
        self.ret_msg = ''
	self.device = ''
        self.snapshots = {}
        self.tags = {"hostname": self.hostname, "device": self.device}

    def snapshot(self):
        try:
            out = {}
	    conn = boto.ec2.connect_to_region(self.region, aws_access_key_id=self.aws_access_key, aws_secret_access_key=self.aws_secret_key)
            volumes = [v for v in conn.get_all_volumes() if v.attach_data.instance_id == self.instance_id]
            for vol in volumes:
                snapshot = vol.create_snapshot()
                self.snapshots[vol.id] = snapshot.id
		self.tags["device"] = vol.attach_data.device
		for key, value in self.tags.items():
	                snapshot.add_tag(key,value=value)

        except Exception as e:
            self.ret_msg = e
            self.success = False

    def check_run(self):
        if self.success:
            self.module.exit_json(msg=self.ret_msg, snapshots=self.snapshots)
        else:
            self.module.fail_json(msg=self.ret_msg)


if __name__ == '__main__':
    ebs = Snapshotmodule()
    ebs.snapshot()
    ebs.check_run()
