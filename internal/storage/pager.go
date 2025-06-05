// TODO: Implement this file
// internal/storage/pager.go
package storage

import (
    "fmt"
    "os"
    "sync"
)

const (
    // PageSize is the fixed size of each page in bytes.
    PageSize = 4096
)

// Pager manages pages in a file and a free list of unused pages.
type Pager struct {
    file      *os.File
    pageCount uint64
    freeList  []uint64
    mu        sync.Mutex
}

// NewPager opens or creates the file and initializes the pager.
func NewPager(path string) (*Pager, error) {
    file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0664)
    if err != nil {
        return nil, err
    }
    fi, err := file.Stat()
    if err != nil {
        file.Close()
        return nil, err
    }
    // Compute number of existing pages.
    pageCount := uint64(fi.Size() / PageSize)
    return &Pager{file: file, pageCount: pageCount}, nil
}

// ReadPage reads the page at pageNum.
func (p *Pager) ReadPage(pageNum uint64) ([]byte, error) {
    p.mu.Lock()
    defer p.mu.Unlock()
    if pageNum >= p.pageCount {
        return nil, fmt.Errorf("invalid page number")
    }
    buf := make([]byte, PageSize)
    _, err := p.file.ReadAt(buf, int64(pageNum*PageSize))
    if err != nil {
        return nil, err
    }
    return buf, nil
}

// WritePage writes data (up to PageSize) to the given page.
func (p *Pager) WritePage(pageNum uint64, data []byte) error {
    p.mu.Lock()
    defer p.mu.Unlock()
    if len(data) > PageSize {
        return fmt.Errorf("data too large for page")
    }
    if pageNum >= p.pageCount {
        // Extend page count if writing beyond current size.
        p.pageCount = pageNum + 1
    }
    _, err := p.file.WriteAt(data, int64(pageNum*PageSize))
    if err != nil {
        return err
    }
    return nil
}

// AllocatePage returns a page number: either reusing a free page or new.
func (p *Pager) AllocatePage() uint64 {
    p.mu.Lock()
    defer p.mu.Unlock()
    n := len(p.freeList)
    if n > 0 {
        pageNum := p.freeList[n-1]
        p.freeList = p.freeList[:n-1]
        return pageNum
    }
    pageNum := p.pageCount
    p.pageCount++
    return pageNum
}

// FreePage marks a page as reusable.
func (p *Pager) FreePage(pageNum uint64) {
    p.mu.Lock()
    defer p.mu.Unlock()
    p.freeList = append(p.freeList, pageNum)
}

// Sync flushes all writes to disk (fsync).
func (p *Pager) Sync() error {
    p.mu.Lock()
    defer p.mu.Unlock()
    return p.file.Sync()
}

// Close closes the underlying file.
func (p *Pager) Close() error {
    return p.file.Close()
}
