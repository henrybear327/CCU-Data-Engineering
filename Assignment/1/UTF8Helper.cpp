#include "UTF8Helper.h"
#include "color.h"
#include <cstdio>
#include <cstdlib>
#include <unistd.h>

#define BUFFER_SIZE 100000

UTF8Helper::UTF8Helper(int _fd, bool needLoading)
{
    fd = _fd;
    isLoaded = false;

    if (needLoading)
        loadFileToMemory();
}

void UTF8Helper::loadFileToMemory()
{
    isLoaded = !isLoaded;

    unsigned char buffer[BUFFER_SIZE];

    while (1) {
        int ret = read(fd, buffer, BUFFER_SIZE);
        if (ret == -1) {
            // error
            perror("Read() error");
            exit(1);
        } else if (ret == 0) {
            // EOF
            break;
        }

        for (int i = 0; i < ret; i++) {
            originalData.push_back(buffer[i]);
        }
    }
}

inline unsigned char UTF8Helper::getNext()
{
    // printf("original %d\n", originalData[currentPosition]);
    return originalData[currentPosition++];
}

int UTF8Helper::determineWordLength()
{
    unsigned char buffer[4] = {0};
    buffer[0] = getNext();

    int decodedWord = 0;
    int bytes;
    if (((buffer[0] >> 7) & 1) == 0) {
        // printf("Number of bytes %d\n", 1);

        decodedWord = buffer[0];
        bytes = 1;
    } else if (((buffer[0] >> 5) & 7) == 6) {
        // printf("Number of bytes %d\n", 2);

        buffer[1] = getNext();

        decodedWord = ((buffer[0] & 0x1F) << 6) | (buffer[1] & 0x3F);
        bytes = 2;
    } else if (((buffer[0] >> 4) & 15) == 14) {
        // printf("Number of bytes %d\n", 3);

        buffer[1] = getNext();
        buffer[2] = getNext();

        decodedWord = ((buffer[0] & 0x0F) << 12) | ((buffer[1] & 0x3F) << 6) |
                      (buffer[2] & 0x3f);
        bytes = 3;
        // printf("%d %d %d: %d\n", buffer[2], buffer[1], buffer[0], decodedWord);
    } else if (((buffer[0] >> 3) & 31) == 30) {
        // printf("Number of bytes %d\n", 4);

        buffer[1] = getNext();
        buffer[2] = getNext();
        buffer[3] = getNext();

        decodedWord = ((buffer[0] & 0x07) << 18) | ((buffer[1] & 0x3F) << 12) |
                      ((buffer[2] & 0x3f) << 6) | ((buffer[3] & 0x3f));
        bytes = 4;
    } else {
        printf(ANSI_COLOR_RED
               "Error decoding utf8 (invalid prefix)\n" ANSI_COLOR_RESET);
        exit(1);
    }

    Word word(buffer, bytes);
    dictionary[decodedWord] = word;
    // printf("decoded %d\n", decodedWord);
    return decodedWord;
}

int UTF8Helper::extractWord()
{
    return determineWordLength();
}

void UTF8Helper::clearOriginalData()
{
    originalData.clear();
    originalData.resize(0);
}