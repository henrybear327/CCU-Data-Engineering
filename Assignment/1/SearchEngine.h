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
        loadText();
    }

    void performTextSearch();

private:
    FileManager *fileManage;

    void loadKeywords();
    void loadText();

    // returns the longest keyword that is matched
    int searchKeywordFromIndex(int fd, int offset);

    std::map<std::vector<int>, int> match;
    std::vector<int> text;
};
#endif
