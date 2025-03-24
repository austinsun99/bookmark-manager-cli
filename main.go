package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const help string = `Welcome to bookmark-manager, a CLI tool to help manage bookmarks and open multiple bookmarks at once

  [ex] represents the command to run the exectuable (depends on OS, and what you named your file). For example, if you built the file with
  go build -o bm. Then the [ex] command would be ./bm on linux to run the bm exectuable.


  To get started, type:

    !!! important command
    [ex] browser [command to open browser]
    Sets the command to open the browser. For example, if you use firefox, the command to open firefox is "firefox". The command
    you use must support this format of opening many urls at once.

    [ex] add -b [name-of-bookmark] [url-of-bookmark] 
      This adds a bookmark with the given name, and url to which the bookmark points to

    [ex] add -f [name-of-folder] [name-of-bookmarks]...
      This creates a folder with the given name, and the list of names of the bookmarks, which you will need to have already created with the above command.
      This creates a folder with the bookmarks. Folders are useful because you can open a bunch of bookmarks at once
    
    [ex] remove -b [name-of-bookmark]
      Removes the bookmark with the given name

    [ex] remove -f [name-of-folder]
      Removes the folder with the given name

    [ex] run -b [name-of-bookmark]
      Opens the bookmark with the given name in the given browser.

    [ex] run [name-of-folder]
      Opens all the bookmarks within the folder with the given browser.

    [ex] list
    [ex] ls
      Lists all the bookmarks and folders you have

    [ex] help
    [ex] -h
      Opens this page!
  `

func main() {
	args := os.Args[1:]
	switch args[0] {
  case "browser":
    SetBrowser(args[1])
	case "help":
		fallthrough
	case "-h":
		fmt.Println(help)
	case "add":
		if args[1] == "-b" {
			AddBookmark(args[2], args[3])
		} else if args[1] == "-f" {
			foundAllBookmarks, unfoundBookmarks := AddBookmarkFolder(args[2], args[3:])
			if !foundAllBookmarks {
				fmt.Println("Was unable to find: " + strings.Join(unfoundBookmarks, ", "))
			}
		}
	case "remove":
		if args[1] == "-b" {
			RemoveBookmark(args[2])
		} else if args[1] == "-f" {
			RemoveBookmarkFolder(args[2])
		}
	case "run":
		if args[1] == "-b" {
			bookmarks := ReadBookmarkInfo().Bookmarks
			for _, mark := range bookmarks {
				if mark.Alias == args[2] {
					cmd := exec.Command(GetBrowser(), mark.URL)
					cmd.Run()
				}
			}
		} else {
			urls := GetFolderURLs(args[1])
			if urls == nil {
				fmt.Println("Folder is empty or folder doesn't exist!")
				break
			}
			cmd := exec.Command(GetBrowser(), urls...)
			cmd.Run()
		}
	case "ls":
		fallthrough
	case "list":
		bookmarkInfo := ReadBookmarkInfo()
		fmt.Println("Bookmarks: ")
		for _, mark := range bookmarkInfo.Bookmarks {
			fmt.Println("Name: " + mark.Alias + " URL: " + mark.URL)
		}
		fmt.Println("\nFolders: ")
		for _, folder := range bookmarkInfo.BookmarkFolders {
			fmt.Println(folder.Name + "{ ")
			for _, mark := range folder.Bookmarks {
				fmt.Println("\tName: " + mark.Alias + " URL: " + mark.URL)
			}
			fmt.Println("}")
		}
	default:
		fmt.Println("Unable to find command, type -h to for help")
	}

}
