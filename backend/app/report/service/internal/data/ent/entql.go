// Code generated by ent, DO NOT EDIT.

package ent

import (
	"kratos-bi/app/report/service/internal/data/ent/acceptancestatus"
	"kratos-bi/app/report/service/internal/data/ent/realtimewarehousing"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 2)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   acceptancestatus.Table,
			Columns: acceptancestatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: acceptancestatus.FieldID,
			},
		},
		Type: "AcceptanceStatus",
		Fields: map[string]*sqlgraph.FieldSpec{
			acceptancestatus.FieldCreateTime: {Type: field.TypeInt64, Column: acceptancestatus.FieldCreateTime},
			acceptancestatus.FieldUpdateTime: {Type: field.TypeInt64, Column: acceptancestatus.FieldUpdateTime},
			acceptancestatus.FieldDeleteTime: {Type: field.TypeInt64, Column: acceptancestatus.FieldDeleteTime},
			acceptancestatus.FieldName:       {Type: field.TypeString, Column: acceptancestatus.FieldName},
			acceptancestatus.FieldStatus:     {Type: field.TypeString, Column: acceptancestatus.FieldStatus},
			acceptancestatus.FieldAppID:      {Type: field.TypeString, Column: acceptancestatus.FieldAppID},
			acceptancestatus.FieldAppKey:     {Type: field.TypeString, Column: acceptancestatus.FieldAppKey},
			acceptancestatus.FieldRemark:     {Type: field.TypeString, Column: acceptancestatus.FieldRemark},
			acceptancestatus.FieldCreatorID:  {Type: field.TypeUint32, Column: acceptancestatus.FieldCreatorID},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   realtimewarehousing.Table,
			Columns: realtimewarehousing.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: realtimewarehousing.FieldID,
			},
		},
		Type: "RealtimeWarehousing",
		Fields: map[string]*sqlgraph.FieldSpec{
			realtimewarehousing.FieldCreateTime: {Type: field.TypeInt64, Column: realtimewarehousing.FieldCreateTime},
			realtimewarehousing.FieldUpdateTime: {Type: field.TypeInt64, Column: realtimewarehousing.FieldUpdateTime},
			realtimewarehousing.FieldDeleteTime: {Type: field.TypeInt64, Column: realtimewarehousing.FieldDeleteTime},
			realtimewarehousing.FieldName:       {Type: field.TypeString, Column: realtimewarehousing.FieldName},
			realtimewarehousing.FieldStatus:     {Type: field.TypeString, Column: realtimewarehousing.FieldStatus},
			realtimewarehousing.FieldAppID:      {Type: field.TypeString, Column: realtimewarehousing.FieldAppID},
			realtimewarehousing.FieldAppKey:     {Type: field.TypeString, Column: realtimewarehousing.FieldAppKey},
			realtimewarehousing.FieldRemark:     {Type: field.TypeString, Column: realtimewarehousing.FieldRemark},
			realtimewarehousing.FieldCreatorID:  {Type: field.TypeUint32, Column: realtimewarehousing.FieldCreatorID},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (asq *AcceptanceStatusQuery) addPredicate(pred func(s *sql.Selector)) {
	asq.predicates = append(asq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the AcceptanceStatusQuery builder.
func (asq *AcceptanceStatusQuery) Filter() *AcceptanceStatusFilter {
	return &AcceptanceStatusFilter{config: asq.config, predicateAdder: asq}
}

// addPredicate implements the predicateAdder interface.
func (m *AcceptanceStatusMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the AcceptanceStatusMutation builder.
func (m *AcceptanceStatusMutation) Filter() *AcceptanceStatusFilter {
	return &AcceptanceStatusFilter{config: m.config, predicateAdder: m}
}

// AcceptanceStatusFilter provides a generic filtering capability at runtime for AcceptanceStatusQuery.
type AcceptanceStatusFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *AcceptanceStatusFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *AcceptanceStatusFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(acceptancestatus.FieldID))
}

// WhereCreateTime applies the entql int64 predicate on the create_time field.
func (f *AcceptanceStatusFilter) WhereCreateTime(p entql.Int64P) {
	f.Where(p.Field(acceptancestatus.FieldCreateTime))
}

// WhereUpdateTime applies the entql int64 predicate on the update_time field.
func (f *AcceptanceStatusFilter) WhereUpdateTime(p entql.Int64P) {
	f.Where(p.Field(acceptancestatus.FieldUpdateTime))
}

// WhereDeleteTime applies the entql int64 predicate on the delete_time field.
func (f *AcceptanceStatusFilter) WhereDeleteTime(p entql.Int64P) {
	f.Where(p.Field(acceptancestatus.FieldDeleteTime))
}

// WhereName applies the entql string predicate on the name field.
func (f *AcceptanceStatusFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(acceptancestatus.FieldName))
}

// WhereStatus applies the entql string predicate on the status field.
func (f *AcceptanceStatusFilter) WhereStatus(p entql.StringP) {
	f.Where(p.Field(acceptancestatus.FieldStatus))
}

// WhereAppID applies the entql string predicate on the app_id field.
func (f *AcceptanceStatusFilter) WhereAppID(p entql.StringP) {
	f.Where(p.Field(acceptancestatus.FieldAppID))
}

// WhereAppKey applies the entql string predicate on the app_key field.
func (f *AcceptanceStatusFilter) WhereAppKey(p entql.StringP) {
	f.Where(p.Field(acceptancestatus.FieldAppKey))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *AcceptanceStatusFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(acceptancestatus.FieldRemark))
}

// WhereCreatorID applies the entql uint32 predicate on the creator_id field.
func (f *AcceptanceStatusFilter) WhereCreatorID(p entql.Uint32P) {
	f.Where(p.Field(acceptancestatus.FieldCreatorID))
}

// addPredicate implements the predicateAdder interface.
func (rwq *RealtimeWarehousingQuery) addPredicate(pred func(s *sql.Selector)) {
	rwq.predicates = append(rwq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the RealtimeWarehousingQuery builder.
func (rwq *RealtimeWarehousingQuery) Filter() *RealtimeWarehousingFilter {
	return &RealtimeWarehousingFilter{config: rwq.config, predicateAdder: rwq}
}

// addPredicate implements the predicateAdder interface.
func (m *RealtimeWarehousingMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the RealtimeWarehousingMutation builder.
func (m *RealtimeWarehousingMutation) Filter() *RealtimeWarehousingFilter {
	return &RealtimeWarehousingFilter{config: m.config, predicateAdder: m}
}

// RealtimeWarehousingFilter provides a generic filtering capability at runtime for RealtimeWarehousingQuery.
type RealtimeWarehousingFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *RealtimeWarehousingFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *RealtimeWarehousingFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(realtimewarehousing.FieldID))
}

// WhereCreateTime applies the entql int64 predicate on the create_time field.
func (f *RealtimeWarehousingFilter) WhereCreateTime(p entql.Int64P) {
	f.Where(p.Field(realtimewarehousing.FieldCreateTime))
}

// WhereUpdateTime applies the entql int64 predicate on the update_time field.
func (f *RealtimeWarehousingFilter) WhereUpdateTime(p entql.Int64P) {
	f.Where(p.Field(realtimewarehousing.FieldUpdateTime))
}

// WhereDeleteTime applies the entql int64 predicate on the delete_time field.
func (f *RealtimeWarehousingFilter) WhereDeleteTime(p entql.Int64P) {
	f.Where(p.Field(realtimewarehousing.FieldDeleteTime))
}

// WhereName applies the entql string predicate on the name field.
func (f *RealtimeWarehousingFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(realtimewarehousing.FieldName))
}

// WhereStatus applies the entql string predicate on the status field.
func (f *RealtimeWarehousingFilter) WhereStatus(p entql.StringP) {
	f.Where(p.Field(realtimewarehousing.FieldStatus))
}

// WhereAppID applies the entql string predicate on the app_id field.
func (f *RealtimeWarehousingFilter) WhereAppID(p entql.StringP) {
	f.Where(p.Field(realtimewarehousing.FieldAppID))
}

// WhereAppKey applies the entql string predicate on the app_key field.
func (f *RealtimeWarehousingFilter) WhereAppKey(p entql.StringP) {
	f.Where(p.Field(realtimewarehousing.FieldAppKey))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *RealtimeWarehousingFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(realtimewarehousing.FieldRemark))
}

// WhereCreatorID applies the entql uint32 predicate on the creator_id field.
func (f *RealtimeWarehousingFilter) WhereCreatorID(p entql.Uint32P) {
	f.Where(p.Field(realtimewarehousing.FieldCreatorID))
}
