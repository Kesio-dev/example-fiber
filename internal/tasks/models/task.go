package models

type TaskModel struct {
	ID          string `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Img         string `db:"img"`
	CreatedAt   string `db:"created_at"`
}

type TaskWithPointsModel struct {
	ID          string `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Img         string `db:"img"`
	CreatedAt   string `db:"created_at"`
	TotalAmount int    `db:"total_amount"`
}

type TaskJSON struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Img         string `json:"img"`
	CreatedAt   string `json:"created_at"`
}

type TaskWithPointsJSON struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Img         string `json:"img"`
	CreatedAt   string `json:"created_at"`
	TotalAmount int    `json:"total_amount"`
}

type TaskItemModel struct {
	ID             string  `db:"id"`
	TaskID         string  `db:"task_id"`
	Title          string  `db:"title"`
	Description    string  `db:"description"`
	Points         float64 `db:"points"`
	ExternalID     string  `db:"external_id"`
	Url            string  `db:"url"`
	ItemType       string  `db:"task_type"`
	CreatedAt      string  `db:"created_at"`
	PhotoUrl       string  `db:"photo_url"`
	ItemStatus     string  `db:"item_status"`
	UserTaskItemID string  `db:"user_task_item_id"`
	UserID         string  `db:"user_id"`
}

type TaskItemJSON struct {
	ID             string  `json:"id"`
	TaskID         string  `json:"task_id"`
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	Points         float64 `json:"points"`
	ExternalID     string  `json:"external_id"`
	Url            string  `json:"url"`
	ItemType       string  `json:"task_type"`
	CreatedAt      string  `json:"created_at"`
	PhotoUrl       string  `json:"photo_url"`
	ItemStatus     string  `json:"item_status"`
	UserTaskItemId string  `json:"user_task_item_id"`
	UserID         string  `json:"user_id"`
}
