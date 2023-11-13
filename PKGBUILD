# Maintainer: Sascha Teske <sascha.teske@microprojects.de>
pkgname=('editor-git')
pkgbase="editor"
pkgver=r4.415990a
pkgrel=1
provides=(e)
arch=('i686' 'x86_64' 'armv6h' 'armv7h' 'aarch64')
url="https://github.com/slaxor/editor"
license=('GPL3')
makedepends=('go')
depends=('neovim')
source=("${srcdir}/${pkgname}-${pkgver}::git+https://github.com/slaxor/${pkgbase}.git")
sha256sums=('SKIP')

pkgver() {
    cd "${srcdir}/${pkgname}-${pkgver}"
    if git describe --long --tags >/dev/null 2>&1; then
        git describe --long --tags | sed 's/^v//;s/\([^-]*-g\)/r\1/;s/-/./g'
    else
        printf 'r%s.%s' "$(git rev-list --count HEAD)" "$(git describe --always)"
    fi
}

# prepare() {
  # cd "$pkgbase-$pkgver"
  # sed -i 's|/usr/bin/env go|/usr/bin/go|g' e
# }

build() {
  cd "${srcdir}/${pkgname}-${pkgver}"
  go get -u -v
  go build -o e .
}

package() {
  # pkgdesc=""
  # depends=('go')
  # provides=("${pkgname}")
  # conflicts=("${pkgname}")

  # cd "${pkgname}-${pkgver}"
  # install -s -m 755 -t /usr/bin e
  cd "${srcdir}/${pkgname}-${pkgver}"
  install -Dm755 -t "${pkgdir}/usr/bin" "e"
}
