#goSimpleServer

A go implementation of python's SimpleHTTPServer

##Install:

`go install github.com/worg/goSimpleServer`

##Basic Usage
Just run `goSimpleServer` to start serving files in the current directory on port `8080`

###Flags
Flag       | Default |  Behaviour
-----------|:-------:| ------
`-p`       | `8080`  | Port<br>[may be overridden also with the environment variable PORT]
`--path`   | `.`     | Specifies the path to be served
`--nocache`| `true`  | Specifies if no-cache header is sent*
>*Note: nocache flag needs '=' preceding the value  
>example: `goSimpleServer --nocache=false`   
>otherwise it won't work.

###License
```
This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
If a copy of the MPL was not distributed with this file,
You can obtain one at http://mozilla.org/MPL/2.0/.

© 2014 Hiram J. Pérez
```