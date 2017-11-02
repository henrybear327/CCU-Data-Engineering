#include <bits/stdc++.h>

using namespace std;

#define BUFFER_SIZE 100000000

char buffer[BUFFER_SIZE];
const char *delimiter = "\n\t ";

unordered_map<string, int> cnt;
typedef pair<int, string> ii;

struct classcomp {
    bool operator()(const ii &lhs, const ii &rhs) const
    {
        if (lhs.first == rhs.first)
            return lhs.second < rhs.second;
        return lhs.first > rhs.first;
    }
};
set<ii, classcomp> ans;

int main()
{
    while (fgets(buffer, BUFFER_SIZE, stdin) != NULL) {
        char *tok = strtok(buffer, delimiter);
        while (tok != NULL) {
            cnt[tok]++;

            tok = strtok(NULL, delimiter);
        }
    }

    for (auto i : cnt) {
        // printf("%20s %10d\n", i.first.c_str(), i.second);
        ans.insert(ii(i.second, i.first));
    }

    for (auto i : ans) {
        // printf("%10d %20s\n", i.first, i.second.c_str());
        printf("%d %s\n", i.first, i.second.c_str());
    }
    
    return 0;
}
