# Maintainer: Sascha Teske <sascha.teske@microprojects.de>
pkgname=('e')
pkgbase=e
pkgver=1.0
pkgrel=1
arch=('i686' 'x86_64' 'armv6h' 'armv7h' 'aarch64')
url="https://github.com/slaxor/e"
license=('GPL3')
makedepends=('go')
depends=('neovim')
source=("$pkgbase-$pkgver.tar.gz::https://github.com/slaxor/e/archive/v$pkgver.tar.gz")
sha256sums=('')

build() {
  cd "$pkgbase-$pkgver"
  go get -u -v
  go build -o e .
}

package() {
  pkgdesc=""
  depends=('go')
  provides=("$pkgbase")
  conflicts=("$pkgbase")

  cd "$pkgbase-$pkgver"
  install -s -m 755 -t /usr/bin e
}
