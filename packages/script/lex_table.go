// MIT License
//
// Copyright (c) 2016-2018 GACHAIN
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package script

// This file was generated with /tools/lextable.go

var (
	alphabet = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 1, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 2, 20, 4, 14, 22, 0, 12, 0, 6, 7, 21, 24, 16, 25, 15, 26, 28,
		29, 29, 29, 29, 29, 29, 29, 29, 29, 0, 5, 17, 19, 18, 0, 23, 30, 30, 30, 30, 30, 30, 30, 30,
		30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 8, 27, 9, 0, 31, 3,
		30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
		30, 30, 10, 13, 11, 0, 0, 32,
	}
	lexTable = [][33]uint32{
		{0xff0000, 0x501, 0x1, 0x30003, 0xe0003, 0x501, 0x101, 0x101, 0x101, 0x101, 0x101, 0x101, 0x40003, 0x50003, 0x101, 0xa0003, 0x101, 0xf0003, 0xf0003, 0x60003, 0xf0003, 0x201, 0x80003, 0x80003, 0x201, 0x201, 0x70003, 0xff0000, 0xc0003, 0xc0003, 0x80003, 0x80003, 0x80003},
		{0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0x405, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000},
		{0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0x705, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001},
		{0x30001, 0x30001, 0x30001, 0x605, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001, 0x30001},
		{0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0x205, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000},
		{0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0x205, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000, 0xff0000},
		{0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x205, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104},
		{0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0xd0001, 0x204, 0x204, 0x204, 0x204, 0x90005, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204},
		{0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x404, 0x80001, 0x80001, 0x80001, 0x80001, 0x80001},
		{0x90001, 0x0, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001, 0x90001},
		{0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x10001, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104, 0x104},
		{0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001},
		{0x304, 0x304, 0x304, 0x304, 0x304, 0x304, 0x304, 0x304, 0x304, 0x304, 0x304, 0x304, 0x304, 0x304, 0x304, 0xc0001, 0x304, 0x304, 0x304, 0x304, 0x304, 0x304, 0x304, 0x304, 0x304, 0x304, 0x304, 0x304, 0xc0001, 0xc0001, 0xff0000, 0xff0000, 0xff0000},
		{0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0x20001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001, 0xd0001},
		{0xe0001, 0xe0001, 0xe0001, 0xe0001, 0x605, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xb0008, 0xe0001, 0xe0001, 0xe0001, 0xe0001, 0xe0001},
		{0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x205, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204, 0x204},
	}
)