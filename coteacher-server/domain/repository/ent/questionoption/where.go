// Code generated by ent, DO NOT EDIT.

package questionoption

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldLTE(FieldID, id))
}

// QuestionID applies equality check predicate on the "question_id" field. It's identical to QuestionIDEQ.
func QuestionID(v uuid.UUID) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldQuestionID, v))
}

// OptionText applies equality check predicate on the "option_text" field. It's identical to OptionTextEQ.
func OptionText(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldOptionText, v))
}

// Order applies equality check predicate on the "order" field. It's identical to OrderEQ.
func Order(v int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldOrder, v))
}

// QuestionIDEQ applies the EQ predicate on the "question_id" field.
func QuestionIDEQ(v uuid.UUID) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldQuestionID, v))
}

// QuestionIDNEQ applies the NEQ predicate on the "question_id" field.
func QuestionIDNEQ(v uuid.UUID) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNEQ(FieldQuestionID, v))
}

// QuestionIDIn applies the In predicate on the "question_id" field.
func QuestionIDIn(vs ...uuid.UUID) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldIn(FieldQuestionID, vs...))
}

// QuestionIDNotIn applies the NotIn predicate on the "question_id" field.
func QuestionIDNotIn(vs ...uuid.UUID) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNotIn(FieldQuestionID, vs...))
}

// OptionTextEQ applies the EQ predicate on the "option_text" field.
func OptionTextEQ(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldOptionText, v))
}

// OptionTextNEQ applies the NEQ predicate on the "option_text" field.
func OptionTextNEQ(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNEQ(FieldOptionText, v))
}

// OptionTextIn applies the In predicate on the "option_text" field.
func OptionTextIn(vs ...string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldIn(FieldOptionText, vs...))
}

// OptionTextNotIn applies the NotIn predicate on the "option_text" field.
func OptionTextNotIn(vs ...string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNotIn(FieldOptionText, vs...))
}

// OptionTextGT applies the GT predicate on the "option_text" field.
func OptionTextGT(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldGT(FieldOptionText, v))
}

// OptionTextGTE applies the GTE predicate on the "option_text" field.
func OptionTextGTE(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldGTE(FieldOptionText, v))
}

// OptionTextLT applies the LT predicate on the "option_text" field.
func OptionTextLT(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldLT(FieldOptionText, v))
}

// OptionTextLTE applies the LTE predicate on the "option_text" field.
func OptionTextLTE(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldLTE(FieldOptionText, v))
}

// OptionTextContains applies the Contains predicate on the "option_text" field.
func OptionTextContains(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldContains(FieldOptionText, v))
}

// OptionTextHasPrefix applies the HasPrefix predicate on the "option_text" field.
func OptionTextHasPrefix(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldHasPrefix(FieldOptionText, v))
}

// OptionTextHasSuffix applies the HasSuffix predicate on the "option_text" field.
func OptionTextHasSuffix(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldHasSuffix(FieldOptionText, v))
}

// OptionTextEqualFold applies the EqualFold predicate on the "option_text" field.
func OptionTextEqualFold(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEqualFold(FieldOptionText, v))
}

// OptionTextContainsFold applies the ContainsFold predicate on the "option_text" field.
func OptionTextContainsFold(v string) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldContainsFold(FieldOptionText, v))
}

// OrderEQ applies the EQ predicate on the "order" field.
func OrderEQ(v int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldEQ(FieldOrder, v))
}

// OrderNEQ applies the NEQ predicate on the "order" field.
func OrderNEQ(v int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNEQ(FieldOrder, v))
}

// OrderIn applies the In predicate on the "order" field.
func OrderIn(vs ...int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldIn(FieldOrder, vs...))
}

// OrderNotIn applies the NotIn predicate on the "order" field.
func OrderNotIn(vs ...int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldNotIn(FieldOrder, vs...))
}

// OrderGT applies the GT predicate on the "order" field.
func OrderGT(v int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldGT(FieldOrder, v))
}

// OrderGTE applies the GTE predicate on the "order" field.
func OrderGTE(v int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldGTE(FieldOrder, v))
}

// OrderLT applies the LT predicate on the "order" field.
func OrderLT(v int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldLT(FieldOrder, v))
}

// OrderLTE applies the LTE predicate on the "order" field.
func OrderLTE(v int) predicate.QuestionOption {
	return predicate.QuestionOption(sql.FieldLTE(FieldOrder, v))
}

// HasQuestion applies the HasEdge predicate on the "question" edge.
func HasQuestion() predicate.QuestionOption {
	return predicate.QuestionOption(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, QuestionTable, QuestionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasQuestionWith applies the HasEdge predicate on the "question" edge with a given conditions (other predicates).
func HasQuestionWith(preds ...predicate.Question) predicate.QuestionOption {
	return predicate.QuestionOption(func(s *sql.Selector) {
		step := newQuestionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSelectedOption applies the HasEdge predicate on the "selectedOption" edge.
func HasSelectedOption() predicate.QuestionOption {
	return predicate.QuestionOption(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SelectedOptionTable, SelectedOptionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSelectedOptionWith applies the HasEdge predicate on the "selectedOption" edge with a given conditions (other predicates).
func HasSelectedOptionWith(preds ...predicate.SelectedOption) predicate.QuestionOption {
	return predicate.QuestionOption(func(s *sql.Selector) {
		step := newSelectedOptionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.QuestionOption) predicate.QuestionOption {
	return predicate.QuestionOption(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.QuestionOption) predicate.QuestionOption {
	return predicate.QuestionOption(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.QuestionOption) predicate.QuestionOption {
	return predicate.QuestionOption(sql.NotPredicates(p))
}
