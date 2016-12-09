package graph

import (
	"fmt"

	"github.com/gin-gonic/gin"
	m "github.com/masato25/owl_backend/app/model/graph"
	"github.com/masato25/owl_backend/app/utils"
)

func EndpointRegexpQuery(c *gin.Context) {
	q := c.DefaultQuery("q", "")
	if q == "" {
		c.JSON(400, gin.H{
			"error": "q is missing",
		})
	} else {
		var endpoint []m.Endpoint
		db.Graph.Table("endpoint").Select("endpoint, id").Where("endpoint regexp ?", q).Scan(&endpoint)
		endpoints := []map[string]interface{}{}
		for _, e := range endpoint {
			endpoints = append(endpoints, map[string]interface{}{"id": e.ID, "endpoint": e.Endpoint})
		}
		c.JSON(200, endpoints)
	}
	return
}

func EndpointCounterRegexpQuery(c *gin.Context) {
	eid := c.DefaultQuery("eid", "")
	metricQuery := c.DefaultQuery("metricQuery", ".+")
	if eid == "" {
		c.JSON(400, gin.H{
			"error": "eid is missing",
		})
	} else {
		eids := utils.ConverIntStringToList(eid)
		if eids == "" {
			c.JSON(400, gin.H{
				"error": "input error, please check your input info.",
			})
			return
		} else {
			eids = fmt.Sprintf("(%s)", eids)
		}
		var counters []m.EndpointCounter
		db.Graph.Table("endpoint_counter").Select("counter").Where(fmt.Sprintf("endpoint_id IN %s AND counter regexp '%s' ", eids, metricQuery)).Scan(&counters)
		countersResp := []interface{}{}
		for _, c := range counters {
			countersResp = append(countersResp, c.Counter)
		}
		c.JSON(200, utils.Set(countersResp))
	}
	return
}
