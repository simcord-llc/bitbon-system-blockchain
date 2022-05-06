// Code generated by "enumer -type=OperationType"; DO NOT EDIT.

package dto

import (
	"fmt"
	"strings"
)

const (
	_OperationTypeName_0      = "UndefinedOperationType"
	_OperationTypeLowerName_0 = "undefinedoperationtype"
	_OperationTypeName_1      = "QuickTransferOperationTypeQuickTransferAllOperationTypeCreateSafeTransferOperationTypeCreateSafeTransferAllOperationTypeCreateWpcSafeTransferOperationTypeCreateWpcSafeTransferAllOperationType"
	_OperationTypeLowerName_1 = "quicktransferoperationtypequicktransferalloperationtypecreatesafetransferoperationtypecreatesafetransferalloperationtypecreatewpcsafetransferoperationtypecreatewpcsafetransferalloperationtype"
	_OperationTypeName_2      = "CreateAssetboxOperationType"
	_OperationTypeLowerName_2 = "createassetboxoperationtype"
)

var (
	_OperationTypeIndex_0 = [...]uint8{0, 22}
	_OperationTypeIndex_1 = [...]uint8{0, 26, 55, 86, 120, 154, 191}
	_OperationTypeIndex_2 = [...]uint8{0, 27}
)

func (i OperationType) String() string {
	switch {
	case i == 0:
		return _OperationTypeName_0
	case 101 <= i && i <= 106:
		i -= 101
		return _OperationTypeName_1[_OperationTypeIndex_1[i]:_OperationTypeIndex_1[i+1]]
	case i == 201:
		return _OperationTypeName_2
	default:
		return fmt.Sprintf("OperationType(%d)", i)
	}
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _OperationTypeNoOp() {
	var x [1]struct{}
	_ = x[UndefinedOperationType-(0)]
	_ = x[QuickTransferOperationType-(101)]
	_ = x[QuickTransferAllOperationType-(102)]
	_ = x[CreateSafeTransferOperationType-(103)]
	_ = x[CreateSafeTransferAllOperationType-(104)]
	_ = x[CreateWpcSafeTransferOperationType-(105)]
	_ = x[CreateWpcSafeTransferAllOperationType-(106)]
	_ = x[CreateAssetboxOperationType-(201)]
}

var _OperationTypeValues = []OperationType{UndefinedOperationType, QuickTransferOperationType, QuickTransferAllOperationType, CreateSafeTransferOperationType, CreateSafeTransferAllOperationType, CreateWpcSafeTransferOperationType, CreateWpcSafeTransferAllOperationType, CreateAssetboxOperationType}

var _OperationTypeNameToValueMap = map[string]OperationType{
	_OperationTypeName_0[0:22]:         UndefinedOperationType,
	_OperationTypeLowerName_0[0:22]:    UndefinedOperationType,
	_OperationTypeName_1[0:26]:         QuickTransferOperationType,
	_OperationTypeLowerName_1[0:26]:    QuickTransferOperationType,
	_OperationTypeName_1[26:55]:        QuickTransferAllOperationType,
	_OperationTypeLowerName_1[26:55]:   QuickTransferAllOperationType,
	_OperationTypeName_1[55:86]:        CreateSafeTransferOperationType,
	_OperationTypeLowerName_1[55:86]:   CreateSafeTransferOperationType,
	_OperationTypeName_1[86:120]:       CreateSafeTransferAllOperationType,
	_OperationTypeLowerName_1[86:120]:  CreateSafeTransferAllOperationType,
	_OperationTypeName_1[120:154]:      CreateWpcSafeTransferOperationType,
	_OperationTypeLowerName_1[120:154]: CreateWpcSafeTransferOperationType,
	_OperationTypeName_1[154:191]:      CreateWpcSafeTransferAllOperationType,
	_OperationTypeLowerName_1[154:191]: CreateWpcSafeTransferAllOperationType,
	_OperationTypeName_2[0:27]:         CreateAssetboxOperationType,
	_OperationTypeLowerName_2[0:27]:    CreateAssetboxOperationType,
}

var _OperationTypeNames = []string{
	_OperationTypeName_0[0:22],
	_OperationTypeName_1[0:26],
	_OperationTypeName_1[26:55],
	_OperationTypeName_1[55:86],
	_OperationTypeName_1[86:120],
	_OperationTypeName_1[120:154],
	_OperationTypeName_1[154:191],
	_OperationTypeName_2[0:27],
}

// OperationTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func OperationTypeString(s string) (OperationType, error) {
	if val, ok := _OperationTypeNameToValueMap[s]; ok {
		return val, nil
	}
	s = strings.ToLower(s)
	if val, ok := _OperationTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to OperationType values", s)
}

// OperationTypeValues returns all values of the enum
func OperationTypeValues() []OperationType {
	return _OperationTypeValues
}

// OperationTypeStrings returns a slice of all String values of the enum
func OperationTypeStrings() []string {
	strs := make([]string, len(_OperationTypeNames))
	copy(strs, _OperationTypeNames)
	return strs
}

// IsAOperationType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i OperationType) IsAOperationType() bool {
	for _, v := range _OperationTypeValues {
		if i == v {
			return true
		}
	}
	return false
}