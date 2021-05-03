package valueObject

type Repo struct {
	Id          uint16
	Name        string
	Description string
	Url         string
}

func NewRepo(id uint16, name string, description string, url string) *Repo {
	repo := new(Repo)
	repo.Id = id
	repo.Name = name
	repo.Description = description
	repo.Url = url
	return repo
}
