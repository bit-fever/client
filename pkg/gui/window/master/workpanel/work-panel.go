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

package workpanel

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bit-fever/client/pkg/core/eventbus"
	"github.com/bit-fever/client/pkg/gui/panel"
	"github.com/bit-fever/client/pkg/gui/window/master/workpanel/portfolio"
	"github.com/bit-fever/client/pkg/model/event"
)

//=============================================================================

func init() {
}

//=============================================================================

func NewWorkPanel() fyne.CanvasObject {
	cp := panel.NewCardPanel()
	cp.AddPanel(event.NavMenu_Portfolio_TradingSystems, portfolio.NewTradingSystems())
	cp.AddPanel(event.NavMenu_Portfolio_Monitoring, widget.NewLabel("Collection Widgets"))
	cp.AddPanel("-", container.NewMax())
	//cp.ShowPanel("-")
	cp.ShowPanel(event.NavMenu_Portfolio_TradingSystems)

	eventbus.AddHandler(event.Any, func(e string) {
		cp.ShowPanel(e)
	})

	return cp
}

//=============================================================================
