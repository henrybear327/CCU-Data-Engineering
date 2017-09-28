CC=/usr/local/Cellar/gcc/7.2.0/bin/g++-7
CXXFLAGS=-Wall -Wextra -Wshadow -std=c++11

SOURCE=$(wildcard *.cpp)
HEADER=$(wildcard *.h)

all: searchEngine fileManager utf8reader main
	$(CC) $(CXXFLAGS) *.o -o main
	
	clang-format -style=LLVM -i *.cpp *.h
	astyle --style=linux *.cpp *.h
	rm *.orig
	rm *.gch

searchEngine: SearchEngine.cpp
	$(CC) $(CXXFLAGS) -c SearchEngine.cpp

fileManager: FileManager.h
	$(CC) $(CXXFLAGS) -c FileManager.h

utf8reader: UTF8Reader.h UTF8Reader.cpp
	$(CC) $(CXXFLAGS) -c UTF8Reader.cpp

main: main.cpp
	$(CC) $(CXXFLAGS) -c main.cpp

clean: 
	rm -f *.o
	rm -f main
