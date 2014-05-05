import sys
import random
import string

cmds = ["GET", "POST", "HASH"]
prob = dict()
for cmd in cmds:
    prob[cmd] = 0

num_args = len(sys.argv) - 1
alphabet = string.lowercase

def randomWord(length = 0):
    if (length == 0):
        length = random.randint(8, 12)
    return ''.join(random.choice(alphabet) for n in xrange(length))

if (num_args != 3):
    print "python generateCommandFile.py [OUTPUT FILE] [NUM COMMANDS] ['hash'/'basic']"
    exit(0)

filename = sys.argv[1]
numCommands = int(sys.argv[2])
cmdType = sys.argv[3]

#... just... keep the probabilities to two digits. 
if cmdType == 'hash':
    print "HASH -- 5% GET, 5% POST, 90% HASH"
    prob["GET"] = 0.05
    prob["POST"] = 0.05
    prob["HASH"] = 0.90
elif cmdType == 'basic':
    print "BASIC -- 40% GET, 40$ POST, 20% HASH"
    prob["GET"] = 0.40
    prob["POST"] = 0.40
    prob["HASH"] = 0.20

cmdPadded = []
for key in prob:
    for i in xrange(int(prob[key] * 100)):
        cmdPadded.append(key)

f = open(filename, 'w')

for i in xrange(numCommands):
    cmd = random.choice(cmdPadded)
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



