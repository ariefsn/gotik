package entities

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/ariefsn/gotik/constant"
)

type Video struct {
	ID             string   `json:"id"`
	Height         int      `json:"height"`
	Width          int      `json:"width"`
	Duration       int      `json:"duration"`
	Ratio          string   `json:"ratio"`
	Cover          string   `json:"cover"`
	OriginCover    string   `json:"originCover"`
	DynamicCover   string   `json:"dynamicCover"`
	PlayAddr       string   `json:"playAddr"`
	PlayAddrPublic string   `json:"playAddrPublic"`
	ShareCover     []string `json:"shareCover"`
	ReflowCover    string   `json:"reflowCover"`
	Bitrate        int      `json:"bitrate"`
	EncodedType    string   `json:"encodedType"`
	Format         string   `json:"format"`
	VideoQuality   string   `json:"videoQuality"`
	EncodeUserTag  string   `json:"encodeUserTag"`
}

func (v *Video) SetPlayAddrPublic(token string) {
	// encode base64
	raw := fmt.Sprintf("%s&token=%s", v.PlayAddr, token)
	v.PlayAddrPublic = base64.StdEncoding.EncodeToString([]byte(raw))
}

type Author struct {
	ID              string `json:"id"`
	UniqueID        string `json:"uniqueId"`
	Nickname        string `json:"nickname"`
	AvatarThumb     string `json:"avatarThumb"`
	AvatarMedium    string `json:"avatarMedium"`
	AvatarLarger    string `json:"avatarLarger"`
	Signature       string `json:"signature"`
	Verified        bool   `json:"verified"`
	SecUID          string `json:"secUid"`
	Secret          bool   `json:"secret"`
	Ftc             bool   `json:"ftc"`
	Relation        int    `json:"relation"`
	OpenFavorite    bool   `json:"openFavorite"`
	CommentSetting  int    `json:"commentSetting"`
	DuetSetting     int    `json:"duetSetting"`
	StitchSetting   int    `json:"stitchSetting"`
	PrivateAccount  bool   `json:"privateAccount"`
	DownloadSetting int    `json:"downloadSetting"`
}

type Extra struct {
	Now             int64         `json:"now"`
	Logid           string        `json:"logid"`
	FatalItemIds    []interface{} `json:"fatal_item_ids"`
	SearchRequestID string        `json:"search_request_id"`
	APIDebugInfo    interface{}   `json:"api_debug_info"`
}

type LogPb struct {
	ImprID string `json:"impr_id"`
}

type GlobalDoodleConfig struct {
	FeedbackSurvey interface{} `json:"feedback_survey"`
}

type Music struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	PlayURL     string `json:"playUrl"`
	CoverThumb  string `json:"coverThumb"`
	CoverMedium string `json:"coverMedium"`
	CoverLarge  string `json:"coverLarge"`
	AuthorName  string `json:"authorName"`
	Original    bool   `json:"original"`
	Duration    int    `json:"duration"`
	Album       string `json:"album"`
}

type DuetInfo struct {
	DuetFromID string `json:"duetFromId"`
}

type Stats struct {
	DiggCount    int `json:"diggCount"`
	ShareCount   int `json:"shareCount"`
	CommentCount int `json:"commentCount"`
	PlayCount    int `json:"playCount"`
	CollectCount int `json:"collectCount"`
}

type TextExtra struct {
	AwemeID      string `json:"awemeId"`
	Start        int    `json:"start"`
	End          int    `json:"end"`
	HashtagName  string `json:"hashtagName"`
	HashtagID    string `json:"hashtagId"`
	Type         int    `json:"type"`
	UserID       string `json:"userId"`
	IsCommerce   bool   `json:"isCommerce"`
	UserUniqueID string `json:"userUniqueId"`
	SecUID       string `json:"secUid"`
	SubType      int    `json:"subType"`
}

type AuthorStats struct {
	FollowingCount int `json:"followingCount"`
	FollowerCount  int `json:"followerCount"`
	HeartCount     int `json:"heartCount"`
	VideoCount     int `json:"videoCount"`
	DiggCount      int `json:"diggCount"`
	Heart          int `json:"heart"`
}

type Challenge struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	ProfileThumb  string `json:"profileThumb"`
	ProfileMedium string `json:"profileMedium"`
	ProfileLarger string `json:"profileLarger"`
	CoverThumb    string `json:"coverThumb"`
	CoverMedium   string `json:"coverMedium"`
	CoverLarger   string `json:"coverLarger"`
	IsCommerce    bool   `json:"isCommerce"`
}

type Item struct {
	ID                string      `json:"id"`
	Desc              string      `json:"desc"`
	CreateTime        int         `json:"createTime"`
	Video             Video       `json:"video"`
	Author            Author      `json:"author"`
	Music             Music       `json:"music"`
	Challenges        []Challenge `json:"challenges"`
	Stats             Stats       `json:"stats"`
	DuetInfo          DuetInfo    `json:"duetInfo"`
	OriginalItem      bool        `json:"originalItem"`
	OfficalItem       bool        `json:"officalItem"`
	TextExtra         []TextExtra `json:"textExtra"`
	Secret            bool        `json:"secret"`
	ForFriend         bool        `json:"forFriend"`
	Digged            bool        `json:"digged"`
	ItemCommentStatus int         `json:"itemCommentStatus"`
	ShowNotPass       bool        `json:"showNotPass"`
	Vl1               bool        `json:"vl1"`
	ItemMute          bool        `json:"itemMute"`
	AuthorStats       AuthorStats `json:"authorStats"`
	PrivateItem       bool        `json:"privateItem"`
	DuetEnabled       bool        `json:"duetEnabled"`
	StitchEnabled     bool        `json:"stitchEnabled"`
	ShareEnabled      bool        `json:"shareEnabled"`
	IsAd              bool        `json:"isAd"`
	Collected         bool        `json:"collected"`
}

type Common struct {
	DocIDStr string `json:"doc_id_str"`
}

type VideoItem struct {
	Type   int    `json:"type"`
	Item   Item   `json:"item"`
	Common Common `json:"common"`
}

type TiktokSearchResponse struct {
	StatusCode         int                    `json:"status_code"`
	Data               []*VideoItem           `json:"data"`
	Qc                 string                 `json:"qc"`
	Cursor             int                    `json:"cursor"`
	HasMore            int                    `json:"has_more"`
	AdInfo             map[string]interface{} `json:"ad_info"`
	Extra              Extra                  `json:"extra"`
	LogPb              `json:"log_pb"`
	GlobalDoodleConfig GlobalDoodleConfig `json:"global_doodle_config"`
	Backtrace          string             `json:"backtrace"`
}

type TiktokVideoStats struct {
	Stats  constant.TiktokQueueStatus `json:"stats"`
	Token  string                     `json:"token"`
	Expiry int64                      `json:"expiry"`
}

type TiktokService interface {
	Search(ctx context.Context, keyword string) ([]*VideoItem, error)
	Get(ctx context.Context, keyword string) ([]*VideoItem, TiktokVideoStats, error)
	GetById(ctx context.Context, id string) (*VideoItem, error)
	SendQueue(ctx context.Context, keyword string)
	GetNotif(ctx context.Context, keyword string) ([]*VideoItem, error)
	GetQueueStatus(ctx context.Context, keyword string) (TiktokVideoStats, error)
	SetQueueStatus(ctx context.Context, keyword string, status constant.TiktokQueueStatus) error
}

type TiktokRepository interface {
	Search(ctx context.Context, keyword string) ([]*VideoItem, error)
	Store(ctx context.Context, keyword string, videos []*VideoItem)
	GetNotif(ctx context.Context, keyword string) ([]*VideoItem, error)
	GetById(ctx context.Context, id string) (*VideoItem, error)
	GetQueueStatus(ctx context.Context, keyword string) (TiktokVideoStats, error)
	SetQueueStatus(ctx context.Context, keyword string, stats TiktokVideoStats) error
}
