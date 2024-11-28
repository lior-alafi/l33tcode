package models

type Language struct {
	Id          string
	Version     string
	Name        string
	DockerImage string
	CSommand    []string
	Mainfunc    string
}
