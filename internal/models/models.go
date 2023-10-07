package models

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
	Issue_code    int
	BookCode      int
	Issue_date    string
	Delivery_date string
	Ticket        int
}
