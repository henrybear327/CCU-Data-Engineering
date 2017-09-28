#include "SearchEngine.h"
#include "FileManager.h"
#include "UTF8Helper.h"

#include <cstdio>
#include <cstdlib>

void SearchEngine::performTextSearch()
{
    // TODO: add timer
    printf("Performing keyword searching...\n");

    for (int i = 0; i < (int)text.size(); i++) {
        // printf("%d: %d\n", i, text[i]);

        std::vector<int> candidate;
        for (int j = 0; i + j < (int)text.size() && j < 7; j++) {
            candidate.push_back(text[i + j]);
        }

        for (int j = (int)candidate.size(); j >= 2; j--) {
            // for (auto k : candidate) {
            //     printf("%d ", k);
            //     // getchar();
            // }
            // puts("");

            if (match.find(candidate) != match.end()) {
                printf("Matched at %d\n", i);

                i += (j - 1);
                break;
            }
            candidate.pop_back();
        }
    }

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
        match[tmp] = 0;
    }

    printf("Done\n\n");
}

void SearchEngine::loadText()
{
    printf("Loading text...\n");

    for (int code = fileManage->textHelper->extractWord(); code != 0;
         code = fileManage->textHelper->extractWord()) {
        text.push_back(code);
    }

    printf("Done\n\n");
}