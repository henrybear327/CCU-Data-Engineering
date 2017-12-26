import re

repeat = int(input("n = "))

pattern = "a?" * repeat + "a" * repeat
string = "a" * 2 * repeat

print(f'"{pattern}" "{string}"')

ret = re.match(pattern, string)
print(ret[0])
