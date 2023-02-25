package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile() string {
	body, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("ERROR: invalid data format, unable to read file:\n%v", err)
	}
	formated := strings.Replace(string(body), "\n", ", ", -1)
	return formated
}

func allRooms(formated string) []string {
	var res []string
	re := regexp.MustCompile(`([a-zA-z_0-9]+) \d+ \d+`)
	all := re.FindAllStringSubmatch(formated, -1)
	if len(all) > 0 {
		for _, group := range all {
			res = append(res, group[1])
		}
		return res
	}
	return res
}

func antNum(formated string, g *Graph) int {
	var res string
	re := regexp.MustCompile(`^\d+`)
	if len(re.FindStringIndex(formated)) > 0 {
		res = re.FindString(formated)
	} else {
		return 0
	}
	num, _ := strconv.Atoi(res)
	g.ants = num
	return num
}

func startRoom(formated string, g *Graph) string {
	var res string
	re := regexp.MustCompile(`##start, [a-zA-z_0-9]+`)
	all := re.FindStringIndex(formated)
	if len(all) > 0 {
		match := re.FindString(formated)
		res = strings.TrimPrefix(match, "##start, ")
	}
	g.start = res
	return res
}

func endRoom(formated string, g *Graph) string {
	var res string
	re := regexp.MustCompile(`##end, [a-zA-z_0-9]+`)
	all := len(re.FindStringIndex(formated))
	if all > 0 {
		match := re.FindString(formated)
		res = strings.TrimPrefix(match, "##end, ")
	}
	g.end = res
	return res
}

func theLinks(formated string) []string {
	var res []string
	re := regexp.MustCompile(`[a-zA-z_0-9]+-[a-zA-z_0-9]+`)
	all := re.FindAllString(formated, -1)
	if len(all) > 0 {
		for _, group := range all {
			edge := strings.Replace(group, "-", ", ", -1)
			res = append(res, edge)
		}
	}
	return res
}

