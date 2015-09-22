package main

// STARTMAIN OMIT
func (t Thing) loop() {
	t.query = make(chan chan Foo) // HL
	for {
		select {
		case resp := <-t.query: // HL
			resp <- t.foo // HL
		}
	}
}

func (t Thing) GetFoo() Foo {
	resp := make(chan Foo) // HL
	t.query <- resp        // HL
	return <-resp          // HL
}

// ENDMAIN OMIT
