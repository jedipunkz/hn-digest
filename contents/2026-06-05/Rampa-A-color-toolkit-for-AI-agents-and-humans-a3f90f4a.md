---
source: "https://rampa.design/"
hn_url: "https://news.ycombinator.com/item?id=48411570"
title: "Rampa – A color toolkit for AI agents and humans"
article_title: "Rampa — A color toolkit for AI agents and humans"
author: "eustoria"
captured_at: "2026-06-05T13:08:53Z"
capture_tool: "hn-digest"
hn_id: 48411570
score: 1
comments: 0
posted_at: "2026-06-05T12:37:19Z"
tags:
  - hacker-news
  - translated
---

# Rampa – A color toolkit for AI agents and humans

- HN: [48411570](https://news.ycombinator.com/item?id=48411570)
- Source: [rampa.design](https://rampa.design/)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T12:37:19Z

## Translation

タイトル: Rampa – AI エージェントと人間のためのカラー ツールキット
記事のタイトル: Rampa — AI エージェントと人間のためのカラー ツールキット
説明: 1 つのツールでパレットの生成、コントラストの分析、色の混合、および任意の形式へのエクスポートを行います。プラグインや依存関係はありません。

記事本文:
AI エージェントと人間のためのカラー ツールキット
知覚的に均一で正確なカラー ランプを端末から直接生成します。 CLI、SDK、Web エディターが利用可能です。
ウェブエディター
ランパ
CLI + SDK
探索と統合
オクラホ/ラボ
知覚的に均一
APCA/WCAG
組み込みのコントラスト分析
スロープ
HSL の開始:終了値
色空間
カラーエッジを定義する
特長
必要なものすべて
プロダクションカラーシステム
1 つのツールでパレットの生成、コントラストの分析、色の混合、および任意の形式へのエクスポートを行います。プラグインや依存関係はありません。
明度、彩度、色相シフトを開始と終了の範囲として構成します。
相補的、三項的、その他の色相シフトのバリエーションを生成します。
不透明度を調整できるオーバーレイとして、完璧なブランドの色合いのニュートラルを取得します。
いくつかのサンプルカラーからフルカラーパレットを導き出すための強力なテーマオプション。
hex、rgb、hsl、または oklch に変換します。構造化された値を含む Text/JSON/CSS を返します。
w3c 仕様を使用して、任意の色のペアのコントラストを確認します。
端末から色を探索するか、SDK を使用して堅牢なカラー システムを構築します。すべての CLI フラグは、チェーン可能な SDK メソッドにマップされます。
AIエージェントに教える
色彩理論
Copilot、Claude Code、Codex、Cursor、およびスキルをサポートするエージェント用の 7 つのインストール可能なスキル。テーマの作成、ステータス カラー、データ Viz パレット、アクセス可能なコントラスト。
着色されたニュートラル/skill.md
# 着色されたニュートラル
ブランドカラーの微妙なヒントを伝えるニュートラルなパレットを作成します。純粋なグレーの代わりに、デザイン システムと一体感のある温かみのあるまたは冷たい中間色を選択してください。
## いつ使用するか
- 「私のブランドにマッチするグレーを作成してください」
- 「暖かい/冷たいニュートラルが欲しい」
- 「ニュートラルパレットをブランドカラーで染める」
## レシピ
### 1. ピュアニュートラル (彩度の低いブランド)
Rampa -C "<ブランド>" -L 98:5 -S 3:8 --size=10 -O css --name=neutral
### 2. 着色されたニュートラル
ブランドカラーを微妙なオーバーレイの色合いとして適用します。
ランパ -C "<br

そして>" -L 98:5 -S 3:8 --tint-color="<ブランド>" --tint-opacity=8 --tint-blend=overlay --size=10
アクセシブル-コントラスト/スキル.md
アクセス可能なコントラスト ペアの数
アクセシブルなテキストと背景の組み合わせ向けに設計されたカラー ランプを生成します。 11 段階のスケールを使用して、予測可能な WCAG 準拠のペアリングを実現します。
## 11 ステップの戦略
なぜ11ステップなのか？予測可能なペアリング計算:
最大 → 0 + 10 → 最高のコントラスト
AAA → 1 + 9 → 7:1+ 比率
AA → 3 + 7 → 4.5:1+ 比率
AA 大 → 4 + 6 → 3:1+ 比率
## レシピ
Rampa -C "<ブランド>" -L 98:5 --size=11 -O css --name=color
ペアは対称です: 0+10、1+9、2+8 など。
data-viz-palette/skill.md
# データ視覚化パレット
チャート、グラフ、ダッシュボードに最適化された色を生成します。色は最大限の区別と均等な視覚的重みを実現するために数学的に配置されています。
## 重要な原則
1. 明度の修正 - 視覚的な重みが等しい
2. 固定彩度 — 一貫した鮮やかさ
3. 最大色相間隔 — 調和を使用する
4. 各色につき 1 つのシェード — ランプではなく、明確な色合い
### 3 色 (トライアディック)
Rampa -C "<ブランド>" --add=triadic --size=2 -L 50:50 -S 70:70 -O css
### 4色(スクエア)
Rampa -C "<ブランド>" --add=square --size=2 -L 50:50 -S 70:70 -O css
アクセントからのステータス/skill.md
# アクセントからのステータスカラー
四角形調和を使用して、ブランドに数学的に関連する成功、警告、危険、および情報の色を生成します。
## スクエアハーモニーアプローチ
カラーホイール上の 90 度間隔の 4 色:
Rampa -C "<ブランド>" --add=square -L 95:15 --size=10 -O css
ベース → あなたのブランド
[切り捨てられた]
CLI を試して簡単に調べるか、SDK をインストールして堅牢なカラー システムを構築します。
npm install -g @basiclines/rampa
ドキュメント
CLIリファレンス →
貢献する
github.com/basiclines/rampa-studio
npm / バン
npm インストール @basiclines/rampa-sdk
ドキュメント
SDKリファレンス→
貢献する
github.com/basiclines/rampa-st

オーディオ
ランパ
ウェブエディター
GitHub
リリース
クレジット
個人および開発者は無料 · MIT ライセンス
テーマカラー
テーマ
色
ニュートラル
前景
背景
色合い
赤
緑
ブルー
シアン
マゼンタ
黄色
デフォルトにリセットする
ランパ
AI エージェントと人間のためのカラー ツールキット
React · React Three Fiber · Three.js · Zustand · chroma-js
Radix UI · shadcn/ui · Tailwind CSS · Vite
chroma-js · citty · TypeScript · Bun
知覚的に均一なパレットジェネレーター
このようなツールを可能にするために
丁寧に作られています · MIT ライセンス · 2024–2026
Claude Opus 4.6 コードの生成とレビュー

## Original Extract

One tool generates palettes, analyzes contrast, mixes colors, and exports to any format. No plugins, no dependencies.

A color toolkit for AI agents and humans
Generate perceptually uniform and precise color ramps straight from your terminal. CLI, SDK and web editor available.
Web editor
rampa
CLI + SDK
Explore and integrate
OKLCH/LAB
Perceptually uniform
APCA/WCAG
Built-in contrast analysis
Ramps
HSL start:end values
Color Spaces
Define color edges
Features
Everything you need for
production color systems
One tool generates palettes, analyzes contrast, mixes colors, and exports to any format. No plugins, no dependencies.
Configure lightness, saturation, and hue shift as start:end ranges.
Generate complementary, triadic and more hue shift variations.
Get perfect brand-tinted neutrals as an overlay with adjustable opacity.
Powerful theming options to derive full color palettes from a few sample colors.
Convert to hex, rgb, hsl, or oklch. Returns Text/JSON/CSS with structured values.
Check contrast in any color pair with w3c specs.
Explore colors from the terminal or build robust color systems with the SDK. Every CLI flag maps to a chainable SDK method.
Teach your AI agent
color theory
7 installable skills for Copilot, Claude Code, Codex, Cursor, and any agent that supports skills. Theme creation, status colors, data viz palettes, accessible contrast.
tinted-neutrals/skill.md
# Tinted Neutrals
Create neutral palettes that carry a subtle hint of your brand color. Instead of pure grays, get warm or cool neutrals that feel cohesive with your design system.
## When to Use
- "Create grays that match my brand"
- "I want warm/cool neutrals"
- "Tint my neutral palette with brand color"
## Recipe
### 1. Pure Neutral (Desaturated Brand)
rampa -C "<brand>" -L 98:5 -S 3:8 --size=10 -O css --name=neutral
### 2. Tinted Neutral
Apply your brand color as a subtle overlay tint:
rampa -C "<brand>" -L 98:5 -S 3:8 --tint-color="<brand>" --tint-opacity=8 --tint-blend=overlay --size=10
accessible-contrast/skill.md
# Accessible Contrast Pairs
Generate color ramps designed for accessible text/background combinations. Uses an 11-step scale for predictable WCAG-compliant pairing.
## The 11-Step Strategy
Why 11 steps? Predictable pairing math:
Maximum → 0 + 10 → Highest contrast
AAA → 1 + 9 → 7:1+ ratio
AA → 3 + 7 → 4.5:1+ ratio
AA Large → 4 + 6 → 3:1+ ratio
## Recipe
rampa -C "<brand>" -L 98:5 --size=11 -O css --name=color
Pairs are symmetrical: 0+10, 1+9, 2+8, etc.
data-viz-palette/skill.md
# Data Visualization Palette
Generate colors optimized for charts, graphs, and dashboards. Colors are mathematically spaced for maximum distinction and equal visual weight.
## Key Principles
1. Fixed lightness — equal visual weight
2. Fixed saturation — consistent vibrancy
3. Maximum hue spacing — use harmonies
4. Single shade per color — distinct hues, not ramps
### 3 Colors (Triadic)
rampa -C "<brand>" --add=triadic --size=2 -L 50:50 -S 70:70 -O css
### 4 Colors (Square)
rampa -C "<brand>" --add=square --size=2 -L 50:50 -S 70:70 -O css
status-from-accent/skill.md
# Status Colors from Accent
Generate success, warning, danger, and info colors mathematically related to your brand using square harmony.
## The Square Harmony Approach
4 colors at 90° intervals on the color wheel:
rampa -C "<brand>" --add=square -L 95:15 --size=10 -O css
base → your brand
[truncated]
Try the CLI for quick exploration or install the SDK to build robust color systems.
npm install -g @basiclines/rampa
Docs
CLI reference →
Contribute
github.com/basiclines/rampa-studio
npm / bun
npm install @basiclines/rampa-sdk
Docs
SDK reference →
Contribute
github.com/basiclines/rampa-studio
rampa
Web editor
GitHub
Releases
Credits
Free for individuals & developers · MIT License
Theme Colors
Themes
Colors
Neutral
Foreground
Background
Hues
Red
Green
Blue
Cyan
Magenta
Yellow
Reset to defaults
Rampa
A color toolkit for AI agents and humans
React · React Three Fiber · Three.js · Zustand · chroma-js
Radix UI · shadcn/ui · Tailwind CSS · Vite
chroma-js · citty · TypeScript · Bun
Perceptually uniform palette generator
For making tools like these possible
Made with care · MIT License · 2024–2026
Claude Opus 4.6 Code generation & review
