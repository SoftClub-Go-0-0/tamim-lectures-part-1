package models

type DB struct {
	Tutors   []Person
	Groups   []Group
	Students []Student
}

func (d *DB) GetTutors() []Person {
	return d.Tutors
}

func (d *DB) GetTutor(id uint) Person {
	for i := 0; i < len(d.Tutors); i++ {
		if d.Tutors[i].ID == id {
			return d.Tutors[i]
		}
	}
	return Person{}
}

func (d *DB) SearchTutor(query string) Person {
	for i := 0; i < len(d.Tutors); i++ {
		if d.Tutors[i].Name == query {
			return d.Tutors[i]
		}
	}
	return Person{}
}

func (d *DB) GetStudents() []Student {
	return d.Students
}

func (d *DB) GetStudent(id uint) Student {
	for i := 0; i < len(d.Students); i++ {
		if d.Students[i].Person.ID == id {
			return d.Students[i]
		}
	}
	return Student{}
}

func (d *DB) SearchStudent(query string) Student {
	for i := 0; i < len(d.Students); i++ {
		if d.Students[i].Person.Name == query {
			return d.Students[i]
		}
	}
	return Student{}
}

func (d *DB) GetGroups() []Group {
	return d.Groups
}

func (d *DB) GetGroup(id uint) Group {
	for i := 0; i < len(d.Groups); i++ {
		if uint(d.Groups[i].ID) == id {
			return d.Groups[i]
		}
	}
	return Group{}
}

func (d *DB) SearchGroup(query string) Group {
	for i := 0; i < len(d.Groups); i++ {
		if d.Groups[i].Title == query {
			return d.Groups[i]
		}
	}
	return Group{}
}
