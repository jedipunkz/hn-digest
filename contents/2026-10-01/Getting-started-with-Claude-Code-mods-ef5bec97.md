---
source: "https://claude.dev/blog/getting-started-with-claude-code-mods/"
hn_url: "https://news.ycombinator.com/item?id=49926243"
title: "Getting started with Claude Code mods"
article_title: "Getting started with Claude Code mods / claude.dev Blog"
image: "https://claude.dev/blog/getting-started-with-claude-code-mods/og.png"
author: "mfiguiere"
captured_at: "2026-10-01T20:10:06Z"
capture_tool: "hn-digest"
hn_id: 49926243
score: 2
comments: 0
posted_at: "2026-10-01T19:48:39Z"
tags:
  - hacker-news
---

# Getting started with Claude Code mods

- HN: [49926243](https://news.ycombinator.com/item?id=49926243)
- Source: [claude.dev](https://claude.dev/blog/getting-started-with-claude-code-mods/)
- Score: 2
- Comments: 0
- Posted: 2026-10-01T19:48:39Z

## Translation

Title: Getting started with Claude Code mods
Article title: Getting started with Claude Code mods / claude.dev Blog
Description: Mods are hooks that ship inside plugins and run inside your Claude Code session. Build one from an empty folder, then tour two larger mods.

Article text:
Getting started with Claude Code mods / claude.dev Blog [ H ] HOME [ M ] MODS [ D ] DOCS [ T ] TERMINAL Try Claude Code [ H ] HOME [ M ] MODS [ D ] DOCS [ T ] TERMINAL Tutorials Getting started with Claude Code mods
Build your first Claude Code mod from an empty folder, then see what else the API can do.
Claude Code already lets you change a lot about how it behaves: settings, permission rules, slash commands, skills and a status line. Mods go further. Mods can rewrite or replace what Claude Code does, and can even draw custom UI. Under the hood, mods are hooks, and they ship inside plugins. Each one is a small JavaScript or TypeScript module that runs inside your session and sees every event as it happens.
That makes mods a way to fit Claude Code to how you work. You can add a readout you check all the time, put a guard in front of the commands that make you nervous, or build a review view for how you like to read changes.
This guide builds one mod from an empty folder, Token Weather , a live forecast of the context window drawn above the prompt. It's about 80 lines. Then it tours two larger mods, Blast Radius and Replay Theater , to show what else the API can do.
Claude Code 2.1.287 or later. Mods are on by default, so there's nothing to turn on. The API can change between releases. Each time Claude Code loads a mod, it writes the type declarations for your build into the mod's .claude-plugin/types/ folder, and those are the authority for your version.
A mod is a Claude Code plugin whose behavior lives in a JavaScript or TypeScript module:
The folder is a normal plugin, with a .claude-plugin/plugin.json manifest.
hooks/hooks.json names one module under modules .
The module exports register(on, options) . Inside it, on(event, matcher?, hook) adds a hook.
Every hook has the same shape:
on ( "tool.call" , { tool : "Bash" }, async ($, e, next) => {
// $ the mods API: ui, session, state, store, fs, process, clock, http, tool, command, model, ...
// e this event's input, as plain data
// next passes e to the other plugins and then to Claude Code's own behavior
return next (e);
});
Hooks form a chain, like middleware. Yours runs, next(e) hands the event to the next plugin, and at the bottom Claude Code does what it would have done anyway. A hook can do one of three things:
next(e)
result
next(e)
result
answer: return { deny } without next
observe: await next(e), then look
rewrite: next({ ...e, command })
EVENT
tool.call
e = { command }
your hook
($, e, next)
other plugins
same shape
Claude Code
runs the tool
next(e)
result
next(e)
result
answer
answer: return { deny } without next
observe: await next(e), then look
rewrite: next({ ...e, command })
Move How Example Observe const r = await next(e); /* look */ return r Record every file edit. Take a reading after each turn. Rewrite return next({ ...e, command: safer }) Change what the rest of the chain sees. Answer return { deny: "…" } without calling next Refuse a tool call. Serve a command or a tool yourself.
The events cover tool calls, the prompt as submitted, turns starting and finishing, the session starting and ending, slash commands, and ui.render : every piece of the interface as it's drawn. The module runs in a sandbox of its own, with no DOM and no Node, so everything outside it goes through $ .
How this differs from settings hooks. A settings hook runs a shell command for each event and passes JSON over stdin and stdout. A mod is loaded once and stays in the session. It can keep state, draw UI that updates as events happen, and call back into Claude Code: open a pane, run a process, register a slash command, or register a tool the model can call.
Claude Code uses them itself. Some of Claude Code's own features are built as mods, including AGENTS.md support and the /diff pane beside the conversation. Their source, with tests, is in the public anthropics/claude-code repository under mods/ , so you can read how the team builds them.
BUILD YOUR FIRST MOD: TOKEN WEATHER
Token Weather reads how full the context window is after each turn and draws one line above the prompt: a weather icon, the percentage, the tokens used out of the window, a small chart of recent turns, and how much the last turn added.
Here it is in a real session. Each turn reads more files, and the band fills from ☀ Clear to ☂ Showers to ☇ Storm:
The shortcut: let Claude build it
You can skip the six steps. Claude Code knows how to write mods, so you can describe the one you want and let it do the work. Start a session with claude and paste the prompt below:
Make me a Claude Code mod called token-weather: a live forecast of my context window, shown in the band above the prompt.
What it should show, on one line:
- A weather icon and word for how full the context window is: under 25% ☀ Clear (yellow), 25–49% ☁ Cloudy (cyan), 50–74% ☂ Showers (blue), 75–89% ☇ Storm (magenta), 90% and up ↯ Compact soon (red).
- The percentage used, then the tokens used out of the window, like "134.4k / 200k".
- A small chart of the last 12 turns, drawn with ▁▂▃▄▅▆▇█.
- How much the last turn added, like "▲ +98.3k last turn".
It should update after every turn.
Claude asks once whether to turn on hot reloading for the session. Allow it, and the band appears above the prompt when Claude's turn ends. From then on, every change reloads in place, so you can keep asking for tweaks ("make Storm start at 70%", "add the dollar cost at the end") and watch the band change. The mod loads only in this session, and its folder is cleaned up later, so to keep it, copy the folder out and install it like any plugin ( Step 6 ).
Notice that the prompt only describes what you want to see. You don't need to know the API to write one. Claude Code's built-in guide for writing mods covers the how: where to keep state so it survives a reload, how to check the plugin with claude plugin validate , and which events to hook. Change the "What it should show" lines and it's your mod, not ours.
If you'd rather see how it's put together first, or want to check what Claude wrote, read on.
Check that your Claude Code is new enough:
claude --version # 2.1.287 or later
Create this layout:
token-weather/
├── .claude-plugin/
│ ├── plugin.json
│ └── types/ (written by Claude Code when it loads the mod)
├── hooks/
│ ├── hooks.json
│ └── token-weather.mjs
├── types/
│ └── index.d.ts (added in step 3)
└── tests/
└── token-weather.test.ts (added in step 5)
.claude-plugin/plugin.json is a standard plugin manifest:
{
"name" : "token-weather" ,
"version" : "0.1.0" ,
"description" : "A live forecast of the context window, drawn above the prompt." ,
"author" : { "name" : "You" }
}
hooks/hooks.json points at the module. A mod has exactly one:
{
"modules" : [ "./token-weather.mjs" ]
}
Step 2: Draw something
The band directly above the prompt is a component called AbovePrompt . Claude Code draws nothing there itself, so it's a good first target. Hook its ui.render event and return a tree of elements:
// hooks/token-weather.mjs
export function register ( on ) {
on ( "ui.render" , { component : "AbovePrompt" }, ( $, e, next ) => {
const { Box , Text } = $.ui. resolve (e);
return Box ({
paddingX : 1 ,
children : [ Text ({ color : "yellow" , bold : true , children : "☀ Clear skies" })],
});
});
}
The elements aren't globals. $.ui.resolve(e) returns the constructors for the surface being drawn, because each surface Claude Code draws on supports a slightly different set. JSX works too, with h as the factory.
Start a session with the plugin loaded:
claude --plugin-dir ./token-weather
"☀ Clear skies" appears above the prompt. Keep the session open. The folder is watched, so every save reloads the module in place, with no restart. That quick feedback loop is most of what makes mods fun to write.
Tip: Once you know the shape, describe the next mod to Claude the way the shortcut does. It writes the plugin to a folder that hot-reloads in the same session.
Step 3: Read real numbers and keep them in $.state
$.session.usage() returns the same figures as the status line. context.tokens is the input the last response was answered over, context.window is the model's window, and context.percent is one over the other. The call is free: it only sends a token-count request if you ask for a breakdown .
Take a reading when the session starts and after every turn:
on ( "session.start" , async ($, e, next) => {
const result = await next (e);
await takeReading ($);
return result;
});
on ( "turn.complete" , async ($, e, next) => {
const result = await next (e);
if (!e. agentId ) {
await takeReading ($); // main-loop turns only, not subagents
}
return result;
});
Both hooks call next(e) first and then observe. Neither one changes what happens.
Where to keep the readings. A module-level let readings = [] looks like the obvious choice, but a hot reload is a fresh load: register runs again, session.start fires again, and module variables start over. Put the history in $.state instead. It holds named values in the host for the whole session, and they survive reloads.
// Held by the host, so the history survives a hot reload of this file.
const readings = { plugin : "token-weather" , key : "readings" };
async function takeReading ( $ ) {
const { context } = await $.session. usage ();
if (!context?. window ) return ;
const tokens = context. tokens ?? 0 ;
const percent = context. percent ?? Math . round ((tokens / context. window ) * 100 );
const { value : history = [] } = await $.state. get (readings);
await $.state. set (readings, [...history, { tokens, window : context. window , percent }]. slice (-HISTORY));
}
A state value is declared in the plugin's type contract , which is a small .d.ts file the manifest points to. Add types/index.d.ts :
export type TokenWeatherReading = { tokens : number ; window : number ; percent : number };
declare module "claude-code" {
interface PluginState {
"token-weather" : { readings : TokenWeatherReading [] };
}
}
Then add "types": "./types/index.d.ts" to plugin.json . If you skip this step, claude plugin validate stops you with an error that names the fix: token-weather.readings is not declared: the manifest's types contract must name it in interface PluginState { … } .
In return, you get redraws for free. A $.state.get made while a render hook runs subscribes that drawing, so every later $.state.set redraws the band. You never call $.ui.invalidate .
// Token Weather: a live forecast of the context window, above the prompt.
const HISTORY = 12 ;
const BARS = "▁▂▃▄▅▆▇█" ;
const FORECAST = [
{ upTo : 25 , icon : "☀" , word : "Clear" , color : "yellow" },
{ upTo : 50 , icon : "☁" , word : "Cloudy" , color : "cyan" },
{ upTo : 75 , icon : "☂" , word : "Showers" , color : "blue" },
{ upTo : 90 , icon : "☇" , word : "Storm" , color : "magenta" },
{ upTo : Infinity , icon : "↯" , word : "Compact soon" , color : "red" },
];
// Held by the host, so the history survives a hot reload of this file.
const readings = { plugin : "token-weather" , key : "readings" };
export function register ( on ) {
on ( "session.start" , async ($, e, next) => {
const result = await next (e);
await takeReading ($);
return result;
});
on ( "turn.complete" , async ($, e, next) => {
const result = await next (e);
if (!e. agentId ) {
await takeReading ($); // main-loop turns only, not subagents
}
return result;
});
on ( "ui.render" , { component : "AbovePrompt" }, async ($, e, next) => {
const { value : history = [] } = await $.state. get (readings);
if (e. props . hasSurvey || history. length === 0 ) {
return next (e);
}
const { Box , Text } = $.ui. resolve (e);
return band ( Box , Text , history, e. props . bodyColumns );
});
}
async function takeReading ( $ ) {
const { context } = await $.session. usage ();
if (!context?. window ) return ;
const tokens = context. tokens ?? 0 ;
const percent = context. percent ?? Math . round ((tokens / context. window ) * 100 );
const { value : history = [] } = await $.state. get (readings);
await $.state. set (readings, [...hi

[truncated]

## Original Extract

Mods are hooks that ship inside plugins and run inside your Claude Code session. Build one from an empty folder, then tour two larger mods.

Getting started with Claude Code mods / claude.dev Blog [ H ] HOME [ M ] MODS [ D ] DOCS [ T ] TERMINAL Try Claude Code [ H ] HOME [ M ] MODS [ D ] DOCS [ T ] TERMINAL Tutorials Getting started with Claude Code mods
Build your first Claude Code mod from an empty folder, then see what else the API can do.
Claude Code already lets you change a lot about how it behaves: settings, permission rules, slash commands, skills and a status line. Mods go further. Mods can rewrite or replace what Claude Code does, and can even draw custom UI. Under the hood, mods are hooks, and they ship inside plugins. Each one is a small JavaScript or TypeScript module that runs inside your session and sees every event as it happens.
That makes mods a way to fit Claude Code to how you work. You can add a readout you check all the time, put a guard in front of the commands that make you nervous, or build a review view for how you like to read changes.
This guide builds one mod from an empty folder, Token Weather , a live forecast of the context window drawn above the prompt. It's about 80 lines. Then it tours two larger mods, Blast Radius and Replay Theater , to show what else the API can do.
Claude Code 2.1.287 or later. Mods are on by default, so there's nothing to turn on. The API can change between releases. Each time Claude Code loads a mod, it writes the type declarations for your build into the mod's .claude-plugin/types/ folder, and those are the authority for your version.
A mod is a Claude Code plugin whose behavior lives in a JavaScript or TypeScript module:
The folder is a normal plugin, with a .claude-plugin/plugin.json manifest.
hooks/hooks.json names one module under modules .
The module exports register(on, options) . Inside it, on(event, matcher?, hook) adds a hook.
Every hook has the same shape:
on ( "tool.call" , { tool : "Bash" }, async ($, e, next) => {
// $ the mods API: ui, session, state, store, fs, process, clock, http, tool, command, model, ...
// e this event's input, as plain data
// next passes e to the other plugins and then to Claude Code's own behavior
return next (e);
});
Hooks form a chain, like middleware. Yours runs, next(e) hands the event to the next plugin, and at the bottom Claude Code does what it would have done anyway. A hook can do one of three things:
next(e)
result
next(e)
result
answer: return { deny } without next
observe: await next(e), then look
rewrite: next({ ...e, command })
EVENT
tool.call
e = { command }
your hook
($, e, next)
other plugins
same shape
Claude Code
runs the tool
next(e)
result
next(e)
result
answer
answer: return { deny } without next
observe: await next(e), then look
rewrite: next({ ...e, command })
Move How Example Observe const r = await next(e); /* look */ return r Record every file edit. Take a reading after each turn. Rewrite return next({ ...e, command: safer }) Change what the rest of the chain sees. Answer return { deny: "…" } without calling next Refuse a tool call. Serve a command or a tool yourself.
The events cover tool calls, the prompt as submitted, turns starting and finishing, the session starting and ending, slash commands, and ui.render : every piece of the interface as it's drawn. The module runs in a sandbox of its own, with no DOM and no Node, so everything outside it goes through $ .
How this differs from settings hooks. A settings hook runs a shell command for each event and passes JSON over stdin and stdout. A mod is loaded once and stays in the session. It can keep state, draw UI that updates as events happen, and call back into Claude Code: open a pane, run a process, register a slash command, or register a tool the model can call.
Claude Code uses them itself. Some of Claude Code's own features are built as mods, including AGENTS.md support and the /diff pane beside the conversation. Their source, with tests, is in the public anthropics/claude-code repository under mods/ , so you can read how the team builds them.
BUILD YOUR FIRST MOD: TOKEN WEATHER
Token Weather reads how full the context window is after each turn and draws one line above the prompt: a weather icon, the percentage, the tokens used out of the window, a small chart of recent turns, and how much the last turn added.
Here it is in a real session. Each turn reads more files, and the band fills from ☀ Clear to ☂ Showers to ☇ Storm:
The shortcut: let Claude build it
You can skip the six steps. Claude Code knows how to write mods, so you can describe the one you want and let it do the work. Start a session with claude and paste the prompt below:
Make me a Claude Code mod called token-weather: a live forecast of my context window, shown in the band above the prompt.
What it should show, on one line:
- A weather icon and word for how full the context window is: under 25% ☀ Clear (yellow), 25–49% ☁ Cloudy (cyan), 50–74% ☂ Showers (blue), 75–89% ☇ Storm (magenta), 90% and up ↯ Compact soon (red).
- The percentage used, then the tokens used out of the window, like "134.4k / 200k".
- A small chart of the last 12 turns, drawn with ▁▂▃▄▅▆▇█.
- How much the last turn added, like "▲ +98.3k last turn".
It should update after every turn.
Claude asks once whether to turn on hot reloading for the session. Allow it, and the band appears above the prompt when Claude's turn ends. From then on, every change reloads in place, so you can keep asking for tweaks ("make Storm start at 70%", "add the dollar cost at the end") and watch the band change. The mod loads only in this session, and its folder is cleaned up later, so to keep it, copy the folder out and install it like any plugin ( Step 6 ).
Notice that the prompt only describes what you want to see. You don't need to know the API to write one. Claude Code's built-in guide for writing mods covers the how: where to keep state so it survives a reload, how to check the plugin with claude plugin validate , and which events to hook. Change the "What it should show" lines and it's your mod, not ours.
If you'd rather see how it's put together first, or want to check what Claude wrote, read on.
Check that your Claude Code is new enough:
claude --version # 2.1.287 or later
Create this layout:
token-weather/
├── .claude-plugin/
│ ├── plugin.json
│ └── types/ (written by Claude Code when it loads the mod)
├── hooks/
│ ├── hooks.json
│ └── token-weather.mjs
├── types/
│ └── index.d.ts (added in step 3)
└── tests/
└── token-weather.test.ts (added in step 5)
.claude-plugin/plugin.json is a standard plugin manifest:
{
"name" : "token-weather" ,
"version" : "0.1.0" ,
"description" : "A live forecast of the context window, drawn above the prompt." ,
"author" : { "name" : "You" }
}
hooks/hooks.json points at the module. A mod has exactly one:
{
"modules" : [ "./token-weather.mjs" ]
}
Step 2: Draw something
The band directly above the prompt is a component called AbovePrompt . Claude Code draws nothing there itself, so it's a good first target. Hook its ui.render event and return a tree of elements:
// hooks/token-weather.mjs
export function register ( on ) {
on ( "ui.render" , { component : "AbovePrompt" }, ( $, e, next ) => {
const { Box , Text } = $.ui. resolve (e);
return Box ({
paddingX : 1 ,
children : [ Text ({ color : "yellow" , bold : true , children : "☀ Clear skies" })],
});
});
}
The elements aren't globals. $.ui.resolve(e) returns the constructors for the surface being drawn, because each surface Claude Code draws on supports a slightly different set. JSX works too, with h as the factory.
Start a session with the plugin loaded:
claude --plugin-dir ./token-weather
"☀ Clear skies" appears above the prompt. Keep the session open. The folder is watched, so every save reloads the module in place, with no restart. That quick feedback loop is most of what makes mods fun to write.
Tip: Once you know the shape, describe the next mod to Claude the way the shortcut does. It writes the plugin to a folder that hot-reloads in the same session.
Step 3: Read real numbers and keep them in $.state
$.session.usage() returns the same figures as the status line. context.tokens is the input the last response was answered over, context.window is the model's window, and context.percent is one over the other. The call is free: it only sends a token-count request if you ask for a breakdown .
Take a reading when the session starts and after every turn:
on ( "session.start" , async ($, e, next) => {
const result = await next (e);
await takeReading ($);
return result;
});
on ( "turn.complete" , async ($, e, next) => {
const result = await next (e);
if (!e. agentId ) {
await takeReading ($); // main-loop turns only, not subagents
}
return result;
});
Both hooks call next(e) first and then observe. Neither one changes what happens.
Where to keep the readings. A module-level let readings = [] looks like the obvious choice, but a hot reload is a fresh load: register runs again, session.start fires again, and module variables start over. Put the history in $.state instead. It holds named values in the host for the whole session, and they survive reloads.
// Held by the host, so the history survives a hot reload of this file.
const readings = { plugin : "token-weather" , key : "readings" };
async function takeReading ( $ ) {
const { context } = await $.session. usage ();
if (!context?. window ) return ;
const tokens = context. tokens ?? 0 ;
const percent = context. percent ?? Math . round ((tokens / context. window ) * 100 );
const { value : history = [] } = await $.state. get (readings);
await $.state. set (readings, [...history, { tokens, window : context. window , percent }]. slice (-HISTORY));
}
A state value is declared in the plugin's type contract , which is a small .d.ts file the manifest points to. Add types/index.d.ts :
export type TokenWeatherReading = { tokens : number ; window : number ; percent : number };
declare module "claude-code" {
interface PluginState {
"token-weather" : { readings : TokenWeatherReading [] };
}
}
Then add "types": "./types/index.d.ts" to plugin.json . If you skip this step, claude plugin validate stops you with an error that names the fix: token-weather.readings is not declared: the manifest's types contract must name it in interface PluginState { … } .
In return, you get redraws for free. A $.state.get made while a render hook runs subscribes that drawing, so every later $.state.set redraws the band. You never call $.ui.invalidate .
// Token Weather: a live forecast of the context window, above the prompt.
const HISTORY = 12 ;
const BARS = "▁▂▃▄▅▆▇█" ;
const FORECAST = [
{ upTo : 25 , icon : "☀" , word : "Clear" , color : "yellow" },
{ upTo : 50 , icon : "☁" , word : "Cloudy" , color : "cyan" },
{ upTo : 75 , icon : "☂" , word : "Showers" , color : "blue" },
{ upTo : 90 , icon : "☇" , word : "Storm" , color : "magenta" },
{ upTo : Infinity , icon : "↯" , word : "Compact soon" , color : "red" },
];
// Held by the host, so the history survives a hot reload of this file.
const readings = { plugin : "token-weather" , key : "readings" };
export function register ( on ) {
on ( "session.start" , async ($, e, next) => {
const result = await next (e);
await takeReading ($);
return result;
});
on ( "turn.complete" , async ($, e, next) => {
const result = await next (e);
if (!e. agentId ) {
await takeReading ($); // main-loop turns only, not subagents
}
return result;
});
on ( "ui.render" , { component : "AbovePrompt" }, async ($, e, next) => {
const { value : history = [] } = await $.state. get (readings);
if (e. props . hasSurvey || history. length === 0 ) {
return next (e);
}
const { Box , Text } = $.ui. resolve (e);
return band ( Box , Text , history, e. props . bodyColumns );
});
}
async function takeReading ( $ ) {
const { context } = await $.session. usage ();
if (!context?. window ) return ;
const tokens = context. tokens ?? 0 ;
const percent = context. percent ?? Math . round ((tokens / context. window ) * 100 );
const { value : history = [] } = await $.state. get (readings);
await $.state. set (readings, [...hi

[truncated]
