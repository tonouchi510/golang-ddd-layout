# golang-ddd-layout
GolangでDDD実践するときのプロジェクトレイアウトとサンプルコード（WIP）

※2025/05追記：こちらのリポジトリはかなり古く良くない実装も多いので、こちらを参照してください（最新版）  
=> https://github.com/tonouchi510/application-arch-blueprint

## 参照
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- [ドメイン駆動設計入門](https://www.amazon.co.jp/dp/479815072X)
  - こちらのC++のコード例をGoで実装してみた

## ディレクトリ構成

```
$ tree -L 3 -d ./internal
./internal
├── application           # アプリケーション層の実装
│   ├── circles
│   └── users
├── domain                # ドメイン層の実装
│   ├── models              # ドメインモデル。内部で集約ごとにパッケージ分割して実装。
│   │   ├── circles           # サークル集約の実装
│   │   └── users             # ユーザ集約の実装
│   ├── services            # ドメインサービスの実装
│   └── shared              # 共通モジュール
├── infrastructure        # インフラ層の実装
│   └── sqlboiler           # 集約ごとにパッケージ分割
│       ├── circles
│       ├── models
│       └── users
└── query                 # CQRSでいうQueryの実装
    ├── circles
    └── users
