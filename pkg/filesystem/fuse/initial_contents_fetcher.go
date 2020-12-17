// +build darwin linux

package fuse

// InitialContentsFetcher is called into by InMemoryDirectory when a
// directory whose contents need to be instantiated lazily is accessed.
// The results returned by FetchContents() are used to populate the
// directory.
//
// FetchContents() should be called until it succeeds at most once. It
// may be possible FetchContents() is never called. This may happen if
// the directory in question is never accessed.
type InitialContentsFetcher interface {
	FetchContents() (map[string]InitialContentsFetcher, map[string]NativeLeaf, error)
}
