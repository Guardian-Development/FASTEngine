package errors

const S2 = "[ERR S2] operator is specified for a field of a type to which the operator is not applicable"
const S3 = "[ERR S3] initial value specified by the value attribute in the concrete syntax cannot be converted to a value of the type of the field"
const S4 = "[ERR S4] no initial value is specified for a constant operator"
const S5 = "[ERR S5] no initial value is specified for a default operator on a mandatory field"

const D5 = "[ERR D5] mandatory field is not present in the stream, has an undefined previous value and there is no initial value in the instruction context"
const D6 = "[ERR D6] mandatory field is not present in the stream and has an empty previous value"
const D7 = "[ERR D7] subtraction length exceeds the length of the base value or if it does not fall in the value rang of an int32"
const D9 = "[ERR D9] decoder cannot find a template associated with a template identifier appearing in the stream"

const R1 = "[ERR R1] decimal must be represented by an exponent in the range [-63 ... 63] and the mantissa must fit in an int64"
const R4 = "[ERR R4] value of an integer type cannot be represented in the target integer type in a conversion"
const R6 = "[ERR R6] read integer does not fit into target type (overlong encoding)"

