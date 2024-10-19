package models

type Cource struct {
	CourseId    int  `json:"id"`
	CourseName  string  `json:"name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

var Cources []Cource = []Cource{
	{
		CourseId:    1,
		CourseName:  "Go Programming",
		CoursePrice: 100,
		Author: &Author{
			Name: "Satwik Biswas",
		},
	},
	{
		CourseId:    2,
		CourseName:  "Python Programming",
		CoursePrice: 120,
		Author: &Author{
			Name: "Jane Smith",
		},
	},
	{
		CourseId:    3,
		CourseName:  "Java Programming",
		CoursePrice: 110,
		Author: &Author{
			Name: "Alice Johnson",
		},
	},
	{
		CourseId:    4,
		CourseName:  "C++ Programming",
		CoursePrice: 130,
		Author: &Author{
			Name: "Bob Brown",
		},
	},
}
