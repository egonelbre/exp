// Type your code here, or load an example.
// Your function name should start with a capital letter.
package main

import (
	"strconv"
	"testing"
	"time"
)

type Private struct {
	Name     string
	Created  int64
	Modified int64
	User     []KeyValue
}

type KeyValue struct{ Key, Value string }

func convert(values []KeyValue) map[string]string {
	if len(values) == 0 {
		return nil
	}
	xs := map[string]string{}
	for _, kv := range values {
		xs[kv.Key] = kv.Value
	}
	return xs
}

type Object struct {
	Name     string
	Created  time.Time
	Modified time.Time
	User     map[string]string
}
type List struct {
	Prefix string
	Items  []Private
	Index  int
}

func (p *List) Next() bool {
	p.Index++
	return p.Index <= len(p.Items)
}

func (p *List) Item() Object {
	item := &p.Items[p.Index-1]
	return Object{
		Name:     p.Prefix + item.Name,
		Created:  time.Unix(item.Created, 0),
		Modified: time.Unix(item.Modified, 0),
		User:     convert(item.User),
	}
}

func (p *List) ItemPtr() *Object {
	item := &p.Items[p.Index-1]
	return &Object{
		Name:     p.Prefix + item.Name,
		Created:  time.Unix(item.Created, 0),
		Modified: time.Unix(item.Modified, 0),
		User:     convert(item.User),
	}
}

var data = func() []Private {
	var xs []Private
	for i := 0; i < 100; i++ {
		xs = append(xs, Private{
			Name:     strconv.Itoa(i),
			Created:  int64(i)*1000 + time.Now().Unix(),
			Modified: int64(i)*1000 + time.Now().Unix(),
			User:     []KeyValue{{"A", "B"}},
		})
	}
	return xs
}()

func BenchmarkStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := List{
			Prefix: "xyz/",
			Index:  0,
			Items:  data,
		}
		var size int
		for list.Next() {
			size += len(list.Item().Name)
		}
		_ = size
	}
}

func BenchmarkPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := List{
			Prefix: "xyz/",
			Index:  0,
			Items:  data,
		}
		var size int
		for list.Next() {
			size += len(list.ItemPtr().Name)
		}
		_ = size
	}
}
