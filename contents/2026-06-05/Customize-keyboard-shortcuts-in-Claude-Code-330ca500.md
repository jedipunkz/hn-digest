---
source: "https://code.claude.com/docs/en/keybindings"
hn_url: "https://news.ycombinator.com/item?id=48410022"
title: "Customize keyboard shortcuts in Claude Code"
article_title: "Customize keyboard shortcuts - Claude Code Docs"
author: "ankitg12"
captured_at: "2026-06-05T10:25:30Z"
capture_tool: "hn-digest"
hn_id: 48410022
score: 1
comments: 0
posted_at: "2026-06-05T09:18:39Z"
tags:
  - hacker-news
  - translated
---

# Customize keyboard shortcuts in Claude Code

- HN: [48410022](https://news.ycombinator.com/item?id=48410022)
- Source: [code.claude.com](https://code.claude.com/docs/en/keybindings)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T09:18:39Z

## Translation

タイトル: Claude Code でキーボード ショートカットをカスタマイズする
記事のタイトル: キーボード ショートカットのカスタマイズ - Claude Code Docs
説明: キーバインド構成ファイルを使用して、Claude Code のキーボード ショートカットをカスタマイズします。

記事本文:
キーボード ショートカットのカスタマイズ - Claude Code Docs Documentation Index
/docs/llms.txt で完全なドキュメントのインデックスを取得します。
さらに探索する前に、このファイルを使用して利用可能なすべてのページを検出します。
メイン コンテンツにスキップ Claude Code Docs ホーム ページ 英語 検索... ⌘ K Ask Assistant Claude 開発者プラットフォーム
検索... ナビゲーション インターフェイス キーボード ショートカットのカスタマイズ はじめに クロード コードによる構築 管理 構成リファレンス エージェント SDK 新機能 リソース 設定と権限
高速モードで応答を高速化
ページをコピー キーバインド構成ファイルを使用して、Claude Code のキーボード ショートカットをカスタマイズします。
ページをコピー カスタマイズ可能なキーボード ショートカットには、Claude Code v2.1.18 以降が必要です。 claude --version でバージョンを確認します。
Claude Code は、カスタマイズ可能なキーボード ショートカットをサポートしています。 /keybindings を実行して、 ~/.claude/keybindings.json に構成ファイルを作成するか開きます。
設定ファイル
キーバインド ファイルへの変更は自動的に検出され、Claude Code を再起動しなくても適用されます。
フィールド 説明 $schema エディターのオートコンプリート用のオプションの JSON スキーマ URL $docs オプションのドキュメント URL バインディング コンテキスト別のバインディング ブロックの配列
この例では、 Ctrl+E をバインドしてチャット コンテキストで外部エディタを開き、 Ctrl+U のバインドを解除します。
{
"$schema" : "https://www.schemastore.org/claude-code-keybindings.json" ,
"$docs" : "https://code.claude.com/docs/en/keybindings" ,
「バインディング」: [
{
"context" : "チャット" ,
"バインディング" : {
"ctrl+e" : "チャット:外部エディタ" ,
"ctrl+u" : null
}
}
]
}
コンテキスト
alt 、 opt 、 option 、または meta - Windows および Linux では Alt キー、macOS では Option キー
cmd 、 command 、 super 、または win - macOS では Command キー、Windows では Windows キー、Linux では Super キー
Ctrl+K Ctrl+K
Shift+Tab Shift+Tab
meta+p macOS では Option + P、Alt + P e

どこか
ctrl+shift+c 複数の修飾子
大文字
Ctrl+K Ctrl+S Ctrl+K を押して放し、Ctrl+S を押します。
特殊キー
上、下、左、右 - 矢印キー
backspace 、 delete - 削除キー
{
「バインディング」: [
{
"context" : "チャット" ,
"バインディング" : {
"ctrl+s" : null
}
}
]
}
これはコードバインディングにも機能します。プレフィックスを共有するすべてのコードのバインドを解除すると、そのプレフィックスが解放されて単一キー バインディングとして使用できるようになります。
{
「バインディング」: [
{
"context" : "チャット" ,
"バインディング" : {
"ctrl+x ctrl+k" : null 、
"ctrl+x ctrl+e" : null 、
"ctrl+x" : "チャット:改行"
}
}
]
}
プレフィックス上のすべてではなく一部のコードのバインドを解除した場合でも、プレフィックスを押すと、残りのバインディングに対してコード待機モードに入ります。
予約されたショートカット
Vim モードはテキスト入力レベル (カーソルの移動、モード、モーション) で入力を処理します。
キーバインドは、コンポーネント レベルでアクションを処理します (ToDo の切り替え、送信など)。
vim モードで Escape キーを押すと、INSERT が NORMAL モードに切り替わります。チャットはトリガーされません:キャンセル
ほとんどの Ctrl+キー ショートカットは vim モードを介してキーバインド システムに渡されます。
vim NORMAL モードでは、?ヘルプメニューを表示します (vim の動作)
vim NORMAL モードでは、標準モードの Ctrl+R と同じように、/ で履歴検索が開きます。
解析エラー (無効な JSON または構造)
ターミナルマルチプレクサの競合
同じコンテキスト内でバインディングが重複する
はい いいえ ステータス行をカスタマイズします ⌘ I Claude Code Docs ホーム ページ x linkedin Company

## Original Extract

Customize keyboard shortcuts in Claude Code with a keybindings configuration file.

Customize keyboard shortcuts - Claude Code Docs Documentation Index
Fetch the complete documentation index at: /docs/llms.txt
Use this file to discover all available pages before exploring further.
Skip to main content Claude Code Docs home page English Search... ⌘ K Ask Assistant Claude Developer Platform
Search... Navigation Interface Customize keyboard shortcuts Getting started Build with Claude Code Administration Configuration Reference Agent SDK What's New Resources Settings and permissions
Speed up responses with fast mode
Copy page Customize keyboard shortcuts in Claude Code with a keybindings configuration file.
Copy page Customizable keyboard shortcuts require Claude Code v2.1.18 or later. Check your version with claude --version .
Claude Code supports customizable keyboard shortcuts. Run /keybindings to create or open your configuration file at ~/.claude/keybindings.json .
​ Configuration file
Changes to the keybindings file are automatically detected and applied without restarting Claude Code.
Field Description $schema Optional JSON Schema URL for editor autocompletion $docs Optional documentation URL bindings Array of binding blocks by context
This example binds Ctrl+E to open an external editor in the chat context, and unbinds Ctrl+U :
{
"$schema" : "https://www.schemastore.org/claude-code-keybindings.json" ,
"$docs" : "https://code.claude.com/docs/en/keybindings" ,
"bindings" : [
{
"context" : "Chat" ,
"bindings" : {
"ctrl+e" : "chat:externalEditor" ,
"ctrl+u" : null
}
}
]
}
​ Contexts
alt , opt , option , or meta - Alt key on Windows and Linux, Option key on macOS
cmd , command , super , or win - Command key on macOS, Windows key on Windows, Super key on Linux
ctrl+k Ctrl + K
shift+tab Shift + Tab
meta+p Option + P on macOS, Alt + P elsewhere
ctrl+shift+c Multiple modifiers
​ Uppercase letters
ctrl+k ctrl+s Press Ctrl+K, release, then Ctrl+S
​ Special keys
up , down , left , right - Arrow keys
backspace , delete - Delete keys
{
"bindings" : [
{
"context" : "Chat" ,
"bindings" : {
"ctrl+s" : null
}
}
]
}
This also works for chord bindings. Unbinding every chord that shares a prefix frees that prefix for use as a single-key binding:
{
"bindings" : [
{
"context" : "Chat" ,
"bindings" : {
"ctrl+x ctrl+k" : null ,
"ctrl+x ctrl+e" : null ,
"ctrl+x" : "chat:newline"
}
}
]
}
If you unbind some but not all chords on a prefix, pressing the prefix still enters chord-wait mode for the remaining bindings.
​ Reserved shortcuts
Vim mode handles input at the text input level (cursor movement, modes, motions)
Keybindings handle actions at the component level (toggle todos, submit, etc.)
The Escape key in vim mode switches INSERT to NORMAL mode; it does not trigger chat:cancel
Most Ctrl+key shortcuts pass through vim mode to the keybinding system
In vim NORMAL mode, ? shows the help menu (vim behavior)
In vim NORMAL mode, / opens history search, the same as Ctrl+R in standard mode
Parse errors (invalid JSON or structure)
Terminal multiplexer conflicts
Duplicate bindings in the same context
Yes No Customize status line ⌘ I Claude Code Docs home page x linkedin Company
