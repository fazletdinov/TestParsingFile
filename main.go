package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type CountWords struct {
	Key   string
	Value int
}

func main() {
	paths := parseCommand()
	var wg sync.WaitGroup
	for i := 0; i < len(paths); i++ {
		wg.Add(1)
		go readFile(paths[i], &wg)
	}
	wg.Wait()

}

func parseCommand() []string {
	flag.Parse()
	paths := flag.Args()
	if len(paths) == 0 {
		log.Fatal("нужно указать путь к файлу или файлам")
	}
	return paths
}

func readFile(path string, wg *sync.WaitGroup) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("ошибка при открытие файла %v", err)
	}

	defer file.Close()
	defer wg.Done()

	addCountWords(file)
}

func addCountWords(file *os.File) {
	data := make(map[string]int, 100)
	reader := bufio.NewScanner(file)
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")

	for reader.Scan() {
		array := strings.Split(reader.Text(), " ")
		for _, value := range array {
			if value == "" {
				continue
			}
			data[reg.ReplaceAllString(strings.ToLower(value), "")]++
		}
	}

	sorted_struct := make([]CountWords, 0, len(data))

	for key, value := range data {
		sorted_struct = append(sorted_struct, CountWords{key, value})
	}

	sort.Slice(sorted_struct, func(i, j int) bool {
		return sorted_struct[i].Value > sorted_struct[j].Value
	})

	printCountWords(sorted_struct, file)
}

func printCountWords(data []CountWords, file *os.File) {
	fileName, err := file.Stat()
	if err != nil {
		log.Fatalf("ошибка получения метаданных файла %e", err)
	}

	fmt.Println(fileName.Name())

	for i := 0; i < 10; i++ {
		fmt.Printf("\"%s\" mentioned %d\n", data[i].Key, data[i].Value)
	}
	fmt.Println()
}
