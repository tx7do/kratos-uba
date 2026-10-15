package task

// CampaignSendTaskType 是营销触达群发任务的类型常量。
const CampaignSendTaskType = "campaign_send"

// CampaignSendTaskData 是营销触达群发任务的载荷。
type CampaignSendTaskData struct {
	CampaignID uint32 `json:"campaignId"`
}
