package entity

type Label struct {
	Id          uint16
	Url         string
	Name        string
	Description string
	Color       string
}

func NewLabel(id uint16, url string, name string, description string, color string) *Label {
	label := new(Label)
	label.Id = id
	label.Url = url
	label.Name = name
	label.Description = description
	label.Color = color
	return label
}
