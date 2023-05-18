// Code generated by ent, DO NOT EDIT.

package ent

import (
	"kratos-bi/app/core/service/internal/data/ent/application"
	"kratos-bi/app/core/service/internal/data/ent/user"

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
			Table:   application.Table,
			Columns: application.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: application.FieldID,
			},
		},
		Type: "Application",
		Fields: map[string]*sqlgraph.FieldSpec{
			application.FieldCreateTime: {Type: field.TypeInt64, Column: application.FieldCreateTime},
			application.FieldUpdateTime: {Type: field.TypeInt64, Column: application.FieldUpdateTime},
			application.FieldDeleteTime: {Type: field.TypeInt64, Column: application.FieldDeleteTime},
			application.FieldName:       {Type: field.TypeString, Column: application.FieldName},
			application.FieldStatus:     {Type: field.TypeString, Column: application.FieldStatus},
			application.FieldAppID:      {Type: field.TypeString, Column: application.FieldAppID},
			application.FieldAppKey:     {Type: field.TypeString, Column: application.FieldAppKey},
			application.FieldRemark:     {Type: field.TypeString, Column: application.FieldRemark},
			application.FieldCreatorID:  {Type: field.TypeUint32, Column: application.FieldCreatorID},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: user.FieldID,
			},
		},
		Type: "User",
		Fields: map[string]*sqlgraph.FieldSpec{
			user.FieldCreateTime:  {Type: field.TypeInt64, Column: user.FieldCreateTime},
			user.FieldUpdateTime:  {Type: field.TypeInt64, Column: user.FieldUpdateTime},
			user.FieldDeleteTime:  {Type: field.TypeInt64, Column: user.FieldDeleteTime},
			user.FieldUserName:    {Type: field.TypeString, Column: user.FieldUserName},
			user.FieldPassword:    {Type: field.TypeString, Column: user.FieldPassword},
			user.FieldNickName:    {Type: field.TypeString, Column: user.FieldNickName},
			user.FieldRealName:    {Type: field.TypeString, Column: user.FieldRealName},
			user.FieldEmail:       {Type: field.TypeString, Column: user.FieldEmail},
			user.FieldPhone:       {Type: field.TypeString, Column: user.FieldPhone},
			user.FieldAvatar:      {Type: field.TypeString, Column: user.FieldAvatar},
			user.FieldDescription: {Type: field.TypeString, Column: user.FieldDescription},
			user.FieldRoleID:      {Type: field.TypeUint32, Column: user.FieldRoleID},
			user.FieldAuthority:   {Type: field.TypeEnum, Column: user.FieldAuthority},
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
func (aq *ApplicationQuery) addPredicate(pred func(s *sql.Selector)) {
	aq.predicates = append(aq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ApplicationQuery builder.
func (aq *ApplicationQuery) Filter() *ApplicationFilter {
	return &ApplicationFilter{config: aq.config, predicateAdder: aq}
}

// addPredicate implements the predicateAdder interface.
func (m *ApplicationMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ApplicationMutation builder.
func (m *ApplicationMutation) Filter() *ApplicationFilter {
	return &ApplicationFilter{config: m.config, predicateAdder: m}
}

// ApplicationFilter provides a generic filtering capability at runtime for ApplicationQuery.
type ApplicationFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *ApplicationFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *ApplicationFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(application.FieldID))
}

// WhereCreateTime applies the entql int64 predicate on the create_time field.
func (f *ApplicationFilter) WhereCreateTime(p entql.Int64P) {
	f.Where(p.Field(application.FieldCreateTime))
}

// WhereUpdateTime applies the entql int64 predicate on the update_time field.
func (f *ApplicationFilter) WhereUpdateTime(p entql.Int64P) {
	f.Where(p.Field(application.FieldUpdateTime))
}

// WhereDeleteTime applies the entql int64 predicate on the delete_time field.
func (f *ApplicationFilter) WhereDeleteTime(p entql.Int64P) {
	f.Where(p.Field(application.FieldDeleteTime))
}

// WhereName applies the entql string predicate on the name field.
func (f *ApplicationFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(application.FieldName))
}

// WhereStatus applies the entql string predicate on the status field.
func (f *ApplicationFilter) WhereStatus(p entql.StringP) {
	f.Where(p.Field(application.FieldStatus))
}

// WhereAppID applies the entql string predicate on the app_id field.
func (f *ApplicationFilter) WhereAppID(p entql.StringP) {
	f.Where(p.Field(application.FieldAppID))
}

// WhereAppKey applies the entql string predicate on the app_key field.
func (f *ApplicationFilter) WhereAppKey(p entql.StringP) {
	f.Where(p.Field(application.FieldAppKey))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *ApplicationFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(application.FieldRemark))
}

// WhereCreatorID applies the entql uint32 predicate on the creator_id field.
func (f *ApplicationFilter) WhereCreatorID(p entql.Uint32P) {
	f.Where(p.Field(application.FieldCreatorID))
}

// addPredicate implements the predicateAdder interface.
func (uq *UserQuery) addPredicate(pred func(s *sql.Selector)) {
	uq.predicates = append(uq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the UserQuery builder.
func (uq *UserQuery) Filter() *UserFilter {
	return &UserFilter{config: uq.config, predicateAdder: uq}
}

// addPredicate implements the predicateAdder interface.
func (m *UserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the UserMutation builder.
func (m *UserMutation) Filter() *UserFilter {
	return &UserFilter{config: m.config, predicateAdder: m}
}

// UserFilter provides a generic filtering capability at runtime for UserQuery.
type UserFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *UserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *UserFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(user.FieldID))
}

// WhereCreateTime applies the entql int64 predicate on the create_time field.
func (f *UserFilter) WhereCreateTime(p entql.Int64P) {
	f.Where(p.Field(user.FieldCreateTime))
}

// WhereUpdateTime applies the entql int64 predicate on the update_time field.
func (f *UserFilter) WhereUpdateTime(p entql.Int64P) {
	f.Where(p.Field(user.FieldUpdateTime))
}

// WhereDeleteTime applies the entql int64 predicate on the delete_time field.
func (f *UserFilter) WhereDeleteTime(p entql.Int64P) {
	f.Where(p.Field(user.FieldDeleteTime))
}

// WhereUserName applies the entql string predicate on the user_name field.
func (f *UserFilter) WhereUserName(p entql.StringP) {
	f.Where(p.Field(user.FieldUserName))
}

// WherePassword applies the entql string predicate on the password field.
func (f *UserFilter) WherePassword(p entql.StringP) {
	f.Where(p.Field(user.FieldPassword))
}

// WhereNickName applies the entql string predicate on the nick_name field.
func (f *UserFilter) WhereNickName(p entql.StringP) {
	f.Where(p.Field(user.FieldNickName))
}

// WhereRealName applies the entql string predicate on the real_name field.
func (f *UserFilter) WhereRealName(p entql.StringP) {
	f.Where(p.Field(user.FieldRealName))
}

// WhereEmail applies the entql string predicate on the email field.
func (f *UserFilter) WhereEmail(p entql.StringP) {
	f.Where(p.Field(user.FieldEmail))
}

// WherePhone applies the entql string predicate on the phone field.
func (f *UserFilter) WherePhone(p entql.StringP) {
	f.Where(p.Field(user.FieldPhone))
}

// WhereAvatar applies the entql string predicate on the avatar field.
func (f *UserFilter) WhereAvatar(p entql.StringP) {
	f.Where(p.Field(user.FieldAvatar))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *UserFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(user.FieldDescription))
}

// WhereRoleID applies the entql uint32 predicate on the role_id field.
func (f *UserFilter) WhereRoleID(p entql.Uint32P) {
	f.Where(p.Field(user.FieldRoleID))
}

// WhereAuthority applies the entql string predicate on the authority field.
func (f *UserFilter) WhereAuthority(p entql.StringP) {
	f.Where(p.Field(user.FieldAuthority))
}