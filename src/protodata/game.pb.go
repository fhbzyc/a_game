// Code generated by protoc-gen-go.
// source: game.proto
// DO NOT EDIT!

package protodata

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// *************************************获取玩家数据**********************************
type UserDataRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *UserDataRequest) Reset()         { *m = UserDataRequest{} }
func (m *UserDataRequest) String() string { return proto.CompactTextString(m) }
func (*UserDataRequest) ProtoMessage()    {}

type UserDataResponse struct {
	Role             *RoleData             `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	Items            []*ItemData           `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	Generals         []*GeneralData        `protobuf:"bytes,3,rep,name=generals" json:"generals,omitempty"`
	SignReward       *SignRewardData       `protobuf:"bytes,4,opt,name=signReward" json:"signReward,omitempty"`
	Chapters         []*ChapterData        `protobuf:"bytes,5,rep,name=chapters" json:"chapters,omitempty"`
	TempItemDiamonds []int32               `protobuf:"varint,6,rep,name=tempItemDiamonds" json:"tempItemDiamonds,omitempty"`
	CoinProducts     []*CoinProductData    `protobuf:"bytes,7,rep,name=coinProducts" json:"coinProducts,omitempty"`
	DiamondProducts  []*DiamondProductData `protobuf:"bytes,8,rep,name=diamondProducts" json:"diamondProducts,omitempty"`
	LeaderId         *int32                `protobuf:"varint,9,opt,name=leaderId" json:"leaderId,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *UserDataResponse) Reset()         { *m = UserDataResponse{} }
func (m *UserDataResponse) String() string { return proto.CompactTextString(m) }
func (*UserDataResponse) ProtoMessage()    {}

func (m *UserDataResponse) GetRole() *RoleData {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *UserDataResponse) GetItems() []*ItemData {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *UserDataResponse) GetGenerals() []*GeneralData {
	if m != nil {
		return m.Generals
	}
	return nil
}

func (m *UserDataResponse) GetSignReward() *SignRewardData {
	if m != nil {
		return m.SignReward
	}
	return nil
}

func (m *UserDataResponse) GetChapters() []*ChapterData {
	if m != nil {
		return m.Chapters
	}
	return nil
}

func (m *UserDataResponse) GetTempItemDiamonds() []int32 {
	if m != nil {
		return m.TempItemDiamonds
	}
	return nil
}

func (m *UserDataResponse) GetCoinProducts() []*CoinProductData {
	if m != nil {
		return m.CoinProducts
	}
	return nil
}

func (m *UserDataResponse) GetDiamondProducts() []*DiamondProductData {
	if m != nil {
		return m.DiamondProducts
	}
	return nil
}

func (m *UserDataResponse) GetLeaderId() int32 {
	if m != nil && m.LeaderId != nil {
		return *m.LeaderId
	}
	return 0
}

// *************************************设置队长**********************************
type SetLeaderRequest struct {
	LeaderId         *int32 `protobuf:"varint,1,opt,name=leaderId" json:"leaderId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SetLeaderRequest) Reset()         { *m = SetLeaderRequest{} }
func (m *SetLeaderRequest) String() string { return proto.CompactTextString(m) }
func (*SetLeaderRequest) ProtoMessage()    {}

func (m *SetLeaderRequest) GetLeaderId() int32 {
	if m != nil && m.LeaderId != nil {
		return *m.LeaderId
	}
	return 0
}

type SetLeaderResponse struct {
	LeaderId         *int32 `protobuf:"varint,1,opt,name=leaderId" json:"leaderId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SetLeaderResponse) Reset()         { *m = SetLeaderResponse{} }
func (m *SetLeaderResponse) String() string { return proto.CompactTextString(m) }
func (*SetLeaderResponse) ProtoMessage()    {}

func (m *SetLeaderResponse) GetLeaderId() int32 {
	if m != nil && m.LeaderId != nil {
		return *m.LeaderId
	}
	return 0
}

// *************************************获取签到信息(签到)**********************************
type SignRewardRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SignRewardRequest) Reset()         { *m = SignRewardRequest{} }
func (m *SignRewardRequest) String() string { return proto.CompactTextString(m) }
func (*SignRewardRequest) ProtoMessage()    {}

type SignRewardResponse struct {
	SignReward       *SignRewardData `protobuf:"bytes,1,opt,name=signReward" json:"signReward,omitempty"`
	Role             *RoleData       `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
	General          *GeneralData    `protobuf:"bytes,3,opt,name=general" json:"general,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *SignRewardResponse) Reset()         { *m = SignRewardResponse{} }
func (m *SignRewardResponse) String() string { return proto.CompactTextString(m) }
func (*SignRewardResponse) ProtoMessage()    {}

func (m *SignRewardResponse) GetSignReward() *SignRewardData {
	if m != nil {
		return m.SignReward
	}
	return nil
}

func (m *SignRewardResponse) GetRole() *RoleData {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *SignRewardResponse) GetGeneral() *GeneralData {
	if m != nil {
		return m.General
	}
	return nil
}

// *************************************设置名字*************************************
type SetUpNameRequest struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SetUpNameRequest) Reset()         { *m = SetUpNameRequest{} }
func (m *SetUpNameRequest) String() string { return proto.CompactTextString(m) }
func (*SetUpNameRequest) ProtoMessage()    {}

func (m *SetUpNameRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type SetUpNameResponse struct {
	Role             *RoleData `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *SetUpNameResponse) Reset()         { *m = SetUpNameResponse{} }
func (m *SetUpNameResponse) String() string { return proto.CompactTextString(m) }
func (*SetUpNameResponse) ProtoMessage()    {}

func (m *SetUpNameResponse) GetRole() *RoleData {
	if m != nil {
		return m.Role
	}
	return nil
}

// *************************************随机名字*************************************
type RandomNameRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RandomNameRequest) Reset()         { *m = RandomNameRequest{} }
func (m *RandomNameRequest) String() string { return proto.CompactTextString(m) }
func (*RandomNameRequest) ProtoMessage()    {}

type RandomNameResponse struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RandomNameResponse) Reset()         { *m = RandomNameResponse{} }
func (m *RandomNameResponse) String() string { return proto.CompactTextString(m) }
func (*RandomNameResponse) ProtoMessage()    {}

func (m *RandomNameResponse) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

// *************************************补充体力**********************************
type BuyStaminaRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *BuyStaminaRequest) Reset()         { *m = BuyStaminaRequest{} }
func (m *BuyStaminaRequest) String() string { return proto.CompactTextString(m) }
func (*BuyStaminaRequest) ProtoMessage()    {}

type BuyStaminaResponse struct {
	Role             *RoleData `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	Stamina          *int32    `protobuf:"varint,2,opt,name=stamina" json:"stamina,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *BuyStaminaResponse) Reset()         { *m = BuyStaminaResponse{} }
func (m *BuyStaminaResponse) String() string { return proto.CompactTextString(m) }
func (*BuyStaminaResponse) ProtoMessage()    {}

func (m *BuyStaminaResponse) GetRole() *RoleData {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *BuyStaminaResponse) GetStamina() int32 {
	if m != nil && m.Stamina != nil {
		return *m.Stamina
	}
	return 0
}

// *************************************补充金币**********************************
type BuyCoinRequest struct {
	ProductIndex     *int32 `protobuf:"varint,1,opt,name=productIndex" json:"productIndex,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BuyCoinRequest) Reset()         { *m = BuyCoinRequest{} }
func (m *BuyCoinRequest) String() string { return proto.CompactTextString(m) }
func (*BuyCoinRequest) ProtoMessage()    {}

func (m *BuyCoinRequest) GetProductIndex() int32 {
	if m != nil && m.ProductIndex != nil {
		return *m.ProductIndex
	}
	return 0
}

type BuyCoinResponse struct {
	Role             *RoleData `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	Coin             *int32    `protobuf:"varint,2,opt,name=coin" json:"coin,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *BuyCoinResponse) Reset()         { *m = BuyCoinResponse{} }
func (m *BuyCoinResponse) String() string { return proto.CompactTextString(m) }
func (*BuyCoinResponse) ProtoMessage()    {}

func (m *BuyCoinResponse) GetRole() *RoleData {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *BuyCoinResponse) GetCoin() int32 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

// *************************************补充钻石**********************************
type BuyDiamondRequest struct {
	ProductIndex     *int32 `protobuf:"varint,1,opt,name=productIndex" json:"productIndex,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BuyDiamondRequest) Reset()         { *m = BuyDiamondRequest{} }
func (m *BuyDiamondRequest) String() string { return proto.CompactTextString(m) }
func (*BuyDiamondRequest) ProtoMessage()    {}

func (m *BuyDiamondRequest) GetProductIndex() int32 {
	if m != nil && m.ProductIndex != nil {
		return *m.ProductIndex
	}
	return 0
}

type BuyDiamondResponse struct {
	OrderId          *string `protobuf:"bytes,1,opt,name=orderId" json:"orderId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BuyDiamondResponse) Reset()         { *m = BuyDiamondResponse{} }
func (m *BuyDiamondResponse) String() string { return proto.CompactTextString(m) }
func (*BuyDiamondResponse) ProtoMessage()    {}

func (m *BuyDiamondResponse) GetOrderId() string {
	if m != nil && m.OrderId != nil {
		return *m.OrderId
	}
	return ""
}

// *************************************英雄升级**********************************
type GeneralLevelUpRequest struct {
	GeneralId        *int32 `protobuf:"varint,1,opt,name=generalId" json:"generalId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GeneralLevelUpRequest) Reset()         { *m = GeneralLevelUpRequest{} }
func (m *GeneralLevelUpRequest) String() string { return proto.CompactTextString(m) }
func (*GeneralLevelUpRequest) ProtoMessage()    {}

func (m *GeneralLevelUpRequest) GetGeneralId() int32 {
	if m != nil && m.GeneralId != nil {
		return *m.GeneralId
	}
	return 0
}

type GeneralLevelUpResponse struct {
	Role             *RoleData    `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	General          *GeneralData `protobuf:"bytes,2,opt,name=general" json:"general,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *GeneralLevelUpResponse) Reset()         { *m = GeneralLevelUpResponse{} }
func (m *GeneralLevelUpResponse) String() string { return proto.CompactTextString(m) }
func (*GeneralLevelUpResponse) ProtoMessage()    {}

func (m *GeneralLevelUpResponse) GetRole() *RoleData {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *GeneralLevelUpResponse) GetGeneral() *GeneralData {
	if m != nil {
		return m.General
	}
	return nil
}

// *************************************购买英雄*************************************
type BuyGeneralRequest struct {
	GeneralId        *int32 `protobuf:"varint,1,opt,name=generalId" json:"generalId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BuyGeneralRequest) Reset()         { *m = BuyGeneralRequest{} }
func (m *BuyGeneralRequest) String() string { return proto.CompactTextString(m) }
func (*BuyGeneralRequest) ProtoMessage()    {}

func (m *BuyGeneralRequest) GetGeneralId() int32 {
	if m != nil && m.GeneralId != nil {
		return *m.GeneralId
	}
	return 0
}

type BuyGeneralResponse struct {
	Role             *RoleData    `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
	General          *GeneralData `protobuf:"bytes,1,opt,name=general" json:"general,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *BuyGeneralResponse) Reset()         { *m = BuyGeneralResponse{} }
func (m *BuyGeneralResponse) String() string { return proto.CompactTextString(m) }
func (*BuyGeneralResponse) ProtoMessage()    {}

func (m *BuyGeneralResponse) GetRole() *RoleData {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *BuyGeneralResponse) GetGeneral() *GeneralData {
	if m != nil {
		return m.General
	}
	return nil
}

// *************************************获取邮件列表**********************************
type MailRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MailRequest) Reset()         { *m = MailRequest{} }
func (m *MailRequest) String() string { return proto.CompactTextString(m) }
func (*MailRequest) ProtoMessage()    {}

type MailResponse struct {
	Mails            []*MailData `protobuf:"bytes,1,rep,name=mails" json:"mails,omitempty"`
	Cnnouncement     *string     `protobuf:"bytes,2,opt,name=cnnouncement" json:"cnnouncement,omitempty"`
	Time             *string     `protobuf:"bytes,3,opt,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *MailResponse) Reset()         { *m = MailResponse{} }
func (m *MailResponse) String() string { return proto.CompactTextString(m) }
func (*MailResponse) ProtoMessage()    {}

func (m *MailResponse) GetMails() []*MailData {
	if m != nil {
		return m.Mails
	}
	return nil
}

func (m *MailResponse) GetCnnouncement() string {
	if m != nil && m.Cnnouncement != nil {
		return *m.Cnnouncement
	}
	return ""
}

func (m *MailResponse) GetTime() string {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return ""
}

// *************************************领取邮件奖励**********************************
type MailRewardRequest struct {
	MailId           *int32 `protobuf:"varint,1,opt,name=mailId" json:"mailId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MailRewardRequest) Reset()         { *m = MailRewardRequest{} }
func (m *MailRewardRequest) String() string { return proto.CompactTextString(m) }
func (*MailRewardRequest) ProtoMessage()    {}

func (m *MailRewardRequest) GetMailId() int32 {
	if m != nil && m.MailId != nil {
		return *m.MailId
	}
	return 0
}

type MailRewardResponse struct {
	Role             *RoleData   `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	Reward           *RewardData `protobuf:"bytes,2,opt,name=reward" json:"reward,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *MailRewardResponse) Reset()         { *m = MailRewardResponse{} }
func (m *MailRewardResponse) String() string { return proto.CompactTextString(m) }
func (*MailRewardResponse) ProtoMessage()    {}

func (m *MailRewardResponse) GetRole() *RoleData {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *MailRewardResponse) GetReward() *RewardData {
	if m != nil {
		return m.Reward
	}
	return nil
}

// *************************************道具升级**********************************
type ItemLevelUpRequest struct {
	ItemId           *int32 `protobuf:"varint,1,opt,name=itemId" json:"itemId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ItemLevelUpRequest) Reset()         { *m = ItemLevelUpRequest{} }
func (m *ItemLevelUpRequest) String() string { return proto.CompactTextString(m) }
func (*ItemLevelUpRequest) ProtoMessage()    {}

func (m *ItemLevelUpRequest) GetItemId() int32 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

type ItemLevelUpResponse struct {
	Role             *RoleData `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	Item             *ItemData `protobuf:"bytes,2,opt,name=item" json:"item,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *ItemLevelUpResponse) Reset()         { *m = ItemLevelUpResponse{} }
func (m *ItemLevelUpResponse) String() string { return proto.CompactTextString(m) }
func (*ItemLevelUpResponse) ProtoMessage()    {}

func (m *ItemLevelUpResponse) GetRole() *RoleData {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *ItemLevelUpResponse) GetItem() *ItemData {
	if m != nil {
		return m.Item
	}
	return nil
}

// ************************************战斗初始化************************************
type FightInitRequest struct {
	FightMode        *int32  `protobuf:"varint,1,opt,name=fightMode" json:"fightMode,omitempty"`
	ChapterId        *int32  `protobuf:"varint,2,opt,name=chapterId" json:"chapterId,omitempty"`
	SectionId        *int32  `protobuf:"varint,3,opt,name=sectionId" json:"sectionId,omitempty"`
	TempItems        []int32 `protobuf:"varint,4,rep,name=tempItems" json:"tempItems,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FightInitRequest) Reset()         { *m = FightInitRequest{} }
func (m *FightInitRequest) String() string { return proto.CompactTextString(m) }
func (*FightInitRequest) ProtoMessage()    {}

func (m *FightInitRequest) GetFightMode() int32 {
	if m != nil && m.FightMode != nil {
		return *m.FightMode
	}
	return 0
}

func (m *FightInitRequest) GetChapterId() int32 {
	if m != nil && m.ChapterId != nil {
		return *m.ChapterId
	}
	return 0
}

func (m *FightInitRequest) GetSectionId() int32 {
	if m != nil && m.SectionId != nil {
		return *m.SectionId
	}
	return 0
}

func (m *FightInitRequest) GetTempItems() []int32 {
	if m != nil {
		return m.TempItems
	}
	return nil
}

type FightInitResponse struct {
	// 	optional int32 sceneId = 1;		//场景id 1:城市,2:沙漠,3:冰川
	// 	optional int32 sceneType = 2;		//场景类型 1:中间障碍,2:怪物障碍,3:巨型障碍
	// 	optional int32 winType = 3;		//过关方式 1:宝箱,2:普通,3,计时
	// 	optional int32 pointsNum = 4;		//过关总进度数
	// 	optional string chestData = 6;		//宝箱关卡数据
	Role             *RoleData `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	FightMode        *int32    `protobuf:"varint,2,opt,name=fightMode" json:"fightMode,omitempty"`
	BattleData       *string   `protobuf:"bytes,3,opt,name=battleData" json:"battleData,omitempty"`
	TempItems        []int32   `protobuf:"varint,4,rep,name=tempItems" json:"tempItems,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *FightInitResponse) Reset()         { *m = FightInitResponse{} }
func (m *FightInitResponse) String() string { return proto.CompactTextString(m) }
func (*FightInitResponse) ProtoMessage()    {}

func (m *FightInitResponse) GetRole() *RoleData {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *FightInitResponse) GetFightMode() int32 {
	if m != nil && m.FightMode != nil {
		return *m.FightMode
	}
	return 0
}

func (m *FightInitResponse) GetBattleData() string {
	if m != nil && m.BattleData != nil {
		return *m.BattleData
	}
	return ""
}

func (m *FightInitResponse) GetTempItems() []int32 {
	if m != nil {
		return m.TempItems
	}
	return nil
}

// *************************************战斗结算**********************************
type FightEndRequest struct {
	// optional int32 generalId = 1;		//奖励英雄Id
	KillNum          *int32 `protobuf:"varint,2,opt,name=killNum" json:"killNum,omitempty"`
	CoinNum          *int32 `protobuf:"varint,3,opt,name=coinNum" json:"coinNum,omitempty"`
	DiamondNum       *int32 `protobuf:"varint,4,opt,name=diamondNum" json:"diamondNum,omitempty"`
	IsWin            *bool  `protobuf:"varint,5,opt,name=isWin" json:"isWin,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FightEndRequest) Reset()         { *m = FightEndRequest{} }
func (m *FightEndRequest) String() string { return proto.CompactTextString(m) }
func (*FightEndRequest) ProtoMessage()    {}

func (m *FightEndRequest) GetKillNum() int32 {
	if m != nil && m.KillNum != nil {
		return *m.KillNum
	}
	return 0
}

func (m *FightEndRequest) GetCoinNum() int32 {
	if m != nil && m.CoinNum != nil {
		return *m.CoinNum
	}
	return 0
}

func (m *FightEndRequest) GetDiamondNum() int32 {
	if m != nil && m.DiamondNum != nil {
		return *m.DiamondNum
	}
	return 0
}

func (m *FightEndRequest) GetIsWin() bool {
	if m != nil && m.IsWin != nil {
		return *m.IsWin
	}
	return false
}

type FightEndResponse struct {
	Role             *RoleData      `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	Reward           *RewardData    `protobuf:"bytes,2,opt,name=reward" json:"reward,omitempty"`
	Chapter          []*ChapterData `protobuf:"bytes,3,rep,name=chapter" json:"chapter,omitempty"`
	General          *GeneralData   `protobuf:"bytes,4,opt,name=general" json:"general,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *FightEndResponse) Reset()         { *m = FightEndResponse{} }
func (m *FightEndResponse) String() string { return proto.CompactTextString(m) }
func (*FightEndResponse) ProtoMessage()    {}

func (m *FightEndResponse) GetRole() *RoleData {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *FightEndResponse) GetReward() *RewardData {
	if m != nil {
		return m.Reward
	}
	return nil
}

func (m *FightEndResponse) GetChapter() []*ChapterData {
	if m != nil {
		return m.Chapter
	}
	return nil
}

func (m *FightEndResponse) GetGeneral() *GeneralData {
	if m != nil {
		return m.General
	}
	return nil
}

// *************************************获取好友列表**********************************
type FriendListRequest struct {
	SnsIds           []string `protobuf:"bytes,1,rep,name=snsIds" json:"snsIds,omitempty"`
	PlatId           *int32   `protobuf:"varint,2,opt,name=platId" json:"platId,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *FriendListRequest) Reset()         { *m = FriendListRequest{} }
func (m *FriendListRequest) String() string { return proto.CompactTextString(m) }
func (*FriendListRequest) ProtoMessage()    {}

func (m *FriendListRequest) GetSnsIds() []string {
	if m != nil {
		return m.SnsIds
	}
	return nil
}

func (m *FriendListRequest) GetPlatId() int32 {
	if m != nil && m.PlatId != nil {
		return *m.PlatId
	}
	return 0
}

type FriendListResponse struct {
	FriendList1      []*FriendData `protobuf:"bytes,1,rep,name=friendList1" json:"friendList1,omitempty"`
	FriendList2      []*FriendData `protobuf:"bytes,2,rep,name=friendList2" json:"friendList2,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *FriendListResponse) Reset()         { *m = FriendListResponse{} }
func (m *FriendListResponse) String() string { return proto.CompactTextString(m) }
func (*FriendListResponse) ProtoMessage()    {}

func (m *FriendListResponse) GetFriendList1() []*FriendData {
	if m != nil {
		return m.FriendList1
	}
	return nil
}

func (m *FriendListResponse) GetFriendList2() []*FriendData {
	if m != nil {
		return m.FriendList2
	}
	return nil
}

// *************************************送体力给好友**********************************
type GiveStaminaRequest struct {
	Uid              *int64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GiveStaminaRequest) Reset()         { *m = GiveStaminaRequest{} }
func (m *GiveStaminaRequest) String() string { return proto.CompactTextString(m) }
func (*GiveStaminaRequest) ProtoMessage()    {}

func (m *GiveStaminaRequest) GetUid() int64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

type GiveStaminaResponse struct {
	Uid              *int64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GiveStaminaResponse) Reset()         { *m = GiveStaminaResponse{} }
func (m *GiveStaminaResponse) String() string { return proto.CompactTextString(m) }
func (*GiveStaminaResponse) ProtoMessage()    {}

func (m *GiveStaminaResponse) GetUid() int64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func init() {
}
