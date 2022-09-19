package utils

import (
	"strings"
)

const MENUURL = "http://n3.datasn.io/data/api/v1/n3_chennan/caipu_daquan_1/main/list/"

type Menu struct {
	Name string
	Info string
}

func GetMenuInfo() (err error) {
	var menu_lists []Menu
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
		var menu_info_s Menu
		if find := strings.Contains(v, "<td class="); find {
			list := strings.Split(v, "<td class=")
			if len(list) == 0 {
				continue
			}
			result := GetBetweenStr(list[2], ">", "<")
			name := result[1:]
			var values, menu_infos string

			if find_div := strings.Contains(list[3], "div"); find_div {
				values = GetBetweenStr(list[3], "<div class=\"collapsible collapsed\" title=\"Click to show / hide\">", "</div>")
				menu_infos = values[64:]
			} else {
				values = GetBetweenStr(list[3], ">", "<")
				menu_infos = values[1:]
			}
			menu_info_s.Name = name
			menu_info_s.Info = menu_infos
			menu_lists = append(menu_lists, menu_info_s)
		}
	}

	return err
}
