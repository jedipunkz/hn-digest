---
source: "https://code.claude.com/docs/en/desktop-linux"
hn_url: "https://news.ycombinator.com/item?id=48734754"
title: "Claude Desktop is now available on Linux (in beta)"
article_title: "Claude Desktop on Linux (beta) - Claude Code Docs"
author: "adocomplete"
captured_at: "2026-06-30T16:43:56Z"
capture_tool: "hn-digest"
hn_id: 48734754
score: 4
comments: 1
posted_at: "2026-06-30T16:09:10Z"
tags:
  - hacker-news
  - translated
---

# Claude Desktop is now available on Linux (in beta)

- HN: [48734754](https://news.ycombinator.com/item?id=48734754)
- Source: [code.claude.com](https://code.claude.com/docs/en/desktop-linux)
- Score: 4
- Comments: 1
- Posted: 2026-06-30T16:09:10Z

## Translation

タイトル: Claude デスクトップが Linux で利用可能になりました (ベータ版)
記事のタイトル: Linux 上の Claude デスクトップ (ベータ版) - Claude Code Docs
説明: Ubuntu および Debian に Claude デスクトップ アプリをインストールして更新します。

記事本文:
Linux 上の Claude デスクトップ (ベータ版) - Claude Code Docs Documentation Index
/docs/llms.txt で完全なドキュメントのインデックスを取得します。
さらに探索する前に、このファイルを使用して利用可能なすべてのページを検出します。
メイン コンテンツにスキップ Claude Code Docs ホーム ページ 英語 検索... ⌘ K Ask Assistant Claude 開発者プラットフォーム
検索... ナビゲーション デスクトップ上の Claude Code Linux 上の Claude デスクトップ (ベータ版) はじめに Claude Code による構築 管理 構成リファレンス エージェント SDK 新機能 リソース はじめに
指示と記憶を保存する
デスクトップ上の Claude Code
インストール ダウンロードしたファイルからインストールします
Linux ベータ版にまだ含まれていないもの
Linux 上のクロード デスクトップ (ベータ版)
ページをコピーする Ubuntu および Debian に Claude デスクトップ アプリをインストールして更新する
ページをコピーする Claude デスクトップ アプリの Linux サポートはベータ版です。 [チャット]、[共同作業]、および [コード] タブはすべて使用できます。
Linux 上のデスクトップ アプリでは、macOS や Windows と同じチャット、コワーク、クロード コードのエクスペリエンス (並列セッション、視覚的な差分レビュー、統合されたターミナルとエディター、ライブ アプリ プレビュー) が提供されます。完全な機能のリファレンスについては、「Claude Code Desktop の使用」を参照してください。
要件
Ubuntu 22.04以降、またはDebian 12以降
Anthropic の apt リポジトリを追加する
sudoカール -fsSLo /usr/share/keyrings/claude-desktop-archive-keyring.asc https://downloads.claude.ai/claude-desktop/key.asc
リポジトリを登録します。 echo "deb [arch=amd64,arm64 signed-by=/usr/share/keyrings/claude-desktop-archive-keyring.asc] https://downloads.claude.ai/claude-desktop/apt/stablesteadal main" | sudo tee /etc/apt/sources.list.d/claude-desktop.list
2 パッケージをインストールする
sudo apt update && sudo apt install クロードデスクトップ
3 起動してサインインする
gpg --show-keys /usr/share/keyrings/claude-desktop-archive-keyring.asc
指紋は 31DD である必要があります

DE24 DDFA B679 F42D 7BD2 BAA9 29FF 1A7E CACE 。
ダウンロードしたファイルからインストールする
sudo apt install ./claude-desktop_ * .deb
この方法でインストールされた .deb は更新を受け取りません。 apt を通じて更新を取得するには、上記のようにリポジトリを追加するか、パッケージが /etc/apt/sources.list.d/claude-desktop.list に書き込むプレースホルダー エントリ内の deb 行のコメントを解除します。
アップデート
sudo apt アップデート && sudo apt アップグレード
ディストリビューションのグラフィカル ソフトウェア アップデーターも新しいバージョンを取得します。
アンインストール
sudo apt はクロードデスクトップを削除します
これにより、アプリとともに署名キーが削除されるため、インストール中にリポジトリ エントリを追加した場合は、それも削除します。
sudo rm /etc/apt/sources.list.d/claude-desktop.list
Linux ベータ版にまだ含まれていないもの
コンピューターの使用: アプリと画面の制御は Linux では利用できません。
ディクテーション : Linux デスクトップ アプリでは音声入力は利用できません。代わりに CLI で音声ディクテーションを使用してください。
クイック エントリ グローバル ホットキー : X11 で動作します。ネイティブ Wayland では、デスクトップ環境の GlobalShortcuts ポータルが必要です。
Fedora および RHEL : 現在、Debian ベースのディストリビューションのみがサポートされています。将来的には、追加のディストリビューションもサポートされる予定です。
はい いいえ 参照 スケジュールされたタスク ⌘ I Claude Code Docs ホーム ページ x linkedin Company

## Original Extract

Install and update the Claude desktop app on Ubuntu and Debian

Claude Desktop on Linux (beta) - Claude Code Docs Documentation Index
Fetch the complete documentation index at: /docs/llms.txt
Use this file to discover all available pages before exploring further.
Skip to main content Claude Code Docs home page English Search... ⌘ K Ask Assistant Claude Developer Platform
Search... Navigation Claude Code on desktop Claude Desktop on Linux (beta) Getting started Build with Claude Code Administration Configuration Reference Agent SDK What's New Resources Getting started
Store instructions and memories
Claude Code on desktop Get started
Install Install from a downloaded file
What’s not in the Linux beta yet
Claude Desktop on Linux (beta)
Copy page Install and update the Claude desktop app on Ubuntu and Debian
Copy page Linux support for the Claude desktop app is in beta. The Chat, Cowork, and Code tabs are all available.
The desktop app on Linux gives you the same Chat, Cowork, and Claude Code experience as macOS and Windows: parallel sessions, visual diff review, an integrated terminal and editor, and live app preview. See Use Claude Code Desktop for the full feature reference.
​ Requirements
Ubuntu 22.04 or later, or Debian 12 or later
Add Anthropic's apt repository
sudo curl -fsSLo /usr/share/keyrings/claude-desktop-archive-keyring.asc https://downloads.claude.ai/claude-desktop/key.asc
Register the repository: echo "deb [arch=amd64,arm64 signed-by=/usr/share/keyrings/claude-desktop-archive-keyring.asc] https://downloads.claude.ai/claude-desktop/apt/stable stable main" | sudo tee /etc/apt/sources.list.d/claude-desktop.list
2 Install the package
sudo apt update && sudo apt install claude-desktop
3 Launch and sign in
gpg --show-keys /usr/share/keyrings/claude-desktop-archive-keyring.asc
The fingerprint should be 31DD DE24 DDFA B679 F42D 7BD2 BAA9 29FF 1A7E CACE .
​ Install from a downloaded file
sudo apt install ./claude-desktop_ * .deb
A .deb installed this way doesn’t receive updates. To get updates through apt, add the repository as shown above, or uncomment the deb line in the placeholder entry the package writes to /etc/apt/sources.list.d/claude-desktop.list .
​ Update
sudo apt update && sudo apt upgrade
Your distribution’s graphical software updater will also pick up new versions.
​ Uninstall
sudo apt remove claude-desktop
This removes the signing key along with the app, so if you added the repository entry during install, remove it too:
sudo rm /etc/apt/sources.list.d/claude-desktop.list
​ What’s not in the Linux beta yet
Computer Use : app and screen control isn’t available on Linux.
Dictation : voice input isn’t available in the Linux desktop app. Use voice dictation in the CLI instead.
Quick Entry global hotkey : works on X11. On native Wayland it requires your desktop environment’s GlobalShortcuts portal.
Fedora and RHEL : only Debian-based distributions are supported today. Support for additional distributions is coming in the future.
Yes No Reference Scheduled tasks ⌘ I Claude Code Docs home page x linkedin Company
