package database

type File struct {
	ID        int    `gorm:"primaryKey"`
	Ip        string `gorm:"size:32;not null"`
	CreatedAt int64  `gorm:"autoCreateTime;index"`
	FileName  string `gorm:"size:255;not null;unique"`
	UserAgent string `gorm:"size:255;not null"`
	MimeType  string `gorm:"size:255;not null"`
}
