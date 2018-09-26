package smtcp

type LambdaSet struct {
	lambdas map[string]func(*Request)
}

func NewLambdaSet() *LambdaSet {
	return &LambdaSet{lambdas: make(map[string]func(*Request))}
}

func (ls *LambdaSet) Set(name string, f func(*Request)) *LambdaSet {
	ls.lambdas[name] = f
	return ls
}

func (ls *LambdaSet) Get(name string) func(*Request) {
	return ls.lambdas[name]
}
