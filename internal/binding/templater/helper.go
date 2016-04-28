package templater

import (
	"runtime"
	"sort"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
)

func functionIsSupported(_ *parser.Class, f *parser.Function) bool {

	switch {
	case
		(f.Class() == "QAccessibleObject" || f.Class() == "QAccessibleInterface" || f.Class() == "QAccessibleWidget" || //QAccessible::State -> quint64
			f.Class() == "QAccessibleStateChangeEvent") && (f.Name == "state" || f.Name == "changedStates" || f.Meta == parser.CONSTRUCTOR),

		f.Fullname == "QPixmapCache::find" && f.OverloadNumber == "3", //Qt::Key -> int
		(f.Fullname == "QPixmapCache::remove" || f.Fullname == "QPixmapCache::insert") && f.OverloadNumber == "2",
		f.Fullname == "QPixmapCache::replace",

		f.Class() == "QSimpleXmlNodeModel" && f.Meta == parser.CONSTRUCTOR,

		f.Fullname == "QSGMaterialShader::attributeNames",

		f.Class() == "QVariant" && (f.Name == "value" || f.Name == "canConvert"), //needs template

		f.Fullname == "QNdefRecord::isRecordType", f.Fullname == "QScriptEngine::scriptValueFromQMetaObject", //needs template
		f.Fullname == "QScriptEngine::fromScriptValue", f.Fullname == "QJSEngine::fromScriptValue",

		f.Class() == "QMetaType" && //needs template
			(f.Name == "hasRegisteredComparators" || f.Name == "registerComparators" ||
				f.Name == "hasRegisteredConverterFunction" || f.Name == "registerConverter" ||
				f.Name == "registerEqualsComparator"),

		parser.ClassMap[f.Class()].Module == parser.MOC && f.Name == "metaObject", //needed for qtmoc

		f.Fullname == "QSignalBlocker::QSignalBlocker" && f.OverloadNumber == "2", //undefined symbol

		(f.Class() == "QCoreApplication" || f.Class() == "QGuiApplication" || f.Class() == "QApplication" ||
			f.Class() == "QAudioInput" || f.Class() == "QAudioOutput") && f.Name == "notify", //redeclared (name collision with QObject)

		f.Fullname == "QGraphicsItem::isBlockedByModalPanel", //** problem

		f.Name == "surfaceHandle", //QQuickWindow && QQuickView //unsupported_cppType(QPlatformSurface)

		f.Name == "readData", f.Name == "QNetworkReply", //TODO: char*

		strings.Contains(f.Access, "unsupported"), strings.ContainsAny(f.Signature, "<>"):

		{
			f.Access = "unsupported_isBlockedFunction"
			return false
		}
	}

	return true
}

func classIsSupported(c *parser.Class) bool {

	switch c.Name {
	case
		"QString", "QStringList", "QByteArray", //mapped to primitive

		"QExplicitlySharedDataPointer", "QFuture", "QDBusPendingReply", "QDBusReply", "QFutureSynchronizer", //needs template
		"QGlobalStatic", "QMultiHash", "QQueue", "QMultiMap", "QScopedPointer", "QSharedDataPointer",
		"QScopedArrayPointer", "QSharedPointer", "QThreadStorage", "QScopedValueRollback", "QVarLengthArray",
		"QWeakPointer", "QWinEventNotifier",

		"QFlags", "QException", "QStandardItemEditorCreator", "QSGSimpleMaterialShader", "QGeoCodeReply", "QFutureWatcher", //other
		"QItemEditorCreator", "QGeoCodingManager", "QGeoCodingManagerEngine",

		"QPlatformGraphicsBuffer", "QPlatformSystemTrayIcon", "QRasterPaintEngine", "QSupportedWritingSystems", "QGeoLocation", //file not found or QPA API
		"QAbstractOpenGLFunctions":

		{
			c.Access = "unsupported_isBlockedClass"
			return false
		}
	}

	switch {
	case
		strings.HasPrefix(c.Name, "QOpenGL"), strings.HasPrefix(c.Name, "QPlace"), //file not found or QPA API

		strings.HasPrefix(c.Name, "QAtomic"), //other

		strings.HasSuffix(c.Name, "terator"), strings.Contains(c.Brief, "emplate"): //needs template

		{
			c.Access = "unsupported_isBlockedClass"
			return false
		}
	}

	return true
}

func hasUnimplementedPureVirtualFunctions(className string) bool {
	for _, f := range parser.ClassMap[className].Functions {
		var f = *f
		cppFunction(&f)

		if f.Virtual == parser.PURE && !functionIsSupported(parser.ClassMap[className], &f) {
			return true
		}
	}
	return false
}

func ShouldBuild(module string) bool {
	return true //Build[module]
}

var Build = map[string]bool{
	"Core":              false,
	"AndroidExtras":     false,
	"Gui":               false,
	"Network":           false,
	"Sql":               false,
	"Xml":               false,
	"DBus":              false,
	"Nfc":               false,
	"Script":            false,
	"Sensors":           false,
	"Positioning":       false,
	"Widgets":           false,
	"MacExtras":         false,
	"Qml":               false,
	"WebSockets":        false,
	"XmlPatterns":       false,
	"Bluetooth":         false,
	"WebChannel":        false,
	"Svg":               false,
	"Multimedia":        false,
	"Quick":             true,
	"Help":              true,
	"Location":          true,
	"ScriptTools":       true,
	"MultimediaWidgets": true,
	"UiTools":           true,
}

var Libs = []string{
	"Core",
	"AndroidExtras",
	"Gui",
	"Network",
	"Sql",
	"Xml",
	"DBus",
	"Nfc",
	"Script",
	"Sensors",
	"Positioning",
	"Widgets",
	"MacExtras",
	"Qml",
	"WebSockets",
	"XmlPatterns",
	"Bluetooth",
	"WebChannel",
	"Svg",
	"Multimedia",
	"Quick",
	"Help",
	"Location",
	"ScriptTools",
	"MultimediaWidgets",
	"UiTools",
}

func GetLibs() []string {
	for i := len(Libs) - 1; i >= 0; i-- {
		switch {
		case
			runtime.GOOS != "darwin" && Libs[i] == "MacExtras",
			runtime.GOOS != "windows" && Libs[i] == "WinExtras",
			runtime.GOOS == "android" && (Libs[i] == "WebEngine" || Libs[i] == "Designer"):
			{
				Libs = append(Libs[:i], Libs[i+1:]...)
			}
		}
	}

	return Libs
}

var LibDeps = map[string][]string{
	"Core":              []string{"Widgets"},
	"AndroidExtras":     []string{"Core"},
	"Gui":               []string{"Core", "Widgets"},
	"Network":           []string{"Core"},
	"Sql":               []string{"Core", "Widgets"},
	"Xml":               []string{"Core", "XmlPatterns"},
	"DBus":              []string{"Core"},
	"Nfc":               []string{"Core"},
	"Script":            []string{"Core"},
	"Sensors":           []string{"Core"},
	"Positioning":       []string{"Core"},
	"Widgets":           []string{"Core", "Gui"},
	"MacExtras":         []string{"Core", "Gui"},
	"Qml":               []string{"Core", "Network"},
	"WebSockets":        []string{"Core", "Network"},
	"XmlPatterns":       []string{"Core", "Network"},
	"Bluetooth":         []string{"Core", "Concurrent"},
	"WebChannel":        []string{"Core", "Network", "Qml"},
	"Svg":               []string{"Core", "Gui", "Widgets"},
	"Multimedia":        []string{"Core", "Gui", "Network", "MultimediaWidgets"},
	"Quick":             []string{"Core", "Gui", "Network", "Widgets", "Qml", "QuickWidgets"},
	"Help":              []string{"Core", "Gui", "Network", "Sql", "CLucene", "Widgets"},
	"Location":          []string{"Core", "Gui", "Network", "Positioning", "Qml", "Quick"},
	"ScriptTools":       []string{"Core", "Gui", "Script", "Widgets"},
	"MultimediaWidgets": []string{"Core", "Gui", "Network", "Widgets", "OpenGL", "Multimedia"},
	"UiTools":           []string{"Core", "Gui", "Widgets"},

	/*
		CLucene
		Designer
		OpenGL
		Concurrent
		WinExtras
	*/

	parser.MOC: make([]string, 0),
}

func isGeneric(f *parser.Function) bool {

	if f.Class() == "QAndroidJniObject" {
		switch f.Name {
		case
			"callMethod",
			"callStaticMethod",

			"getField",
			//"setField", -> uses interface{} if not generic

			"getStaticField",
			//"setStaticField", -> uses interface{} if not generic

			"getObjectField",

			"getStaticObjectField",

			"callObjectMethod",
			"callStaticObjectMethod":
			{
				return true
			}

		case "setStaticField":
			{
				if f.OverloadNumber == "2" || f.OverloadNumber == "4" {
					return true
				}
			}
		}
	}

	return false
}

func needsCallbackFunctions(class *parser.Class) bool {
	for _, function := range class.Functions {
		if function.Virtual == parser.IMPURE || function.Virtual == parser.PURE || function.Meta == parser.SIGNAL || function.Meta == parser.SLOT {
			return true
		}
	}

	return false
}

func shortModule(module string) string {
	return strings.ToLower(strings.TrimPrefix(module, "Qt"))
}

func getSortedClassNamesForModule(module string) []string {
	var output = make([]string, 0)
	for _, class := range parser.ClassMap {
		if class.Module == module {
			output = append(output, class.Name)
		}
	}
	sort.Stable(sort.StringSlice(output))
	return output
}

func getSortedClassesForModule(module string) []*parser.Class {
	var (
		classNames = getSortedClassNamesForModule(module)
		output     = make([]*parser.Class, len(classNames))
	)
	for i, name := range classNames {
		output[i] = parser.ClassMap[name]
	}
	return output
}

func addCallbackNameFunctions(c *parser.Class) {
	if !c.IsQObjectSubClass() && needsCallbackFunctions(c) {
		c.Functions = append(c.Functions, &parser.Function{
			Name:     "objectNameAbs",
			Fullname: c.Name + "::" + "objectNameAbs",
			Access:   "public",
			Meta:     parser.PLAIN,
			Output:   "QString",
		})
		c.Functions = append(c.Functions, &parser.Function{
			Name:     "setObjectNameAbs",
			Fullname: c.Name + "::" + "setObjectNameAbs",
			Access:   "public",
			Meta:     parser.PLAIN,
			Output:   parser.VOID,
			Parameters: []*parser.Parameter{&parser.Parameter{
				Name:  "name",
				Value: "QString",
			}},
		})
	}
}

func manualWeakLink(module string) {
	for _, class := range getSortedClassesForModule(module) {
		class.WeakLink = make(map[string]bool)

		switch class.Module {
		case "QtCore":
			{
				class.WeakLink["QtWidgets"] = true
			}

		case "QtGui":
			{
				class.WeakLink["QtWidgets"] = true
				class.WeakLink["QtMultimedia"] = true
			}

		case "QtMultimedia":
			{
				class.WeakLink["QtMultimediaWidgets"] = true
			}
		}
	}
}
