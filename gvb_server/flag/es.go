package flag

import "gvb_server/models"

func ESCreateIndex() {
	// models.ArticleModel{}.CreateIndex()
	models.FullTextModel{}.CreateIndex()

}
func ESRemoveIndex() {
	models.ArticleModel{}.RemoveIndex()
	models.FullTextModel{}.RemoveIndex()
}