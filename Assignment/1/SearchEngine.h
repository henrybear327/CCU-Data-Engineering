#ifndef SEARCH_ENGINE_H
#define SEARCH_ENGINE_H

#include "FileManager.h"

struct SearchEngine {
public:
    SearchEngine(FileManager *_fileManager)
    {
        fileManage = _fileManager;
    }

    void performTextSearch();

    // returns the longest keyword that is matched
    int searchKeywordFromIndex(int fd, int offset);

private:
    FileManager *fileManage;
};
#endif
