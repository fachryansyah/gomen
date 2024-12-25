package database

func (c *DBModel) NewMigration(models ...interface{}) error {
	db, err := c.OpenDB()
	if err != nil {
		return *err
	}

	return db.AutoMigrate(models...)
}
