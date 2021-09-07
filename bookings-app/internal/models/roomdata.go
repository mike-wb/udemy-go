package models

// RoomData holds data for a single room
type RoomData struct {
	Name        string
	Id          string
	Template    string
	Image       string
	Description string
}

var Rooms = map[string]RoomData{
	"generals-quarters": {
		Name:     "General's Quarters",
		Id:       "generals-quarters",
		Template: "generals.page.tmpl",
		Image:    "images/Rosso.jpg",
		Description: `Named after General so and so, the General's Quarters room has a different flavor. Being the only one 
                      with a gres porcelained flooring, is the couples’ favorite, because of its passion expressing colors.
                      Very spacious: combines elegance and handiness.`,
	},
	"majors-suite": {
		Name:     "Major's Suite",
		Id:       "majors-suite",
		Template: "majors.page.tmpl",
		Image:    "images/Viola.jpg",
		Description: `The Major's Suite is one of the largest and spacious rooms of the bed and breakfast. Equipped with
                      beautiful four meters high ceilings, "Roman vaults" wisely illuminated and particularly silent.`,
	},
}
