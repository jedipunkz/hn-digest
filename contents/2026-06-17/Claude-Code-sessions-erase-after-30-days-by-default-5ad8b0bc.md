---
source: "https://code.claude.com/docs/en/settings"
hn_url: "https://news.ycombinator.com/item?id=48571536"
title: "Claude Code sessions erase after 30 days by default"
article_title: "Claude Code settings - Claude Code Docs"
author: "markrogersjr"
captured_at: "2026-06-17T16:22:36Z"
capture_tool: "hn-digest"
hn_id: 48571536
score: 2
comments: 1
posted_at: "2026-06-17T15:05:48Z"
tags:
  - hacker-news
  - translated
---

# Claude Code sessions erase after 30 days by default

- HN: [48571536](https://news.ycombinator.com/item?id=48571536)
- Source: [code.claude.com](https://code.claude.com/docs/en/settings)
- Score: 2
- Comments: 1
- Posted: 2026-06-17T15:05:48Z

## Translation

タイトル: クロード コードのセッションはデフォルトで 30 日後に消去されます
記事のタイトル: Claude Code の設定 - Claude Code Docs
説明: グローバルおよびプロジェクトレベルの設定、および環境変数を使用してクロードコードを構成します。

記事本文:
Claude Code の設定 - Claude Code Docs Documentation Index
/docs/llms.txt で完全なドキュメントのインデックスを取得します。
さらに探索する前に、このファイルを使用して利用可能なすべてのページを検出します。
メイン コンテンツにスキップ Claude Code Docs ホーム ページ 英語 検索... ⌘ K Ask Assistant Claude 開発者プラットフォーム
検索... ナビゲーション 設定と権限 クロード コードの設定 はじめに クロード コードで構築 管理 構成リファレンス エージェント SDK 新機能 リソース 設定と権限
高速モードで応答を高速化
アドバイザ ツールを使用して難しい意思決定をエスカレーションする
構成スコープ 使用可能なスコープ
設定ファイル 編集が有効になるタイミング
管理設定の無効なエントリ
ポリシーヘルパーを使用して管理設定を計算する
構成システムの重要なポイント
プラグイン設定 プラグイン設定
ページをコピー グローバルおよびプロジェクト レベルの設定、および環境変数を使用してクロード コードを構成します。
ページをコピーする Claude Code には、ニーズに合わせて動作を構成するためのさまざまな設定が用意されています。対話型 REPL の使用時に /config コマンドを実行すると、クロード コードを構成できます。これにより、タブ付きの設定インターフェイスが開き、ステータス情報を表示したり、構成オプションを変更したりできます。
設定範囲
組織全体で適用する必要があるセキュリティ ポリシー
上書きできないコンプライアンス要件
IT/DevOps によって導入された標準化された構成
どこにでも必要な個人的な設定 (テーマ、エディタ設定)
すべてのプロジェクトで使用するツールとプラグイン
API キーと認証 (安全に保管)
チーム共有設定 (権限、フック、MCP サーバー)
チーム全体が持つべきプラグイン
共同作業者間でのツールの標準化
特定のプロジェクトの個人的なオーバーライド
と共有する前に構成をテストする

チーム
他のマシンでは機能しないマシン固有の設定
マネージド (最高) - 何も上書きできません
コマンドライン引数 - 一時的なセッションオーバーライド
ローカル - プロジェクトとユーザー設定を上書きします
プロジェクト - ユーザー設定を上書きします
ユーザー (最低) - 他に何も設定が指定されていない場合に適用されます
ユーザー設定は ~/.claude/settings.json で定義され、すべてに適用されます
プロジェクト。
プロジェクト設定はプロジェクト ディレクトリに保存されます。
ソース管理にチェックインされ、チームと共有される設定の .claude/settings.json
チェックインされていない設定の .claude/settings.local.json。個人的な好みや実験に役立ちます。 Claude Code が .claude/settings.local.json を作成するとき、ファイルを無視するように git を構成します。ファイルを自分で作成した場合は、手動で gitignore に追加します。
管理設定: 一元管理が必要な組織のために、Claude Code は管理設定の複数の配信メカニズムをサポートしています。すべて同じ JSON 形式を使用し、ユーザーまたはプロジェクトの設定によって上書きすることはできません。
サーバー管理の設定 : Claude.ai 管理コンソールを介して Anthropic のサーバーから配信されます。 「サーバー管理の設定」 を参照してください。
MDM/OS レベルのポリシー: macOS および Windows のネイティブ デバイス管理を通じて配信されます。
macOS: com.anthropic.claudecode 管理対象環境設定ドメイン。 plist のトップレベルのキーは、マネージド設定.json をミラーリングし、ネストされた設定を辞書として、配列を plist 配列として表示します。 Jamf、Iru (Kandji)、または同様の MDM ツールの構成プロファイルを介してデプロイします。
Windows: JSON (グループ ポリシーまたは Intune 経由で展開) を含む設定値 (REG_SZ または REG_EXPAND_SZ) を持つ HKLM\SOFTWARE\Policies\ClaudeCode レジストリ キー
Windows (ユーザーレベル): HKCU\SOFTWARE\Policies\ClaudeCode (ポリシーの優先順位が最も低く、管理者レベルのソースがない場合にのみ使用されます)

存在します）
ファイルベース: システム ディレクトリにデプロイされた管理対象設定.json および管理対象 mcp.json:
macOS: /ライブラリ/アプリケーション サポート/ClaudeCode/
Linux および WSL: /etc/claude-code/
Windows: C:\Program Files\ClaudeCode\
従来の Windows パス C:\ProgramData\ClaudeCode\managed-settings.json は、v2.1.75 以降サポートされなくなりました。その場所に設定を展開した管理者は、ファイルを C:\Program Files\ClaudeCode\managed-settings.json に移行する必要があります。
ファイルベースの管理設定は、 manage-settings.json と並んで同じシステム ディレクトリにある manage-settings.d/ のドロップイン ディレクトリもサポートします。これにより、単一のファイルに対する編集を調整することなく、別々のチームが独立したポリシー フラグメントを展開できるようになります。
systemd の規則に従い、managed-settings.json が最初にベースとしてマージされ、次にドロップイン ディレクトリ内のすべての *.json ファイルがアルファベット順に並べ替えられ、一番上にマージされます。スカラー値に関しては、後のファイルが以前のファイルをオーバーライドします。配列は連結され、重複が排除されます。オブジェクトは深く結合されます。で始まる隠しファイル。無視されます。
数値プレフィックスを使用してマージ順序を制御します (例: 10-telemetry.json および 20-security.json )。
管理された展開では、次を使用してプラグイン マーケットプレイスの追加を制限することもできます。
strictKnownMarketplaces 。詳細については、「マネージド マーケットプレイスの制限事項」を参照してください。
他の設定は ~/.claude.json に保存されます。このファイルには、OAuth セッション、ユーザーおよびローカル スコープの MCP サーバー構成、プロジェクトごとの状態 (許可されたツール、信頼設定)、およびさまざまなキャッシュが含まれています。プロジェクト スコープの MCP サーバーは、 .mcp.json に個別に保存されます。
Claude Code は、構成ファイルのタイムスタンプ付きバックアップを自動的に作成し、データ損失を防ぐために最新の 5 つのバックアップを保持します。
settings.json の例 {
"$schema" : "https://json.schemastore.org/claude-code-s

設定.json" ,
「権限」: {
"許可" : [
"Bash(npm run lint)" ,
"Bash(npm run test *)" ,
「読み取り(~/.zshrc)」
]、
「拒否」 : [
"Bash(curl *)" ,
"読み取り(./.env)" ,
"Read(./.env.*)" ,
「読み取り(./secrets/**)」
】
}、
"環境" : {
"CLAUDE_CODE_ENABLE_TELEMETRY" : "1" 、
"OTEL_METRICS_EXPORTER" : "otlp"
}、
「会社のお知らせ」 : [
「アクメ社へようこそ！ docs.acme.com でコード ガイドラインを確認してください。" ,
"リマインダー: すべての PR にはコード レビューが必要です" ,
「新しいセキュリティポリシーが発効中」
】
}
上の例の $schema 行は、クロード コード設定の公式 JSON スキーマを指します。これを settings.json に追加すると、VS Code、Cursor、および JSON スキーマ検証をサポートするその他のエディターでオートコンプリートとインライン検証が有効になります。
公開されたスキーマは定期的に更新され、最新の CLI リリースで追加された設定が含まれていない可能性があるため、最近文書化されたフィールドの検証警告は、必ずしも構成が無効であることを意味するわけではありません。
編集が有効になる時期
model : /model を使用してセッション中に切り替える
OutputStyle : システム プロンプトの一部。 /clear または再起動時に再構築されます。
管理設定の無効なエントリ
対話型セッションでは、起動時に無効なエントリをリストするダイアログが表示されます。
ヘッドレスは -p を指定して実行し、概要を標準エラー出力に出力します。
クロード・ドクターは、無効なエントリをそのソースとフィールドとともにリストします。
v2.1.119 より前のバージョンでは、theme、verbose、editorMode、autoCompactEnabled、preferredNotifChannel など、多数の /config 設定キーも settings.json ではなくここに保存されます。
キー 説明 例 autoConnectIde クロード コードが外部端末から起動されるときに、実行中の IDE に自動的に接続します。デフォルト: false 。 VS Code または JetBrains ターミナルの外で実行している場合、/config に IDE (外部ターミナル) への自動接続として表示されます。 CLAUDE_CODE_

true に設定すると、AUTO_CONNECT_IDE 環境変数がこれをオーバーライドします。 autoInstallIdeExtension VS Code ターミナルから実行するときに、Claude Code IDE 拡張機能を自動的にインストールします。デフォルト: true 。 VS Code または JetBrains ターミナル内で実行する場合、自動インストール IDE 拡張機能として /config に表示されます。 CLAUDE_CODE_IDE_SKIP_AUTO_INSTALL 環境変数を false に設定することもできます。 externalEditorContext Ctrl+G で外部エディタを開いたときに、クロードの以前の応答を # -コメント付きコンテキストとして先頭に追加します。デフォルト: false 。 /config に「外部エディターで最後の応答を表示」として表示されます。 true TeammateDefaultModel 生成プロンプトでチームメイトが指定されていない場合の、エージェント チームのチームメイトのデフォルト モデル。 "sonnet" などのモデル エイリアスに設定するか、リードの現在の /model 選択を継承する場合は null に設定します。 /config にデフォルトのチームメイト モデル「sonnet」として表示されます。
ワークツリーの設定
サンドボックスパスプレフィックス
filesystem.allowWrite 、 filesystem.denyWrite 、 filesystem.denyRead 、および filesystem.allowRead のパスは、次のプレフィックスをサポートしています。
プレフィックス 意味 例 / ファイルシステム ルートからの絶対パス /tmp/build は /tmp/build のままになります ~/ ホーム ディレクトリを基準に ~/.kube が $HOME/.kube ./ になるか、プレフィックスなし プロジェクト設定の場合はプロジェクト ルートを基準、ユーザー設定の場合は ~/.claude を基準にします .claude/settings.json の ./output は <project-root>/output に解決されます
絶対パスの古い //path プレフィックスは引き続き機能します。以前にプロジェクト相対の解決を期待して単一スラッシュ /path を使用していた場合は、 ./path に切り替えます。この構文は、絶対パスに //path を使用し、プロジェクト相対パスに /path を使用する読み取りおよび編集権限ルールとは異なります。サンドボックス ファイルシステムのパスは標準の規則を使用します: /tmp/build は絶対パスです。
構成例:
{
「サンドボックス」: {
"有効" : true 、
"autoAllowBashIfSandboxed" : true 、
「除外C」

コマンド" : [ "docker *" ],
「ファイルシステム」: {
"allowWrite" : [ "/tmp/build" , "~/.kube" ],
"denyRead" : [ "~/.aws/credentials" ]
}、
「ネットワーク」: {
"allowedDomains" : [ "github.com" , "*.npmjs.org" , "registry.yarnpkg.com" ],
"deniedDomains" : [ "uploads.github.com" ],
"allowUnixSockets" : [
「/var/run/docker.sock」
]、
"allowLocalBinding" : true
}
}
}
ファイルシステムとネットワークの制限は、結合する 2 つの方法で構成できます。
Sandbox.filesystem 設定 (上に表示): OS レベルのサンドボックス境界でパスを制御します。これらの制限は、Claude のファイル ツールだけでなく、すべてのサブプロセス コマンド (例: kubectl 、 terraform 、 npm ) に適用されます。
許可ルール: 編集許可/拒否ルールを使用してクロードのファイル ツール アクセスを制御し、読み取り拒否ルールを使用して読み取りをブロックし、WebFetch 許可/拒否ルールを使用してネットワーク ドメインを制御します。これらのルールからのパスもサンドボックス構成にマージされます。
コミットではデフォルトで git トレーラー ( Co-Authored-By など) が使用されます。これはカスタマイズまたは無効にすることができます
プルリクエストの説明はプレーンテキストです
共著: Claude Sonnet 4.6 <noreply@anthropic.com>
トレーラー内のモデル名は、セッションのアクティブなモデルを反映しています。
デフォルトのプルリクエストの属性:
🤖 [クロードコード](https://claude.com/claude-code)で生成
例:
{
「帰属」: {
"commit" : "AI で生成 \n\n 共著者: AI <ai@example.com>" ,
"pr" : ""
}
}
アトリビューション設定は、非推奨の includeCoAuthoredBy 設定よりも優先されます。すべての帰属を非表示にするには、commit と pr を空の文字列に設定します。
ファイル提案設定
{
"ファイルの提案" : {
"タイプ" : "コマンド" ,
"コマンド" : "~/.claude/file-suggestion.sh"
}
}
このコマンドは、 CLAUDE_PROJECT_DIR などのフックと同じ環境変数を使用して実行されます。クエリ フィールドを含む JSON を標準入力経由で受信します。
{ "く

ery" : "ソース/comp" }
改行で区切られたファイル パスを stdout に出力します (現在は 15 に制限されています)。
src/components/Button.tsx
src/components/Modal.tsx
src/コンポーネント/Form.tsx
例:
#!/bin/bash
クエリ = $( cat | jq -r '.query' )
# your-repo-file-index を独自のファイル検索コマンドに置き換えます
your-repo-file-index --query " $query " |頭 -20
フッターリンクバッジ
{
"フッターリンク正規表現" : [
{
"タイプ" : "正規表現" ,
"パターン" : " \\ b(?<key>PROJ- \\ d+) \\ b" ,
"url" : "https://issues.example.com/browse/{key}" ,
"ラベル" : "{キー}"
}
】
}
これを設定すると、ツールの結果またはクロードの応答に PROJ-1234 が表示されると、 https://issues.example.com/browse/PROJ-1234 にリンクするフッターに PROJ-1234 チップが表示されます。
次の制約が各エントリに適用されます。
制約動作 URL の起点 キャプチャされた値は URL エンコードされており、構築された URL はテンプレートのリテラルの起点を共有する必要があります。キャプチャではパス セグメントまたはクエリ値を埋めることはできますが、リンクが指す場所は変更できません URL 長 2048 文字を超える構築された URL は削除されます URL スキーム https 、 http 、または認識されたエディターまたはワークスペースのディープリンク スキームである必要があります: vscode 、 vscode-insider 、cursor 、 Windsurf 、 zed 、 Jetbrains 、 idea 、lack 、 Linear 、 notion 、 figma ラベルのデフォルト

[切り捨てられた]

## Original Extract

Configure Claude Code with global and project-level settings, and environment variables.

Claude Code settings - Claude Code Docs Documentation Index
Fetch the complete documentation index at: /docs/llms.txt
Use this file to discover all available pages before exploring further.
Skip to main content Claude Code Docs home page English Search... ⌘ K Ask Assistant Claude Developer Platform
Search... Navigation Settings and permissions Claude Code settings Getting started Build with Claude Code Administration Configuration Reference Agent SDK What's New Resources Settings and permissions
Speed up responses with fast mode
Escalate hard decisions with the advisor tool
Configuration scopes Available scopes
Settings files When edits take effect
Invalid entries in managed settings
Compute managed settings with a policy helper
Key points about the configuration system
Plugin configuration Plugin settings
Copy page Configure Claude Code with global and project-level settings, and environment variables.
Copy page Claude Code offers a variety of settings to configure its behavior to meet your needs. You can configure Claude Code by running the /config command when using the interactive REPL, which opens a tabbed Settings interface where you can view status information and modify configuration options.
​ Configuration scopes
Security policies that must be enforced organization-wide
Compliance requirements that can’t be overridden
Standardized configurations deployed by IT/DevOps
Personal preferences you want everywhere (themes, editor settings)
Tools and plugins you use across all projects
API keys and authentication (stored securely)
Team-shared settings (permissions, hooks, MCP servers)
Plugins the whole team should have
Standardizing tooling across collaborators
Personal overrides for a specific project
Testing configurations before sharing with the team
Machine-specific settings that won’t work for others
Managed (highest) - can’t be overridden by anything
Command line arguments - temporary session overrides
Local - overrides project and user settings
Project - overrides user settings
User (lowest) - applies when nothing else specifies the setting
User settings are defined in ~/.claude/settings.json and apply to all
projects.
Project settings are saved in your project directory:
.claude/settings.json for settings that are checked into source control and shared with your team
.claude/settings.local.json for settings that are not checked in, useful for personal preferences and experimentation. When Claude Code creates .claude/settings.local.json , it configures git to ignore the file. If you create the file yourself, add it to your gitignore manually.
Managed settings : For organizations that need centralized control, Claude Code supports multiple delivery mechanisms for managed settings. All use the same JSON format and cannot be overridden by user or project settings:
Server-managed settings : delivered from Anthropic’s servers via the Claude.ai admin console. See server-managed settings .
MDM/OS-level policies : delivered through native device management on macOS and Windows:
macOS: com.anthropic.claudecode managed preferences domain. The plist’s top-level keys mirror managed-settings.json , with nested settings as dictionaries and arrays as plist arrays. Deploy via configuration profiles in Jamf, Iru (Kandji), or similar MDM tools.
Windows: HKLM\SOFTWARE\Policies\ClaudeCode registry key with a Settings value (REG_SZ or REG_EXPAND_SZ) containing JSON (deployed via Group Policy or Intune)
Windows (user-level): HKCU\SOFTWARE\Policies\ClaudeCode (lowest policy priority, only used when no admin-level source exists)
File-based : managed-settings.json and managed-mcp.json deployed to system directories:
macOS: /Library/Application Support/ClaudeCode/
Linux and WSL: /etc/claude-code/
Windows: C:\Program Files\ClaudeCode\
The legacy Windows path C:\ProgramData\ClaudeCode\managed-settings.json is no longer supported as of v2.1.75. Administrators who deployed settings to that location must migrate files to C:\Program Files\ClaudeCode\managed-settings.json .
File-based managed settings also support a drop-in directory at managed-settings.d/ in the same system directory alongside managed-settings.json . This lets separate teams deploy independent policy fragments without coordinating edits to a single file.
Following the systemd convention, managed-settings.json is merged first as the base, then all *.json files in the drop-in directory are sorted alphabetically and merged on top. Later files override earlier ones for scalar values; arrays are concatenated and de-duplicated; objects are deep-merged. Hidden files starting with . are ignored.
Use numeric prefixes to control merge order, for example 10-telemetry.json and 20-security.json .
Managed deployments can also restrict plugin marketplace additions using
strictKnownMarketplaces . For more information, see Managed marketplace restrictions .
Other configuration is stored in ~/.claude.json . This file contains your OAuth session, MCP server configurations for user and local scopes, per-project state (allowed tools, trust settings), and various caches. Project-scoped MCP servers are stored separately in .mcp.json .
Claude Code automatically creates timestamped backups of configuration files and retains the five most recent backups to prevent data loss.
Example settings.json {
"$schema" : "https://json.schemastore.org/claude-code-settings.json" ,
"permissions" : {
"allow" : [
"Bash(npm run lint)" ,
"Bash(npm run test *)" ,
"Read(~/.zshrc)"
],
"deny" : [
"Bash(curl *)" ,
"Read(./.env)" ,
"Read(./.env.*)" ,
"Read(./secrets/**)"
]
},
"env" : {
"CLAUDE_CODE_ENABLE_TELEMETRY" : "1" ,
"OTEL_METRICS_EXPORTER" : "otlp"
},
"companyAnnouncements" : [
"Welcome to Acme Corp! Review our code guidelines at docs.acme.com" ,
"Reminder: Code reviews required for all PRs" ,
"New security policy in effect"
]
}
The $schema line in the example above points to the official JSON schema for Claude Code settings. Adding it to your settings.json enables autocomplete and inline validation in VS Code, Cursor, and any other editor that supports JSON schema validation.
The published schema is updated periodically and may not include settings added in the most recent CLI releases, so a validation warning on a recently documented field does not necessarily mean your configuration is invalid.
​ When edits take effect
model : use /model to switch mid-session
outputStyle : part of the system prompt, which is rebuilt on /clear or restart
​ Invalid entries in managed settings
Interactive sessions show a dialog at startup listing the invalid entries.
Headless runs with -p print a summary to stderr.
claude doctor lists each invalid entry with its source and field.
Versions before v2.1.119 also store a number of /config preference keys here instead of in settings.json , including theme , verbose , editorMode , autoCompactEnabled , and preferredNotifChannel .
Key Description Example autoConnectIde Automatically connect to a running IDE when Claude Code starts from an external terminal. Default: false . Appears in /config as Auto-connect to IDE (external terminal) when running outside a VS Code or JetBrains terminal. The CLAUDE_CODE_AUTO_CONNECT_IDE environment variable overrides this when set true autoInstallIdeExtension Automatically install the Claude Code IDE extension when running from a VS Code terminal. Default: true . Appears in /config as Auto-install IDE extension when running inside a VS Code or JetBrains terminal. You can also set the CLAUDE_CODE_IDE_SKIP_AUTO_INSTALL environment variable false externalEditorContext Prepend Claude’s previous response as # -commented context when you open the external editor with Ctrl+G . Default: false . Appears in /config as Show last response in external editor true teammateDefaultModel Default model for agent team teammates when the spawn prompt doesn’t specify one. Set to a model alias such as "sonnet" , or null to inherit the lead’s current /model selection. Appears in /config as Default teammate model "sonnet"
​ Worktree settings
Sandbox path prefixes
Paths in filesystem.allowWrite , filesystem.denyWrite , filesystem.denyRead , and filesystem.allowRead support these prefixes:
Prefix Meaning Example / Absolute path from filesystem root /tmp/build stays /tmp/build ~/ Relative to home directory ~/.kube becomes $HOME/.kube ./ or no prefix Relative to the project root for project settings, or to ~/.claude for user settings ./output in .claude/settings.json resolves to <project-root>/output
The older //path prefix for absolute paths still works. If you previously used single-slash /path expecting project-relative resolution, switch to ./path . This syntax differs from Read and Edit permission rules , which use //path for absolute and /path for project-relative. Sandbox filesystem paths use standard conventions: /tmp/build is an absolute path.
Configuration example:
{
"sandbox" : {
"enabled" : true ,
"autoAllowBashIfSandboxed" : true ,
"excludedCommands" : [ "docker *" ],
"filesystem" : {
"allowWrite" : [ "/tmp/build" , "~/.kube" ],
"denyRead" : [ "~/.aws/credentials" ]
},
"network" : {
"allowedDomains" : [ "github.com" , "*.npmjs.org" , "registry.yarnpkg.com" ],
"deniedDomains" : [ "uploads.github.com" ],
"allowUnixSockets" : [
"/var/run/docker.sock"
],
"allowLocalBinding" : true
}
}
}
Filesystem and network restrictions can be configured in two ways that are merged together:
sandbox.filesystem settings (shown above): Control paths at the OS-level sandbox boundary. These restrictions apply to all subprocess commands (e.g., kubectl , terraform , npm ), not just Claude’s file tools.
Permission rules : Use Edit allow/deny rules to control Claude’s file tool access, Read deny rules to block reads, and WebFetch allow/deny rules to control network domains. Paths from these rules are also merged into the sandbox configuration.
Commits use git trailers (like Co-Authored-By ) by default, which can be customized or disabled
Pull request descriptions are plain text
Co-Authored-By: Claude Sonnet 4.6 <noreply@anthropic.com>
The model name in the trailer reflects the active model for the session.
Default pull request attribution:
🤖 Generated with [Claude Code](https://claude.com/claude-code)
Example:
{
"attribution" : {
"commit" : "Generated with AI \n\n Co-Authored-By: AI <ai@example.com>" ,
"pr" : ""
}
}
The attribution setting takes precedence over the deprecated includeCoAuthoredBy setting. To hide all attribution, set commit and pr to empty strings.
​ File suggestion settings
{
"fileSuggestion" : {
"type" : "command" ,
"command" : "~/.claude/file-suggestion.sh"
}
}
The command runs with the same environment variables as hooks , including CLAUDE_PROJECT_DIR . It receives JSON via stdin with a query field:
{ "query" : "src/comp" }
Output newline-separated file paths to stdout (currently limited to 15):
src/components/Button.tsx
src/components/Modal.tsx
src/components/Form.tsx
Example:
#!/bin/bash
query = $( cat | jq -r '.query' )
# Replace your-repo-file-index with your own file search command
your-repo-file-index --query " $query " | head -20
​ Footer link badges
{
"footerLinksRegexes" : [
{
"type" : "regex" ,
"pattern" : " \\ b(?<key>PROJ- \\ d+) \\ b" ,
"url" : "https://issues.example.com/browse/{key}" ,
"label" : "{key}"
}
]
}
With this configured, when PROJ-1234 appears in a tool result or in Claude’s reply, a PROJ-1234 chip appears in the footer linking to https://issues.example.com/browse/PROJ-1234 .
The following constraints apply to each entry:
Constraint Behavior URL origin Captured values are URL-encoded and the constructed URL must share the template’s literal origin. A capture can fill a path segment or query value but cannot change where the link points URL length Constructed URLs longer than 2048 characters are dropped URL scheme Must be https , http , or a recognized editor or workspace deep-link scheme: vscode , vscode-insiders , cursor , windsurf , zed , jetbrains , idea , slack , linear , notion , figma Label Defaults

[truncated]
