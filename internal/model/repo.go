package model

type Repo struct {
	ID        int
	Name      string
	Path      string
	Namespace string
}
type Repos []Repo
