package auto

import (
	"e.coding.net/mmstudio/blade/server/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"e.coding.net/mmstudio/blade/server/excel"
)

var	equipEnchantEntries	*EquipEnchantEntries	//equipEnchant.xlsx全局变量

// equipEnchant.xlsx属性表
type EquipEnchantEntry struct {
	Id        	int32     	`json:"Id,omitempty"`	//id        
	EquipPos  	int32     	`json:"EquipPos,omitempty"`	//装备位置      
	AttId     	int32     	`json:"AttId,omitempty"`	//属性id      
}

// equipEnchant.xlsx属性表集合
type EquipEnchantEntries struct {
	Rows      	map[int32]*EquipEnchantEntry	`json:"Rows,omitempty"`	//          
}

func  init()  {
	excel.AddEntries("equipEnchant.xlsx", equipEnchantEntries)
}

func (e *EquipEnchantEntries) Load(excelFileRaw *excel.ExcelFileRaw) error {
	
	equipEnchantEntries = &EquipEnchantEntries{
		Rows: make(map[int32]*EquipEnchantEntry),
	}

	for _, v := range excelFileRaw.CellData {
		entry := &EquipEnchantEntry{}
	 	err := mapstructure.Decode(v, entry)
	 	if event, pass := utils.ErrCheck(err, v); !pass {
			event.Msg("decode excel data to struct failed")
	 		return err
	 	}

	 	equipEnchantEntries.Rows[entry.Id] = entry
	}

	log.Info().Str("excel_file", excelFileRaw.Filename).Msg("excel load success")
	return nil
	
}

func  GetEquipEnchantEntry(id int32) (*EquipEnchantEntry, bool) {
	entry, ok := equipEnchantEntries.Rows[id]
	return entry, ok
}

func  GetEquipEnchantSize() int32 {
	return int32(len(equipEnchantEntries.Rows))
}
