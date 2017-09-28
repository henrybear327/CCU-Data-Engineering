#include "SearchEngine.h"
#include "FileManager.h"
#include "UTF8Helper.h"

#include <cstdio>
#include <cstdlib>

void SearchEngine::performTextSearch()
{
    // TODO: add timer
    printf("Performing keyword searching...\n");

    int code;
    do {
        code = fileManage->textHelper->extractWord();
        printf("%d\n", code);
    } while (code != 0);
    puts("");

    printf("Done\n");
}

void SearchEngine::loadKeywords()
{
    printf("Loading keywords...\n");

    while (1) {
        bool terminate = false;
        std::vector<int> tmp;
        for (int i = 0, code = fileManage->keywordHelper->extractWord(); code != 10;
             i++, code = fileManage->keywordHelper->extractWord()) {
            if (code == 0) {
                terminate = true;
                break;
            }

            tmp.push_back(code);
        }

        if (terminate)
            break;

        printf("Keyword added\n");
        cnt[tmp] = 0;
    }

    printf("Done\n\n");
}