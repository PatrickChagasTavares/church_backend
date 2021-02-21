package do

// Result é o modelo padrão para retorno da func Do
type Result struct {
	Data  interface{}
	Error error
}

// ChanResult canal para retorno do resultado
type ChanResult chan *Result

// Do deixa a camada trabalhando de maneira async
func Do(fn func(*Result)) ChanResult {
	ch := make(chan *Result)

	go func() {
		r := new(Result)

		fn(r)

		ch <- r
	}()

	return ch
}
