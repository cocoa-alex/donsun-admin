package routers

type Login struct {
	BaseRouter
}

func (l *Login) Login() {
	l.TplName = "login.html"
	l.Render()
}
