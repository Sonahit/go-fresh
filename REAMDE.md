# Турнир, 1 этап
описание ...

## **25 баллов**
### Сортировка
```
POST /sort
{
    asc=true,
    array=[{"_id": "80808080", "name": "Anton"}, {"_id": "90909090", "name": "Gosha"}]
}
```
### Поиск + Поиск без регистра
```
> GET /search?string=texttext&search=text&match=any
[0, 4]
```
```
> GET /search?string=sampletext&search=text&match=exact
[]
```
### Парсинг сайта
Найти всё по заданному параметру search, на pep8
```
> GET /parse?search=import
```
**Example**
```
PEP:	        8
Title:	        Style Guide for Python Code
Author:	        Guido van Rossum <guido at python.org>, Barry Warsaw <barry at python.org>, Nick Coghlan <ncoghlan at gmail.com>
Status:	        Active
Type:	        Process
Created:	    05-Jul-2001
Post-History:	05-Jul-2001
```
```
{
    "p": ["A style guide is about consistency. Consistency with this style guide is important. Consistency within a project is more important. Consistency within one module or function is the most important."],
    "a": ["Imports"],
    "pre": {"# Correct: import os import sys", "# Wrong: import sys, os"]
    ...
}
```
Выводить начальную информацию.
Выводить ту информацию, которую передали.

## Бонус:
Бонус засчитывается при условии, что выполнены основные задания.
### Парсинг сайта + скачивание (25 баллов)
Сайт python, скачиваем конкретную версию.
```
> GET /download?python=3.8
```
Скачивает себе на АПИ питон последней версии. Генерирует динамическую URL для скачивания и передаёт в ответе ссылку на этот файл
## Как отправить задание на проверку???
Нажмите на кнопку Fork в правом, верхнем углу.
![Click to fork](https://i.imgur.com/oXsDTRI.png)
Склонируйте **только что полученный репозиторий** к себе, и создайте папку, название - ваш ник.
После чего запушьте изменения на гит.
Создайте pull requests для этого репозитария.
![pull requests](https://i.imgur.com/xPYvnj1.png)
Всё.
