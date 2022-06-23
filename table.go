// MIT License

// Copyright (c) 2022 Tree Xie

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package charts

type tableChart struct {
	p   *Painter
	opt *TableChartOption
}

func NewTableChart(p *Painter, opt TableChartOption) *tableChart {
	if opt.Theme == nil {
		opt.Theme = defaultTheme
	}
	return &tableChart{
		p:   p,
		opt: &opt,
	}
}

type TableChartOption struct {
	// The theme
	Theme ColorPalette
	// The padding of table cell
	Padding Box
	// The header data of table
	HeaderData []string
	// The data of table
	Data [][]string
}

func (c *tableChart) Render() (Box, error) {
	return BoxZero, nil
}