## Bookmanager-cli

I made this quick project as my first experience with Go.  

This CLI will allow you to add bookmarks, add bookmark folders, which is a collection of bookmarks.  
You will have the ability to open bookmark folders, which will open a bunch of bookmarks you've chosen at once.  

Usage Example: You may have a bunch of different websites for all your courses that are confusing to manage. You may forget to check one of them!  
Well, you can just add them all into a folder and open them all up at once, so that you'll never miss it!  

Clone the repository into a folder of your choice. You will only be able to operate this CLI within that folder (which is recommended), unless you add the executable to the path.  
The executable is called bm by default. If you move the executable to another location, you will need to create a folder called "data" within the same location.

#### Getting started
Just clone the repo into a folder of your choice!. Run the executable, with the argument "-h" to open the help menu. For example on linux:  
1. Open the folder with the executable  
2. Run ./bm -h  

If you would like to build it yourself, then you can build it with
```go build -o <name-of-executable>```
