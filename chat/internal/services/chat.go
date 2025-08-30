package services

import (
	"chat/config"
	"chat/internal/models"
	"chat/internal/services/ai"
	"fmt"
)

// CreateSession 创建新会话
func CreateSession() (models.Session, error) {
	sql := "INSERT INTO session(title) VALUES(?)"
	// 执行SQL
	result, err := config.DB.Exec(sql, "新会话")
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
	sql := "select id,title,create_time,update_time from session where id = ?;"
	// 执行SQL
	err := config.DB.QueryRow(sql, id).Scan(
		&session.Id,
		&session.Title,
		&session.CreateTime,
		&session.UpdateTime,
	)
	if err != nil {
		fmt.Println("query session by id err:", err)
		return models.Session{}, err
	}
	return session, nil
}

// QueryAllSessions 查询所有会话
func QueryAllSessions() ([]models.Session, error) {
	var sessions []models.Session
	// 分页逻辑
	sql := "select id,title,create_time,update_time from session;"
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
	var sql = "delete from session where id = ?;"
	_, err := config.DB.Exec(sql, sessionId)
	if err != nil {
		fmt.Println("del session err:", err)
		return false, nil
	}
	return true, nil
}

// Ask 处理聊天请求
func Ask(reqContent models.RequestContent) []byte {
	return ai.QueryQWen(reqContent)
}
