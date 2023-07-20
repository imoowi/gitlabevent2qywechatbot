package services

import (
	"encoding/json"
	"fmt"
	"gitlab2wechatbot/apps/bot/models"
	"gitlab2wechatbot/global"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
)

var LatestSendTime time.Time
var LatestSendCounter int

func LoopReadMsgChan() {
	limit_time := global.Config.GetInt(`wechatbot.limit_time`)
	limit_rate := global.Config.GetInt(`wechatbot.limit_rate`)
	for {
		fmt.Println(`LatestSendTime=`, LatestSendTime, `;LatestSendCounter=`, LatestSendCounter, `;msgChan.len=`, len(global.MsgChan))
		msg := <-global.MsgChan
		if msg == nil {
			fmt.Println(`msg=`, msg)
			continue
		}
		if LatestSendTime.Add(time.Duration(limit_time) * time.Second).Before(time.Now()) {
			LatestSendTime = time.Now()
			LatestSendCounter = 0
			Send2WechatBotService(msg)
		} else {
			if LatestSendCounter < limit_rate-1 {
				LatestSendCounter++
				Send2WechatBotService(msg)
			} else {
				//管道回收
				go Add2MsgChan(msg)
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func Send2WechatBotService(msg *models.Msg) {
	if msg == nil {
		return
	}
	urlStr := global.Config.GetString(`wechatbot.url`)
	dataByte, _ := json.Marshal(msg)
	payload := strings.NewReader(string(dataByte))

	client := &http.Client{}
	req, err := http.NewRequest(`POST`, urlStr, payload)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// req.Header.Add("version", "1.0.0")
	req.Header.Add("Content-Type", "application/json")

	result, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer result.Body.Close()
}
func Insert2Chan(bot map[string]any) {
	botJson, err := json.Marshal(bot)
	if err != nil {
		return
	}
	botStr := string(botJson)
	msgContent := ``
	mentionedMobileList := make([]string, 0)

	object_kind := gjson.Get(botStr, `object_kind`)
	project_name := gjson.Get(botStr, `project.name`).String()
	path_with_namespace := gjson.Get(botStr, `project.path_with_namespace`).String()
	namespacedotPath := strings.Replace(path_with_namespace, `/`, `.`, -1)
	replaceFromKey := `gitlab.projects.` + namespacedotPath + `.url.replace.from`
	replaceFrom := global.Config.GetString(replaceFromKey)
	replaceToKey := `gitlab.projects.` + namespacedotPath + `.url.replace.to`
	replaceTo := global.Config.GetString(replaceToKey)
	switch object_kind.String() {
	case `push`:
		total_commits_count := gjson.Get(botStr, `total_commits_count`).String()
		if cast.ToInt(total_commits_count) <= 0 {
			return
		}
		msgContent = gjson.Get(botStr, `user_name`).String() + ` 对【` + project_name + `】有新提交,目标分支:` + gjson.Get(botStr, `ref`).String() + `. commit：` + gjson.Get(botStr, `commits.0.title`).String() + `. url：` + strings.Replace(gjson.Get(botStr, `commits.0.url`).String(), replaceFrom, replaceTo, -1) // + `. bot=` + botStr
	case `merge_request`:
		mergeState := gjson.Get(botStr, `object_attributes.state`).String()
		switch mergeState {
		case `opened`:
			action := gjson.Get(botStr, `object_attributes.action`).String()

			msgContent = gjson.Get(botStr, `user.name`).String() + ` 对【` + project_name + `】合并请求-` + action + `, last_commit：` + gjson.Get(botStr, `object_attributes.last_commit.title`).String() + `. url：` + strings.Replace(gjson.Get(botStr, `object_attributes.url`).String(), replaceFrom, replaceTo, -1) // + `. bot=` + botStr
			// gjson.Get(botStr, `object_kind`).String()
			mentionUserMobile := global.Config.GetString(`gitlab.users.` + gjson.Get(botStr, `assignees.0.username`).String())
			mentionedMobileList = append(mentionedMobileList, mentionUserMobile)
		case `merged`:
			msgContent = gjson.Get(botStr, `user.name`).String() + ` 对【` + project_name + `】合并请求-已合并, last_commit：` + gjson.Get(botStr, `object_attributes.last_commit.title`).String() + `. url：` + strings.Replace(gjson.Get(botStr, `object_attributes.url`).String(), replaceFrom, replaceTo, -1) // + `. bot=` + botStr
			// gjson.Get(botStr, `object_kind`).String()
			mentionUserMobile := global.Config.GetString(`gitlab.users.` + gjson.Get(botStr, `assignees.0.username`).String())
			mentionedMobileList = append(mentionedMobileList, mentionUserMobile)
		case `pending`:
			msgContent = gjson.Get(botStr, `user.name`).String() + ` 对【` + project_name + `】提交了新的-pending, last_commit：` + gjson.Get(botStr, `object_attributes.last_commit.title`).String() + `. url：` + strings.Replace(gjson.Get(botStr, `object_attributes.url`).String(), replaceFrom, replaceTo, -1) // + `. bot=` + botStr
			// gjson.Get(botStr, `object_kind`).String()
			mentionUserMobile := global.Config.GetString(`gitlab.users.` + gjson.Get(botStr, `assignees.0.username`).String())
			mentionedMobileList = append(mentionedMobileList, mentionUserMobile)
		default:
			return
		}
	case `pipeline`:
		buildsStatus := gjson.Get(botStr, `builds.0.status`).String()
		if buildsStatus == `success` {
			msgContent = `【` + project_name + `】流水线[` + gjson.Get(botStr, `object_attributes.id`).String() + `]执行成功. url：` + strings.Replace(gjson.Get(botStr, `project.git_http_url`).String(), `.git`, `/-/pipelines/`+gjson.Get(botStr, `object_attributes.id`).String(), -1) // + `. bot=` + botStr
		} else {
			return
		}

	default:
		return
	}
	msg := models.Msg{
		MsgType: `text`,
		Text: models.Text{
			Content:             msgContent,
			MentionedMobileList: mentionedMobileList,
		},
	}
	Add2MsgChan(&msg)
}
func Add2MsgChan(msg *models.Msg) {
	global.MsgChan <- msg
}
