package models

import "encoding/binary"

//Page carries pagination info to aid in knowing whether any given page has a
//next or previous page, and to know its page number
type Page struct {
	Prev    bool
	PrevVal int

	Next    bool
	NextVal int

	NextURL string

	pages int
	Pages []string
	Total int
	Count int
	Skip  int
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

//SearchPagination returns a page strict which carries details about the
//pagination of any given search result or pag/Users/SMILECS/Downloads/go-oddjobs-master/functions.goe
func SearchPagination(count int, page int, perPage int) Page {
	var pg Page
	var total int

	if count%perPage != 0 {
		total = count/perPage + 1
	} else {
		total = count / perPage
	}

	if total < page {
		page = total
	}

	if page == 1 {
		pg.Prev = false
		pg.Next = true
	}

	if page != 1 {
		pg.Prev = true
	}

	if total > page {
		pg.Next = true
	}

	if total == page {
		pg.Next = false
	}

	var pgs = make([]string, total)

	//The number of number of documents to skip
	skip := perPage * (page - 1)

	pg.Total = total
	pg.Skip = skip
	pg.Count = count
	pg.NextVal = page + 1
	pg.PrevVal = page - 1
	pg.Pages = pgs

	return pg
}
