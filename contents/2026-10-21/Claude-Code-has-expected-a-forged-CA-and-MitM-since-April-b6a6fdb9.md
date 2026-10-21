---
source: "https://github.com/Andytonglove/claude-code-leak/blob/main/src/upstreamproxy/upstreamproxy.ts"
hn_url: "https://news.ycombinator.com/item?id=50184897"
title: "Claude Code has expected a forged CA and MitM since April"
article_title: "claude-code-leak/src/upstreamproxy/upstreamproxy.ts at main · Andytonglove/claude-code-leak · GitHub"
image: "https://opengraph.githubassets.com/e3d101491b7d66854655fd838f83e3fe2e3e416f03c167c625d6b8d9ad2c4fab/Andytonglove/claude-code-leak"
author: "cududa"
captured_at: "2026-10-21T18:26:10Z"
capture_tool: "hn-digest"
hn_id: 50184897
score: 3
comments: 0
posted_at: "2026-10-21T17:50:06Z"
tags:
  - hacker-news
---

# Claude Code has expected a forged CA and MitM since April

- HN: [50184897](https://news.ycombinator.com/item?id=50184897)
- Source: [github.com](https://github.com/Andytonglove/claude-code-leak/blob/main/src/upstreamproxy/upstreamproxy.ts)
- Score: 3
- Comments: 0
- Posted: 2026-10-21T17:50:06Z

## Translation

Title: Claude Code has expected a forged CA and MitM since April
Article title: claude-code-leak/src/upstreamproxy/upstreamproxy.ts at main · Andytonglove/claude-code-leak · GitHub
Description: Claude Code Snapshot for Research. All original source code is the property of Anthropic. - claude-code-leak/src/upstreamproxy/upstreamproxy.ts at main · Andytonglove/claude-code-leak

Article text:
claude-code-leak/src/upstreamproxy/upstreamproxy.ts at main · Andytonglove/claude-code-leak · GitHub
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
Andytonglove
Copy path Blame More file actions Blame More file actions Latest commit
History History 285 lines (267 loc) · 9.58 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Open symbols panel Edit and raw actions 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100 101 102 103 104 105 106 107 108 109 110 111 112 113 114 115 116 117 118 119 120 121 122 123 124 125 126 127 128 129 130 131 132 133 134 135 136 137 138 139 140 141 142 143 144 145 146 147 148 149 150 151 152 153 154 155 156 157 158 159 160 161 162 163 164 165 166 167 168 169 170 171 172 173 174 175 176 177 178 179 180 181 182 183 184 185 186 187 188 189 190 191 192 193 194 195 196 197 198 199 200 201 202 203 204 205 206 207 208 209 210 211 212 213 214 215 216 217 218 219 220 221 222 223 224 225 226 227 228 229 230 231 232 233 234 235 236 237 238 239 240 241 242 243 244 245 246 247 248 249 250 251 252 253 254 255 256 257 258 259 260 261 262 263 264 265 266 267 268 269 270 271 272 273 274 275 276 277 278 279 280 281 282 283 284 285 /**
* CCR upstreamproxy — container-side wiring.
*
* When running inside a CCR session container with upstreamproxy configured,
* this module:
* 1. Reads the session token from /run/ccr/session_token
* 2. Sets prctl(PR_SET_DUMPABLE, 0) to block same-UID ptrace of the heap
* 3. Downloads the upstreamproxy CA cert and concatenates it with the
* system bundle so curl/gh/python trust the MITM proxy
* 4. Starts a local CONNECT→WebSocket relay (see relay.ts)
* 5. Unlinks the token file (token stays heap-only; file is gone before
* the agent loop can see it, but only after the relay is confirmed up
* so a supervisor restart can retry)
* 6. Exposes HTTPS_PROXY / SSL_CERT_FILE env vars for all agent subprocesses
*
* Every step fails open: any error logs a warning and disables the proxy.
* A broken proxy setup must never break an otherwise-working session.
*
* Design doc: api-go/ccr/docs/plans
[truncated]
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Claude Code Snapshot for Research. All original source code is the property of Anthropic. - claude-code-leak/src/upstreamproxy/upstreamproxy.ts at main · Andytonglove/claude-code-leak

claude-code-leak/src/upstreamproxy/upstreamproxy.ts at main · Andytonglove/claude-code-leak · GitHub
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
Andytonglove
Copy path Blame More file actions Blame More file actions Latest commit
History History 285 lines (267 loc) · 9.58 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Open symbols panel Edit and raw actions 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100 101 102 103 104 105 106 107 108 109 110 111 112 113 114 115 116 117 118 119 120 121 122 123 124 125 126 127 128 129 130 131 132 133 134 135 136 137 138 139 140 141 142 143 144 145 146 147 148 149 150 151 152 153 154 155 156 157 158 159 160 161 162 163 164 165 166 167 168 169 170 171 172 173 174 175 176 177 178 179 180 181 182 183 184 185 186 187 188 189 190 191 192 193 194 195 196 197 198 199 200 201 202 203 204 205 206 207 208 209 210 211 212 213 214 215 216 217 218 219 220 221 222 223 224 225 226 227 228 229 230 231 232 233 234 235 236 237 238 239 240 241 242 243 244 245 246 247 248 249 250 251 252 253 254 255 256 257 258 259 260 261 262 263 264 265 266 267 268 269 270 271 272 273 274 275 276 277 278 279 280 281 282 283 284 285 /**
* CCR upstreamproxy — container-side wiring.
*
* When running inside a CCR session container with upstreamproxy configured,
* this module:
* 1. Reads the session token from /run/ccr/session_token
* 2. Sets prctl(PR_SET_DUMPABLE, 0) to block same-UID ptrace of the heap
* 3. Downloads the upstreamproxy CA cert and concatenates it with the
* system bundle so curl/gh/python trust the MITM proxy
* 4. Starts a local CONNECT→WebSocket relay (see relay.ts)
* 5. Unlinks the token file (token stays heap-only; file is gone before
* the agent loop can see it, but only after the relay is confirmed up
* so a supervisor restart can retry)
* 6. Exposes HTTPS_PROXY / SSL_CERT_FILE env vars for all agent subprocesses
*
* Every step fails open: any error logs a warning and disables the proxy.
* A broken proxy setup must never break an otherwise-working session.
*
* Design doc: api-go/ccr/docs/plans
[truncated]
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
