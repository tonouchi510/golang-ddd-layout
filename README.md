# golang-ddd-layout
GolangでDDD実践するときのプロジェクトレイアウトとサンプルコード

## 参照
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- [ドメイン駆動設計入門](https://www.amazon.co.jp/dp/479815072X)


## 備考
- パッケージ構成
  - 集約を単位とした分割にしておく
  - あとでマイクロサービスに切り出されることを想定した構造にする
- private変数・メソッドに対するテスト
  - testはファイル名を`〇〇_test.go`にすれば十分。パッケージ名を揃えることで、privateなものにもアクセスできるようにする
