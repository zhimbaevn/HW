package jsondb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"orders/internal/app/order"
	"os"
	"strings"
)

type Storage struct {
	Path string
}

func (s Storage) String() string {
	return fmt.Sprintf("Path to db: %s", s.Path)
}

// Инициализация хранилища, создание нужных файлов и папок
func GetStorage(p string) (Storage, error) {
	const op = "jsondb.GetStorage"

	//Проверка существования всех папок до конечного файла
	//Создание всех директорий до файла
	dirPath := strings.Split(p, "/")
	dirPath = dirPath[:len(dirPath)-1]
	dir := strings.Join(dirPath, "/")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return Storage{}, fmt.Errorf("%s: %w", op, err)
		}
	}

	//Проверка существования файла
	//Создание файла, если его нет
	_, err := os.Open(p)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f, err := os.Create(p)
			if err != nil {
				return Storage{}, fmt.Errorf("%s: %w", op, err)
			}
			_, err = f.WriteString("{}")
			if err != nil {
				return Storage{}, fmt.Errorf("%s: %w", op, err)
			}
		} else {
			return Storage{}, fmt.Errorf("%s: %w", op, err)
		}
	}

	return Storage{Path: p}, nil
}

func (s Storage) GetData() (*map[int]order.Order, error) {
	const op = "jsondb.GetData"
	data := make(map[int]order.Order)

	f, err := os.ReadFile(s.Path)
	if err != nil {
		return &data, fmt.Errorf("%s: %w", op, err)
	}

	err = json.Unmarshal(f, &data)
	if err != nil {
		return &data, fmt.Errorf("%s: %w", op, err)
	}

	return &data, nil
}

func (s Storage) SaveData(data *map[int]order.Order) error {
	const op = "jsondb.SaveData"

	// Парсинг JSON из полученных данных
	b, err := json.Marshal(&data)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Открывает файл для дальнейшей записи
	file, err := os.OpenFile(s.Path, os.O_RDWR, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer file.Close()

	// Переходим в начало файла, это нужно, чтоб удалить имеющиеся данные
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Обрезаем файл под ноль
	err = file.Truncate(0)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = file.Write(b)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil

}
