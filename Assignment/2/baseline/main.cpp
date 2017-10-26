#include <bits/stdc++.h>

using namespace std;

#define tmp_files 1000

int main()
{
    FILE* fd = fopen("../testcase/number_500M.in", "r");
    vector<string> inp;
    char buffer[40000];
    while(fgets(buffer, 40000, fd) != NULL) {
        inp.push_back(buffer);
    }

    sort(inp.begin(), inp.end());

    FILE* ans = fopen("../testcase/number_500M.baseline.out", "w");
    for(auto i : inp) {
        fprintf(ans, "%s\n", i.c_str());
    }

    return 0;
}