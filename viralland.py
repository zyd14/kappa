# Viral-Land executes programs with a configuration file config.vl. This
# file should contain information of file path(s), and options

# Configuration class: user can configure commands and file path easily in
# Python console or use a configuration file
import os



class VLConfig:
    def __init__(self, title):
        self.title = str(title)
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


class VelvetPack:
    def __init__(self):
        self.fasta = []
        self.fastq = []
    def addfasta(self, path):
        self.fasta.append(str(path))

class EmbossPack:
    def __init__(self):
        self.fasta = []
        self.fastq = []

class FaaPack:
    def __init__(self):
        self.fasta = []

class PaudaPack:
    def __init__(self):
        self.fasta = []
        self.fastq = []
        self.out = ""

class ParsePack:
    def __init__(self):
        self.inputs = []
        self.outputs = []

class KaronaPack:
    def __init__(self):
        self.output = []

class VisualPack:
    def __init__(self):
        self.paths = []
        self.cmd = ""
    def setplotcmd (self, instr):
        self.cmd = str(instr)
    def addfile (self, path):
        self.paths.append(path)
    def plot(self, item):
        exe = self.cmd + " " + item
        os.system(exe)
    def plotall(self):
        for item in self.paths:
            exe = self.cmd + " " + item
            os.system(exe)
