package services

import (
	"chat/config"
	"chat/internal/models"
	"chat/internal/services/ai"
	"fmt"
	"github.com/Masterminds/squirrel"
	"time"
)

// CreateSession 创建新会话
func CreateSession() (models.Session, error) {
	sql, args, err := squirrel.Insert("session").
		Into("session").
		Columns("title").
		Values("新会话").
		ToSql()
	// 执行SQL
	result, err := config.DB.Exec(sql, args...)
	if err != nil {
		fmt.Println("insert new session err:", err)
	}
	// 获取插入记录的ID
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("get last insert id err:", err)
	}
	newSession := models.Session{
		Id:    lastInsertID,
		Title: "新会话",
	}
	return newSession, err
}

// QuerySessionById 查询会话
func QuerySessionById(id string) (interface{}, error) {
	var session = models.Session{}
	sql, args, err := squirrel.Select("id", "title", "create_time", "update_time").
		From("session").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	// 执行SQL
	errQuerySessionById := config.DB.QueryRow(sql, args...).Scan(
		&session.Id,
		&session.Title,
		&session.CreateTime,
		&session.UpdateTime,
	)
	if errQuerySessionById != nil {
		fmt.Println("query session by id err:", errQuerySessionById)
		return models.Session{}, err
	}
	return session, nil
}

// QueryAllSessions 查询所有会话
func QueryAllSessions() ([]models.Session, error) {
	var sessions []models.Session
	// 分页逻辑
	sql, _, _ := squirrel.Select("id", "title", "create_time", "update_time").
		From("session").
		ToSql()
	rows, err := config.DB.Query(sql)
	if err != nil {
		fmt.Println("query all sessions err:", err)
		return []models.Session{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var session models.Session
		err := rows.Scan(
			&session.Id,
			&session.Title,
			&session.CreateTime,
			&session.UpdateTime,
		)
		if err != nil {
			fmt.Println("query all sessions err:", err)
			return []models.Session{}, err
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}

// DelSession 删除会话
func DelSession(sessionId string) (bool, error) {
	sql, args, _ := squirrel.Delete("session").
		Where(squirrel.Eq{"id": sessionId}).
		ToSql()
	_, err := config.DB.Exec(sql, args...)
	if err != nil {
		fmt.Println("del session err:", err)
		return false, nil
	}
	return true, nil
}

// UpdateSession 更新会话
func UpdateSession(sessionId string, title string) (bool, error) {
	sql, args, _ := squirrel.Update("session").
		Set("title", title).
		Where(squirrel.Eq{"id": sessionId}).
		ToSql()
	_, err := config.DB.Exec(sql, args...)
	if err != nil {
		fmt.Println("update session err:", err)
		return false, nil
	}
	return true, nil
}

// SendMessage 发送消息
func SendMessage(sessionId string, role string, content string) ([]models.Message, error) {
	var messages []models.Message
	// 插入消息
	sql, args, _ := squirrel.Insert("message").
		Columns("session_id", "role", "content", "create_time", "update_time").
		Values(sessionId, role, content, time.Now(), time.Now()).
		ToSql()
	_, err := config.DB.Exec(sql, args...)
	if err != nil {
		fmt.Println("send message err:", err)
		return []models.Message{}, nil
	}
	// 查询消息列表
	sql, _, _ = squirrel.Select("id", "session_id", "role", "content", "create_time", "update_time").
		From("message").
		Where(squirrel.Eq{"session_id": sessionId}).
		ToSql()
	rows, err := config.DB.Query(sql)
	defer rows.Close()
	for rows.Next() {
		var message models.Message
		err := rows.Scan(
			&message.Id,
			&message.SessionId,
			&message.Role,
			&message.Content,
			&message.CreateTime,
			&message.UpdateTime,
		)
		if err != nil {
			fmt.Println("query messages err:", err)
			return []models.Message{}, err
		}
		messages = append(messages, message)
	}
	if err != nil {
		fmt.Println("query messages err:", err)
		return messages, err
	}
	return messages, nil
}

// Ask 处理聊天请求
func Ask(reqContent models.RequestContent) []byte {
	return ai.QueryQWen(reqContent)
}
