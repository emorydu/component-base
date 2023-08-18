// Copyright 2023 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package net

// IsValidPort check if the port is legal. 0 is considered as a non valid port.
func IsValidPort(port int) bool {
	return port > 0 && port < 65535
}
