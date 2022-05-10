/*
Copyright © 2022 Jan Kiesewetter <jan@t3easy.de>

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package internal

import "testing"

func TestMainFunc(t *testing.T) {
	var (
		message  string
		expected string
		name     string
	)
	name = "Go"
	message = GetMessage(name)
	expected = "Hello Go!"
	if message != expected {
		t.Errorf("Main returned %s, expected %s", message, expected)
	}

	name = "Gopher"
	message = GetMessage(name)
	expected = "Hello Gopher!"
	if message != expected {
		t.Errorf("Main returned %s, expected %s", message, expected)
	}
}
