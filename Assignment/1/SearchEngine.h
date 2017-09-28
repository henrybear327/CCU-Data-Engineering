#ifndef SEARCH_ENGINE_H
#define SEARCH_ENGINE_H

#include "FileManager.h"
#include <map>
#include <unistd.h>
#include <vector>

struct SearchEngine {
public:
    SearchEngine(FileManager *_fileManager)
    {
        fileManager = _fileManager;
        loadKeywords();
        loadText();
    }

    void performTextSearch();

private:
    FileManager *fileManager;

    void loadKeywords();
    void loadText();

    // returns the longest keyword that is matched
    int searchKeywordFromIndex(int fd, int offset);

    std::map<std::vector<int>, int> match;
    std::vector<int> text;
};
#endif
