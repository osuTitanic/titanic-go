package templates

type LayoutData struct {
	Title string
}

type HomeView struct {
	Layout  LayoutData
	Heading string
	Message string
}
