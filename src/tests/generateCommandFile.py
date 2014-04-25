import sys
import random
import string

cmds = ["GET", "POST"]

num_args = len(sys.argv) - 1
alphabet = string.lowercase

if (num_args != 2):
    print "python generateCommandFile.py [OUTPUT FILE] [NUM COMMANDS]"
    exit(0)

filename = sys.argv[1]
numCommands = int(sys.argv[2])

f = open(filename, 'w')

for i in xrange(numCommands):
    cmd = random.choice(cmds)
    args = ""
    if cmd == "GET":
        args = str(random.randint(1, 20))
    elif cmd == "POST":
        args = str(random.randint(1, 20)) + " " + random.choice(alphabet)
    elif cmd == "COMPUTE":
        pass
    f.write(cmd + " " + args + "\n")
