package model

type Student struct {
	StudentId uint64 `json:"student_id"`
	Name      string `json:"name"`
}
type Lesson struct {
	LessonId      uint64 `json:"lesson_id"`
	Name          string `json:"name"`
	EnrolledCount int    `json:"enrolled_count"`
	Capacity      int    `json:"capacity"`
}
type Enrollment struct {
	Id        uint64 `json:"id"`
	StudentId uint64 `json:"student_id"`
	LessonId  uint64 `json:"lesson_id"`
}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

/*type Request struct {
	Task []struct {
		Id uint64 `json:"id"`
		StudentId uint64 `json:"student_id"`
		LessonId uint64 `json:"lesson_id"`
	}
}

*/
