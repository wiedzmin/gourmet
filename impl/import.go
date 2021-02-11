package impl

func importBookmarks(sourceType int) error {
	if sourceType == typeBuku {
		return importBookmarksBuku()
	}
	return nil
}

func importBookmarksBuku() error {
	return nil
}
