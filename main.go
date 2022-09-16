package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mytest/cqhttpServer/module"
	"mytest/cqhttpServer/utils"
	"net/http"
)

func main() {
	var group_list module.GroupList
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://127.0.0.1:5700/get_group_list", nil)
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &group_list)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	type TempList struct {
		GroupId       int64
		GroupInfoList module.GroupMemberList
	}
	var temps []TempList
	for _, v := range group_list.GroupLists {
		var temp TempList
		var group_member_list module.GroupMemberList
		url := fmt.Sprintf("http://127.0.0.1:5700/get_group_member_list?group_id=%v&no_cache=false", v.GroupId)
		resp, err := utils.SendRequest(url, nil, nil, "GET")
		if err != nil {
			fmt.Println("err : ", err)
			continue
		}
		err = json.Unmarshal(resp, &group_member_list)
		if err != nil {
			continue
		}
		temp.GroupId = v.GroupId
		temp.GroupInfoList = group_member_list
		temps = append(temps, temp)
	}
	fmt.Println(temps)
}

//func init() {
//	server := &http.Server{
//		Addr:         ":5700",
//		ReadTimeout:  5 * time.Second,
//		WriteTimeout: 5 * time.Second,
//	}
//	err := server.ListenAndServe()
//	if err != nil {
//		fmt.Println("err : ", err)
//		return
//	}
//}
