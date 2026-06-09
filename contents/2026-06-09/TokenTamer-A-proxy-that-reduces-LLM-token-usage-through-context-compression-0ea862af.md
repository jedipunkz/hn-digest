---
source: "https://github.com/borhen68/TokenTamer"
hn_url: "https://news.ycombinator.com/item?id=48458633"
title: "TokenTamer A proxy that reduces LLM token usage through context compression"
article_title: "GitHub - borhen68/TokenTamer: A drop-in proxy that compresses bloated code context in real-time, cutting LLM API costs by 50–80% without losing what the model actually needs to know. · GitHub"
author: "borhensaidi"
captured_at: "2026-06-09T10:18:40Z"
capture_tool: "hn-digest"
hn_id: 48458633
score: 1
comments: 1
posted_at: "2026-06-09T09:17:40Z"
tags:
  - hacker-news
  - translated
---

# TokenTamer A proxy that reduces LLM token usage through context compression

- HN: [48458633](https://news.ycombinator.com/item?id=48458633)
- Source: [github.com](https://github.com/borhen68/TokenTamer)
- Score: 1
- Comments: 1
- Posted: 2026-06-09T09:17:40Z

## Translation

タイトル: TokenTamer コンテキスト圧縮を通じて LLM トークンの使用量を削減するプロキシ
記事のタイトル: GitHub - borhen68/TokenTamer: 肥大化したコード コンテキストをリアルタイムで圧縮するドロップイン プロキシ。モデルが実際に知る必要があるものを失うことなく、LLM API コストを 50 ～ 80% 削減します。 · GitHub
説明: 肥大化したコード コンテキストをリアルタイムで圧縮するドロップイン プロキシ。モデルが実際に知る必要があるものを失うことなく、LLM API コストを 50 ～ 80% 削減します。 - borhen68/トークンテイマー

記事本文:
GitHub - borhen68/TokenTamer: 肥大化したコード コンテキストをリアルタイムで圧縮するドロップイン プロキシ。モデルが実際に知る必要があるものを失うことなく、LLM API コストを 50 ～ 80% 削減します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ボーヘン68
/
トークンテイマー
公共
通知
あなた

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .github/ workflows .github/ workflows テスト テスト token_guard.egg-info token_guard.egg-info token_tamer token_tamer token_tamer_core token_tamer_core .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md CONTRIBUTING.mdライセンス ライセンス README.md README.md config.yaml config.yaml pyproject.toml pyproject.toml run_assembler.py run_assembler.py run_core.py run_core.py sys.md sys.md test_smoke.py test_smoke.py token.md token.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
肥大化したコード コンテキストをリアルタイムで圧縮するドロップイン プロキシにより、プレーン チャット コーディング エージェントで LLM API コストを 50 ～ 80% 削減します。
TokenTamer は、AI コーディング エージェントと LLM API の間に位置するミドルウェア プロキシです。生のペイロードをインターセプトし、AST でコードを解析し、「バックグラウンド」ファイルを構造スケルトンに置き換えます。エージェントは引き続きシグネチャ、クラス、インポートを参照できます。編集していない関数本体に対する支払いが停止されるだけです。
⚠️ アルファ版ソフトウェア。これは、洗練された SaaS ではなく、現在開発中の実際のプロジェクトです。インストールする前に、以下のサポート マトリックスをお読みください。
クライアント
HTTPS インターセプト
圧縮がアクティブです
注意事項
Aider ( --openai-api-base )
✅ 必要ありません
✅ フル
最高にサポートされています。プロキシ URL を直接使用します。
カーソル (カスタムベース URL)
✅ 必要ありません
✅ フル
最高にサポートされています。
プレーンなカール / SDK 呼び出し
✅ 必要ありません
✅ フル
テストに最適です。
クロード コード (ハードコードされたエンドポイント)
✅ 作品
✅ ツール認識
tool_result で読み取られた古いファイルはスケルトン化されます。最新の読み取りはそのまま残ります。
Codex CLI (ハードコードされたエンドポイント)
✅ 作品
✅ ツール認識
/v1/responses 経由の同じエンジン。
ツール対応圧縮の仕組み。 Clのようなエージェント

aude コードは Read(file) を繰り返し呼び出します。会話では、複数回ダンプされた同じファイルが蓄積されます。 TokenTamer は、tool_use → ファイルのマッピングをすべて追跡し、各ファイルの最新の読み取りを 100% そのまま維持しながら、古い tool_result の読み取りをスケルトン化します。 tool_use ブロックとツール定義には決して触れません。
何かが壊れた場合は、キル スイッチを押してください。
token-tamer --ssl --port 443 --passthrough # すべての圧縮を無効にする
# または
token-tamer --ssl --port 443 --no-tool-compression # ツール対応パスのみを無効にする
🚨 既知の制限事項
圧縮は再読み取りに依存します。単一読み取りセッションではツールの節約は行われません (テキストの圧縮のみ)。エージェントがファイルを再読み取りする長時間のセッションでは、最もメリットが得られます。
ヒューリスティックなファイル検出。ツール入力で file_path / path / filename キーを探します。珍しいスキーマを持つエキゾチックなエージェントは見逃される可能性があります。
マルチターン クロスリクエスト キャッシュはまだ実装されていません。
macOS は 1 行の証明書セットアップのみ。 Linux/Windows ユーザーは、CA を手動で信頼する必要があります。
本番環境のベンチマークはまだありません。節約の数値は、実際の長いクロード コード セッションではなく、合成ペイロードを使用した単体テストから得られます。
v0.2 — ツール対応圧縮 (✅ 出荷時)
v0.3 — マルチターン/クロスリクエストキャッシュにより、繰り返しコンテンツが再送信されない
v0.4 — 適切な多言語 AST のためのツリーシッター (現在の C スタイルのサポートは中括弧バランスのヒューリスティックです)
v0.5 — ファイルごとの圧縮ヒートマップを備えた Web ダッシュボード
🔌 ドロップイン プロキシ — コーディング エージェントを変更する必要はありません。 API のベース URL を変更するだけです。
🧠 スマートなアクティブ ファイル検出 — 作業中のファイルを自動的に識別し、それらを 100% そのまま残します。
🌳 AST ベースの圧縮 — シグネチャ、インポート、クラス構造を維持しながら関数本体を削除します。
💰 リアルタイムのコスト追跡 — 美しいターミナルダッシュボードの表示

トークンもお金も節約できます。
🔄 完全なストリーミング サポート — OpenAI と Anthropic API の両方に対する透過的な SSE ストリーミング。
⚡ 遅延オーバーヘッドゼロ — 圧縮はミリ秒単位でローカルに行われます。
Python 3.9 以降 ( python3 --version )
macOS、Linux、または Windows (Windows = 手動証明書信頼ステップ)
openssl (macOS およびほとんどの Linux にプリインストールされています)
git clone https://github.com/borhen68/TokenTamer.git
cd トークンテイマー
# 推奨: システム Python の操作を避けるために仮想環境を使用する
python3 -m venv venv
ソース venv/bin/activate # Windows: venv\Scripts\activate
pip install -e 。
インストールされていることを確認します。
トークンテイマー --version
# → TokenTamer 0.2.0
2. パスを選択してください
👉 パス A — Aider、Cursor、または独自の SDK コード (SSL セットアップは必要ありません):
token-tamer --port 8000 --no-dashboard
次に、ツールの API ベース URL を http://127.0.0.1:8000/v1 に指定します。
aider --openai-api-base http://127.0.0.1:8000/v1
カーソルの場合: [設定] → [モデル] → [OpenAI API ベース] → http://127.0.0.1:8000/v1 。終わり。 ✅
👉 パス B — Claude Code または Codex CLI (SSL セットアップ、1 回限り):
これらのツールは API URL をハードコードします。 HTTPS インターセプトを使用します。
# ステップ 1 — ローカル証明書を生成します (実行して終了するだけです)
token-tamer --ssl --port 8443 --no-dashboard &
スリープ 2 && %1 をキル
# ステップ 2 — 証明書を信頼する (macOS)
sudo security add-trusted-cert -d -r trustRoot \
-k /ライブラリ/キーチェーン/System.keychain \
~ /.config/token-tamer/certs/ca-cert.pem
# ステップ 3 — API ドメインをローカルホストにリダイレクトする
echo " 127.0.0.1 api.openai.com " | sudo tee -a /etc/hosts
エコー「127.0.0.1 api.anthropic.com」 | sudo tee -a /etc/hosts
# ステップ 4 — ポート 443 で TokenTamer を実行します (ポートが低い場合は sudo が必要です)
sudo $( what token-tamer ) --ssl --port 443 --no-dashboard
そのターミナルを開いたままにし、新しいターミナルで次のようにします。
クロード「蛇を作りますが」

私「＃または
コーデックス「このモジュールをリファクタリング」
これで、インターセプトと圧縮が行われます。 🎉
# パス A のチェック:
カール http://127.0.0.1:8000/health
# パス B のチェック:
curl https://api.openai.com/health # OpenAI ではなく、TokenTamer の JSON を返す必要があります
どちらも次を返すはずです:
{ "ステータス" : " ok " 、 "バージョン" : " 0.2.0 " 、 "requests_processed" : 0 、 "tokens_saved" : 0 }
4.クリーンアップ（アンインストール）
# /etc/hosts エントリを削除します
sudo sed -i.bak ' /api.openai.com/d;/api.anthropic.com/d ' /etc/hosts
# 証明書を信頼しない
sudo security remove-trusted-cert -d ~ /.config/token-tamer/certs/ca-cert.pem
# パッケージをアンインストールする
pip アンインストール token-tamer
🆘 トラブルシューティング
症状
修正
コマンドが見つかりません: token-tamer
venv をアクティブ化します:source venv/bin/activate
ModuleNotFoundError: 「uvicorn」という名前のモジュールがありません
同じ — venv がアクティブではありません
アドレスはポート 8000 ですでに使用されています
lsof -ti :8000 | xargs キル -9
ポート 443 で許可が拒否されました
1024 未満のポートには sudo を使用するか、より高いポートを選択してください
CURL からの SSL 証明書の問題
security add-trusted-cert ステップを再実行し、新しいターミナルを開きます
クロード コードのハング/エラー
Kill スイッチを押します: --passthrough で再起動します
圧縮で何かが壊れた
--no-tool-compression を指定して再起動し、問題を報告します
APIキー
TokenTamer は、次の優先順位で API キーを解決します。
リクエストヘッダー — エージェントによって送信されたキー (デフォルトの動作、設定は必要ありません)
環境変数 — OPENAI_API_KEY / ANTHROPIC_API_KEY
エージェント TokenTamer LLM API
│ │ │
│── 100,000 トークンペイロード ──────▶│ │
│ │── アクティブなファイルを特定する │
│ │── 背景をスケルトン化 │
│ │── 15,000 トークンペイロード ─ ││
│ │ │
│ │◀──── ストリーミング応答 ────│
│◀── ストリーミング応答 ─────│ │
│ │ │
│ │

── ダッシュボード: 2.45 ドル節約されました! 💰 │
以前 (トークンコストが重い):
def Calculate_tax (金額: float 、領域: str) -> float :
"""複雑なロジックに基づいて地方税率を計算します。"""
rate = get_base_rate (地域)
調整 = fetch_adjustments (領域、金額)
量 > しきい値の場合:
率 *= 1.05
# ... さらに 50 行の複雑な数学 ...
Final_tax を返す
後 (軽量スケルトン):
# [TOKEN-GUARD: 圧縮 — 構造スケルトンのみ]
def Calculate_tax (金額: float 、領域: str) -> float: ...
LLM は、calculate_tax の存在とその呼び出し方法をまだ認識していますが、実装を読み取るトークンを無駄にしません。
作業ディレクトリに config.yaml を作成します。
代理人：
ホスト：「127.0.0.1」
ポート: 8000
上流:
openai_url : " https://api.openai.com "
anthropic_url : " https://api.anthropic.com "
コンテキスト:
repo_path : " /path/to/your/codebase " # セマンティックなアクティブ ファイル検出を有効にします
スケルトン化者:
keep_docstrings : false # 関数 docstrings を保持しますか?
keep_class_attrs : true # クラスレベルの属性を保持しますか?
価格設定: コスト見積もりのための 100 万トークンあたりの数
gpt-4o :
入力：2.50
出力：10.00
クロード・ソネット-4-20250514 :
入力：3.00
出力：15.00
🌐 多言語サポート
TokenTamer は Python だけではなくスケルトン化します。
🧠 セマンティックアクティブファイル検出
config.yaml で repo_path を指定し、 Sentence-transformers をインストールすると、
TokenTamer は埋め込みを使用して、どのファイルが意味的に関連しているかを検出します。
名前を挙げなくても、クエリを実行できます。
pip インストール文-transformers scikit-learn
🛠 サポートされている API
プロバイダー
エンドポイント
ステータス
OpenAI
/v1/チャット/コンプリート
✅
OpenAI
/v1/completions
✅ (パススルー)
OpenAI
/v1/モデル
✅ (パススルー)
人間的
/v1/メッセージ
✅
🧪 テスト
pip install -e " .[dev] "
pytest テスト/ -v
📋 要件

s
依存関係: FastAPI、uvicorn、httpx、tiktoken、rich、pyyaml
肥大化したコード コンテキストをリアルタイムで圧縮するドロップイン プロキシ。モデルが実際に認識する必要があるものを失うことなく、LLM API コストを 50 ～ 80% 削減します。
読み込み中にエラーが発生しました。このページをリロードしてください。
3
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A drop-in proxy that compresses bloated code context in real-time, cutting LLM API costs by 50–80% without losing what the model actually needs to know. - borhen68/TokenTamer

GitHub - borhen68/TokenTamer: A drop-in proxy that compresses bloated code context in real-time, cutting LLM API costs by 50–80% without losing what the model actually needs to know. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
borhen68
/
TokenTamer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .github/ workflows .github/ workflows tests tests token_guard.egg-info token_guard.egg-info token_tamer token_tamer token_tamer_core token_tamer_core .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md config.yaml config.yaml pyproject.toml pyproject.toml run_assembler.py run_assembler.py run_core.py run_core.py sys.md sys.md test_smoke.py test_smoke.py token.md token.md View all files Repository files navigation
A drop-in proxy that compresses bloated code context in real-time, cutting LLM API costs by 50–80% on plain-chat coding agents.
TokenTamer is a middleware proxy that sits between an AI coding agent and the LLM API. It intercepts raw payloads, parses code with AST, and replaces "background" files with structural skeletons. The agent still sees signatures, classes, and imports — it just stops paying for function bodies it isn't editing.
⚠️ Alpha software. This is a real project in active development, not a polished SaaS. Please read the support matrix below before installing.
Client
HTTPS interception
Compression active
Notes
Aider ( --openai-api-base )
✅ Not needed
✅ Full
Best supported. Use the proxy URL directly.
Cursor (custom base URL)
✅ Not needed
✅ Full
Best supported.
Plain curl / SDK calls
✅ Not needed
✅ Full
Great for testing.
Claude Code (hardcoded endpoint)
✅ Works
✅ Tool-aware
Stale file reads in tool_result get skeletonized; latest read stays intact.
Codex CLI (hardcoded endpoint)
✅ Works
✅ Tool-aware
Same engine via /v1/responses .
How tool-aware compression works. Agents like Claude Code call Read(file) repeatedly. The conversation accumulates the same file dumped multiple times. TokenTamer tracks every tool_use → file mapping, then skeletonizes the older tool_result reads while keeping the most recent read of each file 100% intact. tool_use blocks and tool definitions are never touched.
If something ever breaks, hit the kill switch:
token-tamer --ssl --port 443 --passthrough # disable all compression
# or
token-tamer --ssl --port 443 --no-tool-compression # disable only tool-aware path
🚨 Known Limitations
Compression depends on re-reads. Single-read sessions get no tool savings (just text compression). Long sessions where the agent re-reads files benefit the most.
Heuristic file detection. We look for file_path / path / filename keys in tool inputs. Exotic agents with unusual schemas may be missed.
Multi-turn cross-request caching is not yet implemented.
macOS only for the one-line cert setup. Linux/Windows users need to trust the CA manually.
No production benchmarks yet. Savings numbers come from unit tests with synthetic payloads, not real long Claude Code sessions.
v0.2 — Tool-aware compression (✅ shipped)
v0.3 — Multi-turn / cross-request cache so repeat content isn't re-sent
v0.4 — Tree-sitter for proper multi-language AST (current C-style support is a brace-balance heuristic)
v0.5 — Web dashboard with per-file compression heatmap
🔌 Drop-in proxy — No changes needed to your coding agent. Just change the API base URL.
🧠 Smart active file detection — Automatically identifies which files you're working on and leaves them 100% intact.
🌳 AST-based compression — Strips function bodies while preserving signatures, imports, and class structures.
💰 Real-time cost tracking — Beautiful terminal dashboard showing tokens saved and money saved.
🔄 Full streaming support — Transparent SSE streaming for both OpenAI and Anthropic APIs.
⚡ Zero latency overhead — Compression happens locally in milliseconds.
Python 3.9 or newer ( python3 --version )
macOS, Linux, or Windows (Windows = manual cert trust step)
openssl (pre-installed on macOS & most Linux)
git clone https://github.com/borhen68/TokenTamer.git
cd TokenTamer
# Recommended: use a virtual environment to avoid messing with system Python
python3 -m venv venv
source venv/bin/activate # Windows: venv\Scripts\activate
pip install -e .
Verify it installed:
token-tamer --version
# → TokenTamer 0.2.0
2. Choose Your Path
👉 Path A — Aider, Cursor, or your own SDK code (no SSL setup needed):
token-tamer --port 8000 --no-dashboard
Then point your tool's API base URL at http://127.0.0.1:8000/v1 :
aider --openai-api-base http://127.0.0.1:8000/v1
For Cursor: Settings → Models → OpenAI API Base → http://127.0.0.1:8000/v1 . Done. ✅
👉 Path B — Claude Code or Codex CLI (SSL setup, one-time):
These tools hardcode the API URL. We use HTTPS interception:
# Step 1 — Generate the local certificate (just runs and exits)
token-tamer --ssl --port 8443 --no-dashboard &
sleep 2 && kill %1
# Step 2 — Trust the certificate (macOS)
sudo security add-trusted-cert -d -r trustRoot \
-k /Library/Keychains/System.keychain \
~ /.config/token-tamer/certs/ca-cert.pem
# Step 3 — Redirect API domains to localhost
echo " 127.0.0.1 api.openai.com " | sudo tee -a /etc/hosts
echo " 127.0.0.1 api.anthropic.com " | sudo tee -a /etc/hosts
# Step 4 — Run TokenTamer on port 443 (sudo required for low ports)
sudo $( which token-tamer ) --ssl --port 443 --no-dashboard
Leave that terminal open, then in a new terminal :
claude " create a snake game " # or
codex " refactor this module "
You're now intercepting + compressing. 🎉
# Path A check:
curl http://127.0.0.1:8000/health
# Path B check:
curl https://api.openai.com/health # Should return TokenTamer's JSON, not OpenAI's
Both should return:
{ "status" : " ok " , "version" : " 0.2.0 " , "requests_processed" : 0 , "tokens_saved" : 0 }
4. Cleanup (Uninstall)
# Remove /etc/hosts entries
sudo sed -i.bak ' /api.openai.com/d;/api.anthropic.com/d ' /etc/hosts
# Untrust the cert
sudo security remove-trusted-cert -d ~ /.config/token-tamer/certs/ca-cert.pem
# Uninstall the package
pip uninstall token-tamer
🆘 Troubleshooting
Symptom
Fix
command not found: token-tamer
Activate your venv: source venv/bin/activate
ModuleNotFoundError: No module named 'uvicorn'
Same — venv not active
address already in use on port 8000
lsof -ti :8000 | xargs kill -9
Permission denied on port 443
Use sudo for ports <1024, or pick a higher port
SSL certificate problem from curl
Re-run the security add-trusted-cert step, then open a NEW terminal
Claude Code hangs / errors
Hit the kill switch: restart with --passthrough
Compression broke something
Restart with --no-tool-compression and file an issue
API Keys
TokenTamer resolves API keys in this priority order:
Request headers — Keys sent by your agent (default behavior, zero config needed)
Environment variables — OPENAI_API_KEY / ANTHROPIC_API_KEY
Your Agent TokenTamer LLM API
│ │ │
│── 100k token payload ──────▶│ │
│ │── Identify active files │
│ │── Skeletonize background │
│ │── 15k token payload ────────▶│
│ │ │
│ │◀──── Streaming response ─────│
│◀── Streaming response ──────│ │
│ │ │
│ │── Dashboard: saved $2.45! 💰 │
Before (Heavy Token Cost):
def calculate_tax ( amount : float , region : str ) -> float :
"""Calculates regional tax rates based on complex logic."""
rate = get_base_rate ( region )
adjustments = fetch_adjustments ( region , amount )
if amount > THRESHOLD :
rate *= 1.05
# ... 50 more lines of complex math ...
return final_tax
After (Lightweight Skeleton):
# [TOKEN-GUARD: Compressed — structural skeleton only]
def calculate_tax ( amount : float , region : str ) -> float : ...
The LLM still knows calculate_tax exists and how to call it, but doesn't waste tokens reading the implementation.
Create a config.yaml in your working directory:
proxy :
host : " 127.0.0.1 "
port : 8000
upstream :
openai_url : " https://api.openai.com "
anthropic_url : " https://api.anthropic.com "
context :
repo_path : " /path/to/your/codebase " # Enables semantic active-file detection
skeletonizer :
keep_docstrings : false # Preserve function docstrings?
keep_class_attrs : true # Keep class-level attributes?
pricing : # Per 1M tokens for cost estimation
gpt-4o :
input : 2.50
output : 10.00
claude-sonnet-4-20250514 :
input : 3.00
output : 15.00
🌐 Multi-Language Support
TokenTamer skeletonizes more than just Python:
🧠 Semantic Active-File Detection
If you provide a repo_path in config.yaml and install sentence-transformers ,
TokenTamer uses embeddings to detect which files are semantically relevant to your
query — even if you don't mention them by name.
pip install sentence-transformers scikit-learn
🛠 Supported APIs
Provider
Endpoint
Status
OpenAI
/v1/chat/completions
✅
OpenAI
/v1/completions
✅ (pass-through)
OpenAI
/v1/models
✅ (pass-through)
Anthropic
/v1/messages
✅
🧪 Testing
pip install -e " .[dev] "
pytest tests/ -v
📋 Requirements
Dependencies: FastAPI, uvicorn, httpx, tiktoken, rich, pyyaml
A drop-in proxy that compresses bloated code context in real-time, cutting LLM API costs by 50–80% without losing what the model actually needs to know.
There was an error while loading. Please reload this page .
3
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
