# Viraland is a batch-job control program, especially for bioinfomatics data
# file flow control and analysis. Viraland can be run directly in Python 3
# console or be run with a configuration file.
# This program also supply a good amount of API to make data processing more
# flexible.
# author: Yubing Hou
# license: GPL v3

import os
import sys

# Configuration class: user can configure commands and file path easily in
# Python console or use a configuration file
# TODO: Finishi this class
class VLConfig:
    def __init__(self):
        self.title = None
        self.log = None
        self.keep = None
        self.velvetcmd = None
        self.velvetin = None
        self.velvetout = None
        self.embosscmd = None
        self.embossin = None
        self.embossout = None
        self.paudacmd = None
        self.paudain = None
        self.paudaout = None
        self.yparsein = None
        self.yparseout = None
        self.cparsein = None
        self.cparseout = None
        self.visualin = None
        self.visualout = None

    def get_title (self):
        return self.title
    def get_log (self):
        return self.log
    def set_title (self, title):
        self.title = str(title)
    def set_log (self, option):
        self.log = bool (option)
    def set_keep (self, option):
        self.keep = bool(option)
    def set_velvet (self, cmd):
        self.velvetcmd = str(cmd)
    def parse_and_set (self, line):
        if line.startswith("[TITLE]"):
            self.title = line.split("]")[1].strip()
        if line.startswith("[LOG]"):
            option = line.split("]")[1].strip()
            if option == "y":
                self.log = True
            else:
                self.log = False
        if line.startswith("[KEEP]"):
            option = line.split("]")[1].strip()
            if option == "y":
                self.keep = True
            else:
                self.keep = False
        print ("In development")
    def use_config (self, file):
        path = str(file)
        fd = open (path)
        for line in fd:
            self.parse_and_set (line)

# Execution class: instantiate this class by taking a VLConfig instance. This
# allows multiple executions in one session of run.
# TODO: Finish this class
class VLRun:
    def __init__ (self, config):
        self.config = VLConfig (config)
        self.log = None
    def getconfig (self):
        return self.config
    def setlog (self):
        if self.config.getlog() == True:
            fd = open (self.config.get_title())
            self.log = fd
    def execute (self):
        self.setlog ()
        print ("In development")
    def yparse (self):
        print ("In development")

# Class of managing velvet: instantiate this class so that command of velvet
# can match user's system settings.
# TODO: finish this class
class VelvetPack:
    def __init__(self):
        self.fasta = []
        self.fastq = []
    def addfasta(self, path):
        self.fasta.append(str(path))

# Class of managing Emboss: instantiate this class so that command of Emboss
# cand mathc user's system setting
# TODO: finish this class
class EmbossPack:
    #Josh
    def __init__(self, inDirectory):
        self.inFiles = inDirectory

# Class of managing FAA file: instantiate this class so that command of FAA
# cand mathc user's system setting
# TODO: finish this class. Also, what is faa exactly is? a file for fasta/fastq?
class FaaPack:
    def __init__(self):
        self.folder = input('Enter the folder of faas that need to be concatenated')
    def concatFiles(self):
        subprocess.call(['cd', self.folder])
        subprocess.call(['cat', '*', '.faa', '>', '../allTogether.faa'])

# Class of managing Pauda: instantiate this class so that command of Pauda
# cand mathc user's system setting
# TODO: finish this class. Should be expandable to bowtie if needed. Bowtie
# might be more well-known by people than Pauda
class PaudaPack:
    #YB
    def __init__(self):
        self.fasta = []
        self.fastq = []
        self.out = ""

# Class of managing Parsers: instantiate this class so that command of parsers
# cand mathc user's system setting
# TODO: finish this class. Also, parser must follow the same format!
# Parser must be like this: $ ./parser [input file] [output file]
class ParsePack:
    #YB
    def __init__(self, name):
        """
            ParsePack(name) set a parser for files with a job name
        """
        self.name = str(name)
        self.prog = None
        self.inputs = []
        self.outputs = []
    def getname (self):
        return self.name
    def setparser (self, prog):
        """
            parser program must be executed in this format:
            $ [program] [input] [output]
            For example, if [program] is a Java program MyParser, prog should be:
            "java MyParser"
            If [program] is in /usr/bin, for example, /usr/bin/MyParser, prog
            should be:
            "/usr/bin/MyParser"
        """
        self.prog = str(prog)
    def run (self):
        for item in self.inputs:
            out = self.item + "-output.txt"
            exe = self.prog + " " + self.item + " " + out

# Class of managing Karona: instantiate this class so that command of Karona
# can match user's system setting
# TODO: finish this class
class KaronaPack:
    #Jon
    def __init__(self):
        self.output = []

# Class of managing visualization: instantiate this class so that command of
# visualization can match user's system setting
# TODO: finish this class
class VisualPack:
    #Need to get GBKs from user
    #Have user provide GBK directory
    #Use .glob to get all the GBKs
    #Write out to a file, send file to Jons program as argument
    def __init__(self):
        self.paths = []
        self.cmd = None
    def setplotcmd (self, instr):
        self.cmd = str(instr)
    def addfile (self, path):
        self.paths.append(path)
    def plot(self, item):
        if item in self.paths:
            exe = self.cmd + " " + item
            os.system(exe)
        else:
            print ("[System] Cannot find item:", item)
    def plotall(self):
        """
            applying plot command on all input file that is in file list.
            Example:
                >>> visual = VisualPack()
                >>> visual.addfile(myfile.txt)
                >>> visual.setplotcmd ("java viralplot")
                >>> visual.plotall()
        """
        for item in self.paths:
            exe = self.cmd + " " + item
            os.system(exe)
    def reset (self):
        """
            reset all configuration of an instance of VisualPack. Example:
        """
        self.path = []
        self.cmd = None
    def status (self):
        """
            Get the complete configuration of visualization. Example:
            >>> visual = VisualPack()
            >>> visual.setplotcmd ("java viralplot")
            >>> visual.addfile ("data_01.txt")
            >>> visual.addfile ("data_02.txt")
            >>> visual.status()
            command = java viralplot
            file(s):
                data_01.txt
                data_02.txt
            >>>
        """
        print ("command =", self.cmd)
        print ("file(s):")
        for item in self.paths:
            print ("\t", item)
