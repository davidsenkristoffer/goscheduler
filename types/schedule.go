package types

type ScheduleTemplate struct {
	Name     string
	Projects []Project
	Users    []User
}

type Schedule struct {
	Template *ScheduleTemplate
	Values   [][]float32
}

func DummyData() *ScheduleTemplate {

	projects := []Project{{Name: "HVL"}, {Name: "VLFK"}}
	users := []User{{Name: "Per"}, {Name: "Pål"}}
	t := &ScheduleTemplate{Projects: projects, Users: users}
	return t
}

func DummySchedule() *Schedule {
	template := DummyData()
	schedule := &Schedule{
		Template: template,
		Values: [][]float32{
			{1, 1.5}, {2, 3.5},
		},
	}
	return schedule
}
