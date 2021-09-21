package cours

import (
	"fmt"
)

// pointer => variable

var lenomdemonVariablePrive string // nil => null
var lenomdemonVariablePublic string // nil => null

var lenomdemonVariablePriveavecValue = "prive"
var lenomdemonVariablePublicavecValue = "public"

var lenomdemonVariablePriveavecValueType string = "priveType"
var lenomdemonVariablePublicavecValueType string = "publicType"

const lenomdemonConstPrive= "priveConst"
const lenomdemonConstPublic = "publicConst"


const lenomdemonConstPriveavecValueType string = "priveTypeConst"
const lenomdemonConstPublicavecValueType string = "publicTypeConst"

func hello(){
	fmt.Println(lenomdemonConstPublicavecValueType);
}

