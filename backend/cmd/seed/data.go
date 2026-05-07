package main

import "backend/internal"

func GetTestData() []internal.Application {
	return []internal.Application{
		{
			Company:  "Company1",
			Role:     "Role1",
			Location: "Location1",
			Date:     "2026-04-26",
			Status:   "REJECTED",
		},
		{
			Company:  "Company2",
			Role:     "Role2",
			Location: "Location2",
			Date:     "2026-04-27",
			Status:   "REJECTED",
		},
		{
			Company:  "Company3",
			Role:     "Role3",
			Location: "Location3",
			Date:     "2026-04-28",
			Status:   "AWAITING_INTERVIEW",
		},
		{
			Company:  "Company4",
			Role:     "Role4",
			Location: "Location4",
			Date:     "2026-04-29",
			Status:   "ACCEPTED",
		},
		{
			Company:  "Company5",
			Role:     "Role5",
			Location: "Location5",
			Date:     "",
			Status:   "AWAITING_RESPONSE",
		},
		{
			Company:  "Company6",
			Role:     "Role6",
			Location: "Location6",
			Date:     "2026-04-30",
			Status:   "REJECTED",
		},
		{
			Company:  "Company7",
			Role:     "Role7",
			Location: "Location7",
			Date:     "2026-05-02",
			Status:   "REJECTED",
		},
		{
			Company:  "Company8",
			Role:     "Role8",
			Location: "Location8",
			Date:     "2026-05-02",
			Status:   "NO_RESPONSE",
		},
		{
			Company:  "Company9",
			Role:     "Role9",
			Location: "Location9",
			Date:     "2026-05-02",
			Status:   "AWAITING_RESPONSE",
		},
		{
			Company:  "Company10",
			Role:     "Role10",
			Location: "Location10",
			Date:     "2026-05-03",
			Status:   "NO_RESPONSE",
		},
	}
}
