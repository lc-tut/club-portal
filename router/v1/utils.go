package v1

import (
	"github.com/lc-tut/club-portal/models"
	"github.com/lc-tut/club-portal/repos"
	"github.com/lc-tut/club-portal/router/utils"
)

func ValidateToContentArgs(models []models.ClubReqContent) []string {
	contents := make([]string, 0)

	for _, m := range models {
		contents = append(contents, m.Content)
	}

	return contents
}

func ValidateToLinksArgs(models []models.ClubReqLink) []repos.ClubLinkArgs {
	links := make([]repos.ClubLinkArgs, 0)

	for _, m := range models {
		link := repos.ClubLinkArgs{
			Label: m.Label,
			URL:   m.URL,
		}
		links = append(links, link)
	}

	return links
}

func ValidateToScheduleArgs(models []models.ClubReqSchedule) []repos.ClubScheduleArgs {
	schedules := make([]repos.ClubScheduleArgs, 0)

	for _, m := range models {
		schedule := repos.ClubScheduleArgs{
			Month:    m.Month,
			Schedule: m.Schedule,
			Remarks:  utils.NilToEmptyString(m.Remarks),
		}
		schedules = append(schedules, schedule)
	}

	return schedules
}

func ValidateToAchievementArgs(models []models.ClubReqAchievement) []string {
	achieves := make([]string, 0)

	for _, m := range models {
		achieves = append(achieves, m.Achievement)
	}

	return achieves
}

func ValidateToImageArgs(models []models.ClubReqImage) []string {
	images := make([]string, 0)

	for _, m := range models {
		images = append(images, m.Path)
	}

	return images
}

func ValidateToVideoArgs(models []models.ClubReqVideo) []string {
	videos := make([]string, 0)

	for _, m := range models {
		videos = append(videos, m.Path)
	}

	return videos
}
