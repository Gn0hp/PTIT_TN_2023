package query_params

type GetUserParams struct {
	CommonQueryParams
}

func (p *GetUserParams) IsValid() bool {
	return true
}
