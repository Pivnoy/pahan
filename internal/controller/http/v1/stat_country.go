package v1

//
//import (
//	"github.com/gin-gonic/gin"
//	"pahan/internal/entity"
//	"pahan/internal/usecase"
//)
//
//type statCountyRoutes struct {
//	t usecase.Country
//}
//
//func newStatCountryRoutes(handler *gin.RouterGroup, t usecase.Country) {
//	r := engineRoutes{t: t}
//
//	handler.GET("/stat_country")
//}
//
//type StructStatCountryResponse struct {
//	Id        int    `json:"id"`
//	Name      string `json:"name"`
//	CarOrders int    `json:"car_orders"`
//	Gdp       int    `json:"gdp"`
//}
//
//type statCountryResponse struct {
//	Stat []StructStatCountryResponse `json:"stat"`
//}
//
//func (r *statCountyRoutes) prepareResponseStatCountry(country entity.Country) StructStatCountryResponse {
//
//}
//
//func (r *statCountyRoutes) getStatCountry(c *gin.Context) {
//	statResList, err := r.t.StatCountry(c.Request.Context())
//	var result []StructStatCountryResponse
//	for _, stat := range statResList {
//		result = append(result, r.prepareResponseStatCountry(stat))
//	}
//}
