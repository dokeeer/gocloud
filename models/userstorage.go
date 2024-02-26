package models

type UserStorage struct {
	UserID    uint   `json:"user_id"`
	StorageID uint   `json:"storage_id"`
	Role      string `json:"role"`
}
