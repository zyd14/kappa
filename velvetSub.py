#Manipulating velvet with subprocesses

import subprocess

def velvetInit():
	#Set params through user input in shell
	args = getParams()
	velvetInit(args)
def velvetInit(argsIn):
	#Set params according to args given when module function called
	args = argsIn

def getParams():
	args = []
	print('The first thing we need is velvet arguments')
	args.append('echo')
	out = input('What directory would you like to output your files to?')
	args.append(out)
	args.append(input('What kmer size would you like to use?'))
	args.append(input('Enter the file path to the first file'))
	args.append(input('Enter the file path to the second file'))
	args.append(input('What is the format of your files?'))
	args.append(input('What type of reads are your files?'))
	subprocess.call(['echo', 'velveth', '/home/lsb456/Desktop/try', '29', readType, form, '/home/lsb456/Desktop/tryR1.fastq.gz', '/home/lsb456/Desktop/tryR2.fastq.gz'])
	return args

def callVelvet(args):
	subprocess.call(args)
#cmd = getParams()
#subprocess.call(cmd)