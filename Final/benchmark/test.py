import re

repeat = 30

pattern = "a?" * repeat + "a" * repeat
string = "a" * 2 * repeat
# string = "a" * repeat

print(f'"{pattern}" "{string}"')

ret = re.match(pattern, string)
print(ret[0])
