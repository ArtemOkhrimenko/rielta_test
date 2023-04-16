# Rielta_Test_app

Http сервис. Умеющий выводить/увеличивать число

## Параметры запуска

Для того, чтобы запустить проект, необходимо выполнить команду в корне проекта ```docker-compose up```

Сервер запускается на порту  ```:8080```

Для того, чтобы завершить работу проекта ```ctrl + c``` после выполнить: ```docker-compose down```

## Примеры.

## Пример для вывода/прибавления: 

Для вывода числа необходимо выполнить запрос:
```
GET  /number/
```

Для увеличения числа необходимо:
```
PUT  /number/
```
И передать число в строке запроса:
```
{
	"number" : 1
}
```
