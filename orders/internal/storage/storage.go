package storage

import "errors"

var (
	//Ошибки баз данных
	ErrOrderNotFound = errors.New("order not found")     // Возвращается, если заказ с таким id не существует
	ErrOrderExist    = errors.New("order already exist") // Возвращается при попытке занести заказ id которого уже есть в базе

	ErrAttemptIssueFewClients = errors.New("attempt to issue orders to multiple clients") // Возвращается при попытке выдаче заказов в с разными client-id за одну попытку
	ErrNoOneIssed             = errors.New("no orders were issued")                       // Возвращается при попытке выдать заказы, если ни один заказ не был выдан
	ErrClientNotOwner         = errors.New("order does not belong to this client")        // Вовзращается если заказ.клиентid не совпадает с передаваемым клиент.id

	ErrOrderNotIssued = errors.New("order was not issued")                                      // Возвращается при попытке возврата заказа, который не выдан
	ErrOrderRefunded  = errors.New("order already refunded")                                    // Возвращается при попытке вернуть заказ, который уже возвращён
	ErrMoreTwoDays    = errors.New("more than two days have passed since the order was issued") // Возвращается при попытке оформить возврат на заказ, которому больше двух дней с момента выдачи
)
