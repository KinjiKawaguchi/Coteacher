// Code generated by ent, DO NOT EDIT.

package textquestion

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldLTE(FieldID, id))
}

// QuestionID applies equality check predicate on the "question_id" field. It's identical to QuestionIDEQ.
func QuestionID(v uuid.UUID) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldEQ(FieldQuestionID, v))
}

// MaxLength applies equality check predicate on the "max_length" field. It's identical to MaxLengthEQ.
func MaxLength(v int) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldEQ(FieldMaxLength, v))
}

// QuestionIDEQ applies the EQ predicate on the "question_id" field.
func QuestionIDEQ(v uuid.UUID) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldEQ(FieldQuestionID, v))
}

// QuestionIDNEQ applies the NEQ predicate on the "question_id" field.
func QuestionIDNEQ(v uuid.UUID) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldNEQ(FieldQuestionID, v))
}

// QuestionIDIn applies the In predicate on the "question_id" field.
func QuestionIDIn(vs ...uuid.UUID) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldIn(FieldQuestionID, vs...))
}

// QuestionIDNotIn applies the NotIn predicate on the "question_id" field.
func QuestionIDNotIn(vs ...uuid.UUID) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldNotIn(FieldQuestionID, vs...))
}

// MaxLengthEQ applies the EQ predicate on the "max_length" field.
func MaxLengthEQ(v int) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldEQ(FieldMaxLength, v))
}

// MaxLengthNEQ applies the NEQ predicate on the "max_length" field.
func MaxLengthNEQ(v int) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldNEQ(FieldMaxLength, v))
}

// MaxLengthIn applies the In predicate on the "max_length" field.
func MaxLengthIn(vs ...int) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldIn(FieldMaxLength, vs...))
}

// MaxLengthNotIn applies the NotIn predicate on the "max_length" field.
func MaxLengthNotIn(vs ...int) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldNotIn(FieldMaxLength, vs...))
}

// MaxLengthGT applies the GT predicate on the "max_length" field.
func MaxLengthGT(v int) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldGT(FieldMaxLength, v))
}

// MaxLengthGTE applies the GTE predicate on the "max_length" field.
func MaxLengthGTE(v int) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldGTE(FieldMaxLength, v))
}

// MaxLengthLT applies the LT predicate on the "max_length" field.
func MaxLengthLT(v int) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldLT(FieldMaxLength, v))
}

// MaxLengthLTE applies the LTE predicate on the "max_length" field.
func MaxLengthLTE(v int) predicate.TextQuestion {
	return predicate.TextQuestion(sql.FieldLTE(FieldMaxLength, v))
}

// HasQuestion applies the HasEdge predicate on the "question" edge.
func HasQuestion() predicate.TextQuestion {
	return predicate.TextQuestion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, QuestionTable, QuestionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasQuestionWith applies the HasEdge predicate on the "question" edge with a given conditions (other predicates).
func HasQuestionWith(preds ...predicate.Question) predicate.TextQuestion {
	return predicate.TextQuestion(func(s *sql.Selector) {
		step := newQuestionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TextQuestion) predicate.TextQuestion {
	return predicate.TextQuestion(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TextQuestion) predicate.TextQuestion {
	return predicate.TextQuestion(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TextQuestion) predicate.TextQuestion {
	return predicate.TextQuestion(sql.NotPredicates(p))
}