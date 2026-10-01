package boardDto

type Board struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	TutorID int64  `json:"tutor_id"`
	//CreatedAt time.Time `json:"created_at"`
	//UpdatedAt time.Time `json:"updated_at"`
}
