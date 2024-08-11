package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetTable(c *gin.Context) {

	nameSheet := "Наследники" // @todo add switch
	resp, err := http.Get(fmt.Sprintf("https://sheets.googleapis.com/v4/spreadsheets/%s/values/'%s'!A:I?majorDimension=ROWS&key=%s", s.GoogleTable.TableID, nameSheet, s.GoogleTable.Token))
	if err != nil {
		log.Println(err.Error())
		return
	}
	body, _ := io.ReadAll(resp.Body)
	//fmt.Println(string(body))
	var TableValues struct {
		Values [][]string `json:"values"`
	}
	json.Unmarshal(body, &TableValues)

	for i := range TableValues.Values {
		if len(TableValues.Values[1]) > len(TableValues.Values[i]) {
			var ECount int = len(TableValues.Values[1]) - len(TableValues.Values[i])
			for j := 0; j < ECount; j++ {
				TableValues.Values[i] = append(TableValues.Values[i], "")
			}
		}
	}
	var name string
	var result [][]string
	for i, col := range TableValues.Values {
		var data []string
		for j, v := range col {
			// @todo how understand is change
			/*if strings.Contains(v, "\n") {
				v = strings.ReplaceAll(v, "\n", "<br>")
			}*/
			if i == 0 {
				if v == "" {
					v = name
				} else {
					name = v
				}
			} else {
				if j != 0 {
					if v == "" {
						v = "---"
					}
				}
			}
			data = append(data, v)
		}
		result = append(result, data)
	}

	/*data := []Person{
		{"Alice", 30, "Moscow"}, {"Roman", 18, "Belorus"}, {"Mario", 25, "Rim"},
	}
	*/
	c.HTML(200, "page.html", gin.H{
		"Title": "Person Information",
		"Data":  result[2:],
	})
}
