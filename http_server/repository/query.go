package repository

const (
	createUser = `
        INSERT INTO users (id, name, email, password, country, state, role, contact, address, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, ?,?,?,? NOW(), NOW());
    `
	selectUser = `SELECT id, name, email, password, contact, address, created_at, updated_at
		FROM users
		WHERE email = ? LIMIT 1;`
)
