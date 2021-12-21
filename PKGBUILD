pkgname=find-bytes
pkgver=0.1
pkgrel=1
pkgdesc="a tool to find bytes"
arch=('x86_64')
url="https://github.com/kmille/find-bytes"
license=('GPL3')
makedepends=(
  'go'
)
source=(git+https://github.com/kmille/find-bytes)
sha512sums=('SKIP')

build() {
  export GOPATH="$srcdir"/gopath
  export CGO_CPPFLAGS="${CPPFLAGS}"
  export CGO_CFLAGS="${CFLAGS}"
  export CGO_CXXFLAGS="${CXXFLAGS}"
  export CGO_LDFLAGS="${LDFLAGS}"
  export CGO_ENABLED=1

  cd "$srcdir/$pkgname"
  go build
}

package() {
  cd "$srcdir/$pkgname"
  install -Dm 755 find-bytes $pkgdir/usr/local/bin/$pkgname
}
