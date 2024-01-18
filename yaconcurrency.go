// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData [5]string = [5]string{"id1", "id2", "id3", "id4", "id5"}
var dbresults []string = []string{}

func main() {
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbInsert(i)
	}
	wg.Wait()
	fmt.Println("db results", dbresults)
}

func dbInsert(id int) {
	time.Sleep(time.Duration(1) * time.Millisecond)
	fmt.Println("inserting id =>", dbData[id])
	save(id)
	wg.Done()
}

func save(id int) {
	m.Lock()
	dbresults = append(dbresults, dbData[id])
	m.Unlock()
}
