#ifndef SEARCH_ENGINE_H
#define SEARCH_ENGINE_H

#include "FileManager.h"
#include <map>
#include <string>
#include <unistd.h>
#include <unordered_map>
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

#ifdef VECTORBASED
    std::vector<int> text;
    std::map<std::vector<int>, int> match;
#endif

#ifdef STRINGBASED
    std::string text;
    std::unordered_map<std::string, int> match;
#endif

    FileManager *fileManager;
};
#endif
