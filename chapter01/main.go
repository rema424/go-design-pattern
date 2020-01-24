package main

import (
	"fmt"

	"golang.org/x/xerrors"
)

func main() {
	bs := NewBookShelf(4)

	bs.Append(&Book{Name: "Around the World in 80 Days"})
	bs.Append(&Book{Name: "Bible"})
	bs.Append(&Book{Name: "Cinderella"})
	bs.Append(&Book{Name: "Daddy-Long-Legs"})

	it := bs.Iterator()

	for it.HasNext() {
		if v, err := it.Next(); err != nil {
			fmt.Printf("%#v\n", err)
			break
		} else if b := v.(*Book); b != nil {
			fmt.Println(b.Name)
		}
	}
}

type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
}

type Book struct {
	Name string
}

type BookShelf struct {
	books []*Book
}

func NewBookShelf(cap int) *BookShelf {
	return &BookShelf{
		books: make([]*Book, 0, cap),
	}
}

var (
	ErrOutOfIndex = xerrors.New("out of index")
)

func (bs *BookShelf) Get(idx int) (*Book, error) {
	if idx < len(bs.books) {
		return bs.books[idx], nil
	}
	return nil, ErrOutOfIndex
}

func (bs *BookShelf) Append(b *Book) {
	bs.books = append(bs.books, b)
}

func (bs *BookShelf) Length() int {
	return len(bs.books)
}

func (bs *BookShelf) Iterator() Iterator {
	return NewBookShelfIterator(bs)
}

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func NewBookShelfIterator(bs *BookShelf) Iterator {
	return &BookShelfIterator{
		bookShelf: bs,
	}
}

func (bsi *BookShelfIterator) HasNext() bool {
	return bsi.index < bsi.bookShelf.Length()
}

func (bsi *BookShelfIterator) Next() (interface{}, error) {
	b, err := bsi.bookShelf.Get(bsi.index)
	if err != nil {
		return nil, err
	}
	bsi.index++
	return b, nil
}
