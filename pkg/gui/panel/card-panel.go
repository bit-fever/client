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

package panel

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
)

//=============================================================================

type CardPanel struct {
	fyne.Container

	subPanels map[string]fyne.CanvasObject
	currId    string
	currPanel *fyne.CanvasObject
}

//=============================================================================

func NewCardPanel() *CardPanel {
	p := &CardPanel{}
	p.subPanels = map[string]fyne.CanvasObject{}
	p.Layout = layout.NewMaxLayout()
	p.Resize(p.Layout.MinSize(p.Objects))
	p.Layout.Layout(p.Objects, p.Size())
	return p
}

//=============================================================================

func (p *CardPanel) AddPanel(id string, sp fyne.CanvasObject) {
	p.subPanels[id] = sp
	p.Add(sp)
}

//=============================================================================

func (p *CardPanel) ShowPanel(id string) bool {
	if sp, ok := p.subPanels[id]; ok {
		p.Objects = []fyne.CanvasObject{sp}
		p.Refresh()

		return true
	}

	return false
}

//=============================================================================
