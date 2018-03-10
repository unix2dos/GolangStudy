多态

```
type P interface {
	Say()
}
type P1 struct{}
type P2 struct{}

func (p *P1) Say() { fmt.Println("say p1") }
func (p *P2) Say() { fmt.Println("say p2") }

type AAA struct{ P }

func main() {
	//aaa := &AAA{}
	//aaa.P = &P1{}
	//aaa.Say()
	//aaa.P = &P2{}
	//aaa.Say()

	p1 := &P1{}
	p2 := &P2{}

	var p P
	p = p1
	p.Say()
	p = p2
	p.Say()
}
```
