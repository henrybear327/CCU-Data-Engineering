# VSCode setup

* Add `"/usr/local/Cellar/gcc/7.2.0/include/c++/7.2.0",` in the configuration file for c++
    * `"/usr/local/include",`

* 
```
I was able to find a fix for this. It was due to the C++ extension update to v0.11.1.

On GitHub there was an issue reported. If anyone needs to fix this before they release a patch, go to File -> Preferences -> Settings in VS Code and change "C_Cpp.intelliSenseEngine": "Default" to "C_Cpp.intelliSenseEngine": "Tag Parser".
```