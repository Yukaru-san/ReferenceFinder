# ReferenceFinder
Ever got tired of trying to figure out where to find a variable or function in a project or library?<br>
This tool's got your back.<br>

# Useage
``` referenceFinder <command> [path1] [path2] [...] [<flags>]```

Example:
```referenceFinder search "C:\myPath" --find "func main\(\)" --search-sub```
<br>
The above example will show you every file and their respective lines containing the go-typical main function.<br>
Note that the search input works by using RegEx and therefore needs to fit it's rules but also allows for more options when searching across your files.
<br>
<br>
**How do I search for multiple / different strings?**
<br>
As noted above, the seach method is using RegEx. To search for multiple strings it is as simple as<br>
```--find "string1|string2"```
# Commands / Flags

commands
```
search
replace (TODO)
```

flags
```
-h              prints the help msg

-f              files to ignore when searching
-d              directories to ignore when searching (only useful with search-sub)

--search-sub    Search through sub directories
```

Note that you can ignore multiple files / directories by seperating their names using a comma<br>
``` -f "file1,  file2, .exe, .bat" ```
