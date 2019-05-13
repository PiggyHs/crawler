package model

type Food struct {
	DishName    string   //菜名
	DishRes     string   //菜名链接
	Producer    string   //制作人
	Description string   //描述
	Mainrecipe  []Recipe //食材与数量
	Step        []Step   //步骤
}

//步骤
type Step struct {
	Desc    string //步骤描述
	StepNum int    //步骤的顺序
	//ImgRes  string //描述对应的图片路径
}

//食材
type Recipe struct {
	Name string //食材名
	Desc string //描述
}
