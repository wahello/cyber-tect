// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/vicanso/cybertect/ent/schema"
	"github.com/vicanso/cybertect/ent/tcpdetectorresult"
)

// TCPDetectorResult is the model entity for the TCPDetectorResult schema.
type TCPDetectorResult struct {
	config `json:"-" sql:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt,omitempty" sql:"created_at"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty" sql:"updated_at"`
	// Task holds the value of the "task" field.
	Task int `json:"task,omitempty"`
	// Result holds the value of the "result" field.
	Result int8 `json:"result,omitempty"`
	// MaxDuration holds the value of the "maxDuration" field.
	MaxDuration int `json:"maxDuration,omitempty" sql:"max_duration"`
	// Messages holds the value of the "messages" field.
	Messages []string `json:"messages,omitempty"`
	// Addrs holds the value of the "addrs" field.
	Addrs string `json:"addrs,omitempty"`
	// Results holds the value of the "results" field.
	Results schema.TCPDetectorSubResults `json:"results,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TCPDetectorResult) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case tcpdetectorresult.FieldMessages, tcpdetectorresult.FieldResults:
			values[i] = &[]byte{}
		case tcpdetectorresult.FieldID, tcpdetectorresult.FieldTask, tcpdetectorresult.FieldResult, tcpdetectorresult.FieldMaxDuration:
			values[i] = &sql.NullInt64{}
		case tcpdetectorresult.FieldAddrs:
			values[i] = &sql.NullString{}
		case tcpdetectorresult.FieldCreatedAt, tcpdetectorresult.FieldUpdatedAt:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type TCPDetectorResult", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TCPDetectorResult fields.
func (tdr *TCPDetectorResult) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tcpdetectorresult.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tdr.ID = int(value.Int64)
		case tcpdetectorresult.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tdr.CreatedAt = value.Time
			}
		case tcpdetectorresult.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tdr.UpdatedAt = value.Time
			}
		case tcpdetectorresult.FieldTask:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field task", values[i])
			} else if value.Valid {
				tdr.Task = int(value.Int64)
			}
		case tcpdetectorresult.FieldResult:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value.Valid {
				tdr.Result = int8(value.Int64)
			}
		case tcpdetectorresult.FieldMaxDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field maxDuration", values[i])
			} else if value.Valid {
				tdr.MaxDuration = int(value.Int64)
			}
		case tcpdetectorresult.FieldMessages:

			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field messages", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &tdr.Messages); err != nil {
					return fmt.Errorf("unmarshal field messages: %v", err)
				}
			}
		case tcpdetectorresult.FieldAddrs:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field addrs", values[i])
			} else if value.Valid {
				tdr.Addrs = value.String
			}
		case tcpdetectorresult.FieldResults:

			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field results", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &tdr.Results); err != nil {
					return fmt.Errorf("unmarshal field results: %v", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TCPDetectorResult.
// Note that you need to call TCPDetectorResult.Unwrap() before calling this method if this TCPDetectorResult
// was returned from a transaction, and the transaction was committed or rolled back.
func (tdr *TCPDetectorResult) Update() *TCPDetectorResultUpdateOne {
	return (&TCPDetectorResultClient{config: tdr.config}).UpdateOne(tdr)
}

// Unwrap unwraps the TCPDetectorResult entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tdr *TCPDetectorResult) Unwrap() *TCPDetectorResult {
	tx, ok := tdr.config.driver.(*txDriver)
	if !ok {
		panic("ent: TCPDetectorResult is not a transactional entity")
	}
	tdr.config.driver = tx.drv
	return tdr
}

// String implements the fmt.Stringer.
func (tdr *TCPDetectorResult) String() string {
	var builder strings.Builder
	builder.WriteString("TCPDetectorResult(")
	builder.WriteString(fmt.Sprintf("id=%v", tdr.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(tdr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(tdr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", task=")
	builder.WriteString(fmt.Sprintf("%v", tdr.Task))
	builder.WriteString(", result=")
	builder.WriteString(fmt.Sprintf("%v", tdr.Result))
	builder.WriteString(", maxDuration=")
	builder.WriteString(fmt.Sprintf("%v", tdr.MaxDuration))
	builder.WriteString(", messages=")
	builder.WriteString(fmt.Sprintf("%v", tdr.Messages))
	builder.WriteString(", addrs=")
	builder.WriteString(tdr.Addrs)
	builder.WriteString(", results=")
	builder.WriteString(fmt.Sprintf("%v", tdr.Results))
	builder.WriteByte(')')
	return builder.String()
}

// TCPDetectorResults is a parsable slice of TCPDetectorResult.
type TCPDetectorResults []*TCPDetectorResult

func (tdr TCPDetectorResults) config(cfg config) {
	for _i := range tdr {
		tdr[_i].config = cfg
	}
}
