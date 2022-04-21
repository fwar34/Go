package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

const (
	LOG     = "/home/feng/bug/1157238220.log.1"
	USER_ID = "/home/feng/bug/userid.txt"
	RET     = "/home/feng/bug/ret.log"
)

func Test() {
	str := `Alice 20 alice@gmail.com`
	re := regexp.MustCompile(`(?P<name>[a-zA-z]+)\s+.*`)
	match := re.FindStringSubmatch(str)
	groups := re.SubexpNames()

	fmt.Printf("%v, %v, %d, %d\n", match, groups, len(match), len(groups))

	for i, name := range groups {
		fmt.Println(i, ":"+name)
	}
}

func Test2() {
	userid := "32372439"
	re_str_offline := `(?P<time>\d{2}:\d{2}:\d{2}).*userOffline_handle.*` + userid + `.*serviceType=0x1,.*channelType=50c`
	f_log, err := os.Open(LOG)
	if err != nil {
		fmt.Println("Open file err:", err)
		return
	}

	if ret, _ := FindLastStrInLog(f_log, re_str_offline); ret { // 能找到 offline 日志
		fmt.Println("Found")
	} else {
		fmt.Println("Not found")
	}
}

func FindExitInLog(file_log *os.File, user_id string) (bool, string) {
	file_log.Seek(0, io.SeekStart)
	// re := regexp.MustCompile(`(?P<time>[0-9]{2}:[0-9]{2}:[0-9]{2}).*BMS user exit conference success.*` + user_id)
	re := regexp.MustCompile(`(?P<time>\d{2}:\d{2}:\d{2}).*BMS user exit conference success.*` + user_id)

	reader := bufio.NewReader(file_log)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		match := re.FindStringSubmatch(string(line))
		if len(match) > 0 {
			groups := re.SubexpNames()
			if groups[1] == "time" {
				fmt.Println(match[1])
				return true, match[1]
			}
		}
	}
	return false, ""
}

func FindLastStrInLog(file_log *os.File, re_str string) (bool, string) {
	file_log.Seek(0, io.SeekStart)
	var ret string
	re := regexp.MustCompile(re_str)
	reader := bufio.NewReader(file_log)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		match := re.FindStringSubmatch(string(line))
		if len(match) > 0 {
			groups := re.SubexpNames()
			if groups[1] == "time" {
				ret = groups[1]
			}
		}
	}

	if len(ret) == 0 {
		return false, ""
	} else {
		return true, ret
	}
}

func main() {
	Test2()
	f_log, err := os.Open(LOG)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f_log.Close()

	f_user, err := os.Open(USER_ID)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f_user.Close()

	f_ret, err := os.OpenFile(RET, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f_ret.Close()

	user_reader := bufio.NewReader(f_user)
	ret_writer := bufio.NewWriter(f_ret)
	for {
		line, _, err := user_reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
		if ret, exit_time := FindExitInLog(f_log, string(line)); ret {
			// fmt.Println(string(line) + " " + exit_time)
			_, err := ret_writer.WriteString(string(line) + " exit at " + exit_time + "\n")
			if err != nil {
				fmt.Println("write file err:", err)
			}
			// ioutil.WriteFile(RET, line, 644)
		} else { // 找不到退会日志
			re_str_offline := `(?P<time>\d{2}:\d{2}:\d{2}).*userOffline_handle.*` + string(line) + `.*serviceType=0x1,.*channelType=50c`
			if ret, last_offline_time := FindLastStrInLog(f_log, re_str_offline); ret { // 能找到 offline 日志
				re_str_online := `(?P<time>\d{2}:\d{2}:\d{2}).*userOnline_handle.*` + string(line) + `.*serviceType=0x1,.*channelType=50c`
				if ret2, last_online_time := FindLastStrInLog(f_log, re_str_online); ret2 {
					if last_offline_time >= last_online_time {
						ret_writer.WriteString(string(line) + " offline at " + last_offline_time)
					} else {
						ret_writer.WriteString(string(line) + " last online at 11:44:35\n")
					}
				} else {
					fmt.Println("can't find online for " + string(line))
				}
			} else {
				_, err := ret_writer.WriteString(string(line) + " no exit and no offline 11:44:35\n")
				if err != nil {
					fmt.Println("write file err:", err)
				}
			}
		}
	}
	ret_writer.Flush() // 没有 Flush 函数，文件不写内容
}
