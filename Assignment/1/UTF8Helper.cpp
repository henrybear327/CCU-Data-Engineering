#include "UTF8Helper.h"
#include "color.h"
#include <cstdio>
#include <cstdlib>
#include <unistd.h>

UTF8Helper::UTF8Helper(int _fd)
{
    fd = _fd;
}

int UTF8Helper::determineWordLength()
{
    unsigned char buffer[4];
    {
        int ret = read(fd, buffer, 1);
        if (ret == -1) {
            // error
            perror("Read() error");
            exit(1);
        } else if (ret == 0) {
            // EOF
            return 0;
        }
    }

    int decodedWord = 0;
    int bytes;
    if (((buffer[0] >> 7) & 1) == 0) {
        // printf("Number of bytes %d\n", 1);

        decodedWord = buffer[0];
        bytes = 1;
    } else if (((buffer[0] >> 5) & 7) == 6) {
        // printf("Number of bytes %d\n", 2);

        int ret = read(fd, buffer + 1, 1);
        if (ret == -1) {
            // error
            perror("Read() error");
            exit(1);
        } else if (ret == 0) {
            printf(ANSI_COLOR_RED
                   "Error decoding utf8 (missing bytes)\n" ANSI_COLOR_RESET);
            exit(1);
        }

        decodedWord = ((buffer[0] & 0x1F) << 6) | (buffer[1] & 0x3F);
        bytes = 2;
    } else if (((buffer[0] >> 4) & 15) == 14) {
        // printf("Number of bytes %d\n", 3);

        int ret = read(fd, buffer + 1, 2);
        if (ret == -1) {
            // error
            perror("Read() error");
            exit(1);
        } else if (ret == 0) {
            printf(ANSI_COLOR_RED
                   "Error decoding utf8 (missing bytes)\n" ANSI_COLOR_RESET);
            exit(1);
        }

        decodedWord = ((buffer[0] & 0x0F) << 12) | ((buffer[1] & 0x3F) << 6) |
                      (buffer[2] & 0x3f);
        bytes = 3;
        // printf("%d %d %d: %d\n", buffer[2], buffer[1], buffer[0], decodedWord);
    } else if (((buffer[0] >> 3) & 31) == 30) {
        // printf("Number of bytes %d\n", 4);

        int ret = read(fd, buffer + 1, 3);
        if (ret == -1) {
            // error
            perror("Read() error");
            exit(1);
        } else if (ret == 0) {
            printf(ANSI_COLOR_RED
                   "Error decoding utf8 (missing bytes)\n" ANSI_COLOR_RESET);
            exit(1);
        }

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
    return decodedWord;
}

int UTF8Helper::extractWord()
{
    return this->determineWordLength();
}