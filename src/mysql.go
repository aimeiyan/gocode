package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"
	"strconv"
	"sync"
)

type Feed struct {
	Id          int    "id"
	Author      string "author"
	Link        string "link"
	Tags        string "tags"
	UpdatedTs   int    "updated_ts"
	PublishedTs int    "published_ts"
	RssLinkId   int    "rss_link_id"
	SimHash     int64  "simhash"
}

func main() {

	// f := Feed{}

	// var m map[int]string
	// m[1] = "about"

	// log.Println(m, m == nil, m[1], m)

	// log.Println(cachedTypeFields(reflect.TypeOf(f)))

	// passByValue(f)
	// log.Println(f)

	// fs := make([]Feed, 2)
	// fs[0] = f
	// passByValue2(fs)
	// log.Println("aaaaaaaaaa", fs)

	// passByType(func() interface{} {
	// 	return Feed{}
	// })

	con, err := sql.Open("mysql", "feng:@tcp(192.168.1.101:3306)/rssminer")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	rows, err := con.Query("select * from feeds limit 100")
	if err != nil {
		log.Fatal(err)
	}

	results := getRows(rows, Feed{})
	log.Println(len(results))
}

type field struct {
	// name  string
	index int
	ft    reflect.Type
}

var fieldCache struct {
	sync.RWMutex
	m map[reflect.Type]map[string]field
}

func typeFields(t reflect.Type) map[string]field {
	m := make(map[string]field)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := string(f.Tag)
		if name == "" {
			name = f.Name
		}
		m[name] = field{index: i, ft: f.Type}
	}
	return m
}

func cachedTypeFields(t reflect.Type) map[string]field {
	fieldCache.RLock()
	f := fieldCache.m[t]
	fieldCache.RUnlock()
	if f != nil {
		return f
	}

	f = typeFields(t)
	fieldCache.Lock()
	if fieldCache.m == nil {
		fieldCache.m = make(map[reflect.Type]map[string]field)
	}
	fieldCache.m[t] = f
	fieldCache.Unlock()
	return f
}

func getRows(rows *sql.Rows, sample interface{}) []interface{} {
	t := reflect.TypeOf(sample)
	fields := cachedTypeFields(t)

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	results := make([]interface{}, 0, 32)

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		one := reflect.New(t).Elem()
		for i, col := range values {
			if col != nil {
				if f, p := fields[columns[i]]; p {
					value := string(col)
					switch f.ft.Kind() {
					case reflect.String:
						one.Field(f.index).SetString(value)
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
						reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
						reflect.Uint32, reflect.Uint64:
						if v, err := strconv.ParseInt(value, 10, 64); err == nil {
							one.Field(f.index).SetInt(v)
						}
					}
				}
			}
		}
		results = append(results, one.Interface())
	}
	return results
}
