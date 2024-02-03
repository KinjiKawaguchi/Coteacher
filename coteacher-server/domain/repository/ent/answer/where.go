// Code generated by ent, DO NOT EDIT.

package answer

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldLTE(FieldID, id))
}

// QuestionID applies equality check predicate on the "question_id" field. It's identical to QuestionIDEQ.
func QuestionID(v uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldQuestionID, v))
}

// ResponseID applies equality check predicate on the "response_id" field. It's identical to ResponseIDEQ.
func ResponseID(v uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldResponseID, v))
}

// Order applies equality check predicate on the "order" field. It's identical to OrderEQ.
func Order(v int) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldOrder, v))
}

// AnswerText applies equality check predicate on the "answer_text" field. It's identical to AnswerTextEQ.
func AnswerText(v string) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldAnswerText, v))
}

// QuestionIDEQ applies the EQ predicate on the "question_id" field.
func QuestionIDEQ(v uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldQuestionID, v))
}

// QuestionIDNEQ applies the NEQ predicate on the "question_id" field.
func QuestionIDNEQ(v uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldNEQ(FieldQuestionID, v))
}

// QuestionIDIn applies the In predicate on the "question_id" field.
func QuestionIDIn(vs ...uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldIn(FieldQuestionID, vs...))
}

// QuestionIDNotIn applies the NotIn predicate on the "question_id" field.
func QuestionIDNotIn(vs ...uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldNotIn(FieldQuestionID, vs...))
}

// ResponseIDEQ applies the EQ predicate on the "response_id" field.
func ResponseIDEQ(v uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldResponseID, v))
}

// ResponseIDNEQ applies the NEQ predicate on the "response_id" field.
func ResponseIDNEQ(v uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldNEQ(FieldResponseID, v))
}

// ResponseIDIn applies the In predicate on the "response_id" field.
func ResponseIDIn(vs ...uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldIn(FieldResponseID, vs...))
}

// ResponseIDNotIn applies the NotIn predicate on the "response_id" field.
func ResponseIDNotIn(vs ...uuid.UUID) predicate.Answer {
	return predicate.Answer(sql.FieldNotIn(FieldResponseID, vs...))
}

// OrderEQ applies the EQ predicate on the "order" field.
func OrderEQ(v int) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldOrder, v))
}

// OrderNEQ applies the NEQ predicate on the "order" field.
func OrderNEQ(v int) predicate.Answer {
	return predicate.Answer(sql.FieldNEQ(FieldOrder, v))
}

// OrderIn applies the In predicate on the "order" field.
func OrderIn(vs ...int) predicate.Answer {
	return predicate.Answer(sql.FieldIn(FieldOrder, vs...))
}

// OrderNotIn applies the NotIn predicate on the "order" field.
func OrderNotIn(vs ...int) predicate.Answer {
	return predicate.Answer(sql.FieldNotIn(FieldOrder, vs...))
}

// OrderGT applies the GT predicate on the "order" field.
func OrderGT(v int) predicate.Answer {
	return predicate.Answer(sql.FieldGT(FieldOrder, v))
}

// OrderGTE applies the GTE predicate on the "order" field.
func OrderGTE(v int) predicate.Answer {
	return predicate.Answer(sql.FieldGTE(FieldOrder, v))
}

// OrderLT applies the LT predicate on the "order" field.
func OrderLT(v int) predicate.Answer {
	return predicate.Answer(sql.FieldLT(FieldOrder, v))
}

// OrderLTE applies the LTE predicate on the "order" field.
func OrderLTE(v int) predicate.Answer {
	return predicate.Answer(sql.FieldLTE(FieldOrder, v))
}

// AnswerTextEQ applies the EQ predicate on the "answer_text" field.
func AnswerTextEQ(v string) predicate.Answer {
	return predicate.Answer(sql.FieldEQ(FieldAnswerText, v))
}

// AnswerTextNEQ applies the NEQ predicate on the "answer_text" field.
func AnswerTextNEQ(v string) predicate.Answer {
	return predicate.Answer(sql.FieldNEQ(FieldAnswerText, v))
}

// AnswerTextIn applies the In predicate on the "answer_text" field.
func AnswerTextIn(vs ...string) predicate.Answer {
	return predicate.Answer(sql.FieldIn(FieldAnswerText, vs...))
}

// AnswerTextNotIn applies the NotIn predicate on the "answer_text" field.
func AnswerTextNotIn(vs ...string) predicate.Answer {
	return predicate.Answer(sql.FieldNotIn(FieldAnswerText, vs...))
}

// AnswerTextGT applies the GT predicate on the "answer_text" field.
func AnswerTextGT(v string) predicate.Answer {
	return predicate.Answer(sql.FieldGT(FieldAnswerText, v))
}

// AnswerTextGTE applies the GTE predicate on the "answer_text" field.
func AnswerTextGTE(v string) predicate.Answer {
	return predicate.Answer(sql.FieldGTE(FieldAnswerText, v))
}

// AnswerTextLT applies the LT predicate on the "answer_text" field.
func AnswerTextLT(v string) predicate.Answer {
	return predicate.Answer(sql.FieldLT(FieldAnswerText, v))
}

// AnswerTextLTE applies the LTE predicate on the "answer_text" field.
func AnswerTextLTE(v string) predicate.Answer {
	return predicate.Answer(sql.FieldLTE(FieldAnswerText, v))
}

// AnswerTextContains applies the Contains predicate on the "answer_text" field.
func AnswerTextContains(v string) predicate.Answer {
	return predicate.Answer(sql.FieldContains(FieldAnswerText, v))
}

// AnswerTextHasPrefix applies the HasPrefix predicate on the "answer_text" field.
func AnswerTextHasPrefix(v string) predicate.Answer {
	return predicate.Answer(sql.FieldHasPrefix(FieldAnswerText, v))
}

// AnswerTextHasSuffix applies the HasSuffix predicate on the "answer_text" field.
func AnswerTextHasSuffix(v string) predicate.Answer {
	return predicate.Answer(sql.FieldHasSuffix(FieldAnswerText, v))
}

// AnswerTextIsNil applies the IsNil predicate on the "answer_text" field.
func AnswerTextIsNil() predicate.Answer {
	return predicate.Answer(sql.FieldIsNull(FieldAnswerText))
}

// AnswerTextNotNil applies the NotNil predicate on the "answer_text" field.
func AnswerTextNotNil() predicate.Answer {
	return predicate.Answer(sql.FieldNotNull(FieldAnswerText))
}

// AnswerTextEqualFold applies the EqualFold predicate on the "answer_text" field.
func AnswerTextEqualFold(v string) predicate.Answer {
	return predicate.Answer(sql.FieldEqualFold(FieldAnswerText, v))
}

// AnswerTextContainsFold applies the ContainsFold predicate on the "answer_text" field.
func AnswerTextContainsFold(v string) predicate.Answer {
	return predicate.Answer(sql.FieldContainsFold(FieldAnswerText, v))
}

// HasQuestion applies the HasEdge predicate on the "question" edge.
func HasQuestion() predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, QuestionTable, QuestionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasQuestionWith applies the HasEdge predicate on the "question" edge with a given conditions (other predicates).
func HasQuestionWith(preds ...predicate.Question) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := newQuestionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasResponse applies the HasEdge predicate on the "response" edge.
func HasResponse() predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ResponseTable, ResponseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResponseWith applies the HasEdge predicate on the "response" edge with a given conditions (other predicates).
func HasResponseWith(preds ...predicate.Response) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := newResponseStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSelectedOption applies the HasEdge predicate on the "selectedOption" edge.
func HasSelectedOption() predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SelectedOptionTable, SelectedOptionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSelectedOptionWith applies the HasEdge predicate on the "selectedOption" edge with a given conditions (other predicates).
func HasSelectedOptionWith(preds ...predicate.SelectedOption) predicate.Answer {
	return predicate.Answer(func(s *sql.Selector) {
		step := newSelectedOptionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Answer) predicate.Answer {
	return predicate.Answer(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Answer) predicate.Answer {
	return predicate.Answer(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Answer) predicate.Answer {
	return predicate.Answer(sql.NotPredicates(p))
}
