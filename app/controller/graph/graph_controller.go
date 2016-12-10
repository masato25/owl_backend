package graph

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	h "github.com/masato25/owl_backend/app/helper"
	m "github.com/masato25/owl_backend/app/model/graph"
	"github.com/masato25/owl_backend/app/utils"
)

func EndpointRegexpQuery(c *gin.Context) {
	q := c.DefaultQuery("q", "")
	if q == "" {
		h.JSONR(c, http.StatusBadRequest, "q is missing")
	} else {
		var endpoint []m.Endpoint
		db.Graph.Table("endpoint").Select("endpoint, id").Where("endpoint regexp ?", q).Scan(&endpoint)
		endpoints := []map[string]interface{}{}
		for _, e := range endpoint {
			endpoints = append(endpoints, map[string]interface{}{"id": e.ID, "endpoint": e.Endpoint})
		}
		h.JSONR(c, endpoints)
	}
	return
}

func EndpointCounterRegexpQuery(c *gin.Context) {
	eid := c.DefaultQuery("eid", "")
	metricQuery := c.DefaultQuery("metricQuery", ".+")
	if eid == "" {
		h.JSONR(c, http.StatusBadRequest, "eid is missing")
	} else {
		eids := utils.ConverIntStringToList(eid)
		if eids == "" {
			h.JSONR(c, http.StatusBadRequest, "input error, please check your input info.")
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
		h.JSONR(c, utils.UniqSet(countersResp))
	}
	return
}
