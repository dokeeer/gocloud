package models

type UserStorageRelation struct {
	UserID    uint   `json:"user_id"`
	StorageID uint   `json:"storage_id"`
	Role      string `json:"role"`
}
