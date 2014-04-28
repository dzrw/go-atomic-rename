A command-line tool for doing atomic renames of files in Windows.  
This purpose of this project is to explore a potential patch to address 
a bug in bitly/nsq/nsqd on Windows. 


### Cross-compiling for Windows
Use `goxc` to cross-compile for Windows targets.

```Bash
$ goxc -d="./dist" -pv=0.3.0 -bc="windows,amd64"
```

### Usage
```Bash
aren source_file target_file
```
