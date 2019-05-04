package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var dishNameRes = regexp.MustCompile(
	`<h1 class="recipe_De_title"><a href="https://home.meishichina.com/recipe-[1-9]+.html" id="recipe_title" title="[^"]+">([^<]+)</a>`)

var dishImgRes = regexp.MustCompile(
	`<a class="J_photo" title="[^"]+"><span></span><img src="(https://i8.meishichina.com/attachment/recipe/[0-9]+/[0-9]+/[0-9]+/[0-9a-zA-Z]+.jpg\?x-oss-process=style/p800)" alt="[^"]+"> </a>`)

var producerRes = regexp.MustCompile(
	`<span class="userName" id="recipe_username">([^<]+)</span>`)

var descRes = regexp.MustCompile(
	`<div id="block_txt1"><span class="txt_tart">“</span>([^<]+)<span class="txt_end">” </span>`)

var rePicesRes = regexp.MustCompile(
	`<span class="category_s1">
<a target="_blank" href="[^"]+" title="[^"]+"><b>([^<]+)</b></a>
</span>
<span class="category_s2">([^<]+)</span>`)

//var stepsRes = regexp.MustCompile(
//	`<div class="recipeStep_img">
//<img src="(https://i8.meishichina.com/attachment/recipe/[0-9]+/[0-9]+/[0-9]+/[0-9a-zA-Z]+.jpg\?x-oss-process=style/p320)" alt="[^"]+">
//</div>
//<div class="recipeStep_word"><div class="recipeStep_num">([0-9]+)</div>([^<]+)</div>`)

var stepsRes = regexp.MustCompile(
	`<div class="recipeStep_img">
<img src="(https://i8.meishichina.com/attachment/recipe/[^\.]+.jpg\?x-oss-process=style/p320)" alt="[^"]+">
</div>
<div class="recipeStep_word"><div class="recipeStep_num">([^<]+)</div>([^<]+)</div>`)

func ParseRecipe(contents []byte) engine.ParseResult {
	dish := model.Food{}

	name := extractString(contents, dishNameRes)     //菜名
	producer := extractString(contents, producerRes) //制作者
	desc := extractString(contents, descRes)         //描述
	imgRes := extractString(contents, dishImgRes)    //菜的图片路径
	repices := getAllRepice(contents, rePicesRes)    //食材
	steps := getAllStep(contents, stepsRes)          //步骤

	dish.DishName = name
	dish.Producer = producer
	dish.Description = desc
	dish.DishImgRes = imgRes
	dish.Mainrecipe = repices
	dish.Step = steps

	result := engine.ParseResult{
		Items: []interface{}{dish},
	}
	return result
	//profile.
}

func getAllRepice(contents []byte, re *regexp.Regexp) []model.Recipe {
	recipes := []model.Recipe{}
	match := re.FindAllSubmatch(contents, -1)

	for _, m := range match {
		recipe := model.Recipe{}
		recipe.Name = string(m[1])
		recipe.Desc = string(m[2])
		recipes = append(recipes, recipe)
	}
	return recipes
}

func getAllStep(contents []byte, re *regexp.Regexp) []model.Step {
	steps := []model.Step{}
	match := re.FindAllSubmatch(contents, -1)
	for _, m := range match {
		step := model.Step{}
		step.ImgRes = string(m[1])
		num, err := strconv.Atoi(string(m[2]))
		if err == nil {
			step.StepNum = num
		}
		step.Desc = string(m[3])
		steps = append(steps, step)
	}
	return steps
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
