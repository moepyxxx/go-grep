package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)


func main() {
	search := flag.String("search","","name.")
	file := flag.String("file", "", "file.")
	// reg := flag.String("reg", "", "Regexp.")
	kubetsu := flag.Bool("i", false, "大文字と小文字を区別")
    flag.Parse()

	// 1: searchtext処理
	// lineNumber, err := searchKeywordFromFile(*search, *file)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(lineNumber)

	// 2: fileList処理
	// scanner := bufio.NewScanner(os.Stdin)
	// var fileList []string
	// for scanner.Scan() {
	// 	fileList = append(fileList, scanner.Text())
	// }
	// if err := scanner.Err(); err != nil {
	// 	fmt.Printf("error")
	// 	return
	// }
	// fmt.Println(searchFiles(*search, fileList))

	// 3: 正規表現を処理
	// lineNumber, err := searchRegexpFiles(*reg, *file)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(lineNumber)

	// 4: 大文字と小文字を区別するかどうかを入れる処理
	fmt.Println(*kubetsu, "くべつするか？")
	lineNumber, err := searchKeywordFromFileAddKubetsuFlag(*search, *file, *kubetsu)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(lineNumber)
}

func searchRegexpFiles(searchText string, filePath string) (int, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return 0, err
	}
	lines := strings.Split(string(content), "\n")

	pattern := regexp.MustCompile(searchText)
	for i, line := range lines {
		if pattern.MatchString(line) {
			return i + 1, nil
		}
	}
	return -1, nil
}

func searchFiles(searchText string, fileList []string) []string {
	matchingFiles := []string{}
	for _, file := range fileList {
		if strings.Contains(file, searchText) {
			matchingFiles = append(matchingFiles, file)
		}
	}
	return matchingFiles
}

// func searchKeywordFromFile(searchText string, filePath string) (int, error) {
// 	content, err := ioutil.ReadFile(filePath)
// 	if err != nil {
// 		return 0, err
// 	}

// 	lines := strings.Split(string(content), "\n")

// 	for i, line := range lines {
// 		if strings.Contains(line, searchText) {
// 			return i + 1, nil
// 		}
// 	}

// 	return -1, nil
// }

func searchKeywordFromFileAddKubetsuFlag(searchText string, filePath string, kubetsu bool) (int, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(content), "\n")

	if (kubetsu) {
		for i, line := range lines {
			if strings.Contains(line, searchText) {
				return i + 1, nil
			}
		}
	} else {
		for i, line := range lines {
			if strings.Contains(strings.ToLower(line), strings.ToLower(searchText)) {
				return i + 1, nil
			}
		}
	}

	return -1, nil
}
