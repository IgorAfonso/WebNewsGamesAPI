package newsHandler

func createNewsValidations(request CreateNewsObject) error {

	if request.Title == "" {
		return errParamIsRequired("title", "string")
	}

	if request.FirstContent == "" {
		return errParamIsRequired("first-content", "string")
	}

	return nil
}