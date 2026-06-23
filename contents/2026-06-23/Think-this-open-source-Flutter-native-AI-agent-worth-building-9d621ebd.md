---
source: "https://github.com/anasfik/flutter_copilot"
hn_url: "https://news.ycombinator.com/item?id=48648808"
title: "Think this open-source Flutter-native AI agent worth building?"
article_title: "GitHub - anasfik/flutter_copilot: AI agent framework for Flutter. Autonomously navigate and interact with app UIs through the semantics tree using LLMs. · GitHub"
author: "gwhyyy"
captured_at: "2026-06-23T18:39:39Z"
capture_tool: "hn-digest"
hn_id: 48648808
score: 1
comments: 1
posted_at: "2026-06-23T18:02:42Z"
tags:
  - hacker-news
  - translated
---

# Think this open-source Flutter-native AI agent worth building?

- HN: [48648808](https://news.ycombinator.com/item?id=48648808)
- Source: [github.com](https://github.com/anasfik/flutter_copilot)
- Score: 1
- Comments: 1
- Posted: 2026-06-23T18:02:42Z

## Translation

タイトル: このオープンソースの Flutter ネイティブ AI エージェントは構築する価値があると思いますか?
記事タイトル: GitHub - anasfik/flutter_copilot: Flutter 用 AI エージェント フレームワーク。 LLM を使用して、セマンティクス ツリーを通じてアプリの UI を自律的に移動し、操作します。 · GitHub
説明: Flutter 用の AI エージェント フレームワーク。 LLM を使用して、セマンティクス ツリーを通じてアプリの UI を自律的に移動し、操作します。 - anasfik/flutter_copilot

記事本文:
GitHub - anasfik/flutter_copilot: Flutter 用 AI エージェント フレームワーク。 LLM を使用して、セマンティクス ツリーを通じてアプリの UI を自律的に移動し、操作します。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
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
アナスフィク
/
flutter_copilot
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミットの例 example lib lib test test .gitignore .gitignore .pubignore .pubignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md Analysis_options.yaml Analysis_options.yaml pubspec.yaml pubspec.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Flutter アプリに、現在の UI を理解し、アクションを選択し、タップ、入力、スクロール、待機し、ユーザーの目標が完了するまで続行できる目に見えない AI コパイロットを与えます。
flutter_copilot は Flutter のセマンティクス ツリーから動作します。スクリーンショットやオーバーレイ チャット UI は必要ありません。アプリを CopilotApp でラップし、OpenAI 互換の LLM を提供してから、自然言語ゴールを指定して CopilotController.run(...) を呼び出します。
WhatsApp.Video.2026-06-22.at.10.18.08.PM.mp4
注: これはデモを速くするために明らかにバイブコーディングされています。問題を開いてこれについて言及する人を聞きたくないです!
オプション、画面、フォーム、ボタン、タブ、ダイアログ、設定がたくさんあるアプリを想像してみてください。ユーザーにすべてのコントロールを手動で見つけさせる代わりに、flutter_copilot は自然言語からアプリ内のタスクを実行できます。
設定に移動し、ダーク モードをオンにし、毎週の概要メールを有効にして保存します。
現在の Flutter UI を監視し、次に表示されるアクションを決定し、タップ、入力、スクロール、待機し、タスクが完了するまで繰り返します。
サンプル アプリを参照して実際にどのように機能するかを確認し、副操縦士が最も複雑なフローを理解できるように「副操縦士が UI を確認できるようにする」を確認してください。
画面とタブの間を移動します。
ボタン、スイッチ、チェックボックス、リスト項目、およびナビゲーション コントロールをタップします。
表示されているテキストフィールドに入力します。
表示されているスクロール可能な領域をスクロールします。
読み込みまたはアニメーションを待ちます。
EXE

1 つのモデル ステップで、いくつかの独立した目に見えるアクションを実行できます。
イベントを通じて進行状況をレポートできるため、ライブアクティビティログを表示できます。
ブロックされたアクション、破壊的なアクション、または不可能なアクションでは安全に停止します。
LLM と UI が改善されると、さらに多くの効果が得られます。
依存関係 :
flutter_copilot : ^0.9.1
サポートされているプロバイダー
OpenAI API 互換プロバイダー (OpenRouter や、ツール呼び出しをサポートする同様の /chat/completions エンドポイントなど)。
アプリを CopilotApp でラップし、LLM を構成します。
インポート 'パッケージ:flutter/material.dart' ;
import 'パッケージ:flutter_copilot/flutter_copilot.dart' ;
void main() {
runApp (
CopilotApp (
config : CopilotConfig (
// すべての OpenAI および OpenAI API 互換プロバイダーと互換性があります
llm : OpenAILlmAdapter (
apiKey : 'YOUR_KEY_HERE' 、
モデル: 'THE_MODEL_YOU_WANT_TO_USE' 、
// オプション: 異なるエンドポイントで互換性のあるプロバイダーを使用する場合
エンドポイント: Uri 。解析 ( 'https://api.openai.com/v1/chat/completions' )、
）、
デバッグログ: true 、
）、
子: const YourApp ()、
）、
);
}
簡単な使い方
構成した CopilotApp の下の任意の場所で、コントローラーを取得し、最初のプロンプトを実行します。
最終コントローラ = CopilotController 。 （コンテキスト）の;
最終結果 = コントローラーを待ちます。走って（
「設定を開いてダーク モードをオンにしてから、戻ってプロフィールを開き、表示名を Anas に更新して保存します。」 、
);
副操縦士が UI を確認できるようにする
flutter_copilot はセマンティクスを読み取るため、重要なコントロールに明確なラベルがあると、アプリの自動化がはるかに簡単になります。
ほとんどのマテリアル ウィジェットはすでに適切なセマンティクスを公開しています。
SwitchListTile (
タイトル : const Text ( 'ダークモード' )、
値: ダークモード、
onChanged : setDarkMode、
)
カスタム コントロールの場合は、明示的なセマンティクスを追加します。
セマンティクス (
ラベル: 'ダークモード' 、
値: ダークモード ? 'オン' : 'オフ' 、
切り替え：ダークモード、
ボタン: true 、
onTap : () => setDarkMode ( ! darkMode),
子：Exc

ルードセマンティクス (
子: MyCustomSwitch (値: darkMode)、
）、
)
良いラベルは、人が求めるもののように聞こえます。
ツールヒントやセマンティクス ラベルのないアイコンのみのコントロールは避けてください。副操縦士はラベルのないボタンを確実に選択できません。
CopilotController.events をリッスンすることで、いつでもどこでも CopilotApp の下でコパイロットのステータスや進行状況を追跡できます。
コパイロットコントローラー 。 (コンテキスト).イベントの。聞いてください ((イベント) {
print ( 'コパイロットイベント: $event ' );
});
安全性
flutter_copilot には CopilotSafetyPolicy.defaults が付属しており、ターゲット コントロールのラベル、値、ヒントがログアウト、支払い、転送、購入、アカウント削除などの危険なテキストと一致する場合に、計画されたアクションをブロックします。
副操縦士の安全性ポリシー (
BlockedLabels : const <パターン> [
「アカウントを削除」 、
「支払いを確認する」、
「送金」 、
]、
)
計画されたアクションが blockedLabels と一致すると、実行は実行されずに停止します。
GitHub の問題やプル リクエストを通じて貢献を歓迎します。短いガイドについては、CONTRIBUTING.md を参照してください。
flutter_copilot は MIT ライセンスに基づいてリリースされています。 「ライセンス」を参照してください。
Flutter 用の AI エージェント フレームワーク。 LLM を使用して、セマンティクス ツリーを通じてアプリの UI を自律的に移動し、操作します。
pub.dev/packages/flutter_copilot
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AI agent framework for Flutter. Autonomously navigate and interact with app UIs through the semantics tree using LLMs. - anasfik/flutter_copilot

GitHub - anasfik/flutter_copilot: AI agent framework for Flutter. Autonomously navigate and interact with app UIs through the semantics tree using LLMs. · GitHub
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
anasfik
/
flutter_copilot
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits example example lib lib test test .gitignore .gitignore .pubignore .pubignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md analysis_options.yaml analysis_options.yaml pubspec.yaml pubspec.yaml View all files Repository files navigation
Give your Flutter app an invisible AI copilot that can understand the current UI, choose actions, tap, type, scroll, wait, and keep going until a user goal is complete.
flutter_copilot works from Flutter's semantics tree. It does not need screenshots or an overlay chat UI. You wrap your app with CopilotApp , provide an OpenAI-compatible LLM, then call CopilotController.run(...) with a natural-language goal.
WhatsApp.Video.2026-06-22.at.10.18.08.PM.mp4
Note: this is clearly vibe-coded ro make demo faster, don't wanna hear anyone opening an issue mentioning this please!
Imagine an app full of options, screens, forms, buttons, tabs, dialogs, and settings. Instead of making the user find every control manually, flutter_copilot can carry out tasks inside your app from natural language:
Go to settings, turn on dark mode, enable weekly summary emails, and save.
It observes the current Flutter UI, decides the next visible action, taps, types, scrolls, waits, and repeats until the task is done.
See the example app to see how it works in practice, and check out Help the Copilot See Your UI to help the copilot understand your most complex flows.
Navigate between screens and tabs.
Tap buttons, switches, checkboxes, list items, and navigation controls.
Type into visible text fields.
Scroll visible scrollable areas.
Wait for loading or animations.
Execute several independent visible actions in one model step.
Report progress through events so you can show a live activity log.
Stop safely on blocked, destructive, or impossible actions.
Much more as LLMs and your UI improve!
dependencies :
flutter_copilot : ^0.9.1
Supported Providers
OpenAI API-compatible providers, including OpenRouter and similar /chat/completions endpoints that support tool calling.
Wrap your app with CopilotApp and configure the LLM:
import 'package:flutter/material.dart' ;
import 'package:flutter_copilot/flutter_copilot.dart' ;
void main () {
runApp (
CopilotApp (
config : CopilotConfig (
// compatible with all OpenAI and OpenAI API-compatible providers
llm : OpenAILlmAdapter (
apiKey : 'YOUR_KEY_HERE' ,
model : 'THE_MODEL_YOU_WANT_TO_USE' ,
// optional: if using a compatible provider with a different endpoint
endpoint : Uri . parse ( 'https://api.openai.com/v1/chat/completions' ),
),
debugLogging : true ,
),
child : const YourApp (),
),
);
}
Quick Usage
Anywhere below the CopilotApp you configured, get the controller and run your first prompt:
final controller = CopilotController . of (context);
final result = await controller. run (
'Open settings and turn on dark mode, then go back and open profile and update the display name to Anas, then save.' ,
);
Help the Copilot See Your UI
flutter_copilot reads semantics, so your app becomes much easier to automate when important controls have clear labels.
Most Material widgets already expose good semantics:
SwitchListTile (
title : const Text ( 'Dark mode' ),
value : darkMode,
onChanged : setDarkMode,
)
For custom controls, add explicit semantics:
Semantics (
label : 'Dark mode' ,
value : darkMode ? 'On' : 'Off' ,
toggled : darkMode,
button : true ,
onTap : () => setDarkMode ( ! darkMode),
child : ExcludeSemantics (
child : MyCustomSwitch (value : darkMode),
),
)
Good labels sound like what a person would ask for:
Avoid icon-only controls without tooltips or semantics labels. The copilot cannot reliably choose between unlabeled buttons.
You can track the copilot status or progress anytime, anywhere below CopilotApp , by listening to CopilotController.events .
CopilotController . of (context).events. listen ((event) {
print ( 'Copilot event: $ event ' );
});
Safety
flutter_copilot ships with CopilotSafetyPolicy.defaults , which blocks planned actions when the target control label, value, or hint matches risky text like logout, payment, transfer, purchase, or account deletion.
CopilotSafetyPolicy (
blockedLabels : const < Pattern > [
'delete account' ,
'confirm payment' ,
'send money' ,
],
)
When a planned action matches blockedLabels , the run stops instead of executing it.
Contributions are welcome through GitHub issues and pull requests. See CONTRIBUTING.md for the short guide.
flutter_copilot is released under the MIT License. See LICENSE .
AI agent framework for Flutter. Autonomously navigate and interact with app UIs through the semantics tree using LLMs.
pub.dev/packages/flutter_copilot
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
