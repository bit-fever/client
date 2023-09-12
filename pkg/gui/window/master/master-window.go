//=============================================================================
/*
Copyright © 2023 Andrea Carboni andrea.carboni71@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
//=============================================================================

package master

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/bit-fever/client/pkg/gui/window/master/navigator"
	"github.com/bit-fever/client/pkg/gui/window/master/statusbar"
	"github.com/bit-fever/client/pkg/gui/window/master/toolbar"
	"github.com/bit-fever/client/pkg/gui/window/master/workpanel"
)

//=============================================================================

func NewWindow(app fyne.App) fyne.Window {
	w := app.NewWindow("BitFever Client")

	toolBar := toolbar.NewToolBar()
	navMenu := navigator.NewNavigator()
	workPanel := workpanel.NewWorkPanel()
	statusBar := statusbar.NewStatusBar()
	split := container.NewHSplit(navMenu, workPanel)

	mainPanel := container.NewBorder(toolBar, statusBar, nil, nil, split)
	w.SetContent(mainPanel)
	w.SetMaster()
	w.Resize(fyne.NewSize(400, 400))
	return w
}

//=============================================================================
