---
source: "https://github.com/amaldini/maldalang"
hn_url: "https://news.ycombinator.com/item?id=49331723"
title: "MALDA – a language where LLM prompts and tools are syntax"
article_title: "GitHub - amaldini/maldalang: This is official home of MALDA programming language · GitHub"
image: "https://opengraph.githubassets.com/09c7d2128ef6fa5fa289076f3b5801834d8440a542ec71f72cee770f97baf0bf/amaldini/maldalang"
author: "amaldini"
captured_at: "2026-08-17T15:16:46Z"
capture_tool: "hn-digest"
hn_id: 49331723
score: 2
comments: 0
posted_at: "2026-08-17T14:35:13Z"
tags:
  - hacker-news
  - translated
---

# MALDA – a language where LLM prompts and tools are syntax

- HN: [49331723](https://news.ycombinator.com/item?id=49331723)
- Source: [github.com](https://github.com/amaldini/maldalang)
- Score: 2
- Comments: 0
- Posted: 2026-08-17T14:35:13Z

## Translation

タイトル: MALDA – LLM プロンプトとツールが構文である言語
記事タイトル: GitHub - amaldini/maldalang: ここは MALDA プログラミング言語の公式ホームです · GitHub
説明: これは MALDA プログラミング言語 - amaldini/maldalang の公式ホームです

記事本文:
GitHub - amaldini/maldalang: これは MALDA プログラミング言語の公式ホームです · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
アマルディーニ
/
マルダラン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
172 コミット 172 コミット .cursor/ rules .cursor/ rules .github .github .vscode .vscode 例 例 MaldaLang.Compiler MaldaLang.Compiler MaldaLang.DesktopIDE.Tests MaldaLang.DesktopIDE.Tests MaldaLang.DesktopIDE Malda

Lang.DesktopIDE MaldaLang.IDE MaldaLang.IDE MaldaLang.LanguageServer MaldaLang.LanguageServer MaldaLang.TestLib MaldaLang.TestLib MaldaLang.Tests MaldaLang.UIHost MaldaLang.UIHost MaldaLang MaldaLang リファレンスマニュアル リファレンスマニュアル テンプレート テンプレート準拠/ tier0適合/ tier0 ドキュメント ドキュメント パッケージ パッケージ スクリプト スクリプト vscode-malda vscode-malda .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md LICENSE-APACHEライセンス-APACHE ライセンス-MIT ライセンス-MIT ライセンス-ランタイム-例外 ライセンス-ランタイム-例外 MaldaLang.sln MaldaLang.sln README.md README.md SECURITY.md SECURITY.md THIRD-PARTY-NOTICES.md THIRD-PARTY-NOTICES.md TRADEMARK.md TRADEMARK.md build_malda_distribution.bat build_malda_distribution.bat host.html host.html install-malda-extension.malda install-malda-extension.malda llms.txt llms.txt malda-cheat-sheet.html malda-cheat-sheet.html project.html project.html すべてのファイルを表示 リポジトリ ファイルのナビゲーション
MALDA (開発自動化を備えたマルチ エージェント言語) は、エージェント、Web アプリ、耐久性のあるワークフローのための AI ネイティブ プログラミング言語およびプラットフォームです。この名前の自動化は両方の方法で実行されます。コーディング エージェントが MALDA を作成します (docs/llm/ 内の言語パックはそのリーダー用に存在します)。そして、MALDA プログラムが開発作業を順番に自動化します。
このリポジトリは、言語ランタイム、コンパイラ/トランスパイラ、IDE、言語サーバー、サンプル、テンプレート、リファレンス マニュアルなどのオープンソース コアです。
Examples/Basics/first_look.malda からそのまま引用。 APIキーがありません。 -> レビューにより、プロンプトがスキーマにバインドされます。 await がない場合、呼び出しはレンダリングされたテンプレートになります。 await はモデルを呼び出し、 Review に対して JSON を検証します。 validate("Review", …) は同じチェックであり、オフラインで表示されます (

他の場所の型アノテーションは IDE/LSP ヒントであり、実行時チェックではありません)。
スキーマのレビュー {
概要: 文字列;
問題: string[];
}
プロンプト codeReview(コード, 言語) -> レビュー {
system: "あなたは {言語} の専門査読者です。",
ユーザー: 「この {言語} コードを確認してください:\n\n{コード}」
}
var rendered = codeReview("function add(a, b) { return a + b; }", "javascript");
io.print(rendered.user);
var selected = validate("レビュー", {
"summary": "問題ないようです",
「問題」: []
});
if (checked.ok) {
io.print("スキーマ ok: " +checked.data.summary);
} それ以外の場合は {
io.print("スキーマが失敗しました: " +checked.error);
}
マルダの例/基本/first_look.malda
プロンプトは構文です。 await を使用しないと、レンダリングされたプロンプトが表示されます。 await を使用してモデルを呼び出します。
HttpServer の問い合わせフォーム: フィールドに記入、 @POST /submit 、サンキューページ (Examples/ui_contact_form.malda パターン)。
言語とランタイム — オブジェクト、関数、アクター、例外、名前空間付き標準ライブラリ ( io / math / str )
3 つのバックエンド、文書化された重複 - 反復のために解釈。 .NET 実行可能ファイルにトランスパイルします。ブラウザーのサブセットを JavaScript にコンパイルします (JS パス上にエージェント、HTTP サーバー、またはワークフローはありません)。マトリックス: docs/spec/backend-capability-matrix.md
AI — プロンプト宣言、エージェント、ツール、MCP
Web — ホスト ランタイム上の REST デコレータと @PAGE / @AIPAGE
ワークフロー — ローカル SQLite での永続的なステップ / 再試行 / 補正 (クラスターではなく単一のライター)
タイプ — 注釈は IDE/LSP にフィードされます (エディターでは、不一致はデフォルトでエラーになります)。ランタイムは動的に維持されます。 malda コンパイル --mode transpile / public はこれらのエラーでの出力を拒否します ( --lenient-types をスキップします)
ツール — デスクトップ IDE (Windows リファレンス)、Web IDE (ブラウザー プレイグラウンド — デスクトップ パリティではない)、VS Code + LSP
これは、Temporal クラスター、完全な静的型システム、または 3 つの同等のバックエンドではありません。ワークフローは

ローカル SQLite (単一ライター)。 JavaScript はブラウザのサブセットです。 Spec Final 1.0 は言語カーネルです。ツールチェーンは 1.0.0 です (発行は型境界です)。
MALDA は存在しないにもかかわらず、2 つの最大のショーケースはすべてコーディング エージェントによって作成されました。
モデルのトレーニング データ。 docs/llm/ にはそのためのコンパクトな言語パックが同梱されているためです。
リーダー: Examples/Agents/secondbrain_semantic.malda (共有ライブラリを含む最大 7,600 行) および
例/RalphWiggum/ (4,049 行)。ガードテスト、適合マトリックス、および実行
リファレンスマニュアルのスニペットにより、これらの主張をチェックできるようになります。
反対意見と現在の弱点を前もって述べた、より長いバージョン:
docs/payment.md 。
ソースからビルドする場合は .NET 8 SDK
Windows デスクトップ IDE (WPF); CLI、Web IDE、および VS Code 拡張機能は、.NET 8 が行うクロスプラットフォームで動作します。
ソースを複製せずにインストールする
GitHub リリースには、自己完結型の zip が同梱されています (個別の .NET インストールは必要ありません)。
リリースを開いてダウンロードする
malda-<バージョン>-win-x64.zip または malda-<バージョン>-linux-x64.zip 。
# Windows - CLI
malda.bat 例 \B asics \f irst_look.malda
# または: bin\malda\malda.exe Examples\Basics\first_look.malda
# Windows — デスクトップ IDE (win-x64 zip に含まれています)
MaldaDesktop.bat
# Linux — CLI のみ (WPF デスクトップ IDE なし)
./malda 例/基本/first_look.malda
このアーカイブには、Examples/ 、HTML ReferenceManual/ 、Templates/ も含まれています。
( malda new の場合)、 docs/llm/ 、 docs/spec/ の下の言語パック、およびデュアルライセンス
ファイル。コーディング エージェントに AGENTS.md (その後 docs/llm/ ) を指定して、.malda を書き込めるようにします。
git ソースのないプログラム。オプション: bin/malda を PATH に追加します。
これらの zip をローカルで再構築するには:
# Windows (CLI + デスクトップ IDE)
build_malda_distribution.bat
# または: powershell -File scripts/build-oss-dist.ps1 -Runtime all
タグ付きリリース ( v* ) は、.githu を介して CI で zip をビルドします

b/workflows/release.yml 。
git clone https://github.com/amaldini/maldalang.git
CD マルダラン
ドットネット ビルド MaldaLang.sln
dotnet run --project MaldaLang -- 例/基本/first_look.malda
Linux および macOS では、ソリューションの代わりにプロジェクトをビルドします — MaldaLang.sln
WPF デスクトップ IDE が含まれています。これは net8.0-windows をターゲットにしており、他の場所ではビルドできません。
ドットネット ビルド MaldaLang
dotnet build MaldaLang.Compiler
dotnet build MaldaLang.LanguageServer
ドットネット ビルド MaldaLang.IDE
dotnet run --project MaldaLang -- 例/基本/first_look.malda
または、再利用可能な CLI 出力を構築します。
dotnet build MaldaLang -o artifacts/malda-cli
# Windows: malda.exe — Linux/macOS: malda
artifacts/malda-cli/malda 例/基本/first_look.malda
実行可能ファイルにコンパイルします (ワンライナーのスモーク: Examples/Basics/hello_world.malda ):
dotnet run --project MaldaLang --compile Examples/Basics/first_look.malda --mode transpile -o first-look.exe
LLM アクセス (オプション)
first_look.malda や hello_world.malda などのコア言語の例には API キーは必要ありません
そしてネットワーク通話を行わないでください。
AI 機能 (プロンプト、エージェント、@AIPAGE など) は次の順序で使用します。
OPENROUTER_API_KEY が設定されている場合の OpenRouter (または、providers.openrouter.apiKey
~/.malda/config.json )。
それ以外の場合は、ローカル GGUF フォールバック: 最初に MALDA ダウンロードを使用する
qwen2.5-0.5b-instruct-q4_k_m.gguf
Hugging Face からローカルのアプリデータ キャッシュに (~500 MB)
(Windows では %LOCALAPPDATA%\MaldaLang\Models\default)。そのモデルはパイプラインを証明します
オフライン。これは実稼働品質のチャット モデルではありません。
ローカル モデルを意図的にプリフェッチするには:
dotnet run --project MaldaLang -- onboard --download-local-llama
IDE とエディター
これらは同じ製品表面ではありません。
# Web IDE (ブラウザ プレイグラウンド)
dotnet run --project MaldaLang.IDE
# 出力された https://localhost:... URL を開きます
# デスクトップ

IDE (Windows)
dotnet run --project MaldaLang.DesktopIDE
Web IDE (Monaco UX、診断、サンプル ブラウザー) に関するコミュニティ ヘルプは大歓迎です。デスクトップは引き続き参照 IDE です。
私は何年も C#、Java、VB.NET でビジネス ソフトウェア (PDM/PLM システム、CAD) を書いてきました。
統合、ERP やその他のビジネス システムとの統合、データ処理 - どこで作業するか
難しいのはドメインではなく、既存のシステムを相互に一致させることです。
言語の内部。これは私が初めて作成したコンパイラまたはインタプリタです。コーディングあり
エージェント 今では手作業でコードを書くことははるかに少なくなり、言語はプロジェクトとして残ります。
決定事項: エージェントはパーサーを入力しますが、文法が何であるかを誰かが言わなければなりません
という意味です。これらの決定も単独で行われたわけではなく、文法と意味論が議論されました。
モデルを使って開発したため、プロジェクトは最終的に、プログラミングをどのように設計するかという独立したテーマになりました。
現在、エージェントはほとんどのコードをこの言語で記述しています。トライするよう後押ししたのはジェフ・ハントリー氏の言葉だった。
ラルフ・ウィガムのテクニックと
彼がそれを使って構築した呪われた言語 - したがって、
例/RalphWiggum/ 。私が望んでいたのは難解な言語とは異なりました。
主題は私が日々費やしている仕事なので、プロンプト、ツール、エージェント、エンドポイント、そして耐久性
ワークフローはグルー コードではなく構文です。今日はその仕事で MALDA を実行していません - 私の日です
対象となるのは、すでに稼働中の大規模システムであり、それらは実験用に書き換えられることはありません。のために
今は、AI をベースにした構築がどのようなものかを実験する場所です。
パス
目的
エージェント.md
人間とコーディングエージェントのためのマップ
ドキュメント/llm/
.malda プログラムを作成するためのコンパクト パック
docs/architecture.md
パイプラインとプロジェクトのレイアウト
docs/start-here.md
ガイド付きオンボーディング
llms.txt
LLM ツール用のコンパクトなドキュメント インデックス
docs/releases/v1.0.0.md
ツールチェーン 1.0

.0 (publish は型境界です)
docs/releases/v0.1.0.md
最初のパブリックタグに関する注意事項
docs/releases/v0.1.1.md
--embed-folder と Second Brain に関する注意事項
docs/releases/v0.1.2.md
セマンティック セカンド ブレイン、BGE-M3、getProgramDirectory()
docs/releases/v0.1.3.md
埋め込みからの GraphMemory 読み取り専用ロード:
docs/releases/v0.1.4.md
Second Brain ASK Web UI、markdownToHtml、空脳チャット
docs/releases/v0.1.5.md
ポータブル ASK ブランド化、メニュー パック、compileMALDA embedFolder
docs/releases/v0.1.6.md
ロケールセーフな float トランスパイル、Spectre マークアップ フォールバック、リリース バージョン ガード
docs/releases/v0.1.7.md
Agent/CodingAgent トランスパイル、オプションの Agent クライアント、Second Brain コンパイル
例/README.md
完全なサンプル カタログ ( が必要、トラック)
例/基本
言語の基礎
例/プロンプト
プロンプトとエージェント
例/Web
REST、UI、認証、ジョブ
docs/tutorials/fullstack-sessions-auth.md
セッション/CSRF/ジョブのウォークスルー
テンプレート/
マルダ新しいウェバピ / マルダ新しいフルスタック
リファレンスマニュアル/
言語リファレンス (HTML)。英語 · イタリア語
vscode-マルダ/
VS コード拡張機能
ソリューションのレイアウト
プロジェクト
役割
マルダラン
CLI + インタプリタ ( malda )
MaldaLang.コンパイラ
トランスパイラー/パブリッシュ
MaldaLang.IDE
Web IDE (ブラウザ プレイグラウンド)
MaldaLang.DesktopIDE
デスクトップ IDE (完全な Windows IDE)
MaldaLang.L

[切り捨てられた]

## Original Extract

This is official home of MALDA programming language - amaldini/maldalang

GitHub - amaldini/maldalang: This is official home of MALDA programming language · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
amaldini
/
maldalang
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
172 Commits 172 Commits .cursor/ rules .cursor/ rules .github .github .vscode .vscode Examples Examples MaldaLang.Compiler MaldaLang.Compiler MaldaLang.DesktopIDE.Tests MaldaLang.DesktopIDE.Tests MaldaLang.DesktopIDE MaldaLang.DesktopIDE MaldaLang.IDE MaldaLang.IDE MaldaLang.LanguageServer MaldaLang.LanguageServer MaldaLang.TestLib MaldaLang.TestLib MaldaLang.Tests MaldaLang.Tests MaldaLang.UIHost MaldaLang.UIHost MaldaLang MaldaLang ReferenceManual ReferenceManual Templates Templates conformance/ tier0 conformance/ tier0 docs docs packages packages scripts scripts vscode-malda vscode-malda .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT LICENSE-RUNTIME-EXCEPTION LICENSE-RUNTIME-EXCEPTION MaldaLang.sln MaldaLang.sln README.md README.md SECURITY.md SECURITY.md THIRD-PARTY-NOTICES.md THIRD-PARTY-NOTICES.md TRADEMARK.md TRADEMARK.md build_malda_distribution.bat build_malda_distribution.bat host.html host.html install-malda-extension.malda install-malda-extension.malda llms.txt llms.txt malda-cheat-sheet.html malda-cheat-sheet.html program.html program.html View all files Repository files navigation
MALDA (Multi Agent Language with Development Automation) is an AI-native programming language and platform for agents, web apps, and durable workflows. The automation in the name runs both ways: coding agents write MALDA — the language pack in docs/llm/ exists for that reader — and MALDA programs automate development work in turn.
This repository is the open-source core : language runtime, compiler/transpiler, IDEs, language server, examples, templates, and reference manual.
Verbatim from Examples/Basics/first_look.malda . No API key. -> Review binds the prompt to the schema. Without await , the call is a rendered template; await would call the model and validate the JSON against Review . validate("Review", …) is that same check, shown offline (type annotations elsewhere are IDE/LSP hints, not runtime checks).
schema Review {
summary: string;
issues: string[];
}
prompt codeReview(code, language) -> Review {
system: "You are an expert reviewer of {language}.",
user: "Review this {language} code:\n\n{code}"
}
var rendered = codeReview("function add(a, b) { return a + b; }", "javascript");
io.print(rendered.user);
var checked = validate("Review", {
"summary": "Looks fine",
"issues": []
});
if (checked.ok) {
io.print("schema ok: " + checked.data.summary);
} else {
io.print("schema failed: " + checked.error);
}
malda Examples/Basics/first_look.malda
prompt is syntax: without await you get the rendered prompt; with await you call the model.
Contact form on HttpServer : fill fields, @POST /submit , thank-you page ( Examples/ui_contact_form.malda pattern).
Language & runtime — objects, functions, actors, exceptions, namespaced standard library ( io / math / str )
Three backends, documented overlap — interpret for iteration; transpile to a .NET executable; compile a browser subset to JavaScript (no agents, HTTP servers, or workflows on the JS path). Matrix: docs/spec/backend-capability-matrix.md
AI — prompt declarations, agents, tools, MCP
Web — REST decorators and @PAGE / @AIPAGE on the host runtime
Workflows — durable step / retry / compensate on local SQLite (single writer, not a cluster)
Types — annotations feed the IDE/LSP (mismatches are Errors by default in the editor); runtime stays dynamic. malda compile --mode transpile / publish refuses emit on those Errors ( --lenient-types to skip)
Tooling — Desktop IDE (Windows reference), Web IDE (browser playground — not Desktop parity), VS Code + LSP
What this is not: a Temporal cluster, a full static type system, or three equal backends. Workflows are local SQLite (single writer). JavaScript is a browser subset. Spec Final 1.0 is the language kernel; the toolchain is 1.0.0 (publish is the type boundary).
The two largest showcases were written entirely by coding agents, even though MALDA is in no
model's training data, because docs/llm/ ships a compact language pack for that
reader: Examples/Agents/secondbrain_semantic.malda (~7,600 lines with shared libs) and
Examples/RalphWiggum/ (4,049 lines). Guard tests, the conformance matrix, and executed
reference-manual snippets keep those claims checkable.
Longer version, with objections and current weaknesses stated up front:
docs/announcement.md .
.NET 8 SDK if you build from source
Windows for Desktop IDE (WPF); CLI, Web IDE, and VS Code extension work cross-platform where .NET 8 does
Install without cloning the sources
GitHub Releases ship self-contained zips (no separate .NET install needed):
Open Releases and download
malda-<version>-win-x64.zip or malda-<version>-linux-x64.zip .
# Windows — CLI
malda.bat Examples \B asics \f irst_look.malda
# or: bin\malda\malda.exe Examples\Basics\first_look.malda
# Windows — Desktop IDE (included in the win-x64 zip)
MaldaDesktop.bat
# Linux — CLI only (no WPF Desktop IDE)
./malda Examples/Basics/first_look.malda
The archive also includes Examples/ , the HTML ReferenceManual/ , Templates/
(for malda new ), the language pack under docs/llm/ , docs/spec/ , and dual-licence
files. Point coding agents at AGENTS.md (then docs/llm/ ) so they can write .malda
programs without the git sources. Optional: add bin/malda to your PATH .
To rebuild those zips locally:
# Windows (CLI + Desktop IDE)
build_malda_distribution.bat
# or: powershell -File scripts/build-oss-dist.ps1 -Runtime all
Tagged releases ( v* ) build the zips in CI via .github/workflows/release.yml .
git clone https://github.com/amaldini/maldalang.git
cd maldalang
dotnet build MaldaLang.sln
dotnet run --project MaldaLang -- Examples/Basics/first_look.malda
On Linux and macOS , build the projects instead of the solution — MaldaLang.sln
includes the WPF Desktop IDE, which targets net8.0-windows and cannot build elsewhere:
dotnet build MaldaLang
dotnet build MaldaLang.Compiler
dotnet build MaldaLang.LanguageServer
dotnet build MaldaLang.IDE
dotnet run --project MaldaLang -- Examples/Basics/first_look.malda
Or build a reusable CLI output:
dotnet build MaldaLang -o artifacts/malda-cli
# Windows: malda.exe — Linux/macOS: malda
artifacts/malda-cli/malda Examples/Basics/first_look.malda
Compile to an executable (one-liner smoke: Examples/Basics/hello_world.malda ):
dotnet run --project MaldaLang -- compile Examples/Basics/first_look.malda --mode transpile -o first-look.exe
LLM access (optional)
Core language examples such as first_look.malda and hello_world.malda need no API key
and make no network calls .
AI features ( prompt , agents, @AIPAGE , and similar) use, in order:
OpenRouter when OPENROUTER_API_KEY is set (or providers.openrouter.apiKey in
~/.malda/config.json ).
Otherwise a local GGUF fallback : on first use MALDA downloads
qwen2.5-0.5b-instruct-q4_k_m.gguf
(~500 MB) from Hugging Face into the local app-data cache
( %LOCALAPPDATA%\MaldaLang\Models\default on Windows). That model proves the pipeline
offline; it is not a production-quality chat model.
To prefetch the local model deliberately:
dotnet run --project MaldaLang -- onboard --download-local-llama
IDEs and editors
These are not the same product surface:
# Web IDE (browser playground)
dotnet run --project MaldaLang.IDE
# then open the printed https://localhost:... URL
# Desktop IDE (Windows)
dotnet run --project MaldaLang.DesktopIDE
Community help on the Web IDE (Monaco UX, diagnostics, examples browser) is welcome; Desktop remains the reference IDE.
I have spent years writing business software in C#, Java and VB.NET — PDM/PLM systems, CAD
integrations, integrations with ERP and other business systems, data processing — work where
the hard part is the domain and getting existing systems to agree with each other, not
language internals. This is the first compiler or interpreter I have written. With coding
agents I now write far less code by hand, and a language is the project where what remains is
the deciding: an agent will type a parser, but someone still has to say what the grammar
means. Those decisions were not made alone either — the grammar and the semantics were argued
out with models, so the project ended up being its own subject: how you design a programming
language now that agents write most of the code. The push to try came from Geoff Huntley's
Ralph Wiggum technique and the
cursed language he built with it — hence the name of
Examples/RalphWiggum/ . What I wanted was different from an esoteric language: one whose
subject is the work I spend my days around, so prompts, tools, agents, endpoints and durable
workflows are syntax instead of glue code. I am not running MALDA in that work today — my day
job is large systems already in flight, and those do not get rewritten for an experiment. For
now it is where I experiment with what building on AI looks like.
Path
Purpose
AGENTS.md
Map for humans and coding agents
docs/llm/
Compact pack for writing .malda programs
docs/architecture.md
Pipeline and project layout
docs/start-here.md
Guided onboarding
llms.txt
Compact doc index for LLM tools
docs/releases/v1.0.0.md
Toolchain 1.0.0 (publish is the type boundary)
docs/releases/v0.1.0.md
Notes for the first public tag
docs/releases/v0.1.1.md
Notes for --embed-folder and Second Brain
docs/releases/v0.1.2.md
Semantic Second Brain, BGE-M3, getProgramDirectory()
docs/releases/v0.1.3.md
GraphMemory read-only load from embed:
docs/releases/v0.1.4.md
Second Brain ASK web UI, markdownToHtml , empty-brain chat
docs/releases/v0.1.5.md
Portable ASK branding, menu PACK, compileMALDA embedFolder
docs/releases/v0.1.6.md
Locale-safe float transpile, Spectre markup fallback, release version guard
docs/releases/v0.1.7.md
Agent/CodingAgent transpile, optional Agent client, Second Brain compile
Examples/README.md
Full examples catalog ( requires , tracks)
Examples/Basics
Language basics
Examples/Prompts
Prompts and agents
Examples/Web
REST, UI, auth, jobs
docs/tutorials/fullstack-sessions-auth.md
Sessions / CSRF / jobs walkthrough
Templates/
malda new webapi / malda new fullstack
ReferenceManual/
Language reference (HTML); English · Italiano
vscode-malda/
VS Code extension
Solution layout
Project
Role
MaldaLang
CLI + interpreter ( malda )
MaldaLang.Compiler
Transpiler / publish
MaldaLang.IDE
Web IDE (browser playground)
MaldaLang.DesktopIDE
Desktop IDE (full Windows IDE)
MaldaLang.L

[truncated]
