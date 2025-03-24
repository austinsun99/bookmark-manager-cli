package main

import (
	"encoding/json"
	"os"
	"slices"
)

const BOOKMARKS_INFO_FILE_PATH string = "data/bookmarks.json"

type Bookmark struct {
	Alias string `json:"alias"`
	URL   string `json:"url"`
}

type BookmarkFolder struct {
	Name      string     `json:"name"`
	Bookmarks []Bookmark `json:"bookmark"`
}

type BookmarksInfo struct {
	Browser         string           `json:"browser"`
	Bookmarks       []Bookmark       `json:"bookmarks"`
	BookmarkFolders []BookmarkFolder `json:"bookmarkFolders"`
}

func CreateNewBookmarkInfo() {
  WriteBookmarkInfo(&BookmarksInfo{Browser: "", Bookmarks: []Bookmark{}})
}

func ReadBookmarkInfo() BookmarksInfo {
	dat, _ := os.ReadFile(BOOKMARKS_INFO_FILE_PATH)
	var res BookmarksInfo
	json.Unmarshal(dat, &res)
	return res
}

func WriteBookmarkInfo(info *BookmarksInfo) {
	dat, _ := json.MarshalIndent(info, "", "  ")
	f, _ := os.Create(BOOKMARKS_INFO_FILE_PATH)
	defer f.Close()
	f.Write(dat)
}

func GetBrowser() string {
	var currentBookmarksInfo BookmarksInfo = ReadBookmarkInfo()
  return currentBookmarksInfo.Browser
}

func SetBrowser(browser string) {
	var currentBookmarksInfo BookmarksInfo = ReadBookmarkInfo()
  currentBookmarksInfo.Browser = browser
	WriteBookmarkInfo(&currentBookmarksInfo)
}

func AddBookmark(alias string, url string) {
	var currentBookmarksInfo BookmarksInfo = ReadBookmarkInfo()
	currentBookmarksInfo.Bookmarks = append(currentBookmarksInfo.Bookmarks, Bookmark{alias, url})
	WriteBookmarkInfo(&currentBookmarksInfo)
}

func RemoveBookmark(alias string) {
	var currentBookmarksInfo BookmarksInfo = ReadBookmarkInfo()
	indexOfBookmark := GetIndexOfBookmark(alias, &currentBookmarksInfo)
	currentBookmarksInfo.Bookmarks = slices.Delete(currentBookmarksInfo.Bookmarks, indexOfBookmark, indexOfBookmark+1)
	WriteBookmarkInfo(&currentBookmarksInfo)
}

func RemoveBookmarkFolder(name string) {
	var currentBookmarksInfo BookmarksInfo = ReadBookmarkInfo()
	indexOfFolder := GetIndexOfFolder(name, &currentBookmarksInfo)
	currentBookmarksInfo.BookmarkFolders = slices.Delete(currentBookmarksInfo.BookmarkFolders, indexOfFolder, indexOfFolder+1)
	WriteBookmarkInfo(&currentBookmarksInfo)
}

func AddBookmarkFolder(name string, bookmarkNames []string) (bool, []string) {
	bookmarkInfo := ReadBookmarkInfo()

	var bookmarks []Bookmark
	for _, mark := range bookmarkInfo.Bookmarks {
		if slices.Contains(bookmarkNames, mark.Alias) {
			bookmarkNames = slices.DeleteFunc(bookmarkNames, func(name string) bool {
				return name == mark.Alias
			})
			bookmarks = append(bookmarks, mark)
		}
	}

  if len(bookmarkNames) != 0 {
		return false, bookmarkNames
	}

	var folder BookmarkFolder = BookmarkFolder{Name: name, Bookmarks: bookmarks}
	if bookmarkInfo.BookmarkFolders == nil {
		bookmarkInfo.BookmarkFolders = []BookmarkFolder{}
	}
	bookmarkInfo.BookmarkFolders = append(bookmarkInfo.BookmarkFolders, folder)

	dat, _ := json.MarshalIndent(bookmarkInfo, "", "  ")
	f, _ := os.Create(BOOKMARKS_INFO_FILE_PATH)
	defer f.Close()
	f.Write(dat)

	return true, nil

}

func GetFolderURLs(name string) []string {

	bookmarkInfo := ReadBookmarkInfo()
	var bookmarkFolder BookmarkFolder
	var bookmarkURLs []string

	for _, folders := range bookmarkInfo.BookmarkFolders {
		if folders.Name == name {
			bookmarkFolder = folders
			break
		}
	}

	for _, bookmark := range bookmarkFolder.Bookmarks {
		bookmarkURLs = append(bookmarkURLs, bookmark.URL)
	}

	return bookmarkURLs

}

func GetIndexOfBookmark(alias string, info *BookmarksInfo) int {
	for i, mark := range info.Bookmarks {
		if mark.Alias == alias {
			return i
		}
	}
	return -1
}

func GetIndexOfFolder(name string, info *BookmarksInfo) int {
	for i, folder := range info.BookmarkFolders {
		if folder.Name == name {
			return i
		}
	}
	return -1
}
