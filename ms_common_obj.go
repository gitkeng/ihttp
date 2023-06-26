package ihttp

import (
	"github.com/gitkeng/ihttp/util/convutil"
	"github.com/gitkeng/ihttp/util/stringutil"
	"strings"
)

type Response struct {
	RequestId  string         `json:"request_id,omitempty"`
	StatusCode int            `json:"status_code"`
	Code       string         `json:"code,omitempty"`
	Message    string         `json:"message,omitempty"`
	Data       map[string]any `json:"data,omitempty"`
	Error      Errors         `json:"errors,omitempty"`
}

func (resp *Response) String() string {
	return stringutil.Json(*resp)
}

func (resp *Response) ToMap() map[string]any {
	return convutil.Obj2Map(*resp)
}

type Errors []Error

func (errs *Errors) String() string {
	return stringutil.Json(*errs)
}

func (errs *Errors) ToMap() map[string]any {
	errors := make(map[string]any)
	errors["errors"] = errs
	return convutil.Obj2Map(errors)
}

func (errs *Errors) Error() string {
	return stringutil.Json(*errs)
}

type Error struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Fields  map[string]any `json:"fields,omitempty"`
}

func (err *Error) String() string {
	return stringutil.Json(*err)
}

func (err *Error) Error() string {
	return stringutil.Json(*err)
}

func (err *Error) ToMap() map[string]any {
	return convutil.Obj2Map(*err)
}

type Field struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

func NewError(code string, message string, fields ...Field) Error {
	var errorFields map[string]any
	if len(fields) > 0 {
		errorFields = make(map[string]any)
		for _, field := range fields {
			errorFields[field.Key] = field.Value
		}
	}
	return Error{
		Code:    code,
		Message: message,
		Fields:  errorFields,
	}
}

// DB Filter and Order

const (
	Ascending  Direction = "asc"
	Descending Direction = "desc"
)

var (
	QueryOrderDirectionMap = map[string]Direction{
		"asc":  Ascending,
		"desc": Descending,
	}
)

type (
	Direction    string
	IQueryFilter interface {
		GetField() string
		GetValue() any
		GetFromValue() any
		GetToValue() any
	}

	QueryFilter struct {
		Field     string `json:"field,omitempty"`
		Value     any    `json:"value,omitempty"`
		FromValue any    `json:"from_value,omitempty"`
		ToValue   any    `json:"to_value,omitempty"`
	}

	IQueryOption interface {
		GetLimit() int
		GetOffset() int
		GetSort() []IQueryOrder
	}

	QueryOption struct {
		Limit  int          `json:"limit,omitempty"`
		Offset int          `json:"offset,omitempty"`
		Sort   []QueryOrder `json:"sorts,omitempty"`
	}

	IQueryOrder interface {
		GetField() string
		GetOrder() Direction
	}

	QueryOrder struct {
		Field string    `json:"field,omitempty"`
		Order Direction `json:"order,omitempty"`
	}
)

func (filter *QueryFilter) GetField() string {
	return strings.TrimSpace(filter.Field)
}

func (filter *QueryFilter) GetValue() any {
	if value, ok := filter.Value.(string); ok {
		return strings.TrimSpace(value)
	}
	return filter.Value
}

func (filter *QueryFilter) GetFromValue() any {
	if value, ok := filter.FromValue.(string); ok {
		return strings.TrimSpace(value)
	}
	return filter.FromValue
}

func (filter *QueryFilter) GetToValue() any {
	if value, ok := filter.ToValue.(string); ok {
		return strings.TrimSpace(value)
	}
	return filter.ToValue
}

func (option *QueryOption) GetLimit() int {
	return option.Limit
}

func (option *QueryOption) GetOffset() int {
	return option.Offset
}

func (option *QueryOption) GetSort() []IQueryOrder {
	iQueryOrders := make([]IQueryOrder, 0)
	for idx, _ := range option.Sort {
		iQueryOrders = append(iQueryOrders, &option.Sort[idx])
	}
	return iQueryOrders
}

func (order *QueryOrder) GetField() string {
	return strings.TrimSpace(order.Field)
}

func (order *QueryOrder) GetOrder() Direction {
	return order.Order
}

type FilterRequest struct {
	Filters []QueryFilter `json:"filters,omitempty"`
	Option  QueryOption   `json:"option,omitempty"`
}

func (req *FilterRequest) Validate() error {
	return nil
}

func (req *FilterRequest) String() string {
	return stringutil.Json(*req)
}

func (req *FilterRequest) ToMap() map[string]any {
	return convutil.Obj2Map(*req)
}

func (req *FilterRequest) GetFilters() []IQueryFilter {
	var filters []IQueryFilter
	for idx, _ := range req.Filters {
		filters = append(filters, &req.Filters[idx])
	}
	return filters
}

func (req *FilterRequest) GetOption() IQueryOption {
	return &req.Option
}
