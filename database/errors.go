package database

const (
	UniqueViolation     = "23505" // Нарушение уникального ограничения (duplicate key)
	NotNullViolation    = "23502" // Нарушение NOT NULL ограничения
	ForeignKeyViolation = "23503" // Нарушение FOREIGN KEY
	CheckViolation      = "23514" // CHECK constraint violation
	StringTooLong       = "22001" // Превышена длина столбца
	TableNotFound       = "42P01" // Таблица не найдена
)
