package twitter

// CursoredIDs is a cursored collection of friend ids.
type CursoredIDs struct {
	IDs               []int64 `json:"ids"`
	NextCursor        int64   `json:"next_cursor"`
	NextCursorStr     string  `json:"next_cursor_str"`
	PreviousCursor    int64   `json:"previous_cursor"`
	PreviousCursorStr string  `json:"previous_cursor_str"`
}

// CursoredUsers is a cursored collection of friends.
type CursoredUsers struct {
	Users             []User `json:"users"`
	NextCursor        int64  `json:"next_cursor"`
	NextCursorStr     string `json:"next_cursor_str"`
	PreviousCursor    int64  `json:"previous_cursor"`
	PreviousCursorStr string `json:"previous_cursor_str"`
}
