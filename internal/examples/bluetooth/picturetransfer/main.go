//source: http://doc.qt.io/qt-5/qtbluetooth-picturetransfer-example.html

package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

func main() {
	app := gui.NewQGuiApplication(len(os.Args), os.Args)
	//! [Transfer-2]
	fileTransfer := NewFileTransfer(nil)

	view := quick.NewQQuickView(nil)
	view.RootContext().SetContextProperty("fileTransfer", fileTransfer)
	view.RootContext().SetContextProperty2("SystemPictureFolder", core.NewQVariant14(core.QStandardPaths_StandardLocations(core.QStandardPaths__PicturesLocation)[0]))

	fmt.Println(core.QStandardPaths_StandardLocations(core.QStandardPaths__PicturesLocation))

	view.SetSource(core.NewQUrl3("qrc:/bttransfer.qml", 0))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Engine().ConnectQuit(app.Quit)
	view.Show()

	gui.QGuiApplication_Exec()
}
