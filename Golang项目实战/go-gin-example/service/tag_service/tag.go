package tag_service

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/export"
	"github.com/EDDYCJY/go-gin-example/pkg/file"
	"github.com/tealeg/xlsx"
	"io"
	"strconv"
	"time"
)

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

func (t *Tag) Add() error {
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}

func (t *Tag) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = t.ModifiedBy
	data["name"] = t.Name
	if t.State >= 0 {
		data["state"] = t.State
	}

	return models.EditTag(t.ID, data)
}

func (t *Tag) Delete() error {
	return models.DeleteTag(t.ID)
}

func (t *Tag) Count() (int, error) {
	return models.GetTagTotal(t.getMaps())
}

func (t *Tag) GetAll() ([]models.Tag, error) {
	var (
		tags, cacheTags []models.Tag
	)

	//cache := cache_service.Tag{
	//	State: t.State,
	//
	//	PageNum:  t.PageNum,
	//	PageSize: t.PageSize,
	//}
	_ = cacheTags

	//key := cache.GetTagsKey()
	//if gredis.Exists(key) {
	//	data, err := gredis.Get(key)
	//	if err != nil {
	//		fmt.Println(err)
	//	} else {
	//		json.Unmarshal(data, &cacheTags)
	//		return cacheTags, nil
	//	}
	//}

	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}

	//gredis.Set(key, tags, 3600)
	return tags, nil
}
func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State >= 0 {
		maps["state"] = t.State
	}

	return maps
}

func (t *Tag) Export() (string, error) {
	tags, err := t.GetAll()
	if err != nil {
		return "", err
	}

	xlsxFile := xlsx.NewFile()
	sheet, err := xlsxFile.AddSheet("标签信息") // 创建一个新的excel表
	if err != nil {
		return "", err
	}

	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	row := sheet.AddRow() // 为excel表添加新的一行

	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell() // 为一行添加新的元素
		cell.Value = title   // 第一行就是表头的各列列名
	}

	for _, v := range tags { // 遍历所有的tag，每一个tag占据excel的一行
		values := []string{
			strconv.Itoa(v.ID),
			v.Name,
			v.CreatedBy,
			strconv.Itoa(v.CreatedOn),
			v.ModifiedBy,
			strconv.Itoa(v.ModifiedOn),
		}

		row = sheet.AddRow() // 每个tag占据一行
		for _, value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}

	time := strconv.Itoa(int(time.Now().Unix()))
	filename := "tags-" + time + ".xlsx" // excel文件名

	fullPath := export.GetExcelFullPath() + filename
	file.IsNotExistMkDir(export.GetExcelFullPath())
	err = xlsxFile.Save(fullPath) // 保存到指定路径
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (t *Tag) Import(r io.Reader) error {
	xlsx, err := excelize.OpenReader(r) // 读取xlsx文件
	if err != nil {
		return err
	}

	rows := xlsx.GetRows("标签信息") // 获取excel表格中的所有行
	for irow, row := range rows {
		if irow > 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}

			models.AddTag(data[1], 1, data[2]) // 将excel文件中每行数据保存在mysql数据库中(name / state / createBy)
		}
	}

	return nil
}
