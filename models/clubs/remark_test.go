package clubs

import (
	"database/sql"
	"github.com/lc-tut/club-portal/consts"
)

var remarks = []ClubRemark{
	{1, consts.DummyUUID, 1, 1, sql.NullString{placeRemark1, true}, sql.NullString{"", false}},
	{2, consts.DummyUUID, 2, 2, sql.NullString{"", false}, sql.NullString{"time remark1", true}},
	{3, consts.DummyUUID, 3, 3, sql.NullString{placeRemark2, true}, sql.NullString{"time remark2", true}},
}
