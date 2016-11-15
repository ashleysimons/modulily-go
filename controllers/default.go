package controllers

import (
	"github.com/astaxie/beego"
)

const (
	contentTypeText = "Text"
	contentTypeMarkdown = "Markdown"
	contentTypeChar = "Char"
	contentTypeChoice = "Choice"
)

type ContentType struct {
	Name   string `json:"Name"`
	Fields []ContentTypeSpec `json:"Fields"`
}
type ContentTypeSpec struct {
	Name    string `json:"Name"`
	Type    string `json:"Type"`
	Options ContentTypeOptions `json:"Options"`
}
type ContentTypeOptions struct {
	Choices map[string]string `json:"Choices"`
}
/**
 * TypesController
 */
type TypesController struct {
	beego.Controller
}

func (c *TypesController) Get() {
	var contentTypes [3]ContentType
	publishedStates := map[string]string{
		"unpublished": "Unpublished",
		"published"  : "Published",
	}
	contentTypes[0] = ContentType{"Page", []ContentTypeSpec{
		{"Title", contentTypeChar, ContentTypeOptions{}},
		{"Body", contentTypeMarkdown, ContentTypeOptions{}},
		{"Published", contentTypeChoice, ContentTypeOptions{Choices: publishedStates}},
	}}
	contentTypes[1] = ContentType{"Post", []ContentTypeSpec{
		{"Title", contentTypeChar, ContentTypeOptions{}},
		{"Summary", contentTypeText, ContentTypeOptions{}},
		{"Body", contentTypeMarkdown, ContentTypeOptions{}},
		{"Published", contentTypeChoice, ContentTypeOptions{Choices: publishedStates}},
	}}
	contentTypes[2] = ContentType{"Panel", []ContentTypeSpec{
		{"Summary", contentTypeText, ContentTypeOptions{}},
		{"Position", contentTypeText, ContentTypeOptions{Choices: map[string]string{"top":"Top", "right":"Right", "bottom":"Bottom"}}},
	}}
	c.Data["json"] = &contentTypes
	c.ServeJSON()
}
