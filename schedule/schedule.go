package goschedule

type ScheduleTemplate struct {
	Name     string
	Projects []Project
	Users    []User
}

type Schedule struct {
	Template *ScheduleTemplate
	Values   [][]float32
}

func dummyData() *ScheduleTemplate {

	projects := []Project{{Name: "HVL"}, {Name: "VLFK"}}
	users := []User{{Name: "Per"}, {Name: "Pål"}}
	t := &ScheduleTemplate{Projects: projects, Users: users}
	return t
}
