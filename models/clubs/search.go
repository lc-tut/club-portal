package clubs

func (p Pages) ToClubSearchResponse() *ClubSearchResponse {
	externalInfo := p.ToExternalInfo()

	return &ClubSearchResponse{Result: externalInfo}
}

type ClubSearchResponse struct {
	Result []ClubPageExternalInfo `json:"result"`
}
