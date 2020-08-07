package notfound

import (
	. "github.com/golangee/forms"
)

type ContentView struct {
	*VStack
}

func NewContentView(path string) *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("the route '"+path+"' is not available").Style(Font(Headline1)),
		NewButton("Index").AddClickListener(func(v View) {
			v.Context().Navigate("/")
		}),
	)}
}

func FromQuery(q Query) View {
	return NewContentView(q.Path())
}
