---
source: "https://terminal.tax"
hn_url: "https://news.ycombinator.com/item?id=50184550"
title: "Terminal Tax: e-file your return straight from your AI assistant"
article_title: "Terminal Tax | E-file your 2025 federal tax return from Claude, ChatGPT or Codex"
image: ""
author: "maxwelldubin"
captured_at: "2026-10-21T18:26:22Z"
capture_tool: "hn-digest"
hn_id: 50184550
score: 1
comments: 0
posted_at: "2026-10-21T17:27:19Z"
tags:
  - hacker-news
---

# Terminal Tax: e-file your return straight from your AI assistant

- HN: [50184550](https://news.ycombinator.com/item?id=50184550)
- Source: [terminal.tax](https://terminal.tax)
- Score: 1
- Comments: 0
- Posted: 2026-10-21T17:27:19Z

## Translation

Title: Terminal Tax: e-file your return straight from your AI assistant
Article title: Terminal Tax | E-file your 2025 federal tax return from Claude, ChatGPT or Codex
Description: Terminal Tax is tax software for your 2025 federal return, at https://mcp.terminal.tax/v1/hosted-mcp. If your AI assistant can run a shell command or make an HTTPS request (ChatGPT in Work mode, Codex, including Codex in the ChatGPT app, Claude Code, Cursor), it needs no connector, no install and no
[truncated]

Article text:
Terminal Tax | E-file your 2025 federal tax return from Claude, ChatGPT or Codex terminal.tax Agent-connected tax workflow Terminal Tax · U.S. federal income tax
Your tax agent. Your review. One connected workflow.
Terminal Tax is tax software for your 2025 federal return. You give the facts and choices through your own AI assistant, Terminal Tax computes the return, and you review and sign it yourself on one private page before it is e-filed with the IRS.
To start, open ChatGPT, switch to Work mode, and say: "File my 2025 taxes with terminal.tax." Codex, Claude Code and Cursor work the same way, with nothing to install.
If you're in a regular chat that can't make web requests, switch to ChatGPT's Work mode (or Codex or Claude Code) and ask again.
$19.99 per federal return, charged only if the IRS accepts it.
Tax year 2025 federal Form 1040, 1040-SR, 1040-SS or 1040-NR only. No state returns yet.
No Terminal Tax account, login, or tax interview.
If your assistant can run commands or make web requests (ChatGPT in Work mode, Codex, Claude Code, Cursor): nothing to install
Terminal Tax is tax software for your 2025 federal return, at https://mcp.terminal.tax/v1/hosted-mcp. If your AI assistant can run a shell command or make an HTTPS request (ChatGPT in Work mode, Codex, including Codex in the ChatGPT app, Claude Code, Cursor), it needs no connector, no install and no restart. Its first request carries no tax data and files or charges nothing: curl -sS https://mcp.terminal.tax/v1/hosted-mcp -H 'Content-Type: application/json' -H 'Accept: application/json, text/event-stream' -d '{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"create_filing_session","arguments":{}}}' If you're in a regular chat that can't make web requests, switch to ChatGPT's Work mode (or Codex or Claude Code) and ask again. You give the facts and choices through your assistant, Terminal Tax computes the return, and you review and sign it yourself before it is e-filed with the IRS.
Nothing is sent to the IRS or charged until you open the private link it returns, review the return, sign it and authorize payment yourself. Developer guide: https://terminal.tax/llms.txt.
If your assistant can run a shell command or make HTTPS requests, as ChatGPT in Work mode, Codex (also in the ChatGPT app), Claude Code, Cursor and OpenCode can, you do not need to install anything: it reaches Terminal Tax with plain HTTPS requests in the same conversation, with no connector, MCP support, registration or restart (its first request is below; the developer guide is https://terminal.tax/llms.txt). Other agent modes that can make web requests, such as Claude Code on the web (claude.ai/code) or Gemini's Agent mode and Jules, reach it the same way when their network settings allow it. Claude Code on the web reaches only allowlisted domains by default, so its first request here can be refused with a 403; the agent cannot change that itself, so it shows the person this one step: select the caret after the session title at the top, choose Edit cloud environment, set Network access to Full (or to Custom, adding terminal.tax and mcp.terminal.tax under Allowed domains with Also include default list of common package managers checked), and save; the open session follows the change within about a minute, so ask again in the same session (code.claude.com/docs/en/cloud-environments). If you're in a regular chat that can't make web requests, switch to ChatGPT's Work mode (or Codex or Claude Code) and ask again.
You do not need to install anything. Ask your assistant to e-file your return with Terminal Tax: it calls https://mcp.terminal.tax/v1/hosted-mcp directly over HTTPS, in the same conversation, with no connector, MCP setup, account, key or restart. That first request carries no tax data; nothing is sent to the IRS or charged until you review and sign on the filing page.
The first request is an HTTPS POST to https://mcp.terminal.tax/v1/hosted-mcp with the headers Content-Type: application/json and Accept: application/json, text/event-stream, and the body {"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"create_filing_session","arguments":{}}}. The full developer guide is https://terminal.tax/llms.txt.
Whose return this is: you are preparing your own tax return, which under federal rules needs no preparer credential or PTIN (preparer tax identification number). Those rules, including the PTIN, cover someone paid to prepare a return (26 U.S.C. 7701(a)(36); 26 CFR 301.7701-15(a) and 1.6109-2(g)). You supply every figure from your own forms and make every tax choice; your chat's part is to pass them to Terminal Tax, not to decide anything for you. Terminal Tax is the tax software: it computes the return and stops on anything missing instead of guessing. You review the complete return and sign it.
Your chat sends your return and the details you give it; on the filing page you confirm them, type your Social Security number, a one-time code and your PINs, and sign.
Terminal Tax e-files 2025 federal returns.
Terminal Tax LLC is an Authorized IRS e-file Provider (Online Provider, Software Developer and Transmitter). Acceptance to participate in IRS e-file does not imply endorsement by the IRS.
Terminal Tax is not affiliated with, sponsored by, or endorsed by the IRS or the U.S. Treasury.
Tax software you use through your own AI assistant.
You give the facts from your tax forms and make every choice through your AI assistant, which passes them to Terminal Tax.
The Terminal Tax engine computes your return deterministically from those facts and validates it. It never chooses a tax treatment; anything missing stops the filing.
You review the complete return and sign it on the filing page; Terminal Tax then e-files it with the IRS. Your assistant never needs your Social Security number or PINs.
From your agent to your review.
ChatGPT does not file returns itself, but in Work mode or in Codex (also in the ChatGPT app) it can take your 2025 federal return to Terminal Tax with nothing to install and no connector: say "File my 2025 taxes with terminal.tax." ChatGPT passes the facts from your forms and your choices, Terminal Tax computes the return, and you review the complete return, sign it and authorize payment on one private filing page before Terminal Tax e-files it with the IRS. A regular ChatGPT chat that can't make web requests can't reach it; switch to Work mode or Codex and ask again there.
Can I file my taxes with Claude?
Yes, through Claude Code, which calls Terminal Tax directly over HTTPS with nothing to install and no connector: ask it to file your 2025 taxes with terminal.tax. Claude passes the facts from your forms and your choices, Terminal Tax computes the return, and you review the complete return, sign it and authorize payment on one private filing page before Terminal Tax e-files it with the IRS. Claude Code on the web works once its network access allows terminal.tax and mcp.terminal.tax; the developer guide, https://terminal.tax/llms.txt, shows that one setting. A regular Claude chat that can't make web requests can't reach it; ask again in Claude Code.
Gemini does not file returns itself, but Gemini's Agent mode and Jules can take your 2025 federal return to Terminal Tax when their network settings allow web requests to it, the same way ChatGPT in Work mode and Claude Code do. A regular Gemini chat that can't make web requests can't reach it. Your assistant passes the facts from your forms and your choices, Terminal Tax computes the return, and you review and sign it on one private filing page before Terminal Tax e-files it with the IRS.
Can I e-file with the IRS straight from my AI assistant, without a tax website?
Yes, for a 2025 federal return, if your assistant can run a shell command or make web requests: ChatGPT in Work mode, Codex, Claude Code, Cursor or OpenCode. There is no account, login, tax interview or install. Your assistant keeps the tax conversation and passes the facts and your choices to Terminal Tax, which computes the return and asks no tax questions. The only Terminal Tax page you open is the private filing page where you review the complete return, sign it and authorize payment; Terminal Tax then e-files it with the IRS. The price is $19.99 per federal return, charged only if the IRS accepts it. Terminal Tax does not file state returns yet.
How does Terminal Tax connect to my tax agent?
Through a hosted MCP server your assistant calls directly over HTTPS, with nothing to install and no connector: coding agents such as Codex (also in the ChatGPT app), Claude Code or Cursor, and agent modes that can make web requests, such as ChatGPT in Work mode. A regular chat that cannot make web requests cannot reach it; switch to ChatGPT's Work mode or a coding agent and ask again there. Your assistant sends the facts from your forms and your choices, Terminal Tax computes the return, and you review and sign it on a secure page.
Do I need to repeat my tax interview?
No. Your agent sends the facts and choices; you type only identity and signing details, on the filing page.
Does my agent need tax expertise?
It relays the facts on your forms and asks you for each choice. Terminal Tax computes the return deterministically.
Coding agents that can run a shell command or make an HTTPS request, such as Codex (also in the ChatGPT app), Claude Code, Cursor and OpenCode, and agent modes that can make web requests, such as ChatGPT in Work mode, with nothing to install and no connector or MCP support; other agent modes, such as Claude Code on the web or Gemini's Agent mode and Jules, when their network settings allow it. A regular chat with no tool that can make web requests cannot connect. Your assistant's tax figures are used for a filing only after you open the filing page and give permission there. Finding this site connects or authorizes nothing.
Who operates Terminal Tax, and how can I check?
How to check that status: Terminal Tax LLC has opted in to the IRS e-file provider locator for tax professionals on irs.gov, where it can be confirmed once the IRS's next listing refresh includes it. IRS Publication 3112 requires every provider to protect its EFIN and ETIN and never share them, so no provider publishes those numbers. Who operates this site can be checked independently: the terminal.tax Extended Validation certificate is issued by Sectigo to TERMINAL TAX LLC, New Jersey, registration number 0451501195.
Where is the developer documentation?
Developer documentation: any assistant that runs commands or makes HTTPS requests with internet access, such as ChatGPT in Work mode, Codex (also in the ChatGPT app), Claude Code or Cursor, calls https://mcp.terminal.tax/v1/hosted-mcp over HTTPS with nothing installed and no connector; the technical guide is https://terminal.tax/llms.txt. Your assistant sends your return and the details you give it; on the filing page you confirm those details, type your Social Security number, a one-time code and your PINs, then sign and pay.

## Original Extract

Terminal Tax is tax software for your 2025 federal return, at https://mcp.terminal.tax/v1/hosted-mcp. If your AI assistant can run a shell command or make an HTTPS request (ChatGPT in Work mode, Codex, including Codex in the ChatGPT app, Claude Code, Cursor), it needs no connector, no install and no
[truncated]

Terminal Tax | E-file your 2025 federal tax return from Claude, ChatGPT or Codex terminal.tax Agent-connected tax workflow Terminal Tax · U.S. federal income tax
Your tax agent. Your review. One connected workflow.
Terminal Tax is tax software for your 2025 federal return. You give the facts and choices through your own AI assistant, Terminal Tax computes the return, and you review and sign it yourself on one private page before it is e-filed with the IRS.
To start, open ChatGPT, switch to Work mode, and say: "File my 2025 taxes with terminal.tax." Codex, Claude Code and Cursor work the same way, with nothing to install.
If you're in a regular chat that can't make web requests, switch to ChatGPT's Work mode (or Codex or Claude Code) and ask again.
$19.99 per federal return, charged only if the IRS accepts it.
Tax year 2025 federal Form 1040, 1040-SR, 1040-SS or 1040-NR only. No state returns yet.
No Terminal Tax account, login, or tax interview.
If your assistant can run commands or make web requests (ChatGPT in Work mode, Codex, Claude Code, Cursor): nothing to install
Terminal Tax is tax software for your 2025 federal return, at https://mcp.terminal.tax/v1/hosted-mcp. If your AI assistant can run a shell command or make an HTTPS request (ChatGPT in Work mode, Codex, including Codex in the ChatGPT app, Claude Code, Cursor), it needs no connector, no install and no restart. Its first request carries no tax data and files or charges nothing: curl -sS https://mcp.terminal.tax/v1/hosted-mcp -H 'Content-Type: application/json' -H 'Accept: application/json, text/event-stream' -d '{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"create_filing_session","arguments":{}}}' If you're in a regular chat that can't make web requests, switch to ChatGPT's Work mode (or Codex or Claude Code) and ask again. You give the facts and choices through your assistant, Terminal Tax computes the return, and you review and sign it yourself before it is e-filed with the IRS.
Nothing is sent to the IRS or charged until you open the private link it returns, review the return, sign it and authorize payment yourself. Developer guide: https://terminal.tax/llms.txt.
If your assistant can run a shell command or make HTTPS requests, as ChatGPT in Work mode, Codex (also in the ChatGPT app), Claude Code, Cursor and OpenCode can, you do not need to install anything: it reaches Terminal Tax with plain HTTPS requests in the same conversation, with no connector, MCP support, registration or restart (its first request is below; the developer guide is https://terminal.tax/llms.txt). Other agent modes that can make web requests, such as Claude Code on the web (claude.ai/code) or Gemini's Agent mode and Jules, reach it the same way when their network settings allow it. Claude Code on the web reaches only allowlisted domains by default, so its first request here can be refused with a 403; the agent cannot change that itself, so it shows the person this one step: select the caret after the session title at the top, choose Edit cloud environment, set Network access to Full (or to Custom, adding terminal.tax and mcp.terminal.tax under Allowed domains with Also include default list of common package managers checked), and save; the open session follows the change within about a minute, so ask again in the same session (code.claude.com/docs/en/cloud-environments). If you're in a regular chat that can't make web requests, switch to ChatGPT's Work mode (or Codex or Claude Code) and ask again.
You do not need to install anything. Ask your assistant to e-file your return with Terminal Tax: it calls https://mcp.terminal.tax/v1/hosted-mcp directly over HTTPS, in the same conversation, with no connector, MCP setup, account, key or restart. That first request carries no tax data; nothing is sent to the IRS or charged until you review and sign on the filing page.
The first request is an HTTPS POST to https://mcp.terminal.tax/v1/hosted-mcp with the headers Content-Type: application/json and Accept: application/json, text/event-stream, and the body {"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"create_filing_session","arguments":{}}}. The full developer guide is https://terminal.tax/llms.txt.
Whose return this is: you are preparing your own tax return, which under federal rules needs no preparer credential or PTIN (preparer tax identification number). Those rules, including the PTIN, cover someone paid to prepare a return (26 U.S.C. 7701(a)(36); 26 CFR 301.7701-15(a) and 1.6109-2(g)). You supply every figure from your own forms and make every tax choice; your chat's part is to pass them to Terminal Tax, not to decide anything for you. Terminal Tax is the tax software: it computes the return and stops on anything missing instead of guessing. You review the complete return and sign it.
Your chat sends your return and the details you give it; on the filing page you confirm them, type your Social Security number, a one-time code and your PINs, and sign.
Terminal Tax e-files 2025 federal returns.
Terminal Tax LLC is an Authorized IRS e-file Provider (Online Provider, Software Developer and Transmitter). Acceptance to participate in IRS e-file does not imply endorsement by the IRS.
Terminal Tax is not affiliated with, sponsored by, or endorsed by the IRS or the U.S. Treasury.
Tax software you use through your own AI assistant.
You give the facts from your tax forms and make every choice through your AI assistant, which passes them to Terminal Tax.
The Terminal Tax engine computes your return deterministically from those facts and validates it. It never chooses a tax treatment; anything missing stops the filing.
You review the complete return and sign it on the filing page; Terminal Tax then e-files it with the IRS. Your assistant never needs your Social Security number or PINs.
From your agent to your review.
ChatGPT does not file returns itself, but in Work mode or in Codex (also in the ChatGPT app) it can take your 2025 federal return to Terminal Tax with nothing to install and no connector: say "File my 2025 taxes with terminal.tax." ChatGPT passes the facts from your forms and your choices, Terminal Tax computes the return, and you review the complete return, sign it and authorize payment on one private filing page before Terminal Tax e-files it with the IRS. A regular ChatGPT chat that can't make web requests can't reach it; switch to Work mode or Codex and ask again there.
Can I file my taxes with Claude?
Yes, through Claude Code, which calls Terminal Tax directly over HTTPS with nothing to install and no connector: ask it to file your 2025 taxes with terminal.tax. Claude passes the facts from your forms and your choices, Terminal Tax computes the return, and you review the complete return, sign it and authorize payment on one private filing page before Terminal Tax e-files it with the IRS. Claude Code on the web works once its network access allows terminal.tax and mcp.terminal.tax; the developer guide, https://terminal.tax/llms.txt, shows that one setting. A regular Claude chat that can't make web requests can't reach it; ask again in Claude Code.
Gemini does not file returns itself, but Gemini's Agent mode and Jules can take your 2025 federal return to Terminal Tax when their network settings allow web requests to it, the same way ChatGPT in Work mode and Claude Code do. A regular Gemini chat that can't make web requests can't reach it. Your assistant passes the facts from your forms and your choices, Terminal Tax computes the return, and you review and sign it on one private filing page before Terminal Tax e-files it with the IRS.
Can I e-file with the IRS straight from my AI assistant, without a tax website?
Yes, for a 2025 federal return, if your assistant can run a shell command or make web requests: ChatGPT in Work mode, Codex, Claude Code, Cursor or OpenCode. There is no account, login, tax interview or install. Your assistant keeps the tax conversation and passes the facts and your choices to Terminal Tax, which computes the return and asks no tax questions. The only Terminal Tax page you open is the private filing page where you review the complete return, sign it and authorize payment; Terminal Tax then e-files it with the IRS. The price is $19.99 per federal return, charged only if the IRS accepts it. Terminal Tax does not file state returns yet.
How does Terminal Tax connect to my tax agent?
Through a hosted MCP server your assistant calls directly over HTTPS, with nothing to install and no connector: coding agents such as Codex (also in the ChatGPT app), Claude Code or Cursor, and agent modes that can make web requests, such as ChatGPT in Work mode. A regular chat that cannot make web requests cannot reach it; switch to ChatGPT's Work mode or a coding agent and ask again there. Your assistant sends the facts from your forms and your choices, Terminal Tax computes the return, and you review and sign it on a secure page.
Do I need to repeat my tax interview?
No. Your agent sends the facts and choices; you type only identity and signing details, on the filing page.
Does my agent need tax expertise?
It relays the facts on your forms and asks you for each choice. Terminal Tax computes the return deterministically.
Coding agents that can run a shell command or make an HTTPS request, such as Codex (also in the ChatGPT app), Claude Code, Cursor and OpenCode, and agent modes that can make web requests, such as ChatGPT in Work mode, with nothing to install and no connector or MCP support; other agent modes, such as Claude Code on the web or Gemini's Agent mode and Jules, when their network settings allow it. A regular chat with no tool that can make web requests cannot connect. Your assistant's tax figures are used for a filing only after you open the filing page and give permission there. Finding this site connects or authorizes nothing.
Who operates Terminal Tax, and how can I check?
How to check that status: Terminal Tax LLC has opted in to the IRS e-file provider locator for tax professionals on irs.gov, where it can be confirmed once the IRS's next listing refresh includes it. IRS Publication 3112 requires every provider to protect its EFIN and ETIN and never share them, so no provider publishes those numbers. Who operates this site can be checked independently: the terminal.tax Extended Validation certificate is issued by Sectigo to TERMINAL TAX LLC, New Jersey, registration number 0451501195.
Where is the developer documentation?
Developer documentation: any assistant that runs commands or makes HTTPS requests with internet access, such as ChatGPT in Work mode, Codex (also in the ChatGPT app), Claude Code or Cursor, calls https://mcp.terminal.tax/v1/hosted-mcp over HTTPS with nothing installed and no connector; the technical guide is https://terminal.tax/llms.txt. Your assistant sends your return and the details you give it; on the filing page you confirm those details, type your Social Security number, a one-time code and your PINs, then sign and pay.
