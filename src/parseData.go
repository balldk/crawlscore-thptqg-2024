package src

import (
	"html"
	"regexp"
	"strconv"
)

func ParseStudent(htmlBody *string, sbd *string) *Student {
	var r = regexp.MustCompile(`(?m)<tr>\s*<td>(.*)<\/td>\s*<td>(.*)<\/td>\s*<\/tr>`)
	matches := r.FindAllStringSubmatch(*htmlBody, -1)

	if len(matches) == 0 {
		return nil
	}
	// fmt.Println("success", *sbd)

	std := &Student{}
	for _, match := range matches {
		subjectName := html.UnescapeString(match[1])
		score, _ := strconv.ParseFloat(match[2], 64)

		switch subjectName {
		case "Toán":
			std.Score.Math = float32(score)
		case "Lí":
			std.Score.Physics = float32(score)
		case "Hóa":
			std.Score.Chemistry = float32(score)
		case "Sinh":
			std.Score.Biology = float32(score)
		case "Văn":
			std.Score.Literature = float32(score)
		case "Sử":
			std.Score.History = float32(score)
		case "Địa":
			std.Score.Geography = float32(score)
		case "Ngoại ngữ":
			std.Score.Language = float32(score)
		case "GDCD":
			std.Score.Civic = float32(score)
		}
	}
	std.SBD = *sbd

	return std
}
