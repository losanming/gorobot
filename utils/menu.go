package utils

import (
	"fmt"
	"strings"
)

const MENUURL = "http://n3.datasn.io/data/api/v1/n3_chennan/caipu_daquan_1/main/list/"

func GetMenuInfo() (err error) {
	resp, err := SendRequest(MENUURL, nil, nil, "GET")
	if err != nil {
		return err
	}
	menu_list := strings.Split(string(resp), "tbody")
	if len(menu_list) == 0 || len(menu_list) != 3 {
		return nil
	}

	menu_info := strings.Split(menu_list[1], "<td class=\"row_id saa-head\">")
	if len(menu_info) == 0 {
		return err
	}

	for _, v := range menu_info {
		if find := strings.Contains(v, "<td class="); find {
			list := strings.Split(v, "<td class=")
			if len(list) == 0 {
				continue
			}
			result := GetBetweenStr(list[2], ">", "<")
			name := result[1:]

			if find_div := strings.Contains(list[3], "div"); find_div {
				value := GetBetweenStr(list[3], "<div class=\"collapsible collapsed\" title=\"Click to show / hide\">", "</div>")
				menu_infos := value[1:]
			} else {

			}

			fmt.Println(list, name)
		}
	}

	return err
}
