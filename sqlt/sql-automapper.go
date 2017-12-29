package sqlt

import (
	"reflect"
	"errors"
	"strings"
	"regexp"
	"time"
	//"fmt"
)

type TypeChange struct{
	DefaultGoTypeToSql map[string]string
	TagType []string
}

// Shared Resource, It's not safe for multi-thread
var tablesMeta map[string]string

func NewEnergine(driveName string, dataSourceName string) (*DatabaseConn,error) {
    db, err := Open(driveName, dataSourceName)
    if err != nil{
    	return nil, err
	}

    return &DatabaseConn{ db }, nil
}

// RegisterTable .
//No exist, create it.
//Exist, registeCURDSql.
func (db *DatabaseConn)RegisterTable(o interface{}, tableName string) (*SQLAutoMapper ,error) {
	var str string = "CREATE TABLE "
	value := reflect.Indirect(reflect.ValueOf(o))
	Typeofs := value.Type()
	Tc := NewTypeChange()
	tablesMeta  = make(map[string]string)
	if value.Kind() != reflect.Struct {
		return nil, errors.New("Type error: "+ reflect.ValueOf(o).Kind().String())
	}


	//get Table name
	str += "`" + tableName + "`( "

	//get field name
	for i := 0; i < value.NumField(); i++ {
		defaultName := strings.ToLower(Typeofs.Field(i).Name)
		defaultType := Tc.getDefaultType(value.Field(i).Type())

		trueTags := ""
		tags := strings.Split(Typeofs.Field(i).Tag.Get("myorm"), " ")
		for _, v := range tags{
			if Tc.TagExistType(v){
				defaultType = v
				continue
			}
			if match,_ := regexp.MatchString(`'\w+'`, v); match{
				defaultName = string([]rune(v)[1:len(v)-1])
				continue
			}
			trueTags += " " + v
		}
		str += " `" + defaultName + "` " + defaultType + trueTags + ","
	}
	str  = string([]rune(str)[0:len(str)-1]) + ");"
    _, err := db.Sqldb.Exec(str)
    if err == nil {
        tablesMeta[value.Type().String()] = tableName
		return registerCURDSql(db.Sqldb),nil
	}
	if strings.Contains(err.Error(), "Table") && strings.Contains(err.Error(), "already exists"){
		if _, ok := tablesMeta[value.Type().String()]; !ok{
			tablesMeta[value.Type().String()] = tableName
		}
		return registerCURDSql(db.Sqldb),nil
	}
	return nil, err
}

func registerCURDSql(o interface{}) *SQLAutoMapper {
	return &SQLAutoMapper{ NewSQLTemplate(o.(SQLExecer))}
}

//SQLAutoMapper .
type SQLAutoMapper struct {
	tpl SQLTemplate
}

//Save .
func (am *SQLAutoMapper) Save(o interface{}) error {
   var str string = "INSERT INTO " + tablesMeta[reflect.Indirect(reflect.ValueOf(o)).Type().String()] + "("
   var templateVaule string = "values("
   var Field []interface{}
   value := reflect.Indirect(reflect.ValueOf(o))

   for i := 0; i < value.NumField() -1; i++{
	   defaultName := strings.ToLower(value.Type().Field(i).Name)

	   tags := strings.Split(value.Type().Field(i).Tag.Get("myorm"), " ")
	   for _, v := range tags {
		   if match, _ := regexp.MatchString(`'\w+'`, v); match {
			   defaultName = string([]rune(v)[1:len(v)-1])
		   }
	   }
        str += " " + defaultName + ","
        templateVaule += " " + "?,"
        Field = append(Field, value.Field(i).Interface())
   }
   str += " " + strings.ToLower(value.Type().Field(value.NumField() -1).Name) + " ) " + templateVaule + " ? )"
   Field = append(Field, value.Field(value.NumField() -1).Interface())
   err := am.tpl.Insert(str, nil, Field...)
   if err != nil{
   	   return err
   }
	return nil
}

//FindAll .
func (am *SQLAutoMapper) Find(t reflect.Type, subSQL string, args ...interface{}) ([]interface{}, error) {
	if t.Kind() == reflect.Ptr{
		t = t.Elem()
	}
	T := reflect.New(t).Elem()
	count := T.NumField()

	valuePtrs := make([]interface{}, count)
	values := make([]interface{}, count)
    resultList := make([]interface{},0,0)
	if T.NumField() != count{
		return nil, errors.New("TypeError: "+t.String()+" is invalid for this table")
	}

	for i := 0; i < count; i++ {
		valuePtrs[i] = &values[i]
	}
	var findall = func(rscanner RowScanner) error{
		result := make([]interface{}, count)
		err := rscanner.Scan(valuePtrs...)
		if err != nil{
			return err
		}

		for i := 0; i < count; i++{
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)

			if ok {
				v = string(b)
			} else {
				v = val
			}

			if T.Field(i).Type().String() == "*time.Time" || T.Field(i).Type().String() == "time.Time"{
				createtime, err := time.Parse("2006-01-02 15:04:05", string(b))
				if err != nil{
					return err
				}else{
					result[i] = createtime
		        }
			    continue
	        }
	        result[i] = v
	    }
		resultList = append(resultList, result)
		return nil
	}
	err := am.tpl.Select(subSQL, findall, args...)
	if err != nil{
		return nil, err
	}
	return resultList, nil
}

//FindOne
func (am *SQLAutoMapper) FindOne(t reflect.Type, subSQL string, args ...interface{}) (interface{}, error) {
	if t.Kind() == reflect.Ptr{
		t = t.Elem()
	}
	T := reflect.New(t).Elem()
	count := T.NumField()

	valuePtrs := make([]interface{}, count)
	values := make([]interface{}, count)
	result := make([]interface{}, count)
	if T.NumField() != count{
		return nil, errors.New("TypeError: "+t.String()+" is invalid for this table")
	}

	for i := 0; i < count; i++ {
		valuePtrs[i] = &values[i]
	}
	var findone = func(rscanner RowScanner) error{
		err := rscanner.Scan(valuePtrs...)
		if err != nil{
			return err
		}

		for i := 0; i < count; i++{
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)

			if ok {
				v = string(b)
			} else {
				v = val
			}

			if T.Field(i).Type().String() == "*time.Time" || T.Field(i).Type().String() == "time.Time"{
				createtime, err := time.Parse("2006-01-02 15:04:05", string(b))
				if err != nil{
					return err
				}else{
					result[i] = createtime
				}
				continue
			}
			result[i] = v
		}
		return nil
	}
	err := am.tpl.SelectOne(subSQL, findone, args...)
	if err != nil{
		return nil, err
	}
	return result, nil
}

//Update

//initial TypeChange
func NewTypeChange() *TypeChange {
    typechange := &TypeChange{
    	DefaultGoTypeToSql: make(map[string]string),
    }
	typechange.DefaultGoTypeToSql = map[string]string{
		"int":"INT", "int8":"INT", "int16":"INT", "int32":"INT",
		"uint":"INT", "uint8":"INT","uint16":"INT", "uint32":"INT",
        "int64":"BIGINT", "uint64":"BIGINT", "floa32t":"FLOAT", "float64":"DOUBLE",
        "complex64":"VARCHAR(64)", "complex128":"VARCHAR(64)", "[]uint8":"BLOB",
        "array":"TEXT", "slice":"TEXT", "map":"TEXT", "bool":"TINYINT", "string":"Varchar(255)",
        "time.Time":"DATETIME", "cascade struct":"BIGINT", "struct":"TEXT",
	}
	typechange.TagType = []string{"BIT", "INT", "TINYINT", "SMALLINT",
	    "MEDIUMINT", "INTEGER", "BIGINT", "CHAR", "VARCHAR", "TINYTEXT", "TEXT",
	    "MEDIUMTEXT", "LONGTEXT", "BINARY", "VARBINARY", "DATE", "DATETIME",
	    "TIME", "TIMESTAMP", "REAL", "FLOAT", "DOUBLE", "DECIMAL", "NUMERIC",
	    "TINYBLOB", "BLOB", "MEDIUMBLOB", "LONGBLOB"}
	return typechange
}

func (tc *TypeChange)getDefaultType(p reflect.Type) string{
	if p.Kind().String() == "struct" && p.String() == "time.Time"{
		return tc.DefaultGoTypeToSql[p.String()]
	}
	if (p.Kind().String() == "array" || p.Kind().String() == "slice") && p.String() == "[]uint8"{
		return tc.DefaultGoTypeToSql["[]uint8"]
	}
	if v, ok := tc.DefaultGoTypeToSql[p.String()]; ok{
		return v
	}
	return "TEXT"
}

func (tc *TypeChange)TagExistType(s string) bool{
    for _,v := range tc.TagType{
    	reg := regexp.MustCompile(v+"("+"\\d+"+")")
    	if reg.MatchString(s){
    		return true
		}
    	if v == s {
    		return true
		}
	}

	return false
}
