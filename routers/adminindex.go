package routers

type AdminIndex struct {
	BaseRouter
}

func (u *AdminIndex) Index() {
	u.TplName = "admin/index.html"
	u.Render()
}
