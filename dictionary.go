package bencode

import (
	"container/list"
	"sync"
)

type dictionaryItem struct {
	Key   String
	Value Value
}

func (di *dictionaryItem) Bencode() []byte {
	var raw []byte

	raw = append(raw, di.Key.Bencode()...)
	raw = append(raw, di.Value.Bencode()...)

	return raw
}

type Dictionary struct {
	mu sync.RWMutex

	l *list.List
	m map[string]*list.Element
}

func (d *Dictionary) Bencode() []byte {
	d.mu.RLock()
	defer d.mu.RUnlock()

	raw := []byte{'d'}

	if d.l != nil {
		for el := d.l.Front(); el != nil; el = el.Next() {
			di := el.Value.(*dictionaryItem)

			raw = append(raw, di.Bencode()...)
		}
	}

	raw = append(raw, byte('e'))

	return raw
}

func (d *Dictionary) init() {
	if d.l == nil {
		d.l = list.New()
	}

	if d.m == nil {
		d.m = map[string]*list.Element{}
	}
}

func (d *Dictionary) Get(key String) (Value, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if d.m == nil {
		return nil, false
	}

	var bk string

	if len(key) > 0 {
		bk = key.BencodeKey()
	}

	var (
		ok bool
		el *list.Element
	)

	el, ok = d.m[bk]
	if !ok {
		return nil, false
	}

	di := el.Value.(*dictionaryItem)

	return di.Value, true
}

func (d *Dictionary) Set(key String, value Value) {
	d.mu.Lock()
	defer d.mu.Unlock()

	var bk string

	if len(key) > 0 {
		bk = key.BencodeKey()
	}

	d.init()

	var (
		ok bool

		el *list.Element
	)

	di := &dictionaryItem{
		Key:   key,
		Value: value,
	}

	el, ok = d.m[bk]
	if ok {
		el.Value = di

		return
	} else {
		el = nil
	}

	var markBeforeInsert *list.Element

	var (
		currentElement *list.Element = d.l.Front()
		nextElement    *list.Element
	)

	for currentElement != nil {
		nextElement = currentElement.Next()

		cedi := currentElement.Value.(*dictionaryItem)
		cebk := cedi.Key.BencodeKey()

		if nextElement == nil {
			if bk >= cebk {
				markBeforeInsert = currentElement

				break
			}
		} else {
			nedi := nextElement.Value.(*dictionaryItem)
			nebk := nedi.Key.BencodeKey()

			if bk > cebk && bk < nebk {
				markBeforeInsert = currentElement

				break
			}
		}

		currentElement = nextElement
		nextElement = nil
	}

	if markBeforeInsert == nil {
		el = d.l.PushFront(di)
	} else {
		el = d.l.InsertAfter(di, markBeforeInsert)
	}

	d.m[bk] = el
}

func (d *Dictionary) Remove(key String) {
	d.mu.Lock()
	defer d.mu.Unlock()

	var bk string

	if len(key) > 0 {
		bk = key.BencodeKey()
	}

	d.init()

	var (
		ok bool
		el *list.Element
	)

	el, ok = d.m[bk]
	if ok {
		d.l.Remove(el)

		delete(d.m, bk)
	}
}

func NewDictionary() *Dictionary {
	return &Dictionary{}
}
