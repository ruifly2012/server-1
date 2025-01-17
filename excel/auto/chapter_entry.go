package auto

import (
	"github.com/east-eden/server/excel"
	"github.com/east-eden/server/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
)

var chapterEntries *ChapterEntries //Chapter.csv全局变量

// Chapter.csv属性表
type ChapterEntry struct {
	Id            int32   `json:"Id,omitempty"`            // 主键
	PrevChapterId int32   `json:"PrevChapterId,omitempty"` //前置章节id
	ConditionId   int32   `json:"ConditionId,omitempty"`   //解锁条件id
	TotalStar     int32   `json:"TotalStar,omitempty"`     //章节星级总数
	RewardStars   []int32 `json:"RewardStars,omitempty"`   //章节宝箱所需星数
	RewardLootIds []int32 `json:"RewardLootIds,omitempty"` //章节宝箱掉落id
}

// Chapter.csv属性表集合
type ChapterEntries struct {
	Rows map[int32]*ChapterEntry `json:"Rows,omitempty"` //
}

func init() {
	excel.AddEntryLoader("Chapter.csv", (*ChapterEntries)(nil))
}

func (e *ChapterEntries) Load(excelFileRaw *excel.ExcelFileRaw) error {

	chapterEntries = &ChapterEntries{
		Rows: make(map[int32]*ChapterEntry, 100),
	}

	for _, v := range excelFileRaw.CellData {
		entry := &ChapterEntry{}
		err := mapstructure.Decode(v, entry)
		if !utils.ErrCheck(err, "decode excel data to struct failed", v) {
			return err
		}

		chapterEntries.Rows[entry.Id] = entry
	}

	log.Info().Str("excel_file", excelFileRaw.Filename).Msg("excel load success")
	return nil

}

func GetChapterEntry(id int32) (*ChapterEntry, bool) {
	entry, ok := chapterEntries.Rows[id]
	return entry, ok
}

func GetChapterSize() int32 {
	return int32(len(chapterEntries.Rows))
}

func GetChapterRows() map[int32]*ChapterEntry {
	return chapterEntries.Rows
}
