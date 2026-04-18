package routes

import "github.com/osuTitanic/titanic-go/services/stern/internal/templates"

func BuildLayoutData(title string) templates.LayoutData {
	return templates.LayoutData{
		Title: title,
		// TODO: Add website stats, auth, etc.
	}
}
