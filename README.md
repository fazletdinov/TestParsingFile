# TestParsingFile

Сервис принимает на вход несколько файлов и с помощью горутин ведет подсчет слов в указанных файлах, после чего выводит информацию в Stdout
Проект релизован с помощью языка Golang

## Чтобы запустить проект у себя локально, вам необохимо
Склонировать репозиторий

`git clone git@github.com:fazletdinov/TestParsingFile.git`

И запустите команду с нужными вам файлами. Например:

```
go run main.go War_and_Peace.txt Hello.txt
```
Вышеуказанная команда запустит скрипт и выведет в Stdout 10 самых часто встрчающихся слов.

### Пример тела ответа
```
Hello.txt
"you" mentioned 8
"the" mentioned 8
"and" mentioned 7
"i" mentioned 6
"of" mentioned 4
"was" mentioned 4
"she" mentioned 3
"" mentioned 3
"a" mentioned 3
"do" mentioned 3

War_and_Peace.txt
"the" mentioned 8130
"and" mentioned 4948
"to" mentioned 3761
"of" mentioned 3054
"a" mentioned 2610
"he" mentioned 2351
"his" mentioned 2058
"in" mentioned 1976
"was" mentioned 1642
"that" mentioned 1608
```

### Автор
[Idel Fazletdinov - fazletdinov](https://github.com/fazletdinov)
