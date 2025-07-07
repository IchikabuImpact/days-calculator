# Days Calculator

Go 製のシンプルなユーティリティです。**N 日前**の日付を *CLI* でも *REST API* でも取得できます。
作った経緯：dotinstallでマイペースに学習をしていると「N日前に完了済です」のNの値が2517日前などの表示がざらにあったりします。そうするとN日前は何年何月何日なのか気になります。
そういう時はPromptでAIに聞けば良いのですが、ちょっとした実用的なものをAPIなどで小さく作っておくと後で出番があるかもしれません。
このプログラムは一つのことを上手くやるだけのシンプルなものです。

---

## 特長

- **CLI モード** : `-days` フラグまたは位置引数で即座に日付を出力。
- **HTTP API** : `/api/calculate?days=N` で JSON を返却。
- `.env` でポートを切り替え可能（既定 `8080`）。
- Docker イメージ提供でどこでも動作。

---

## 動作環境

| 必須     | バージョン例                  |
| ------ | ----------------------- |
| Go     | 1.22 以上                 |
| Docker | 24.x 以上 (API だけ使う場合は任意) |

---

## 1️⃣ CLI で使う

> **ポイントは 3 行**：
>
> 1. **ビルド**（拡張子は OS に合わせる）
> 2. ``** / **``** or **`` を付けて実行
> 3. 引数が無いと HTTP サーバーになる

### ビルド & 実行（Unix/macOS）

```bash
# プロジェクト直下
# 実行ファイルを作る
go build -o days_calculator ./app

# 7 日前
./days_calculator -days 7
./days_calculator 7    # 位置引数でも可
```

### ビルド & 実行（Windows PowerShell）

```powershell
# プロジェクト直下
# .exe 拡張子が必要
go build -o days_calculator.exe ./app

# 7 日前
.\days_calculator.exe -days 7

# ビルドせずワンショット
# (Unix でも同様)
go run ./app 7
```

| オプション     | 動作                                   |
| --------- | ------------------------------------ |
| `-days N` | **N 日前** を表示                         |
| 位置引数 `N`  | 同上。`-days` と同等                       |
| （引数なし）    | HTTP サーバーを起動 (`PORT` 環境変数または `.env`) |
| `-h`      | ヘルプを表示                               |

---

## 2️⃣ HTTP API を使う

### Go バイナリで起動

```bash
# 引数なしで実行 → サーバーモード
./days_calculator
# => Server started at http://localhost:8080
```

### Docker で起動

```bash
docker build -t days-calculator .
docker run --rm --env-file .env -p 8089:8089 days-calculator
# => Server started at http://localhost:8089
```

`.env` の例:

```env
PORT=8089
```

### エンドポイント

| メソッド | パス               | クエリ         | 説明              |
| ---- | ---------------- | ----------- | --------------- |
| GET  | `/api/calculate` | `days` (整数) | **N 日前** の日付を返す |

#### リクエスト例

```bash
curl "http://localhost:8089/api/calculate?days=2"
```

#### レスポンス例

```json
{"date":"2024/12/21"}
```

---

## 3️⃣ 開発 & テスト

```bash
# 依存関係取得
go mod download

# テストがある場合
go test ./...
```

---

## ライセンス

MIT License  (詳細は `LICENSE` ファイル参照)

---

