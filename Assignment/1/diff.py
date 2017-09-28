import sys


def getList(filename):
    r = list()

    with open(filename, 'r') as f:
        for line in f:
            line = line.replace(' ', '')
            line = line.replace('\n', '')
            s = line.split(':')
            r.append((s[0], s[1]))

    return r

def Sort2cmp(r):
    r1 = sorted(r, key=lambda x: x[0])
    r2 = sorted(r1, key=lambda x: int(x[1]), reverse=True)
    return r2
    
f1 = sys.argv[1]
f2 = sys.argv[2]

r1 = getList(f1)
r2 = getList(f2)

r1 = Sort2cmp(r1)
r2 = Sort2cmp(r2)

if r1 == r2:
    print("True")
else:
    print("False")
