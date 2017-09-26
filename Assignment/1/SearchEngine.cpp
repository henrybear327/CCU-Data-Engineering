#include <cstdio>
#include <cstdlib>
#include "FileManager.h"
#include "SearchEngine.h"

void performTextSearch(FileData& fileData)
{
	// TODO: add timer
	// TODO: add color
	printf("Performing keyword searching...\n");
	if(fileData.isLoaded() == false) {
		printf("File loading error\n");
		exit(1);
	}
	
	printf("Done\n");
}
