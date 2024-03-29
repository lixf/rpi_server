import sys
import random
import string

cmds = ["GET", "POST", "HASH", "PICT"]
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
    print "python generateCommandFile.py [OUTPUT FILE] [NUM COMMANDS] ['hash'/'basic'/'bandwidth']"
    exit(0)

filename = sys.argv[1]
numCommands = int(sys.argv[2])
cmdType = sys.argv[3]

#... just... keep the probabilities to two digits. 
if cmdType == 'hash':
    print "HASH -- 0% GET, 0% POST, 100% HASH"
    prob["GET"] = 0.00
    prob["POST"] = 0.00
    prob["HASH"] = 1.00
    prob["PICT"] = 0.00
elif cmdType == 'basic':
    print "BASIC -- 50% GET, 50$ POST, 0% HASH"
    prob["GET"] = 0.50
    prob["POST"] = 0.50
    prob["HASH"] = 0.00
    prob["PICT"] = 0.00
elif cmdType == 'bandwidth':
    print "Bandwidth -- 0% GET, 0$ POST, 0% HASH, 100% PICT"
    prob["GET"] = 0.00
    prob["POST"] = 0.00
    prob["HASH"] = 0.00
    prob["PICT"] = 1.00


cmdPadded = []
for key in prob:
    for i in xrange(int(prob[key] * 100)):
        cmdPadded.append(key)

f = open(filename, 'w')
local_path = "/home/xiaofan/Desktop/418_final/rpi_server/src"
store_path = "/home/pi/code/rpi_server/src"
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
    #for picture transmission
    elif cmd == "PICT":
        ran = random.randint(1,3)
        if ran == 1:
            local = local_path + "/local_pict/sig.ppm"
            store = store_path + "/store_pict/sig_t.ppm"
        elif ran == 2:
            local = local_path + "/local_pict/rpi.ppm"
            store = store_path + "/store_pict/rpi_t.ppm"
        elif ran == 3:
            local = local_path + "/local_pict/china.ppm"
            store = store_path + "/store_pict/china_t.ppm"
        args = local + " " + store
    f.write(cmd + " " + args + "\n")



