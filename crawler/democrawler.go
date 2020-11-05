package crawler

import (
	"github.com/PuerkitoBio/goquery"
)

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * Demo爬虫
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package crawler
 * @description
 *
 *
 */

type DemoCrawler struct{ Document *goquery.Document }

//func NewHlCrawler(Response *http.Response,Document *goquery.Document) *HlCrawler {
func NewDemoCrawler() *DemoCrawler    { return &DemoCrawler{} }
func (this *DemoCrawler) Match() bool { return true }
func (this *DemoCrawler) Nav()        {}
func (this *DemoCrawler) List()       {}
func (this *DemoCrawler) Live()       {}
