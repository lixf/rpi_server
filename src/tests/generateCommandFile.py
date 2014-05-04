import sys
import random
import string

cmds = ["GET", "POST", "HASH"]

num_args = len(sys.argv) - 1
alphabet = string.lowercase

def randomWord(length = 0):
    if (length == 0):
        length = random.randint(8, 12)
    return ''.join(random.choice(alphabet) for n in xrange(length))

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
        args = randomWord()
    elif cmd == "POST":
        args = randomWord() + " " + randomWord()
    elif cmd == "HASH":
        key = randomWord()
        #First, POST a certain KEY/VALUE pair
        args = key + " " + randomWord()
        f.write("POST " + args + "\n")
        #Next, HASH the previous key with a generic salt and a random cost
        args = key + " salt " + str(random.randint(1, 3))
    f.write(cmd + " " + args + "\n")



