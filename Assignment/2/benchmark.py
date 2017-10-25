import os
import csv
import time
import pathlib
from subprocess import run

in_file = "./testcase/number_{0}M.in"
out_file = "./testcase/number_{0}M.out"
my_out_file = "number.my.out"
fileList = ['50', '100', '200', '300', '400', '500', '600']

def time_bash_sort(inp, out):
    print("\nTiming system sort using {0} as input and {1} as output".format(inp, out))
    cmd = "sort -d {0} -o {1}".format(inp, out)

    start_time = time.time()
    run(cmd, shell=True)
    end_time = time.time()
    return end_time - start_time

def time_my_sort(inp, out, chunks):
    print("Build my go sorting code...")
    build = "go build -o my_sort *.go"
    run(build, shell=True)
    print("Done")

    print("\nTiming my sort using {0} as input and {1} as output".format(inp, out))
    cmd = './my_sort -i {0} -o {1} -chunks {2} -tmp {3}'.format(inp, out, chunks, "./tmp")
    start_time = time.time()
    run(cmd, shell=True)
    end_time = time.time()
    return end_time - start_time

def check(my, ans):
    cmd = "cmp {0} {1}".format(my, ans)
    run(cmd, shell=True)

for i in fileList:
    time_bash_sort(in_file.format(i), out_file.format(i))
    time_my_sort(in_file.format(i), my_out_file, 1000)

print("\n\n=====================================")
check(my_out_file)
print("=====================================")