package converter

import (
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func goOutput(name, value string, f *parser.Function) string {
	var vOld = value

	name = cleanName(name, value)
	value = cleanValue(value)

	switch value {
	case "QStringList":
		{
			return fmt.Sprintf("strings.Split(%v, \"|\")", goOutput(name, "QString", f))
		}

	case "uchar", "char", "QString", "QByteArray":
		{
			return fmt.Sprintf("C.GoString(%v)", name)
		}

	case "int", "long":
		{
			return fmt.Sprintf("int(%v)", name)
		}

	case "bool":
		{
			return fmt.Sprintf("%v != 0", name)
		}

	case "void", "":
		{
			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("unsafe.Pointer(%v)", name)
			}
			return name
		}

	case "T", "JavaVM", "jclass", "jobject":
		{
			switch f.TemplateMode {
			case "Int":
				{
					return fmt.Sprintf("int(%v)", name)
				}

			case "Boolean":
				{
					return fmt.Sprintf("int(%v) != 0", name)
				}

			case "Void":
				{
					return name
				}
			}
			return fmt.Sprintf("unsafe.Pointer(%v)", name)
		}

	case "qreal":
		{
			return fmt.Sprintf("float64(%v)", name)
		}

	case "qint64":
		{
			return fmt.Sprintf("int64(%v)", name)
		}

	case "WId":
		{
			return fmt.Sprintf("uintptr(%v)", name)
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			if c, exists := parser.ClassMap[class(cppEnum(f, value, false))]; exists && module(c.Module) != module(f) && module(c.Module) != "" {
				if parser.ClassMap[f.Class()].WeakLink[c.Module] {
					return fmt.Sprintf("int64(%v)", name)
				}
				return fmt.Sprintf("%v.%v(%v)", module(c.Module), goEnum(f, value), name)
			}
			return fmt.Sprintf("%v(%v)", goEnum(f, value), name)
		}

	case isClass(value):
		{
			if m := module(parser.ClassMap[value].Module); m != module(f) {
				if parser.ClassMap[f.Class()].WeakLink[parser.ClassMap[value].Module] {
					return fmt.Sprintf("unsafe.Pointer(%v)", name)
				}
				return fmt.Sprintf("%v.New%vFromPointer(%v)", m, value, name)
			}

			if f.Meta == "constructor" {
				return fmt.Sprintf("new%vFromPointer(%v)", value, name)
			}

			if f.TemplateMode == "String" {
				return fmt.Sprintf("New%vFromPointer(%v).ToString()", value, name)
			}

			return fmt.Sprintf("New%vFromPointer(%v)", value, name)
		}
	}

	f.Access = fmt.Sprintf("unsupported_goOutput(%v)", value)
	return f.Access
}

func goOutputFailed(value string, f *parser.Function) string {
	var vOld = value

	value = cleanValue(value)

	switch value {
	case "bool":
		{
			return "false"
		}

	case "int", "qreal", "qint64", "WId", "long":
		{
			return "0"
		}

	case "uchar", "char", "QString", "QByteArray":
		{
			return "\"\""
		}

	case "QStringList":
		{
			return "make([]string, 0)"
		}

	case "void", "":
		{
			if strings.Contains(vOld, "*") {
				return "nil"
			}
			return ""
		}

	case "T", "JavaVM", "jclass", "jobject":
		{
			switch f.TemplateMode {
			case "Int":
				{
					return "0"
				}

			case "Boolean":
				{
					return "false"
				}

			case "Void":
				{
					return ""
				}
			}

			return "nil"
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			return "0"
		}

	case isClass(value):
		{
			if f.TemplateMode == "String" {
				return "\"\""
			}

			return "nil"
		}
	}

	f.Access = fmt.Sprintf("unsupported_goOutputFailed(%v)", value)
	return f.Access
}

func cgoOutput(name, value string, f *parser.Function) string {

	var vOld = value

	name = cleanName(name, value)
	value = cleanValue(value)

	switch value {
	case "QStringList":
		{
			return fmt.Sprintf("strings.Split(%v, \"|\")", cgoOutput(name, "QString", f))
		}

	case "uchar", "char", "QString", "QByteArray":
		{
			return fmt.Sprintf("C.GoString(%v)", name)
		}

	case "int", "long":
		{
			return fmt.Sprintf("int(%v)", name)
		}

	case "bool":
		{
			return fmt.Sprintf("%v != 0", cgoOutput(name, "int", f))
		}

	case "void", "":
		{
			if strings.Contains(vOld, "*") {
				return name
			}
			return ""
		}

	case "qreal":
		{
			return fmt.Sprintf("float64(%v)", name)
		}

	case "qint64":
		{
			return fmt.Sprintf("int64(%v)", name)
		}

	case "WId":
		{
			return fmt.Sprintf("uintptr(%v)", name)
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			if c, exists := parser.ClassMap[class(cppEnum(f, value, false))]; exists && module(c.Module) != module(f) && module(c.Module) != "" {
				if parser.ClassMap[f.Class()].WeakLink[c.Module] {
					return fmt.Sprintf("unsafe.Pointer(%v)", name)
				}
				return fmt.Sprintf("%v.%v(%v)", module(c.Module), goEnum(f, value), name)
			}
			return fmt.Sprintf("%v(%v)", goEnum(f, value), name)
		}

	case isClass(value):
		{
			if m := module(parser.ClassMap[value].Module); m != module(f) {
				if parser.ClassMap[f.Class()].WeakLink[parser.ClassMap[value].Module] {
					return fmt.Sprintf("unsafe.Pointer(%v)", name)
				}
				return fmt.Sprintf("%v.New%vFromPointer(%v)", m, value, name)
			}
			return fmt.Sprintf("New%vFromPointer(%v)", value, name)
		}
	}

	f.Access = fmt.Sprintf("unsupported_cgoOutput(%v)", value)
	return f.Access
}

func CppOutput(name, value string, f *parser.Function) string {
	return cppOutput(name, value, f)
}

func cppOutput(name, value string, f *parser.Function) string {

	var vOld = value

	name = cleanName(name, value)
	value = cleanValue(value)

	switch value {
	case "QStringList":
		{
			return cppOutput(fmt.Sprintf("%v.join(\"|\")", name), "QString", f)
		}

	case "QString":
		{
			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("%v->toUtf8().data()", name)
			}
			return fmt.Sprintf("%v.toUtf8().data()", name)
		}

	case "QByteArray", "char":
		{
			return cppOutput(fmt.Sprintf("QString(%v)", name), "QString", f)
		}

	case "bool", "int", "long", "void", "", "T", "JavaVM", "jclass", "jobject":
		{
			if value == "void" {
				if strings.Contains(vOld, "*") {
					if strings.Contains(vOld, "const") {
						return fmt.Sprintf("const_cast<%v*>(%v)", value, name)
					}
					return name
				}
			}

			switch value {
			case "bool", "int", "long":
				{
					if strings.Contains(vOld, "*") {
						return fmt.Sprintf("*%v", name)
					}
				}
			}

			if value == "T" {
				if strings.Contains(vOld, "*") {
					if strings.Contains(vOld, "const") {
						return fmt.Sprintf("const_cast<void*>(%v)", name)
					}
				}
			}

			return name
		}

	case "qreal":
		{
			if strings.Contains(vOld, "*") {
				return fmt.Sprintf("*static_cast<double*>(%v)", name)
			}
			return fmt.Sprintf("static_cast<double>(%v)", name)
		}

	case "qint64":
		{
			return fmt.Sprintf("static_cast<long long>(%v)", name)
		}

	case "WId":
		{
			return fmt.Sprintf("static_cast<unsigned long long>(%v)", name)
		}
	}

	switch {
	case isEnum(f.Class(), value):
		{
			return name
		}

	case isClass(value):
		{
			if strings.Contains(vOld, "*") {
				if strings.Contains(vOld, "const") {
					return fmt.Sprintf("const_cast<%v*>(%v)", value, name)
				}
				return name
			}

			switch value {
			case "QModelIndex", "QMetaMethod", "QItemSelection":
				{
					return fmt.Sprintf("new %v(%v)", value, name)
				}

			case "QAndroidJniObject":
				{
					return fmt.Sprintf("new %v(%v.object())", value, name)
				}

			case "QPoint", "QPointF":
				{
					return fmt.Sprintf("new %v(static_cast<%v>(%v).x(), static_cast<%v>(%v).y())", value, value, name, value, name)
				}

			case "QSize", "QSizeF":
				{
					return fmt.Sprintf("new %v(static_cast<%v>(%v).width(), static_cast<%v>(%v).height())", value, value, name, value, name)
				}

			case "QRect", "QRectF":
				{
					return fmt.Sprintf("new %v(static_cast<%v>(%v).x(), static_cast<%v>(%v).y(), static_cast<%v>(%v).width(), static_cast<%v>(%v).height())", value, value, name, value, name, value, name, value, name)
				}

			case "QLine", "QLineF":
				{
					return fmt.Sprintf("new %v(static_cast<%v>(%v).p1(), static_cast<%v>(%v).p2())", value, value, name, value, name)
				}

			case "QMargins", "QMarginsF":
				{
					return fmt.Sprintf("new %v(static_cast<%v>(%v).left(), static_cast<%v>(%v).top(), static_cast<%v>(%v).right(), static_cast<%v>(%v).bottom())", value, value, name, value, name, value, name, value, name)
				}
			}

			for _, f := range parser.ClassMap[value].Functions {
				if f.Meta == "constructor" {
					if len(f.Parameters) == 1 {
						if cleanValue(f.Parameters[0].Value) == value {
							return fmt.Sprintf("new %v(%v)", value, name)
						}
					}
				}
			}
		}
	}

	f.Access = fmt.Sprintf("unsupported_cppOutput(%v)", value)
	return f.Access
}
