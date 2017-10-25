import os
import csv
import time
import pathlib
from subprocess import run

def time_bash_sort(inp, out):
    print("Timing system sort using {0} as input and {1} as output".format(inp, out))
    cmd = "sort -d {0} -o {1}".format(inp, out)

    start_time = time.time()
    run(cmd, shell=True)
    end_time = time.time()
    return end_time - start_time

def time_my_sort(inp, out, p, type):
    print("Build my go sorting code...")
    build = "go build -o my_sort *.go"
    run(build, shell=True)
    print("Done")

    cmd = './my_sort -i {} -o {} -p {} -a {}'.format(inp, out, p, type)
    start_time = time.time()
    run(cmd, shell=True)
    end_time = time.time()
    return end_time - start_time

folder = pathlib.Path('/tmp/hw2/')
folder.mkdir(exist_ok=True)
inps = sorted(pathlib.Path('../data/C_Chat/').glob('*00'))

results = []
for inp in inps:   
    for p in [100_000 * (1 << exp) for exp in range(5)]:
        out = './out.{}'.format(inp.stem)

        print('*' * 40)
        print('inp =', inp)
        print('out =', out)
        print('p =', p)
        print('*' * 40)

        res1 = time_alg(inp, out, p, 'parti')
        results.append({'size': inp.stem, 'type': 'parti', 'p': p, 'dur': res1})

        res2 = time_alg(inp, out, p, 'merge')
        results.append({'size': inp.stem, 'type': 'merge', 'p': p, 'dur': res2})

        with open('./result.csv', 'w') as f:
            fieldnames = ['size', 'type', 'p', 'dur']
            writer = csv.DictWriter(f, fieldnames=fieldnames)
            writer.writeheader()
            for row in results:
                writer.writerow(row)
