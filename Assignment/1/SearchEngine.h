#include "FileManager.h"

#ifndef SEARCH_ENGINE_H
#define SEARCH_ENGINE_H

void performTextSearch(FileData &fileData);

// returns the longest keyword that is matched
int searchKeywordFromIndex(FILE *fd, int offset);

#endif
