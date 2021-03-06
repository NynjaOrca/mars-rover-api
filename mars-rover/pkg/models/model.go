package models

type NasaResponse struct {
	Photos []Photos `json:"photos"`
}
type Camera struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	RoverID  int    `json:"rover_id"`
	FullName string `json:"full_name"`
}
type Rover struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	LandingDate string `json:"landing_date"`
	LaunchDate  string `json:"launch_date"`
	Status      string `json:"status"`
}
type Photos struct {
	ID        int    `json:"id"`
	Sol       int    `json:"sol"`
	Camera    Camera `json:"camera"`
	ImgSrc    string `json:"img_src"`
	EarthDate string `json:"earth_date"`
	Rover     Rover  `json:"rover"`
}
