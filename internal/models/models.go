package models

//"time"

type Authors struct { // Таблица авторов книг
	AuthorCode int
	Fullname   string
}

type Books struct { // Таблица книг
	BookCode   int
	Book_name  string
	AuthorCode int
	Yearpub    string
	Amount     int
	AuthorName string
}

type Readers struct { // Таблица читателей
	Ticket   int
	Fullname string
	Adress   string
	phone    string
}

type Issues struct { // Таблица выдачи
	Issue_code       int
	BookCode         int
	Issue_date       string
	Return_date      string
	Ticket           int
	Formatted_issue  string
	Formatted_return string
	IsExpired        bool
	Book_name        string
	Reader_Fullname  string
	Reader_Phone     string
}
