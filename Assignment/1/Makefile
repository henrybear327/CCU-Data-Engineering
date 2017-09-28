#CC=/usr/local/Cellar/gcc/7.2.0/bin/g++-7
CC=g++
CXXFLAGS=-Wall -Wextra -Wshadow -std=c++11 -O3

SOURCE=$(wildcard *.cpp)
HEADER=$(wildcard *.h)

all: searchEngine fileManager utf8Helper main
	$(CC) $(CXXFLAGS) *.o -o main
	
	clang-format -style=LLVM -i *.cpp *.h
	astyle --style=linux *.cpp *.h
	rm -f *.orig
	rm -f *.gch

searchEngine: SearchEngine.h SearchEngine.cpp
	$(CC) $(CXXFLAGS) -c SearchEngine.cpp

fileManager: FileManager.h FileManager.cpp
	$(CC) $(CXXFLAGS) -c FileManager.cpp

utf8Helper: UTF8Helper.h UTF8Helper.cpp
	$(CC) $(CXXFLAGS) -c UTF8Helper.cpp

main: main.cpp
	$(CC) $(CXXFLAGS) -c main.cpp

clean: 
	rm -f *.o
	rm -f main
