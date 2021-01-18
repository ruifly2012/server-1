package auto

import (
	"e.coding.net/mmstudio/blade/server/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"e.coding.net/mmstudio/blade/server/excel"
)

var	itemEntries	*ItemEntries	//item.xlsx全局变量

// item.xlsx属性表
type ItemEntry struct {
	Id        	int32     	`json:"Id,omitempty"`	//id        
	Name      	string    	`json:"Name,omitempty"`	//名字        
	Desc      	string    	`json:"Desc,omitempty"`	//描述        
	Icon      	string    	`json:"Icon,omitempty"`	//图标        
	Type      	int32     	`json:"Type,omitempty"`	//类型        
	SubType   	int32     	`json:"SubType,omitempty"`	//子类型       
	Quality   	int32     	`json:"Quality,omitempty"`	//品质        
	MaxStack  	int32     	`json:"MaxStack,omitempty"`	//最大堆叠数     
	EffectType	int32     	`json:"EffectType,omitempty"`	//使用效果      
	EffectValue	[]int32   	`json:"EffectValue,omitempty"`	//效果数值      
	EquipEnchantID	int32     	`json:"EquipEnchantID,omitempty"`	//装备强化id    
}

// item.xlsx属性表集合
type ItemEntries struct {
	Rows      	map[int32]*ItemEntry	`json:"Rows,omitempty"`	//          
}

func  init()  {
	excel.AddEntries("item.xlsx", itemEntries)
}

func (e *ItemEntries) Load(excelFileRaw *excel.ExcelFileRaw) error {
	
	itemEntries = &ItemEntries{
		Rows: make(map[int32]*ItemEntry),
	}

	for _, v := range excelFileRaw.CellData {
		entry := &ItemEntry{}
	 	err := mapstructure.Decode(v, entry)
	 	if event, pass := utils.ErrCheck(err, v); !pass {
			event.Msg("decode excel data to struct failed")
	 		return err
	 	}

	 	itemEntries.Rows[entry.Id] = entry
	}

	log.Info().Str("excel_file", excelFileRaw.Filename).Msg("excel load success")
	return nil
	
}

func  GetItemEntry(id int32) (*ItemEntry, bool) {
	entry, ok := itemEntries.Rows[id]
	return entry, ok
}

func  GetItemSize() int32 {
	return int32(len(itemEntries.Rows))
}
