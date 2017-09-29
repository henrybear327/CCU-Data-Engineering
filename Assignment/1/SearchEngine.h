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
    }

    void loadFilesToMemory();
    void performTextSearch();
    void printFrequencyList();

private:
    void loadKeywords();
    void loadText();

    std::vector<int> text;
    std::map<std::vector<int>, int> match;

    FileManager *fileManager;
};
#endif
