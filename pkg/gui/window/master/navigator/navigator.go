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

package navigator

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/bit-fever/client/pkg/core/eventbus"
)

//=============================================================================

func NewNavigator() fyne.Widget {

	initStruct()

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return navChildren[uid]
		},

		IsBranch: func(uid string) bool {
			children, ok := navChildren[uid]
			return ok && len(children) > 0
		},

		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel(navWidth)
		},

		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			item, ok := navMap[uid]
			if !ok {
				fyne.LogError("Missing navigation item: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(item.Name)
		},

		OnSelected: func(uid string) {
			if _, ok := navMap[uid]; ok {
				//a.Preferences().SetString(preferenceCurrentTutorial, uid)
				eventbus.FireEvent(uid)
			} else {
				fyne.LogError("Unknown event to fire: "+uid, nil)
			}
		},
	}
	return tree
}

//=============================================================================

func initStruct() {
	initMenu("", navMenu)
}

//=============================================================================

func initMenu(parent string, items []navItem) {

	for _, item := range items {
		navMap[item.Id] = item

		list := navChildren[parent]
		if list == nil {
			list = []string{}
		}

		navChildren[parent] = append(list, item.Id)
		initMenu(item.Id, item.items)
	}
}

//=============================================================================
