type T struct {
    val int
}
func (p *T) a() {
    p.val += 1
}
func (p T) b() {
    p.val += 2
}
func main() {
    x := T{5}
    x.a() //makes it 6
    x.b() //makes it 8
    fmt.Println(x.val) //6
}

/*
EXPLANATION :::
  a => is a Pointer Receiver like this.. func (t *Type) Method() {} //pointer receiver 
  But b => is a Value Receiver like this.. func (t Type) Method() {} //value receiver

  If you want to change the state of the receiver in a method, manipulating the value of it, use a pointer receiver.
  It’s not possible with a value receiver, which copies by value.
  Any modification to a value receiver is local to that copy.
*/
