package utils

import (
	"github.com/goccy/go-json"
	clubmodels "github.com/lc-tut/club-portal/models/clubs"
	usermodels "github.com/lc-tut/club-portal/models/users"
	repos "github.com/lc-tut/club-portal/repos/clubs"
	"github.com/lc-tut/club-portal/utils"
)

func ByteSliceToSessionData(b []byte) (*SessionData, error) {
	model := &SessionData{}

	if err := json.Unmarshal(b, model); err != nil {
		return nil, err
	}

	return model, nil
}

func ToUserInfoResponse(users []usermodels.UserInfo) []usermodels.UserResponse {
	res := make([]usermodels.UserResponse, len(users))

	for i, user := range users {
		res[i] = *user.ToUserResponse()
	}

	return res
}

func ValidateToContentArgs(models []clubmodels.ContentRequest) []string {
	contents := make([]string, len(models))

	for i, m := range models {
		contents[i] = m.Content
	}

	return contents
}

func ValidateToLinksArgs(models []clubmodels.LinkRequest) []repos.ClubLinkArgs {
	links := make([]repos.ClubLinkArgs, len(models))

	for i, m := range models {
		link := repos.ClubLinkArgs{
			Label: m.Label,
			URL:   m.URL,
		}
		links[i] = link
	}

	return links
}

func ValidateToScheduleArgs(models []clubmodels.ScheduleRequest) []repos.ClubScheduleArgs {
	schedules := make([]repos.ClubScheduleArgs, len(models))

	for i, m := range models {
		schedule := repos.ClubScheduleArgs{
			Month:    m.Month,
			Schedule: m.Schedule,
		}
		schedules[i] = schedule
	}

	return schedules
}

func ValidateToAchievementArgs(models []clubmodels.AchievementRequest) []string {
	achieves := make([]string, len(models))

	for i, m := range models {
		achieves[i] = m.Achievement
	}

	return achieves
}

func ValidateToImageArgs(models []clubmodels.ImageRequest) []uint32 {
	images := make([]uint32, len(models))

	for i, m := range models {
		images[i] = m.ImageID
	}

	return images
}

func ValidateToVideoArgs(models []clubmodels.VideoRequest) []string {
	videos := make([]string, len(models))

	for i, m := range models {
		videos[i] = m.Path
	}

	return videos
}

func ValidateToTimeArgs(models []clubmodels.ActivityDetailRequest) []repos.ClubTimeArgs {
	times := make([]repos.ClubTimeArgs, len(models))

	for i, m := range models {
		time := repos.ClubTimeArgs{
			TimeID: m.TimeID,
			Date:   m.Date,
			Time:   m.Time,
		}
		times[i] = time
	}

	return times
}

func ValidateToPlaceArgs(models []clubmodels.ActivityDetailRequest) []repos.ClubPlaceArgs {
	places := make([]repos.ClubPlaceArgs, len(models))

	for i, m := range models {
		place := repos.ClubPlaceArgs{
			PlaceID: m.PlaceID,
			Place:   m.Place,
		}
		places[i] = place
	}

	return places
}

func ValidateToTPRemarkArgs(models []clubmodels.ActivityDetailRequest) []repos.ClubTPRemarkArgs {
	remarks := make([]repos.ClubTPRemarkArgs, len(models))

	for i, m := range models {
		remark := repos.ClubTPRemarkArgs{
			TimeID:       m.TimeID,
			PlaceID:      m.PlaceID,
			TimeRemarks:  utils.StringPToString(m.TimeRemark),
			PlaceRemarks: utils.StringPToString(m.PlaceRemark),
		}
		remarks[i] = remark
	}

	return remarks
}

func ValidateToActivityDetailArgs(models []clubmodels.ActivityDetailRequest) []repos.ActivityDetailArgs {
	details := make([]repos.ActivityDetailArgs, len(models))

	for i, m := range models {
		detail := repos.ActivityDetailArgs{
			TimeID:  m.TimeID,
			PlaceID: m.PlaceID,
		}
		details[i] = detail
	}

	return details
}
