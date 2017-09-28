#ifndef SEARCH_ENGINE_H
#define SEARCH_ENGINE_H

#include "FileManager.h"
#include <map>
#include <vector>

struct SearchEngine {
public:
    SearchEngine(FileManager *_fileManager)
    {
        fileManage = _fileManager;
        loadKeywords();
    }

    void performTextSearch();

private:
    FileManager *fileManage;

    void loadKeywords();

    // returns the longest keyword that is matched
    int searchKeywordFromIndex(int fd, int offset);

    std::map<std::vector<int>, int> cnt;
};
#endif
