---
source: "https://github.com/makethisbetter/makethisbetter-js"
hn_url: "https://news.ycombinator.com/item?id=49181691"
title: "Show HN: Feedback widget with screenshots, annotations, and AI triage"
article_title: "GitHub - makethisbetter/makethisbetter-js: Feedback widget with screenshots, annotations, and AI triage · GitHub"
author: "ken2049"
captured_at: "2026-08-05T12:48:50Z"
capture_tool: "hn-digest"
hn_id: 49181691
score: 3
comments: 0
posted_at: "2026-08-05T12:07:42Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Feedback widget with screenshots, annotations, and AI triage

- HN: [49181691](https://news.ycombinator.com/item?id=49181691)
- Source: [github.com](https://github.com/makethisbetter/makethisbetter-js)
- Score: 3
- Comments: 0
- Posted: 2026-08-05T12:07:42Z

## Translation

タイトル: HN を表示: スクリーンショット、注釈、AI トリアージを備えたフィードバック ウィジェット
記事のタイトル: GitHub - makethisbetter/makethisbetter-js: スクリーンショット、注釈、AI トリアージを備えたフィードバック ウィジェット · GitHub
説明: スクリーンショット、注釈、AI トリアージを備えたフィードバック ウィジェット - makethisbetter/makethisbetter-js

記事本文:
GitHub - makethisbetter/makethisbetter-js: スクリーンショット、注釈、AI トリアージを備えたフィードバック ウィジェット · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
これをより良くする
/
makethisbetter-js
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット アセット アセット デモ デモ e2e e2e スクリプト スクリプト src src テスト test .gitignore .gitignore LICE

NSE ライセンス README.md README.md SELF_HOSTING.md SELF_HOSTING.md package-lock.json package-lock.json package.json package.json test-install-methods.sh test-install-methods.sh tsconfig.json tsconfig.json vite.config.ts vite.config.ts vite.iife.config.ts vite.iife.config.ts vite.standalone.config.ts vite.standalone.config.ts vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
スクリーンショット、注釈、AI トリアージを備えたフィードバック ウィジェット。
AI エージェントが付属しています。ユーザーは、開発では決して目にすることのないバグに遭遇します。彼らは去ります。その理由は決して分かりません。
このウィジェットを使用すると、ユーザーは 2 回のクリックで、何が問題になったのか (注釈付きスクリーンショット、コンソール エラー、DOM 状態、ブラウザ情報) を正確に報告することができます。 AI トリアージは、コーディング エージェント (クロード コード、カーソル、コーデックス) が自動的に選択する構造化タスクに変換します。エージェントが修正を配布します。ユーザーに通知が届きます。
「何が起こったのか説明してもらえますか？」という質問はもう必要ありません。 Slack でスクリーンショットを失うことはもうありません。不満を抱いたユーザーから配布された修正に至るまでの完全なループは、コンテキストを切り替えることなく実行されます。
< スクリプト src =" https://unpkg.com/makethisbetter@1 " > </ スクリプト >
< スクリプト >
これをより良くする。 init ( { projectKey : 'mtb_proj_YOUR_KEY' } )
</ スクリプト >
npm
npm インストールでこれを改善する
import { MakeThisBetter } from 'makethisbetter'
これをより良くする。 init ( { projectKey : 'mtb_proj_YOUR_KEY' } )
それだけです。ページにフィードバック タブが表示されます。
ユーザーがフィードバック タブをクリックする
-> 問題に注釈を付けます (クリックして固定、ドラッグして描画)
-> コメントを追加します
-> 送信する
+-- 自動的にキャプチャされたスクリーンショット
+-- コンソールエラーが収集されました
+-- アセンブルされたページ コンテキスト (URL、ブラウザ、OS、セレクター)
+-- この API を改善するために送信されました
+-- AI は必要に応じて明確な質問をします
-> ダッシュボードには構造化されたフィードバックが表示されます
-> AI トリアージがエージェント対応のタスクを生成
-> あなたのコーディング年齢

nt がそれを受け取り、修正を発送します
-> ユーザーに通知が届きます: 修正は有効です
フレームワークガイド
React / Next.js
// app/providers.tsx (アプリルーター) または pages/_app.tsx (ページルーター)
「クライアントを使用」
「反応」からインポート { useEffect }
エクスポート関数 FeedbackProvider ({ user } : { user ?: { id : string , email ?: string } } ) {
useEffect ( ( ) => {
import ( 'makethisbetter' ) 。 then ( ( { MakeThisBetter } ) => {
これをより良くする。初期化 ( {
projectKey : プロセス。環境 。 NEXT_PUBLIC_MTB_KEY ! 、
ユーザー
})
})
return() => {
import ( 'makethisbetter' ) 。 then ( ( { MakeThisBetter } ) => MakeThisBetter . destroy ( ) )
}
} , [ ユーザー ] )
nullを返す
}
Vue / Nuxt
// plugins/makethisbetter.client.ts (Nuxt) または main.ts (Vue)
import { MakeThisBetter } from 'makethisbetter'
デフォルトのdefineNuxtPluginをエクスポート ( ( ) => {
これをより良くする。初期化 ( {
projectKey : useRuntimeConfig() 。パブリック。 mtbキー 、
})
戻り値 {
提供: { mtbDestroy:() => MakeThisBetter 。破壊する ( ) }
}
})
アストロ
<!-- src/components/Feedback.astro -->
< スクリプト >
import { MakeThisBetter } from 'makethisbetter'
MakeThisBetter.init({ projectKey: import.meta.env.PUBLIC_MTB_KEY })
</ スクリプト >
レール
<%# app/views/layouts/application.html.erb %>
<本体>
<%= 収率 %>
<%# Turbo は、アクセスするたびに <body> を置き換え、スクリプトに追加された要素を取得します
それと一緒に。ホストを自分でレンダリングして永続的なマークを付けるか、ウィジェットが
各ナビゲーションや飛行中の何かの後に再構築される - 書きかけの
レポート、画面記録が失われます。 %>
< div id =" mtb-widget-host " data-turbo-permanent > </ div >
<% if current_user &.管理者？ %>
< スクリプト src =" https://unpkg.com/makethisbetter@1 " > </ スクリプト >
< スクリプト >
これをより良くする。初期化 ( {
projectKey : ' <%= Rails 。アプリケーションの資格情報 。 mtb_project_key %> ' ,
ユーザー: { id : ' <%= current_user 。 ID %>

' 、電子メール: ' <%= current_user 。メール %> ' }
})
</ スクリプト >
<% 終了 %>
</本文>
ホストは <body> 内に存在する必要があります — 永続要素を内部の ID でターボペアリングします
ボディのスナップショットなので、<head> に配置されたスナップショットは決して一致せず、何もありません
が起こります。 ID と属性の両方が必要です。どちらか一人で
[切り捨てられた]
これが必要なのはターボ駆動アプリのみです。 React、Vue、Svelte ルーターの内部で再レンダリング
独自のコンテナを使用し、 <body> を置き換えることはありません。そのため、ホストは単独で存続します。
< スクリプト src =" https://unpkg.com/makethisbetter@1 " > </ スクリプト >
< スクリプト >
これをより良くする。 init ( { projectKey : 'mtb_proj_YOUR_KEY' } )
</ スクリプト >
特長
任意の要素をクリックして固定するか、ドラッグしてフリーフォーム ハイライトを描画します。 SDK は、要素の CSS セレクター、テキスト コンテンツ、および位置をキャプチャします。
ツールバーでリプレイ モードに切り替えて、インタラクション リプレイ (最大 60 秒) をキャプチャします。 rrweb DOM の突然変異と相互作用イベントを記録します。画面のビデオや音声はキャプチャされず、ブラウザーのメディア許可も要求されません。レコーダーは遅延ロードされるため、レポーターが再生を開始するまでコストはかかりません。
機密データのフィルタリングは固定されており、SDK 構成を通じて無効にすることはできません。検索クエリ、問題の説明、内部フォーム フィールドなどの通常の値は、問題を再現するために必要になることが多いため、引き続き使用できます。
要素を除外します。 SDK がキャプチャしてはいけないページ領域に rrweb のプライバシー クラスを追加します。これらのクラスは、インタラクション リプレイ、クリックと入力のパンくずリスト、注釈メタデータ、スクリーンショットに一貫して適用されます。
class="rr-block" — フットプリントを維持しながら、マークされた領域全体を非表示にします。
class="rr-mask" — 周囲のレイアウトを維持しながら、テキストとフォーム コントロールのコンテンツを非表示にします。
< div class =" rr-block " > <!-- リプレイには表示されません -->

</ div >
<span class="rr-mask">口座残高:$12,400 </span>
スクリーンショットまたはリプレイ フィルタリングが完了できない場合、その添付ファイルはサイレントに省略され、テキスト フィードバックは引き続き送信されます。自動フィルタリングではサイト固有のシークレットをすべて識別できないため、機密アプリケーション領域では rr-block または rr-mask を使用します。
SDK は、ユーザーが苦労しているというシグナルを監視し、積極的にフィードバックを収集するよう提案します。
fristrationDetection: false で無効にします。
提出前に、AI アシスタントは本当のニーズを明確にするために 1 つの短いフォローアップの質問をすることがあります。ユーザーが実際の問題ではなく、試みた解決策を説明する XY 問題を回避します。交換が完了すると、ウィジェットは自動的にフィードバックを送信します。
すべての提出には次のものが自動的に含まれます。
ページ URL の起点とパス名、ブラウザ、OS、画面解像度
エラーの種類、スクリプトのパス名、行/列の位置 ( window.onerror および window.onunhandledrejection 経由)
ターゲット要素セレクターとテキスト
アップロード前にプライバシー カバーが適用された注釈付きスクリーンショット (html-to-image 経由)
注釈の座標と描画パス
7 言語の組み込みサポート:
これをより良くする。初期化 ( {
// 必須
プロジェクトキー: 'mtb_proj_xxx' 、
// オプション
locale : 'en' , // UI 言語。未設定: <html lang> に戻り、その後 'en' に戻ります。
Position : 'right' , // タブ位置: 'left' | 「そうだね」
tabText : 'Feedback' , // ドッキングされたタブのラベル。未設定: ロケール独自の表現
brandColors : { // 完全なウィジェットのオプションのセマンティックカラー
プライマリ: '#2563eb' 、
ホバー: '#1d4ed8' 、
アクティブ: '#1e40af' 、
onPrimary : '#ffffff' 、
} 、
tabColor : '#2563eb' , // 従来のランチャーのみの色。 brandColors が有効な場合は無視されます
entryMode : 'button' , // 'button' はタブをドッキングします | 「api」は何もレンダリングしません
テーマ : 'auto' , // 'ライト' | '暗い' | 「自動」
イライラ

ationDetection : true , // プロアクティブな不満プロンプト
apiUrl : 'https://...' , // セルフホスト API エンドポイント
// ユーザー ID (推奨)。
// 有効な userToken/userTokenFn JWT が存在する場合は完全に無視されます —
// 以下の本人確認を参照してください。
ユーザー: {
id : 'usr_123' 、
電子メール: 'alex@example.com' 、
名前：「アレックス・チェン」、
} 、
})
ウィジェットのブランディング
brandColors を使用して製品のセマンティック カラーをランチャー全体に適用します。
注釈ツール、フォーカスおよび選択状態、行動喚起、レポーターバブル、
そしてAIによる装飾。 4 つの値すべてを 6 桁の 16 進数の色として指定します。 SDKが使用するのは、
これらは提供されたとおりに正確に提供され、カラー スケールは生成されません。
これをより良くする。初期化 ( {
プロジェクトキー: 'mtb_proj_xxx' 、
tabText : '問題を報告する' ,
ブランドカラー: {
プライマリ: '#2563eb' 、
ホバー: '#1d4ed8' 、
アクティブ: '#1e40af' 、
onPrimary : '#ffffff' 、
} 、
})
成功、エラー、警告、記録の色は、それぞれの状態の意味を保持します。アン
不完全なグループまたは #RRGGBB 以外の値は完全なグループを拒否します
デフォルトのウィジェットの色をそのままにします。
tabColor は、既存のインストールおよび次の製品で引き続き使用できます。
1つのブランドカラーのみを公開します。 6 桁の 16 進数の色を 1 つ受け入れ、のみに影響します。
ドッキングされたランチャー、およびそのランチャーのホバー、アクティブ、およびフォアグラウンドを取得します。
色。両方のオプションが存在し、brandColors が有効な場合、tabColor は
無視されました。 「Make This Better」を緑色のままにするには、両方のオプションを省略します。
/makethisbetter セットアップ スキルは、brandColors が見つかった場合にのみ brandColors を推奨します。
デザイン システム内の完全なセマンティック グループ。プライマリのみが見つかった場合
color の場合は、代わりに tabColor を提供します。欠けている色合いを推測することはありません。
ロケールは、渡したロケール、ページの順序で一度解決されます。
<html lang> 属性、次に en 。完全に一致しないタグは次のように再試行されます。

アウト
そのリージョン ( fr-CA → fr )、それでも一致しないものは en に戻ります。
実行時の言語の変更
MakeThisBetter.setLocale('zh-CN') は、タブと
その後何かが開いた。すでに画面上にあるポップアップでは言語が維持されます
で開かれたため、レポーターが文の途中で再レンダリングされることはありません。の前に呼び出します
レポーターはウィジェットを開きます。たとえば、アプリがウィジェットを適用するのと同じ場所で、
言語の変更。
これをより良くする。 setLocale ( 'zh-CN' )
本人確認
ID 検証により、認証されたユーザーにフィードバックがリンクされ、ユーザーはフィードバック ボードで自分の投稿を確認できるようになります。
レベル 0 -- 匿名 (デフォルト): ユーザー トークンなし。フィードバックは匿名です。
これをより良くする。 init ( { projectKey : 'mtb_proj_xxx' } )
匿名の記者にはフォローアップが提供されます。成功カードにはオプションの回答が表示されます。
電子メール フィールドに入力されたアドレスがレポーター エンドポイントに送信され、
mtb_reporter_email の下の localStorage に保存されるため、フィールドは要求されません
同じブラウザからのその後のレポートでもう一度説明します。次の場合、フィールドは完全にスキップされます。
ユーザーが設定されている場合、または JWT がレポーターをすでに識別している場合。サイトデータの消去
それをクリアします。
レベル 1 -- 静的トークン: 事前に生成された JWT を渡します。シンプルですが、

[切り捨てられた]

## Original Extract

Feedback widget with screenshots, annotations, and AI triage - makethisbetter/makethisbetter-js

GitHub - makethisbetter/makethisbetter-js: Feedback widget with screenshots, annotations, and AI triage · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
makethisbetter
/
makethisbetter-js
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit assets assets demo demo e2e e2e scripts scripts src src test test .gitignore .gitignore LICENSE LICENSE README.md README.md SELF_HOSTING.md SELF_HOSTING.md package-lock.json package-lock.json package.json package.json test-install-methods.sh test-install-methods.sh tsconfig.json tsconfig.json vite.config.ts vite.config.ts vite.iife.config.ts vite.iife.config.ts vite.standalone.config.ts vite.standalone.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
Feedback widget with screenshots, annotations, and AI triage.
You ship with AI agents. Your users hit bugs you never see in dev. They leave. You never find out why.
This widget gives your users a way to report exactly what went wrong — annotated screenshot, console errors, DOM state, browser info — in two clicks. AI triage turns that into a structured task your coding agent (Claude Code, Cursor, Codex) picks up automatically. The agent ships the fix. The user gets notified.
No more "can you describe what happened?" No more lost screenshots in Slack. The full loop, from frustrated user to shipped fix, runs without you context-switching.
< script src =" https://unpkg.com/makethisbetter@1 " > </ script >
< script >
MakeThisBetter . init ( { projectKey : 'mtb_proj_YOUR_KEY' } )
</ script >
npm
npm install makethisbetter
import { MakeThisBetter } from 'makethisbetter'
MakeThisBetter . init ( { projectKey : 'mtb_proj_YOUR_KEY' } )
That's it. A feedback tab appears on your page.
User clicks feedback tab
-> Annotates the problem (click to pin, drag to draw)
-> Adds a comment
-> Submits
+-- Screenshot captured automatically
+-- Console errors collected
+-- Page context assembled (URL, browser, OS, selectors)
+-- Sent to Make This Better API
+-- AI asks a clarifying question if needed
-> Dashboard shows structured feedback
-> AI triage produces an agent-ready task
-> Your coding agent picks it up and ships the fix
-> User gets notified: the fix is live
Framework Guides
React / Next.js
// app/providers.tsx (App Router) or pages/_app.tsx (Pages Router)
'use client'
import { useEffect } from 'react'
export function FeedbackProvider ( { user } : { user ?: { id : string , email ?: string } } ) {
useEffect ( ( ) => {
import ( 'makethisbetter' ) . then ( ( { MakeThisBetter } ) => {
MakeThisBetter . init ( {
projectKey : process . env . NEXT_PUBLIC_MTB_KEY ! ,
user
} )
} )
return ( ) => {
import ( 'makethisbetter' ) . then ( ( { MakeThisBetter } ) => MakeThisBetter . destroy ( ) )
}
} , [ user ] )
return null
}
Vue / Nuxt
// plugins/makethisbetter.client.ts (Nuxt) or main.ts (Vue)
import { MakeThisBetter } from 'makethisbetter'
export default defineNuxtPlugin ( ( ) => {
MakeThisBetter . init ( {
projectKey : useRuntimeConfig ( ) . public . mtbKey ,
} )
return {
provide : { mtbDestroy : ( ) => MakeThisBetter . destroy ( ) }
}
} )
Astro
<!-- src/components/Feedback.astro -->
< script >
import { MakeThisBetter } from 'makethisbetter'
MakeThisBetter.init({ projectKey: import.meta.env.PUBLIC_MTB_KEY })
</ script >
Rails
<%# app/views/layouts/application.html.erb %>
< body >
<%= yield %>
<%# Turbo replaces <body> on every visit, taking any script-appended element
with it. Render the host yourself and mark it permanent, or the widget is
rebuilt after each navigation and anything mid-flight — a half-written
report, a screen recording — is lost. %>
< div id =" mtb-widget-host " data-turbo-permanent > </ div >
<% if current_user &. admin? %>
< script src =" https://unpkg.com/makethisbetter@1 " > </ script >
< script >
MakeThisBetter . init ( {
projectKey : ' <%= Rails . application . credentials . mtb_project_key %> ' ,
user : { id : ' <%= current_user . id %> ' , email : ' <%= current_user . email %> ' }
} )
</ script >
<% end %>
</ body >
The host must sit inside <body> — Turbo pairs permanent elements by id within
the body snapshot, so one placed in <head> is never matched and nothing
happens. It needs both the id and the attribute; either alone
[truncated]
Only Turbo-driven apps need this. React, Vue and Svelte routers re-render inside
their own container and never replace <body> , so the host survives on its own.
< script src =" https://unpkg.com/makethisbetter@1 " > </ script >
< script >
MakeThisBetter . init ( { projectKey : 'mtb_proj_YOUR_KEY' } )
</ script >
Features
Click any element to pin it, or drag to draw a freeform highlight. The SDK captures the element's CSS selector, text content, and position.
Switch to Replay mode in the toolbar to capture an Interaction Replay (up to 60 seconds). It records rrweb DOM mutations and interaction events. It does not capture screen video or audio and does not request browser media permissions. The recorder loads lazily, so there is zero cost until the reporter starts a replay.
Sensitive-data filtering is fixed and cannot be disabled through SDK configuration. Ordinary values such as search queries, issue descriptions, and internal form fields remain available because they are often necessary to reproduce a problem.
Excluding an element. Add rrweb's privacy classes to any page region that the SDK must not capture. These classes apply consistently to Interaction Replay, click and input breadcrumbs, annotation metadata, and screenshots:
class="rr-block" — hides the entire marked region while preserving its footprint.
class="rr-mask" — hides text and form-control content while preserving the surrounding layout.
< div class =" rr-block " > <!-- never appears in a replay --> </ div >
< span class =" rr-mask " > Account balance: $12,400 </ span >
If screenshot or Replay filtering cannot complete, that attachment is silently omitted and the text feedback still submits. Automated filtering cannot identify every site-specific secret, so use rr-block or rr-mask on sensitive application regions.
The SDK watches for signals that a user is struggling and proactively offers to collect feedback:
Disable with frustrationDetection: false .
Before submission, an AI assistant may ask one short follow-up question to clarify the real need — avoiding XY problems where users describe their attempted solution instead of the actual problem. Once the exchange is complete, the widget submits the feedback automatically.
Every submission automatically includes:
Page URL origin and pathname, browser, OS, screen resolution
Error type, script pathname, and line/column location (via window.onerror and window.onunhandledrejection )
Target element selector and text
Annotated screenshot with privacy covers applied before upload (via html-to-image )
Annotation coordinates and draw paths
Built-in support for 7 languages:
MakeThisBetter . init ( {
// Required
projectKey : 'mtb_proj_xxx' ,
// Optional
locale : 'en' , // UI language. Unset: falls back to <html lang>, then 'en'
position : 'right' , // Tab position: 'left' | 'right'
tabText : 'Feedback' , // Label on the docked tab. Unset: the locale's own wording
brandColors : { // Optional semantic colors for the complete Widget
primary : '#2563eb' ,
hover : '#1d4ed8' ,
active : '#1e40af' ,
onPrimary : '#ffffff' ,
} ,
tabColor : '#2563eb' , // Legacy launcher-only color; ignored when brandColors is valid
entryMode : 'button' , // 'button' docks a tab | 'api' renders none
theme : 'auto' , // 'light' | 'dark' | 'auto'
frustrationDetection : true , // Proactive frustration prompts
apiUrl : 'https://...' , // Self-hosted API endpoint
// User identification (recommended).
// Ignored entirely when a valid userToken/userTokenFn JWT is present —
// see Identity Verification below.
user : {
id : 'usr_123' ,
email : 'alex@example.com' ,
name : 'Alex Chen' ,
} ,
} )
Widget branding
Use brandColors to apply your product's semantic colors across the launcher,
annotation tools, focus and selection states, calls to action, Reporter bubbles,
and AI decoration. Supply all four values as six-digit hex colors. The SDK uses
them exactly as provided and does not generate a color scale.
MakeThisBetter . init ( {
projectKey : 'mtb_proj_xxx' ,
tabText : 'Report a problem' ,
brandColors : {
primary : '#2563eb' ,
hover : '#1d4ed8' ,
active : '#1e40af' ,
onPrimary : '#ffffff' ,
} ,
} )
Success, error, warning, and recording colors keep their state meanings. An
incomplete group or any value that is not #RRGGBB rejects the complete group
and leaves the default Widget colors in place.
tabColor remains available for existing installations and for products that
only expose one brand color. It accepts one six-digit hex color, affects only
the docked launcher, and derives its launcher hover, active, and foreground
colors. When both options are present and brandColors is valid, tabColor is
ignored. Omit both options to keep the Make This Better green.
The /makethisbetter setup skill recommends brandColors only when it finds a
complete semantic group in your design system. When it finds only a primary
color, it offers tabColor instead; it never guesses the missing shades.
locale is resolved once, in this order: the locale you pass, then the page's
<html lang> attribute, then en . A tag with no exact match is retried without
its region ( fr-CA → fr ), and anything still unmatched falls back to en .
Changing the language at runtime
MakeThisBetter.setLocale('zh-CN') switches the language for the tab and for
anything opened afterwards. A popup that is already on screen keeps the language it
was opened in, so a reporter is never re-rendered mid-sentence. Call it before the
reporter opens the widget — for example, in the same place your app applies a
language change.
MakeThisBetter . setLocale ( 'zh-CN' )
Identity Verification
Identity verification links feedback to authenticated users and lets them view their own submissions on the feedback board.
Level 0 -- Anonymous (default): No user token. Feedback is anonymous.
MakeThisBetter . init ( { projectKey : 'mtb_proj_xxx' } )
Anonymous reporters are offered a follow-up: the success card shows an optional
email field, and an address entered there is sent to the reporter endpoint and
kept in localStorage under mtb_reporter_email so the field is not asked for
again on later reports from the same browser. The field is skipped entirely when
user is set or when a JWT already identifies the reporter. Clearing site data
clears it.
Level 1 -- Static token : Pass a pre-generated JWT. Simple, but t

[truncated]
