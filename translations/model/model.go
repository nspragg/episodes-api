package model

type CY struct {
	Availability_am                           string `json:"availability.am"`
	Availability_available_for                string `json:"availability.available-for"`
	Availability_available_for_over_a_year    string `json:"availability.available-for-over-a-year"`
	Availability_available_until              string `json:"availability.available-until"`
	Availability_day                          string `json:"availability.day"`
	Availability_days                         string `json:"availability.days"`
	Availability_expires_today                string `json:"availability.expires-today"`
	Availability_expires_today_suffix         string `json:"availability.expires-today-suffix"`
	Availability_expires_tomorrow             string `json:"availability.expires-tomorrow"`
	Availability_expires_tomorrow_suffix      string `json:"availability.expires-tomorrow-suffix"`
	Availability_expires_tonight              string `json:"availability.expires-tonight"`
	Availability_expires_tonight_suffix       string `json:"availability.expires-tonight-suffix"`
	Availability_fri                          string `json:"availability.fri"`
	Availability_hour                         string `json:"availability.hour"`
	Availability_hours                        string `json:"availability.hours"`
	Availability_left                         string `json:"availability.left"`
	Availability_less_than_1_hour_left        string `json:"availability.less-than-1-hour-left"`
	Availability_mon                          string `json:"availability.mon"`
	Availability_month                        string `json:"availability.month"`
	Availability_months                       string `json:"availability.months"`
	Availability_over_a_year_left             string `json:"availability.over-a-year-left"`
	Availability_pm                           string `json:"availability.pm"`
	Availability_sat                          string `json:"availability.sat"`
	Availability_sun                          string `json:"availability.sun"`
	Availability_thu                          string `json:"availability.thu"`
	Availability_tue                          string `json:"availability.tue"`
	Availability_wed                          string `json:"availability.wed"`
	Availability_weeks                        string `json:"availability.weeks"`
	Category_arts                             string `json:"category.arts"`
	Category_audio_described                  string `json:"category.audio-described"`
	Category_cbbc                             string `json:"category.cbbc"`
	Category_cbeebies                         string `json:"category.cbeebies"`
	Category_comedy                           string `json:"category.comedy"`
	Category_comedy_panel_shows               string `json:"category.comedy-panel-shows"`
	Category_comedy_sitcoms                   string `json:"category.comedy-sitcoms"`
	Category_comedy_stand_up                  string `json:"category.comedy-stand-up"`
	Category_documentaries                    string `json:"category.documentaries"`
	Category_documentaries_arts               string `json:"category.documentaries-arts"`
	Category_documentaries_crime_and_justice  string `json:"category.documentaries-crime-and-justice"`
	Category_documentaries_history            string `json:"category.documentaries-history"`
	Category_documentaries_life_stories       string `json:"category.documentaries-life-stories"`
	Category_documentaries_music              string `json:"category.documentaries-music"`
	Category_documentaries_science_and_nature string `json:"category.documentaries-science-and-nature"`
	Category_drama_and_soaps                  string `json:"category.drama-and-soaps"`
	Category_drama_classic_and_period         string `json:"category.drama-classic-and-period"`
	Category_drama_crime                      string `json:"category.drama-crime"`
	Category_drama_sci_fi_and_fantasy         string `json:"category.drama-sci-fi-and-fantasy"`
	Category_entertainment                    string `json:"category.entertainment"`
	Category_films                            string `json:"category.films"`
	Category_food                             string `json:"category.food"`
	Category_history                          string `json:"category.history"`
	Category_lifestyle                        string `json:"category.lifestyle"`
	Category_music                            string `json:"category.music"`
	Category_news                             string `json:"category.news"`
	Category_northern_ireland                 string `json:"category.northern-ireland"`
	Category_science_and_nature               string `json:"category.science-and-nature"`
	Category_scotland                         string `json:"category.scotland"`
	Category_signed                           string `json:"category.signed"`
	Category_sport                            string `json:"category.sport"`
	Category_wales                            string `json:"category.wales"`
	Duration_min                              string `json:"duration.min"`
	Duration_mins                             string `json:"duration.mins"`
	Groups_popular_synopses_small             string `json:"groups.popular.synopses.small"`
	Groups_popular_title                      string `json:"groups.popular.title"`
	Guidance_has_guidance                     string `json:"guidance.has-guidance"`
	Labels_editorial_all_new                  string `json:"labels.editorial.all-new"`
	Labels_editorial_archive                  string `json:"labels.editorial.archive"`
	Labels_editorial_exclusive                string `json:"labels.editorial.exclusive"`
	Labels_editorial_new_series               string `json:"labels.editorial.new-series"`
	Labels_editorial_premiere                 string `json:"labels.editorial.premiere"`
	Labels_time_last_chance                   string `json:"labels.time.last-chance"`
	Month_apr                                 string `json:"month.apr"`
	Month_aug                                 string `json:"month.aug"`
	Month_dec                                 string `json:"month.dec"`
	Month_feb                                 string `json:"month.feb"`
	Month_jan                                 string `json:"month.jan"`
	Month_jul                                 string `json:"month.jul"`
	Month_jun                                 string `json:"month.jun"`
	Month_mar                                 string `json:"month.mar"`
	Month_may                                 string `json:"month.may"`
	Month_nov                                 string `json:"month.nov"`
	Month_oct                                 string `json:"month.oct"`
	Month_sep                                 string `json:"month.sep"`
}

func NewModel(lang string) interface{} {
	return new(CY)
}
