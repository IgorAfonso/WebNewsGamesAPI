package articleshandlers

func createArticleValidations(request CreateArticleRequest) error {

	if request.Title == "" {
		return errParamIsRequired("title", "string")
	}

	if request.Content == "" {
		return errParamIsRequired("content", "string")
	}

	return nil
}