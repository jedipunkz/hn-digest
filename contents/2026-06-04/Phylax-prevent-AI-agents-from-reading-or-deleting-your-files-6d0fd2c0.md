---
source: "https://phylaxx.pages.dev/"
hn_url: "https://news.ycombinator.com/item?id=48394498"
title: "Phylax prevent AI agents from reading or deleting your files"
article_title: "Phylax - OS-level Protection for AI Coding Agents"
author: "usertheinfo"
captured_at: "2026-06-04T07:45:52Z"
capture_tool: "hn-digest"
hn_id: 48394498
score: 1
comments: 1
posted_at: "2026-06-04T05:45:42Z"
tags:
  - hacker-news
  - translated
---

# Phylax prevent AI agents from reading or deleting your files

- HN: [48394498](https://news.ycombinator.com/item?id=48394498)
- Source: [phylaxx.pages.dev](https://phylaxx.pages.dev/)
- Score: 1
- Comments: 1
- Posted: 2026-06-04T05:45:42Z

## Translation

タイトル: Phylax は AI エージェントによるファイルの読み取りまたは削除を防止します
記事のタイトル: Phylax - AI コーディング エージェントの OS レベルの保護
説明: Phylax は、AI エージェントが秘密に触れる前に、危険なファイルへのアクセスをブロックします。 100% ローカル、アカウント、クラウド、テレメトリなし。 Windows ACL の強制。

記事本文:
Phylax Home Docs 開発パス AI コーディング エージェントの Windows セキュリティ層
Phylax は、AI エージェントがプライベート ファイルにアクセスする前に停止します。
実際の OS レベルの保護。カーネルは ACCESS_DENIED を返し、エージェントは 1 バイトも認識しません。
100% ローカル · アカウントなし · クラウドなし · テレメトリなし
問題は現実です。解決策はローカルにあります。
Claude Code、Cursor、OpenCode などの AI エージェントは、ファイル システムに完全にアクセスできます。彼らは何でも読み取り、書き込み、削除できます。
Phylax は、それらと秘密の間に実際の OS レベルの境界を設けます。プロキシもラッパーもありません。カーネルがそれを強制します。
アカウントもクラウドもテレメトリもありません。すべてがマシン上に残ります。ローカル SQLite の監査ログ。完全にオフラインで動作します。
Claude、Cursor、OpenCode、Copilot、Windsurf、Aider などを認識します。プロセス名、環境変数、および子の継承によってエージェントを検出します。
実際の Windows ACL (DENY ACE + Mandatory Integrity Control) を適用します。カーネル自体は ACCESS_DENIED を返します。エージェントはファイルにアクセスしません。
Phylax は、拒否されたすべてのファイルに 3 層の Windows セキュリティを適用します。読み取り/書き込み/削除の DENY ACE、ACL 変更の WRITE_DAC 保護、特権バイパスを停止する必須整合性制御です。
クラウド プロキシ、API キー、ネットワークは必要ありません。すべてがマシン上でローカルに実行されます。
イメージ名、環境変数、コマンドライン検査によって AI エージェント プロセスを識別します。子プロセスはエージェント ラベルを自動的に継承します。
ファイル パスと操作に対して phylax.toml ルールをチェックします。常に否定が勝ちます。優先順位が付けられたバケットにより、すべてのアクセス試行が解決されます。
実際の Windows ACL を適用します。エージェントが 1 バイトにアクセスする前に、カーネルは ACCESS_DENIED を返します。ユーザー空間のトリックを回避することはできません。
AI エージェントが保護されたファイルにアクセスしようとすると、これが起こります。
三層

Windows セキュリティの特徴: DENY ACE はファイル アクセスをブロックし、WRITE_DAC は ACL の変更を防止し、Mandatory Integrity Control は特権バイパスを停止します。カーネルは ACCESS_DENIED を返し、エージェントは 1 バイトも認識しません。
Phylax は、優先順位に従って 6 つの権限バケットを使用します。常に否定が勝ちます。プリセットから始めて、 phylax.toml 経由でカスタマイズします。
保守的なデフォルト
一致するルールがない場合: 読み取り = 許可、書き込み = 要求、削除 = 拒否。
機密ファイルと重要なファイルを保護します。ソース編集は高速です。ロックファイルを変更すると確認が求められます。
.env、.pem、.key をブロックします。 src/** および testing/** を許可します。移行とロックファイルのプロンプトが表示されます。
[プロジェクト]
名前 = "私の-phylax-プロジェクト"
デフォルト = "保守的"
[拒否]
files = [".env", ".env.*", "secrets/**", "*.pem", "*.key", "phylax.toml"]
[尋ねる]
files = ["Cargo.lock"、"package-lock.json"、"migrations/**"]
[書く]
ファイル = ["src/**"、"tests/**"、"docs/**"]
[読む]
files = ["README.md", "docs/**"] 最大制御 厳格
最大限のセキュリティ。すべてのソース編集とロックファイル変更には明示的な承認が必要です。
.env、.pem、.key、.p12、.pfx、secrets/** を拒否します。すべてのソース編集を要求します。デフォルトでは読み取り専用です。
[プロジェクト]
名前 = "phylax-strict"
デフォルト = "保守的"
[拒否]
files = [".env", ".env.*", "secrets/**", "keys/**", "*.pem", "*.key", "*.p12", "phylax.toml"]
[尋ねる]
files = ["src/**"、"tests/**"、"Cargo.lock"、"package-lock.json"、"migrations/**"]
[読む]
files = ["README.md", "docs/**", "src/**", "tests/**"] 低摩擦 高速かつ柔軟
エージェントが自由に編集できるようにします。シークレットとマニフェストのみが保護されます。
.env、.pem、.key、phylax.toml をブロックします。それ以外はすべて書き込み可能です。通常の編集ではプロンプトは表示されません。
[プロジェクト]
名前 = "フィラックスファスト"
デフォルト = "保守的"
[拒否]
files = [".env", ".env.*", "secrets/**", "*.pem", "*.key", "phylax.toml"]
[書く]
ファイル = ["src/**", "テスト

/**"、"docs/**"、"examples/**"、"Cargo.lock"、"package-lock.json"]
[読む]
files = ["README.md", "docs/**", "src/**", "tests/**", "examples/**"] 05 インストール
コマンドは 1 つです。ゼロ構成。
コマンドは 1 つです。アカウントもクラウドもテレメトリもありません。デーモンはバックグラウンドで目に見えずに実行されます。

## Original Extract

Phylax blocks dangerous file access before AI agents touch your secrets. 100% local, no accounts, no cloud, no telemetry. Windows ACL enforcement.

Phylax Home Docs Development Path Windows security layer for AI coding agents
Phylax stops AI agents before they touch your private files.
Real OS-level protection. The kernel returns ACCESS_DENIED, the agent never sees a single byte.
100% local · No accounts · No cloud · No telemetry
The problem is real. The solution is local.
AI agents like Claude Code, Cursor, and OpenCode have full filesystem access . They can read, write, or delete anything.
Phylax puts a real OS-level boundary between them and your secrets. No proxy, no wrapper. The kernel enforces it.
No account, no cloud, no telemetry. Everything stays on your machine. Audit logs in local SQLite. Works fully offline.
Recognizes Claude, Cursor, OpenCode, Copilot, Windsurf, Aider, and more. Detects agents by process name, environment variables, and child inheritance.
Applies real Windows ACLs (DENY ACEs + Mandatory Integrity Control). The kernel itself returns ACCESS_DENIED - the agent never touches the file.
Phylax applies three layers of Windows security to every denied file: DENY ACEs for read/write/delete, WRITE_DAC protection for ACL modification, and Mandatory Integrity Control to stop privilege bypass.
No cloud proxy, no API keys, no network required. Everything runs locally on your machine.
Identifies AI agent processes by image name, environment variables, and command-line inspection. Child processes inherit the agent label automatically.
Checks your phylax.toml rules against the file path and operation. Deny always wins. Priority-ordered buckets resolve every access attempt.
Applies real Windows ACLs. The kernel returns ACCESS_DENIED before the agent touches a single byte. No userspace trick can bypass it.
This is what happens when an AI agent tries to access a protected file.
Three layers of Windows security: DENY ACEs block file access, WRITE_DAC prevents ACL modification, and Mandatory Integrity Control stops privilege bypass. The kernel returns ACCESS_DENIED, the agent never sees a single byte.
Phylax uses six permission buckets ordered by priority. Deny always wins. Start with a preset, then customize via phylax.toml .
Conservative default
When no rule matches: read = Allow, write = Ask, delete = Deny.
Protects secrets and critical files. Source edits are fast. Lockfile changes ask for confirmation.
Blocks .env, .pem, .key. Allows src/** and tests/**. Prompts for migrations and lockfiles.
[project]
name = "my-phylax-project"
default = "conservative"
[deny]
files = [".env", ".env.*", "secrets/**", "*.pem", "*.key", "phylax.toml"]
[ask]
files = ["Cargo.lock", "package-lock.json", "migrations/**"]
[write]
files = ["src/**", "tests/**", "docs/**"]
[read]
files = ["README.md", "docs/**"] Maximum control Strict
Maximum security. Every source edit and lockfile change requires explicit approval.
Denies .env, .pem, .key, .p12, .pfx, secrets/**. Asks for every source edit. Read-only by default.
[project]
name = "phylax-strict"
default = "conservative"
[deny]
files = [".env", ".env.*", "secrets/**", "keys/**", "*.pem", "*.key", "*.p12", "phylax.toml"]
[ask]
files = ["src/**", "tests/**", "Cargo.lock", "package-lock.json", "migrations/**"]
[read]
files = ["README.md", "docs/**", "src/**", "tests/**"] Low friction Fast & Flexible
Lets agents edit freely. Only secrets and the manifest are protected.
Blocks .env, .pem, .key, phylax.toml. Everything else is writable. No prompts for normal edits.
[project]
name = "phylax-fast"
default = "conservative"
[deny]
files = [".env", ".env.*", "secrets/**", "*.pem", "*.key", "phylax.toml"]
[write]
files = ["src/**", "tests/**", "docs/**", "examples/**", "Cargo.lock", "package-lock.json"]
[read]
files = ["README.md", "docs/**", "src/**", "tests/**", "examples/**"] 05 Install
One command. Zero configuration.
One command. No accounts, no cloud, no telemetry. The daemon runs invisibly in the background.
