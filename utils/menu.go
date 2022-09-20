package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const MENUURL = "http://n3.datasn.io/data/api/v1/n3_chennan/caipu_daquan_1/main/list/"

type Menu struct {
	Name string
	Info string
}

func GetMenuInfo(page int) (menu_lists []Menu, err error) {
	url := fmt.Sprintf("http://n3.datasn.io/data/api/v1/n3_chennan/caipu_daquan_1/main/list/%v/", page)
	resp, err := SendRequest(url, nil, nil, "GET")
	if err != nil {
		return menu_lists, err
	}
	menu_list := strings.Split(string(resp), "tbody")
	if len(menu_list) == 0 || len(menu_list) != 3 {
		return menu_lists, nil
	}

	menu_info := strings.Split(menu_list[1], "<td class=\"row_id saa-head\">")
	if len(menu_info) == 0 {
		return menu_lists, err
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

	return menu_lists, err
}

func GetMenusListByFile() (menu_lists []Menu, err error) {
	file, err := ioutil.ReadFile("Menus/menus")
	if err != nil {
		return menu_lists, err
	}
	err = json.Unmarshal(file, &menu_lists)
	if err != nil {
		return menu_lists, err
	}
	return menu_lists, err
}
