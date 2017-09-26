#include <cstdio>
#include <string>

#ifndef FILE_MANAGER_H
#define FILE_MANAGER_H

struct FileData {
public:
	FileData() {
		printf("Loading files...\n");
		// TODO: file descriptor
		
		printf("Done\n");
	}

	bool isLoaded() {
		// TODO: perform checking
		return true;
	}

private:
	std::string keywordFile = "";
	std::string textFile = "";
	std::string resultFile = "";

	FILE* readFileWithName(const char* filename) {
		printf("Attemping to read %s\n", filename);

		printf("Done\n");
		return NULL;
	}

	// TODO: use fd? as long as fd can work
	FILE* keyword;
	FILE* text;
	FILE* result;
};
#endif
