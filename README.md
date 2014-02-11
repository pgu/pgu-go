pgu-go
===
[→ Here](http://pgu-go.appspot.com/)

Testing [Golang](http://golang.org/) on [app engine](https://developers.google.com/appengine/docs/go/gettingstarted/introduction) :squirrel:

```
$GOROOT is set to $golang_gae/goroot
```

Use of

- [pre-commit hook](http://golang.org/misc/git/pre-commit)
- [Intellij typefile](http://golang.org/misc/IntelliJIDEA/Go.xml)
- [chrome extension](http://golang.org/misc/chrome/gophertool)
- [bash completion](http://golang.org/misc/bash/go)

For [VI](https://code.google.com/p/go-wiki/wiki/IDEsAndTextEditorPlugins)
- [Official golang vim](http://golang.org/misc/vim/)
- [Gocode](https://github.com/nsf/gocode)
- [Goimports](http://godoc.org/code.google.com/p/go.tools/cmd/goimports)
- [Gotags](https://github.com/jstemmer/gotags) for Tagbar
- [Vim compiler Go](https://github.com/rjohnsondev/vim-compiler-go)
- [Vim godef](https://github.com/dgryski/vim-godef) with [godef](https://code.google.com/p/rog-go/source/browse/exp/cmd/godef/)
- [Golint](https://github.com/golang/lint) Deactivated for now


[Useful metal3d article](http://www.metal3d.org/ticket/2013/07/07/vim-for-golang) :fr:

### More Resources

Useful links about Golang + [AngularJS](http://angularjs.org/) + AppEngine

- Francesc Campoy Flores Demo
 - [video](http://pivotallabs.com/francesc-flores-go-language/)
 - [slides](http://go-talks.appspot.com/github.com/campoy/todo/talk/talk.slide#1)
 - [code](https://github.com/campoy/todo)

### Sublime
- [Sublime Text 2.0.2](http://www.sublimetext.com/)
- [GoSublime](https://github.com/DisposaBoy/GoSublime)
- [GoImports](https://michaelwhatcott.com/gosublime-goimports/)


### Old steps :older_man:

**Integration with the app engine SDK**
- Add links to the SDK, it fixes the **vim compiler go** :trollface:

```
ln -s $APPENGINE_SDK/goroot/src/pkg/appengine $GOROOT/src/pkg/
ln -s $APPENGINE_SDK/goroot/src/pkg/appengine_internal $GOROOT/src/pkg/
mkdir -p $GOROOT/src/pkg/code.google.com/p/
ln -s $APPENGINE_SDK/goroot/src/pkg/code.google.com/p/goprotobuf $GOROOT/src/pkg/code.google.com/p/
```

- Update **gocode** lib, it fixes the autocompletion for appengine packages

```
gocode set lib-path "$APPENGINE_SDK/goroot/pkg/linux_amd64_appengine"
```

Useful links about this topic: [1](http://stackoverflow.com/questions/21012037/go-cannot-find-package-appengine) [2](https://stackoverflow.com/questions/11286534/test-cases-for-go-and-appengine)


