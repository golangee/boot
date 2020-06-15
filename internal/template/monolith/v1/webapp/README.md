# frontend for {{.BinaryName}}
This module contains the frontend as a web assembly module.
Usually, you do not build and deploy this separately. Use
the *make run* from the service itself, which takes care
of everything.

## run
The frontend does not contain server code to run the frontend. Use the 
root make file:

```bash
cd ..
make run
```

## build
Yet, you can build the frontend independently of the backend. If you
open the module in your IDE, to may need to set *GOOS=js GARCH=wasm*
so that code completion works properly.


## Debugging
Currently, only Chrome supports the DWARF format, however the Go linker does
not yet include the information. See [33503](https://github.com/golang/go/issues/33503)
for more information.

