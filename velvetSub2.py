#Manipulating velvet with subprocesses

import subprocess
import glob
class velvetPak:
	def __init__(self):
		self.fmt = None
		self.out = None
		self.inFiles = None
		self.kmer = None
		self.readType = None
		self.vCmd = None
		self.r1 = None
		self.r2 = None
		self.mult = None

	def getParams(self):
		#Ask user for arguments necessary to run velvet
		args = []
		print('The first thing we need is velvet arguments')

		fmt = str(input('What is the format of your files?'))
		self.fmt = fmt
		readType = str(input('What type of reads are your files?'))
		self.readType = readType

		#Get contents of folder at user specified directory
		files = self.getFolder(fmt)
		self.inFiles = files

		self.vCmd = 'echo'
		self.out = str(input('What directory would you like to output your files to?'))
		self.kmer = str(input('What kmer size would you like to use?'))
		#subprocess.call(['echo', 'velveth', '/home/lsb456/Desktop/try', '29', readType, form, '/home/lsb456/Desktop/tryR1.fastq.gz', '/home/lsb456/Desktop/tryR2.fastq.gz'])
		#return args

	def callVelvet(self, folder):
		if len(folder) > 1 and 'paired' in self.readType:
			print('process as paired reads in folder')
			#process as paired reads
		elif len(folder) > 1:
			print('process multiple single reads in folder')
			#process reads 1 by 1
		elif self.mult == 0:
			print('process a single read')
		#subprocess.call(args)

	def makeArgs(self):
		args = [self.vCmd, self.out, self.kmer, self.readType, self.fmt]
		return args

	def getFolder(self, fmt):
		#Get files from directory specified by user
		direct = input('Please enter the directory of the folder containing your reads')
		folder = glob.glob(direct+'*.' + fmt)
		return folder

a = velvetPak()
a.getParams()
a.callVelvet(a.makeArgs())
#cmd = getParams()
#subprocess.call(cmd)
