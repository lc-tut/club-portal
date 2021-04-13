package v1

import (
	"github.com/lc-tut/club-portal/models"
	"github.com/lc-tut/club-portal/repos"
	"github.com/lc-tut/club-portal/utils"
)

func validateToContentArgs(models []models.ContentRequest) []string {
	contents := make([]string, len(models))

	for i, m := range models {
		contents[i] = m.Content
	}

	return contents
}

func validateToLinksArgs(models []models.LinkRequest) []repos.ClubLinkArgs {
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

func validateToScheduleArgs(models []models.ScheduleRequest) []repos.ClubScheduleArgs {
	schedules := make([]repos.ClubScheduleArgs, len(models))

	for i, m := range models {
		schedule := repos.ClubScheduleArgs{
			Month:    m.Month,
			Schedule: m.Schedule,
			Remarks:  utils.NilToEmptyString(m.Remarks),
		}
		schedules[i] = schedule
	}

	return schedules
}

func validateToAchievementArgs(models []models.AchievementRequest) []string {
	achieves := make([]string, len(models))

	for i, m := range models {
		achieves[i] = m.Achievement
	}

	return achieves
}

func validateToImageArgs(models []models.ImageRequest) []string {
	images := make([]string, len(models))

	for i, m := range models {
		images[i] = m.Path
	}

	return images
}

func validateToVideoArgs(models []models.VideoRequest) []string {
	videos := make([]string, len(models))

	for i, m := range models {
		videos[i] = m.Path
	}

	return videos
}

func validateToTimeArgs(models []models.ActivityDetailRequest) []repos.ClubTimeArgs {
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

func validateToPlaceArgs(models []models.ActivityDetailRequest) []repos.ClubPlaceArgs {
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

func validateToRemarkArgs(models []models.ActivityDetailRequest) []repos.ClubRemarkArgs {
	remarks := make([]repos.ClubRemarkArgs, len(models))

	for i, m := range models {
		remark := repos.ClubRemarkArgs{
			TimeID:       m.TimeID,
			PlaceID:      m.PlaceID,
			TimeRemarks:  utils.NilToEmptyString(m.TimeRemark),
			PlaceRemarks: utils.NilToEmptyString(m.PlaceRemark),
		}
		remarks[i] = remark
	}

	return remarks
}

func validateToActivityDetailArgs(models []models.ActivityDetailRequest) []repos.ActivityDetailArgs {
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
