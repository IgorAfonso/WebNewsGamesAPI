package articleshandlers

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (r *CreateArticleRequest) CreateArticleValidations() error {
	if r.Title == "" && r.Content == "" {
		return fmt.Errorf("request body empty or malformed")
	}

	if r.Title == "" {return errParamIsRequired("title", "string")}
	if r.Content == "" {return errParamIsRequired("content", "string")}

	return nil
}

type UpdateArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

