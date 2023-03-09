package main

import "fmt"

func main() {
	var command int
	for {
		fmt.Println("Введите команду 0-9")
		fmt.Printf("\t0 - выход\n")
		fmt.Printf("\t1 - список всех преподавателей\n\t2 - поиск преподавателя по ID\n\t3 - поиск преподавателя по другим параметрам\n")
		fmt.Printf("\t4 - список всех групп\n\t5 - поиск группы по ID\n\t6 - поиск группы по названию\n")
		fmt.Printf("\t7 - список всех студентов\n\t8 - поиск студента по ID\n\t9 - поиск студента по другим параметрам\n")
		fmt.Scan(&command)

		switch command {
		case 1:
			fmt.Println("Список всех преподавателей центра:")
			for _, tutor := range db.GetTutors() {
				fmt.Println(tutor)
			}
		case 2:
			var id uint
			fmt.Println("Введите ID преподавателя:")
			fmt.Scan(&id)
			fmt.Println("Результат поиска:")
			fmt.Println(db.GetTutor(id))
		case 3:
			var query string
			fmt.Println("Введите имя, фамилию или номера телефона преподавателя:")
			fmt.Scan(&query)
			fmt.Println("Результат поиска:")
			fmt.Println(db.SearchTutor(query))
		case 4:
			fmt.Println("Список всех групп центра:")
			for _, group := range db.GetGroups() {
				fmt.Println(group.ID, group.Title)
			}
		case 5:
			var id uint
			fmt.Println("Введите ID группы:")
			fmt.Scan(&id)
			fmt.Println("Результат поиска:")
			group := db.GetGroup(id)
			fmt.Println(group.ID, group.Title)
			fmt.Println("Преподаватель:", group.Tutor)
			fmt.Println("Студенты:")
			for _, student := range group.Students {
				fmt.Println(student)
			}
		case 6:
			var query string
			fmt.Println("Введите название группы:")
			fmt.Scan(&query)
			fmt.Println("Результат поиска:")
			group := db.SearchGroup(query)
			fmt.Println(group.ID, group.Title)
			fmt.Println("Преподаватель:", group.Tutor)
			fmt.Println("Студенты:")
			for _, student := range group.Students {
				fmt.Println(student)
			}
		case 7:
			fmt.Println("Список всех студентов центра:")
			for _, student := range db.GetStudents() {
				fmt.Println(student)
			}
		case 8:
			var id uint
			fmt.Println("Введите ID студента:")
			fmt.Scan(&id)
			fmt.Println("Результат поиска:")
			fmt.Println(db.GetStudent(id))
		case 9:
			var query string
			fmt.Println("Введите имя, фамилию или номера телефона студента:")
			fmt.Scan(&query)
			fmt.Println("Результат поиска:")
			fmt.Println(db.SearchStudent(query))
		case 0:
			fmt.Println("Bye!")
			return
		}
	}
}

type Person struct {
	ID      uint
	Name    string
	Surname string
	Phone   string
}

type Student struct {
	Person         Person
	GroupID        uint
	Grades         map[string]int
	TopicsToWorkOn []string
}
type Group struct {
	ID       uint
	Title    string
	Tutor    Person
	Students []Student
}

type DB struct {
	tutors   []Person
	groups   []Group
	students []Student
}

func (d *DB) GetTutors() []Person {
	return d.tutors
}

func (d *DB) GetTutor(id uint) Person {
	for i := 0; i < len(d.tutors); i++ {
		if d.tutors[i].ID == id {
			return d.tutors[i]
		}
	}
	return Person{}
}

func (d *DB) SearchTutor(query string) Person {
	for i := 0; i < len(d.tutors); i++ {
		if d.tutors[i].Name == query {
			return d.tutors[i]
		}
	}
	return Person{}
}

// group methods

func (d *DB) GetGroups() []Group {
	return d.groups
}

func (d *DB) GetGroup(id uint) Group {
	for i := 0; i < len(d.groups); i++ {
		if uint(d.groups[i].ID) == id {
			return d.groups[i]
		}
	}
	return Group{}
}

func (d *DB) SearchGroup(query string) Group {
	for i := 0; i < len(d.groups); i++ {
		if d.groups[i].Title == query {
			return d.groups[i]
		}
	}
	return Group{}
}

// student methods

func (d *DB) GetStudents() []Student {
	return d.students
}

func (d *DB) GetStudent(id uint) Student {
	for i := 0; i < len(d.students); i++ {
		if d.students[i].Person.ID == id {
			return d.students[i]
		}
	}
	return Student{}
}

func (d *DB) SearchStudent(query string) Student {
	for i := 0; i < len(d.students); i++ {
		if d.students[i].Person.Name == query {
			return d.students[i]
		}
	}
	return Student{}
}

// инициализация базы данных
var db = DB{
	tutors:   tutors,
	groups:   groups,
	students: append(students, groups[2].Students...),
}

var tutors = []Person{
	{
		ID:      1,
		Name:    "Nurullah",
		Surname: "Sulaymonov",
		Phone:   "987654300",
	},
	{
		ID:      2,
		Name:    "Muhammadjon",
		Surname: "Mirzoev",
		Phone:   "987654301",
	},
	{
		ID:      3,
		Name:    "Siyovush",
		Surname: "Shorakhimzoda",
		Phone:   "987654302",
	},
}

var groups = []Group{
	{
		ID:    1,
		Title: "C#",
		Tutor: tutors[0],
		Students: []Student{
			{
				Person: Person{
					ID:      4,
					Name:    "John",
					Surname: "Doe",
					Phone:   "987654303",
				},
				GroupID: 1,
				Grades: map[string]int{
					"week1": 78,
					"week2": 85,
					"week3": 82,
					"week4": 56,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      5,
					Name:    "Oliver",
					Surname: "William",
					Phone:   "987654304",
				},
				GroupID: 1,
				Grades: map[string]int{
					"week1": 43,
					"week2": 56,
					"week3": 68,
					"week4": 77,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      6,
					Name:    "Jack",
					Surname: "Henry",
					Phone:   "987654305",
				},
				GroupID: 1,
				Grades: map[string]int{
					"week1": 98,
					"week2": 89,
					"week3": 90,
					"week4": 77,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      7,
					Name:    "Jackson",
					Surname: "Mateo",
					Phone:   "987654306",
				},
				GroupID: 1,
				Grades: map[string]int{
					"week1": 35,
					"week2": 62,
					"week3": 48,
					"week4": 80,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      8,
					Name:    "Daniel",
					Surname: "Logan",
					Phone:   "987654307",
				},
				GroupID: 1,
				Grades: map[string]int{
					"week1": 96,
					"week2": 42,
					"week3": 80,
					"week4": 84,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      9,
					Name:    "Samuel",
					Surname: "Jacob",
					Phone:   "987654308",
				},
				GroupID: 1,
				Grades: map[string]int{
					"week1": 43,
					"week2": 85,
					"week3": 29,
					"week4": 88,
				},
				TopicsToWorkOn: nil,
			},
		},
	},
	{
		ID:    2,
		Title: "C++",
		Tutor: tutors[1],
		Students: []Student{
			{
				Person: Person{
					ID:      10,
					Name:    "John",
					Surname: "Joseph",
					Phone:   "987654309",
				},
				GroupID: 2,
				Grades: map[string]int{
					"week1": 54,
					"week2": 85,
					"week3": 33,
					"week4": 56,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      11,
					Name:    "David",
					Surname: "Hudson",
					Phone:   "987654310",
				},
				GroupID: 2,
				Grades: map[string]int{
					"week1": 39,
					"week2": 65,
					"week3": 82,
					"week4": 78,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      12,
					Name:    "Jack",
					Surname: "Henry",
					Phone:   "987654311",
				},
				GroupID: 2,
				Grades: map[string]int{
					"week1": 44,
					"week2": 63,
					"week3": 82,
					"week4": 79,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      13,
					Name:    "Leo",
					Surname: "Matthew",
					Phone:   "987654312",
				},
				GroupID: 2,
				Grades: map[string]int{
					"week1": 56,
					"week2": 85,
					"week3": 90,
					"week4": 67,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      14,
					Name:    "Daniel",
					Surname: "Luke",
					Phone:   "987654313",
				},
				GroupID: 2,
				Grades: map[string]int{
					"week1": 68,
					"week2": 85,
					"week3": 82,
					"week4": 78,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      15,
					Name:    "Carter",
					Surname: "Jacob",
					Phone:   "987654314",
				},
				GroupID: 2,
				Grades: map[string]int{
					"week1": 90,
					"week2": 65,
					"week3": 82,
					"week4": 67,
				},
				TopicsToWorkOn: nil,
			},
		},
	},
	{
		ID:    3,
		Title: "Go",
		Tutor: tutors[2],
		Students: []Student{
			{
				Person: Person{
					ID:      16,
					Name:    "Alisher",
					Surname: "Yoqubov",
					Phone:   "987654315",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 86,
					"week2": 80,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      17,
					Name:    "Amir",
					Surname: "Arifjonov",
					Phone:   "987654316",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 76,
					"week2": 68,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      18,
					Name:    "Behruz",
					Surname: "Shodiev",
					Phone:   "987654317",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 68,
					"week2": 62,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      19,
					Name:    "Farhod",
					Surname: "Olimzoda",
					Phone:   "987654318",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 88,
					"week2": 86,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      20,
					Name:    "Foziljon",
					Surname: "Muminov",
					Phone:   "987654319",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 85,
					"week2": 88,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      21,
					Name:    "Alijon",
					Surname: "Boboyorov",
					Phone:   "987654320",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 89,
					"week2": 74,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      22,
					Name:    "Khurshed",
					Surname: "Rahimov",
					Phone:   "987654321",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 96,
					"week2": 99,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      23,
					Name:    "Mehrdod",
					Surname: "Rahmatov",
					Phone:   "987654322",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 96,
					"week2": 94,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      24,
					Name:    "Muhammad",
					Surname: "Hoshimov",
					Phone:   "987654323",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 83,
					"week2": 56,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      25,
					Name:    "Muhammad",
					Surname: "Khujaev",
					Phone:   "987654324",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 91,
					"week2": 88,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      26,
					Name:    "Nozim",
					Surname: "Safarov",
					Phone:   "987654325",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 99,
					"week2": 90,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      27,
					Name:    "Sunatullo",
					Surname: "Gafurov",
					Phone:   "987654326",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 88,
					"week2": 97,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      28,
					Name:    "Tamim",
					Surname: "Orif",
					Phone:   "987654327",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 89,
					"week2": 71,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: Person{
					ID:      29,
					Name:    "Zohira",
					Surname: "Furmal",
					Phone:   "987654328",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 85,
					"week2": 74,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
		},
	},
}

var students = append(groups[0].Students, groups[1].Students...)

// Start working from here
