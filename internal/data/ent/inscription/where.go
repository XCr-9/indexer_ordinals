// Code generated by ent, DO NOT EDIT.

package inscription

import (
	"github.com/adshao/ordinals-indexer/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Inscription {
	return predicate.Inscription(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Inscription {
	return predicate.Inscription(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Inscription {
	return predicate.Inscription(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Inscription {
	return predicate.Inscription(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Inscription {
	return predicate.Inscription(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Inscription {
	return predicate.Inscription(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Inscription {
	return predicate.Inscription(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldUpdatedAt, v))
}

// InscriptionID applies equality check predicate on the "inscription_id" field. It's identical to InscriptionIDEQ.
func InscriptionID(v int64) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldInscriptionID, v))
}

// UID applies equality check predicate on the "uid" field. It's identical to UIDEQ.
func UID(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldUID, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldAddress, v))
}

// OutputValue applies equality check predicate on the "output_value" field. It's identical to OutputValueEQ.
func OutputValue(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldOutputValue, v))
}

// ContentLength applies equality check predicate on the "content_length" field. It's identical to ContentLengthEQ.
func ContentLength(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldContentLength, v))
}

// ContentType applies equality check predicate on the "content_type" field. It's identical to ContentTypeEQ.
func ContentType(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldContentType, v))
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldTimestamp, v))
}

// GenesisHeight applies equality check predicate on the "genesis_height" field. It's identical to GenesisHeightEQ.
func GenesisHeight(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldGenesisHeight, v))
}

// GenesisFee applies equality check predicate on the "genesis_fee" field. It's identical to GenesisFeeEQ.
func GenesisFee(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldGenesisFee, v))
}

// GenesisTx applies equality check predicate on the "genesis_tx" field. It's identical to GenesisTxEQ.
func GenesisTx(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldGenesisTx, v))
}

// Location applies equality check predicate on the "location" field. It's identical to LocationEQ.
func Location(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldLocation, v))
}

// Output applies equality check predicate on the "output" field. It's identical to OutputEQ.
func Output(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldOutput, v))
}

// Offset applies equality check predicate on the "offset" field. It's identical to OffsetEQ.
func Offset(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldOffset, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldLTE(FieldUpdatedAt, v))
}

// InscriptionIDEQ applies the EQ predicate on the "inscription_id" field.
func InscriptionIDEQ(v int64) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldInscriptionID, v))
}

// InscriptionIDNEQ applies the NEQ predicate on the "inscription_id" field.
func InscriptionIDNEQ(v int64) predicate.Inscription {
	return predicate.Inscription(sql.FieldNEQ(FieldInscriptionID, v))
}

// InscriptionIDIn applies the In predicate on the "inscription_id" field.
func InscriptionIDIn(vs ...int64) predicate.Inscription {
	return predicate.Inscription(sql.FieldIn(FieldInscriptionID, vs...))
}

// InscriptionIDNotIn applies the NotIn predicate on the "inscription_id" field.
func InscriptionIDNotIn(vs ...int64) predicate.Inscription {
	return predicate.Inscription(sql.FieldNotIn(FieldInscriptionID, vs...))
}

// InscriptionIDGT applies the GT predicate on the "inscription_id" field.
func InscriptionIDGT(v int64) predicate.Inscription {
	return predicate.Inscription(sql.FieldGT(FieldInscriptionID, v))
}

// InscriptionIDGTE applies the GTE predicate on the "inscription_id" field.
func InscriptionIDGTE(v int64) predicate.Inscription {
	return predicate.Inscription(sql.FieldGTE(FieldInscriptionID, v))
}

// InscriptionIDLT applies the LT predicate on the "inscription_id" field.
func InscriptionIDLT(v int64) predicate.Inscription {
	return predicate.Inscription(sql.FieldLT(FieldInscriptionID, v))
}

// InscriptionIDLTE applies the LTE predicate on the "inscription_id" field.
func InscriptionIDLTE(v int64) predicate.Inscription {
	return predicate.Inscription(sql.FieldLTE(FieldInscriptionID, v))
}

// UIDEQ applies the EQ predicate on the "uid" field.
func UIDEQ(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldUID, v))
}

// UIDNEQ applies the NEQ predicate on the "uid" field.
func UIDNEQ(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldNEQ(FieldUID, v))
}

// UIDIn applies the In predicate on the "uid" field.
func UIDIn(vs ...string) predicate.Inscription {
	return predicate.Inscription(sql.FieldIn(FieldUID, vs...))
}

// UIDNotIn applies the NotIn predicate on the "uid" field.
func UIDNotIn(vs ...string) predicate.Inscription {
	return predicate.Inscription(sql.FieldNotIn(FieldUID, vs...))
}

// UIDGT applies the GT predicate on the "uid" field.
func UIDGT(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldGT(FieldUID, v))
}

// UIDGTE applies the GTE predicate on the "uid" field.
func UIDGTE(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldGTE(FieldUID, v))
}

// UIDLT applies the LT predicate on the "uid" field.
func UIDLT(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldLT(FieldUID, v))
}

// UIDLTE applies the LTE predicate on the "uid" field.
func UIDLTE(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldLTE(FieldUID, v))
}

// UIDContains applies the Contains predicate on the "uid" field.
func UIDContains(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldContains(FieldUID, v))
}

// UIDHasPrefix applies the HasPrefix predicate on the "uid" field.
func UIDHasPrefix(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldHasPrefix(FieldUID, v))
}

// UIDHasSuffix applies the HasSuffix predicate on the "uid" field.
func UIDHasSuffix(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldHasSuffix(FieldUID, v))
}

// UIDEqualFold applies the EqualFold predicate on the "uid" field.
func UIDEqualFold(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldEqualFold(FieldUID, v))
}

// UIDContainsFold applies the ContainsFold predicate on the "uid" field.
func UIDContainsFold(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldContainsFold(FieldUID, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Inscription {
	return predicate.Inscription(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Inscription {
	return predicate.Inscription(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldContainsFold(FieldAddress, v))
}

// OutputValueEQ applies the EQ predicate on the "output_value" field.
func OutputValueEQ(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldOutputValue, v))
}

// OutputValueNEQ applies the NEQ predicate on the "output_value" field.
func OutputValueNEQ(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldNEQ(FieldOutputValue, v))
}

// OutputValueIn applies the In predicate on the "output_value" field.
func OutputValueIn(vs ...uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldIn(FieldOutputValue, vs...))
}

// OutputValueNotIn applies the NotIn predicate on the "output_value" field.
func OutputValueNotIn(vs ...uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldNotIn(FieldOutputValue, vs...))
}

// OutputValueGT applies the GT predicate on the "output_value" field.
func OutputValueGT(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldGT(FieldOutputValue, v))
}

// OutputValueGTE applies the GTE predicate on the "output_value" field.
func OutputValueGTE(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldGTE(FieldOutputValue, v))
}

// OutputValueLT applies the LT predicate on the "output_value" field.
func OutputValueLT(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldLT(FieldOutputValue, v))
}

// OutputValueLTE applies the LTE predicate on the "output_value" field.
func OutputValueLTE(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldLTE(FieldOutputValue, v))
}

// ContentLengthEQ applies the EQ predicate on the "content_length" field.
func ContentLengthEQ(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldContentLength, v))
}

// ContentLengthNEQ applies the NEQ predicate on the "content_length" field.
func ContentLengthNEQ(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldNEQ(FieldContentLength, v))
}

// ContentLengthIn applies the In predicate on the "content_length" field.
func ContentLengthIn(vs ...uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldIn(FieldContentLength, vs...))
}

// ContentLengthNotIn applies the NotIn predicate on the "content_length" field.
func ContentLengthNotIn(vs ...uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldNotIn(FieldContentLength, vs...))
}

// ContentLengthGT applies the GT predicate on the "content_length" field.
func ContentLengthGT(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldGT(FieldContentLength, v))
}

// ContentLengthGTE applies the GTE predicate on the "content_length" field.
func ContentLengthGTE(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldGTE(FieldContentLength, v))
}

// ContentLengthLT applies the LT predicate on the "content_length" field.
func ContentLengthLT(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldLT(FieldContentLength, v))
}

// ContentLengthLTE applies the LTE predicate on the "content_length" field.
func ContentLengthLTE(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldLTE(FieldContentLength, v))
}

// ContentTypeEQ applies the EQ predicate on the "content_type" field.
func ContentTypeEQ(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldContentType, v))
}

// ContentTypeNEQ applies the NEQ predicate on the "content_type" field.
func ContentTypeNEQ(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldNEQ(FieldContentType, v))
}

// ContentTypeIn applies the In predicate on the "content_type" field.
func ContentTypeIn(vs ...string) predicate.Inscription {
	return predicate.Inscription(sql.FieldIn(FieldContentType, vs...))
}

// ContentTypeNotIn applies the NotIn predicate on the "content_type" field.
func ContentTypeNotIn(vs ...string) predicate.Inscription {
	return predicate.Inscription(sql.FieldNotIn(FieldContentType, vs...))
}

// ContentTypeGT applies the GT predicate on the "content_type" field.
func ContentTypeGT(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldGT(FieldContentType, v))
}

// ContentTypeGTE applies the GTE predicate on the "content_type" field.
func ContentTypeGTE(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldGTE(FieldContentType, v))
}

// ContentTypeLT applies the LT predicate on the "content_type" field.
func ContentTypeLT(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldLT(FieldContentType, v))
}

// ContentTypeLTE applies the LTE predicate on the "content_type" field.
func ContentTypeLTE(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldLTE(FieldContentType, v))
}

// ContentTypeContains applies the Contains predicate on the "content_type" field.
func ContentTypeContains(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldContains(FieldContentType, v))
}

// ContentTypeHasPrefix applies the HasPrefix predicate on the "content_type" field.
func ContentTypeHasPrefix(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldHasPrefix(FieldContentType, v))
}

// ContentTypeHasSuffix applies the HasSuffix predicate on the "content_type" field.
func ContentTypeHasSuffix(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldHasSuffix(FieldContentType, v))
}

// ContentTypeEqualFold applies the EqualFold predicate on the "content_type" field.
func ContentTypeEqualFold(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldEqualFold(FieldContentType, v))
}

// ContentTypeContainsFold applies the ContainsFold predicate on the "content_type" field.
func ContentTypeContainsFold(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldContainsFold(FieldContentType, v))
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldTimestamp, v))
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldNEQ(FieldTimestamp, v))
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldIn(FieldTimestamp, vs...))
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldNotIn(FieldTimestamp, vs...))
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldGT(FieldTimestamp, v))
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldGTE(FieldTimestamp, v))
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldLT(FieldTimestamp, v))
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.Inscription {
	return predicate.Inscription(sql.FieldLTE(FieldTimestamp, v))
}

// GenesisHeightEQ applies the EQ predicate on the "genesis_height" field.
func GenesisHeightEQ(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldGenesisHeight, v))
}

// GenesisHeightNEQ applies the NEQ predicate on the "genesis_height" field.
func GenesisHeightNEQ(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldNEQ(FieldGenesisHeight, v))
}

// GenesisHeightIn applies the In predicate on the "genesis_height" field.
func GenesisHeightIn(vs ...uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldIn(FieldGenesisHeight, vs...))
}

// GenesisHeightNotIn applies the NotIn predicate on the "genesis_height" field.
func GenesisHeightNotIn(vs ...uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldNotIn(FieldGenesisHeight, vs...))
}

// GenesisHeightGT applies the GT predicate on the "genesis_height" field.
func GenesisHeightGT(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldGT(FieldGenesisHeight, v))
}

// GenesisHeightGTE applies the GTE predicate on the "genesis_height" field.
func GenesisHeightGTE(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldGTE(FieldGenesisHeight, v))
}

// GenesisHeightLT applies the LT predicate on the "genesis_height" field.
func GenesisHeightLT(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldLT(FieldGenesisHeight, v))
}

// GenesisHeightLTE applies the LTE predicate on the "genesis_height" field.
func GenesisHeightLTE(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldLTE(FieldGenesisHeight, v))
}

// GenesisFeeEQ applies the EQ predicate on the "genesis_fee" field.
func GenesisFeeEQ(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldGenesisFee, v))
}

// GenesisFeeNEQ applies the NEQ predicate on the "genesis_fee" field.
func GenesisFeeNEQ(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldNEQ(FieldGenesisFee, v))
}

// GenesisFeeIn applies the In predicate on the "genesis_fee" field.
func GenesisFeeIn(vs ...uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldIn(FieldGenesisFee, vs...))
}

// GenesisFeeNotIn applies the NotIn predicate on the "genesis_fee" field.
func GenesisFeeNotIn(vs ...uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldNotIn(FieldGenesisFee, vs...))
}

// GenesisFeeGT applies the GT predicate on the "genesis_fee" field.
func GenesisFeeGT(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldGT(FieldGenesisFee, v))
}

// GenesisFeeGTE applies the GTE predicate on the "genesis_fee" field.
func GenesisFeeGTE(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldGTE(FieldGenesisFee, v))
}

// GenesisFeeLT applies the LT predicate on the "genesis_fee" field.
func GenesisFeeLT(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldLT(FieldGenesisFee, v))
}

// GenesisFeeLTE applies the LTE predicate on the "genesis_fee" field.
func GenesisFeeLTE(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldLTE(FieldGenesisFee, v))
}

// GenesisTxEQ applies the EQ predicate on the "genesis_tx" field.
func GenesisTxEQ(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldGenesisTx, v))
}

// GenesisTxNEQ applies the NEQ predicate on the "genesis_tx" field.
func GenesisTxNEQ(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldNEQ(FieldGenesisTx, v))
}

// GenesisTxIn applies the In predicate on the "genesis_tx" field.
func GenesisTxIn(vs ...string) predicate.Inscription {
	return predicate.Inscription(sql.FieldIn(FieldGenesisTx, vs...))
}

// GenesisTxNotIn applies the NotIn predicate on the "genesis_tx" field.
func GenesisTxNotIn(vs ...string) predicate.Inscription {
	return predicate.Inscription(sql.FieldNotIn(FieldGenesisTx, vs...))
}

// GenesisTxGT applies the GT predicate on the "genesis_tx" field.
func GenesisTxGT(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldGT(FieldGenesisTx, v))
}

// GenesisTxGTE applies the GTE predicate on the "genesis_tx" field.
func GenesisTxGTE(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldGTE(FieldGenesisTx, v))
}

// GenesisTxLT applies the LT predicate on the "genesis_tx" field.
func GenesisTxLT(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldLT(FieldGenesisTx, v))
}

// GenesisTxLTE applies the LTE predicate on the "genesis_tx" field.
func GenesisTxLTE(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldLTE(FieldGenesisTx, v))
}

// GenesisTxContains applies the Contains predicate on the "genesis_tx" field.
func GenesisTxContains(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldContains(FieldGenesisTx, v))
}

// GenesisTxHasPrefix applies the HasPrefix predicate on the "genesis_tx" field.
func GenesisTxHasPrefix(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldHasPrefix(FieldGenesisTx, v))
}

// GenesisTxHasSuffix applies the HasSuffix predicate on the "genesis_tx" field.
func GenesisTxHasSuffix(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldHasSuffix(FieldGenesisTx, v))
}

// GenesisTxEqualFold applies the EqualFold predicate on the "genesis_tx" field.
func GenesisTxEqualFold(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldEqualFold(FieldGenesisTx, v))
}

// GenesisTxContainsFold applies the ContainsFold predicate on the "genesis_tx" field.
func GenesisTxContainsFold(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldContainsFold(FieldGenesisTx, v))
}

// LocationEQ applies the EQ predicate on the "location" field.
func LocationEQ(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldLocation, v))
}

// LocationNEQ applies the NEQ predicate on the "location" field.
func LocationNEQ(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldNEQ(FieldLocation, v))
}

// LocationIn applies the In predicate on the "location" field.
func LocationIn(vs ...string) predicate.Inscription {
	return predicate.Inscription(sql.FieldIn(FieldLocation, vs...))
}

// LocationNotIn applies the NotIn predicate on the "location" field.
func LocationNotIn(vs ...string) predicate.Inscription {
	return predicate.Inscription(sql.FieldNotIn(FieldLocation, vs...))
}

// LocationGT applies the GT predicate on the "location" field.
func LocationGT(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldGT(FieldLocation, v))
}

// LocationGTE applies the GTE predicate on the "location" field.
func LocationGTE(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldGTE(FieldLocation, v))
}

// LocationLT applies the LT predicate on the "location" field.
func LocationLT(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldLT(FieldLocation, v))
}

// LocationLTE applies the LTE predicate on the "location" field.
func LocationLTE(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldLTE(FieldLocation, v))
}

// LocationContains applies the Contains predicate on the "location" field.
func LocationContains(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldContains(FieldLocation, v))
}

// LocationHasPrefix applies the HasPrefix predicate on the "location" field.
func LocationHasPrefix(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldHasPrefix(FieldLocation, v))
}

// LocationHasSuffix applies the HasSuffix predicate on the "location" field.
func LocationHasSuffix(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldHasSuffix(FieldLocation, v))
}

// LocationEqualFold applies the EqualFold predicate on the "location" field.
func LocationEqualFold(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldEqualFold(FieldLocation, v))
}

// LocationContainsFold applies the ContainsFold predicate on the "location" field.
func LocationContainsFold(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldContainsFold(FieldLocation, v))
}

// OutputEQ applies the EQ predicate on the "output" field.
func OutputEQ(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldOutput, v))
}

// OutputNEQ applies the NEQ predicate on the "output" field.
func OutputNEQ(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldNEQ(FieldOutput, v))
}

// OutputIn applies the In predicate on the "output" field.
func OutputIn(vs ...string) predicate.Inscription {
	return predicate.Inscription(sql.FieldIn(FieldOutput, vs...))
}

// OutputNotIn applies the NotIn predicate on the "output" field.
func OutputNotIn(vs ...string) predicate.Inscription {
	return predicate.Inscription(sql.FieldNotIn(FieldOutput, vs...))
}

// OutputGT applies the GT predicate on the "output" field.
func OutputGT(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldGT(FieldOutput, v))
}

// OutputGTE applies the GTE predicate on the "output" field.
func OutputGTE(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldGTE(FieldOutput, v))
}

// OutputLT applies the LT predicate on the "output" field.
func OutputLT(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldLT(FieldOutput, v))
}

// OutputLTE applies the LTE predicate on the "output" field.
func OutputLTE(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldLTE(FieldOutput, v))
}

// OutputContains applies the Contains predicate on the "output" field.
func OutputContains(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldContains(FieldOutput, v))
}

// OutputHasPrefix applies the HasPrefix predicate on the "output" field.
func OutputHasPrefix(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldHasPrefix(FieldOutput, v))
}

// OutputHasSuffix applies the HasSuffix predicate on the "output" field.
func OutputHasSuffix(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldHasSuffix(FieldOutput, v))
}

// OutputEqualFold applies the EqualFold predicate on the "output" field.
func OutputEqualFold(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldEqualFold(FieldOutput, v))
}

// OutputContainsFold applies the ContainsFold predicate on the "output" field.
func OutputContainsFold(v string) predicate.Inscription {
	return predicate.Inscription(sql.FieldContainsFold(FieldOutput, v))
}

// OffsetEQ applies the EQ predicate on the "offset" field.
func OffsetEQ(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldEQ(FieldOffset, v))
}

// OffsetNEQ applies the NEQ predicate on the "offset" field.
func OffsetNEQ(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldNEQ(FieldOffset, v))
}

// OffsetIn applies the In predicate on the "offset" field.
func OffsetIn(vs ...uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldIn(FieldOffset, vs...))
}

// OffsetNotIn applies the NotIn predicate on the "offset" field.
func OffsetNotIn(vs ...uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldNotIn(FieldOffset, vs...))
}

// OffsetGT applies the GT predicate on the "offset" field.
func OffsetGT(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldGT(FieldOffset, v))
}

// OffsetGTE applies the GTE predicate on the "offset" field.
func OffsetGTE(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldGTE(FieldOffset, v))
}

// OffsetLT applies the LT predicate on the "offset" field.
func OffsetLT(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldLT(FieldOffset, v))
}

// OffsetLTE applies the LTE predicate on the "offset" field.
func OffsetLTE(v uint64) predicate.Inscription {
	return predicate.Inscription(sql.FieldLTE(FieldOffset, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Inscription) predicate.Inscription {
	return predicate.Inscription(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Inscription) predicate.Inscription {
	return predicate.Inscription(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Inscription) predicate.Inscription {
	return predicate.Inscription(func(s *sql.Selector) {
		p(s.Not())
	})
}
