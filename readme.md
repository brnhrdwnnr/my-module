### Library
  ```
  buat public repository di github
  buat direktori cek-ganjil-genap
    go mod init github.com/brnhrdwnnr/cek-ganjil-genap
  buat fungsi CekGanjilGenap pada file cek_ganjil_genap.go
    git add .
    git commit -m "add fgo test -v -bench=.unction cekganjilgenap"
    git branch -M main
    git remote add origin https://github.com/brnhrdwnnr/cek-ganjil-genap.git
    git push -u origin main
  lakukan tagging v1.0.0
    git tag v1.0.0
    git tag
    git push origin v1.0.0
  update fungsi CekGanjilGenap menjadi v2.0.0
  tambahkan /v2 pada module di file go.mod
    git add .
    git commit -m "update v2.0.0"
    git push -u origin main
   lakukan tagging v2.0.0
    git tag v2.0.0
    git tag
    git push origin v2.0.0
  ```

### my-module
  ```
  buat direktori my-module
    go mod init my-module
    go get github.com/brnhrdwnnr/cek-ganjil-genap@v1.0.0
  buat fungsi pada file main.go
    go get github.com/brnhrdwnnr/cek-ganjil-genap/v2
  update fungsi pada file main.go
  ```

### unit-test
  ```
  go get github.com/stretchr/testify
  ```