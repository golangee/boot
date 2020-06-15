package helloworld

import (
	"context"
	"github.com/golangee/uuid"
	"time"
)

type Message struct {
	ID        uuid.UUID `ee.sql.Name:"id"`
	CreatedAt time.Time `ee.sql.Name:"created_at"`
	Text      string    `ee.sql.Name:"text"`
}

/**
@ee.sql.Schema("""
	{
		"dialect":"mysql"
	}
	CREATE TABLE IF NOT EXISTS `message` (
	  `id` BINARY(16) NOT NULL COMMENT 'an uuid as the unique key',
	  `text` TEXT NOT NULL COMMENT 'the message text',
	  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'the time of insertion',
	  PRIMARY KEY (`id`))
	ENGINE = InnoDB
	COMMENT = 'contains all hello world messages.'
""")

@ee.stereotype.Repository("msg")
*/
type MessageRepository interface {
	// @ee.sql.Query("SELECT id,created_at,text FROM sms LIMIT :limit")
	FindAll(ctx context.Context, limit int) ([]Message, error)

	// @ee.sql.Query("SELECT id,created_at,text FROM message WHERE id = :id")
	FindById(ctx context.Context, id uuid.UUID) (Message, error)

	// @ee.sql.Query("INSERT INTO message (id,recipient,text) VALUES (:uuid, :text)")
	Create(ctx context.Context, uuid uuid.UUID, text string) error
}