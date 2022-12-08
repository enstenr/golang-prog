package main

func main() {

	type SyncedBuffer struct {
		lock    sync.Mutex
		buffer  bytes.Buffer
	}

	p := new(SyncedBuffer)  // type *SyncedBuffer
var v SyncedBuffer      // type  SyncedBuffer

}