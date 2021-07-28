package main

import "fmt"

type (
	MySQL struct {
	}

	PostgreSQL struct {
	}

	DBCOnn interface {
		Query() interface{}
	}

	UsersRepository struct {
		db DBCOnn
	}
)

func (db *MySQL) Query() interface{} {
	return []string{"sammidev", "sam", "dev"}
}

func (db *PostgreSQL) Query() interface{} {
	return map[string]string{
		"a3f69c2b-d153-48fd-b10c-5b641657477a": "sammidev",
		"a4f69c2b-d153-48fd-b10c-5b641657477b": "sam",
		"a5f69c2b-d153-48fd-b10c-5b641657477c": "dev",
	}
}

func (r UsersRepository) GetUsers() []string {
	var users []string
	res := r.db.Query()

	switch res.(type) {
	case map[string]string:
		for _, u := range res.(map[string]string) {
			users = append(users, u)
		}
		return users
	case []string:
		return res.([]string)
	}

	return []string{}
}
func main() {
	mysqlDB := MySQL{}
	postgreDB := PostgreSQL{}

	repo1 := UsersRepository{db: &mysqlDB}
	repo2 := UsersRepository{db: &postgreDB}

	fmt.Println("mySQL : ", repo1.GetUsers())
	fmt.Println("postgreSQL : ", repo2.GetUsers())
}
