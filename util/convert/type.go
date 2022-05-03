package convert

import "database/sql"

// StringToNullString ...
func StringToNullString(s string) (ns sql.NullString) {
	if s != "" {
		ns.String = s
		ns.Valid = true
	}
	return ns
}

//Float64ToNullFloat64 ...
func Float64ToNullFloat64(f float64) (nf sql.NullFloat64) {
	if f != 0 {
		nf.Float64 = f
		nf.Valid = true
	}
	return nf
}

// EmptyStringToNull also handles the case when user explicity wants to delete photo
func EmptyStringToNull(s string) (res sql.NullString) {
	if s == "" {
		res.Valid = false
		return res
	} else if s == "delete" {
		res.String = ""
		res.Valid = true
		return res
	} else {
		res.String = s
		res.Valid = true
		return res
	}
}

func EmptyArrayStringToNullArray(array []string) []sql.NullString {
	var res []sql.NullString
	if len(array) == 0 {
		return res
	}
	for _, elem := range array {
		if elem == "" {
			res = append(res, sql.NullString{Valid: false})
		} else if elem == "delete" {
			res = append(res, sql.NullString{String: "", Valid: true})
		} else {
			res = append(res, sql.NullString{String: elem, Valid: true})
		}
	}
	return res
}
