#viral crosses Python

import os
import sys
import pipes

#file formats: v#S#
start = 1
v1 = range(1,5)
s1 = range(start,15)
cmd = pipes.Template()


for v in v1:
	for s in s1:
		cmd.append('echo works', '--')
		print('v' + str(v) + 'S' + str(s))

cmd.open