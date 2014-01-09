pgu-go
===

Testing [Golang](http://golang.org/) on [app engine](https://developers.google.com/appengine/docs/go/gettingstarted/introduction) :squirrel:

Use of

- [pre-commit hook](http://golang.org/misc/git/pre-commit)
- [Intellij typefile](http://golang.org/misc/IntelliJIDEA/Go.xml)
- [chrome extension](http://golang.org/misc/chrome/gophertool)
- [bash completion](http://golang.org/misc/bash/go)

For [VI](https://code.google.com/p/go-wiki/wiki/IDEsAndTextEditorPlugins)
- [Official golang vim](http://golang.org/misc/vim/)
- [Gocode](https://github.com/nsf/gocode)
- [Pathogen](https://github.com/tpope/vim-pathogen)
- [Sensible](https://github.com/tpope/vim-sensible)
- [Syntastic](https://github.com/scrooloose/syntastic)
- [Nerdtree](https://github.com/scrooloose/nerdtree)
- [Tagbar](http://majutsushi.github.io/tagbar/) avec dep [Exuberant Ctags](http://ctags.sourceforge.net/) (./configure, make)
- [Gotags](https://github.com/jstemmer/gotags) for Tagbar
- [Vim compiler Go](https://github.com/rjohnsondev/vim-compiler-go)
- [Vim godef](https://github.com/dgryski/vim-godef) with [godef](https://code.google.com/p/rog-go/source/browse/exp/cmd/godef/)

[Useful metal3d article](http://www.metal3d.org/ticket/2013/07/07/vim-for-golang) :fr:

:heavy_exclamation_mark: **Integration with the app engine SDK**
- Add links to the SDK, it fixes the **vim compiler go**

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

Useful links on this topic: [1](http://stackoverflow.com/questions/21012037/go-cannot-find-package-appengine) [2](https://stackoverflow.com/questions/11286534/test-cases-for-go-and-appengine)

:satisfied:

