---
source: "https://github.com/agntflw/claude-dbview"
hn_url: "https://news.ycombinator.com/item?id=50198367"
title: "Claude Mod Database View"
article_title: "GitHub - agntflw/claude-dbview: Custom generic db view build in Claude Code using mod · GitHub"
image: "https://opengraph.githubassets.com/497882569da0935d6a5c0f891fff3e107dbc734a4b11824cb5d2dcb77597b150/agntflw/claude-dbview"
author: "agntflow"
captured_at: "2026-10-22T18:21:10Z"
capture_tool: "hn-digest"
hn_id: 50198367
score: 1
comments: 0
posted_at: "2026-10-22T17:23:55Z"
tags:
  - hacker-news
---

# Claude Mod Database View

- HN: [50198367](https://news.ycombinator.com/item?id=50198367)
- Source: [github.com](https://github.com/agntflw/claude-dbview)
- Score: 1
- Comments: 0
- Posted: 2026-10-22T17:23:55Z

## Translation

Title: Claude Mod Database View
Article title: GitHub - agntflw/claude-dbview: Custom generic db view build in Claude Code using mod · GitHub
Description: Custom generic db view build in Claude Code using mod - agntflw/claude-dbview

Article text:
GitHub - agntflw/claude-dbview: Custom generic db view build in Claude Code using mod · GitHub
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
agntflw
claude-dbview Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Custom generic db view build in Claude Code using mod
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
35 Commits 35 Commits Folders and files
.claude-plugin .claude-plugin dev dev docs/ images docs/ images hooks hooks tests tests types types .gitattributes .gitattributes .gitignore .gitignore LICENSE LICENSE README.md README.md tsconfig.json tsconfig.json View all files Repository files navigation
A PostgreSQL database view for Claude Code, in a pane beside the chat: every database and Postgres container as an IDE-style tree, with columns, constraints, rows and DDL. Read-only.
The pane on the desktop app, as it draws the sample database in dev/seed.sql . On the terminal it uses glyphs and no borders.
Install, once, for every project (user scope):
claude plugin install dbview --marketplace agntflw/claude-dbview --config psql_path=psql
Start a new Claude Code session (one already open does not see the plugin until you run /reload-plugins in it), then type /db . The pane opens on the databases it finds: your DATABASE_URL or PG* variables, and every running Postgres container.
To add another one: /db connect postgres://user:password@host:5432/database [name] .
Press a database, then a table. Switch between Columns, Rows (pick 50, 100, 500 or 1000) and DDL.
Paste this into Claude Code, in any folder:
Install the dbview plugin for Claude Code globally for me (user scope), from github.com/agntflw/claude-dbview.
1. Run: claude plugin install dbview --marketplace agntflw/claude-dbview --scope user --config psql_path=psql
If this Claude Code is too old for --marketplace, run instead:
claude plugin marketplace add agntflw/claude-dbview, then claude plugin install dbview@claude-dbview --scope user
2. Run: claude plugin list, and check dbview is installed and enabled.
3. Check what it needs: psql --version, and docker --version if I use Postgres in Docker.
Tell me what is missing and how to install it on this machine, but do not install it.
4. Tell me to start a new Claude Code session and type /db.
Add MySQL beside Postgres
dbview reads PostgreSQL only. To have Claude add MySQL (8.4 LTS, or pinned to 8.4.5) to your copy, paste this into Claude Code, in any folder:
Add MySQL support to my dbview plugin (github.com/agntflw/claude-dbview), beside PostgreSQL,
without changing how Postgres works.
Find the code first: use the claude-dbview folder that `claude plugin list` shows dbview is read from,
else one already on this machine, else clone the repository to ~/.claude/plugins-src/claude-dbview.
Work in that folder.
Set up a test database first:
- Run MySQL 8.4 in Docker: docker run -d --name mysql84 -e MYSQL_ROOT_PASSWORD=dev
-e MYSQL_DATABASE=shop -p 33306:3306 mysql:8.4 (the latest 8.4 LTS; write mysql:8.4.5 here to pin it).
- Wait until it accepts connections, then load a MySQL version of dev/seed.sql
(schemas become databases; serial becomes AUTO_INCREMENT; keep the same tables, keys and rows).
Then extend the mod, following its layout (types first, then engine, then pane):
1. types/index.d.ts: add an engine field ('postgres' | 'mysql') to DbSource, and keep every request shape.
2. hooks/engine/: add a MySQL driver next to psql.ts. Run the mysql client inside the container
(docker exec -i <container> mysql -N -B ...) or on the host for mysql:// URLs.
Pass the password in MYSQL_PWD, never on the command line.
Every session is read-only: SET SESSION TRANSACTION READ ONLY and max_execution_time=10000.
Build each answer as one JSON document with JSON_OBJECT and JSON_ARRAYAGG
from information_schema (catalog, columns, constraints, indexes).
Use SHOW CREATE TABLE / VIEW for the DDL.
3. hooks/engine/docker.ts: also discover mysql and mariadb containers
(user root with MYSQL_ROOT_PASSWORD, or MYSQL_USER and MYSQL_PASSWORD; database from MYSQL_DATABASE).
4. /db connect: accept mysql:// URLs as well as postgres://.
5. Pane: show the engine and version on each database row and in the switcher
(MySQL 8.4 · docker). MySQL has no schemas inside a database, so skip the schema level for it.
Sequences and materialized views do not exist there.
Prove it:
- Extend dev/engine-e2e.ts with MySQL checks against mysql84
(catalog, columns, a foreign
[truncated]
Each database is a root of the tree, with its schemas as packages inside, then Tables, Views, Materialized views, Sequences and Functions folders, then the objects and their columns. Open a table or view to see:
Columns : name, type, nullable, primary key or → schema.table(column) reference, defaults, then the constraints (keys, foreign keys, checks, not null), indexes, and the foreign keys elsewhere that point at it ( Referenced by ).
Rows : the first 50, 100, 500 or 1000 rows in key order, with the query shown.
DDL : the table as a migration would create it (its sequences, the table with defaults and named constraints, indexes, OWNED BY ), numbered, with Copy.
Foreign keys are links that look like plain text. Press a → schema.table(column) reference, a foreign key constraint, a ← schema.table(column) line under Referenced by, or a key column in the tree, and the table at the other end opens on the same tab, unfolded and scrolled into view in the tree. ← name at the top right goes back, one jump at a time.
Every Postgres container Docker is running shows up by itself, so the databases a migration test leaves behind (say mig_before and mig_after ) sit side by side in the tree to compare.
Everything is read-only: each request runs in a read-only session with a 10 second statement timeout.
At the prompt of a terminal session:
/plugin install dbview --marketplace agntflw/claude-dbview
Answer y to add the marketplace, then pick a scope (user is the first). From a shell:
claude plugin install dbview --marketplace agntflw/claude-dbview
In the desktop app, where /plugin does not run: click + next to the prompt box, then Plugins, Add marketplace with agntflw/claude-dbview , then Add plugin.
On a Claude Code older than v2.1.292, add the marketplace first, then install from it:
claude plugin marketplace add agntflw/claude-dbview
claude plugin install dbview@claude-dbview
To update: claude plugin update dbview@claude-dbview , or turn on auto-update for claude-dbview in /plugin › Marketplaces.
On a claude.ai Team or Enterprise plan, an Owner can also add dbview for everyone in Organization settings › Plugins & skills ; it then reaches each member's Claude Code as dbview@synced .
Start a new session, or run /reload-plugins in the open one. A plugin installed from a shell loads in sessions started after the install.
Check claude --version : mods (plugins with a hooks module, like this one) need a recent Claude Code. Run claude update if it is old.
Run claude plugin list : dbview should say enabled. In a session, /plugin shows any error loading it.
The "userConfig option not yet set" note after install is harmless: psql_path defaults to psql .
/db open the pane on every database it finds
/db connect postgres://u:p@host:5432/db [name] save a database under a name (default: its database name)
/db forget <name> forget a saved database
/db refresh list the databases again and reload the open ones
/db theme [system|light|dark] the pane's colours in the desktop app (default: system)
The pane lists, in order: the databases saved with /db connect , the one named by DATABASE_URL or the PG* variables (unless a saved one is the same), then each running Postgres container. Press a database to unfold it (its catalog loads the first time), or pick one in the switcher at the top. Press a schema or folder to fold it, press a table to open it (press again to fold its columns), and filter the tree by table or column name. On the terminal, with the pane focused: c , w , d switch to Columns, Rows, DDL; r refreshes; y copies the DDL; b goes back after a jump.
The pane follows its width: the tree and the open object sit side by side when it is wide, the object under the tree when it is narrow, and the tree takes about a third of a wide pane. In the Claude desktop app the pane scrolls as one page with the wheel or trackpad, and a table opens beside (or, when narrow, right under) the row you picked; once you stop scrolling, the open table comes back to the top of the window (scrolling down over the table itself leaves it, so its rows can be read). In a terminal each panel scrolls on its own: the wheel moves the part under the pointer, and over the bar under the open object it moves the columns sideways. The ▴ ▾ ◂ ▸ steps at the ends of the bars do the same with a click; the ◂ ▸ bar over a wide grid works in the desktop app too.
In the Claude desktop app (and VS Code) the pane follows the app's light or dark theme by itself: system , the default, draws its text in the app's own colour and its panels, pills and lines as faint see-through layers, so the same pane sits on either theme and changes with it. /db theme light or /db theme dark paints the panels in fixed warm off-whites or warm greys instead; pick the one that matches your app, since buttons always take the app's own text colour. /db theme system goes back. The choice is kept across sessions. The terminal keeps its own colours whatever the setting.
claude plugin marketplace update claude-dbview
claude plugin update dbview@claude-dbview
Then run /reload-plugins in an open session, or start a new one. Installed from a local folder instead: git pull in that folder, then /reload-plugins .
For saved and DATABASE_URL databases: psql (the PostgreSQL client) on the machine running Claude Code. The engine looks on the PATH and in the usual Homebrew, Postgres.app and Linux locations; set another path with the psql_path option in /config .
For containers: docker on that machine. The engine runs docker exec … psql inside the container as its POSTGRES_USER on its POSTGRES_DB , so it needs no published port, password or host psql.
types/index.d.ts : the contract between the two sides (sources, catalog, relation detail, preview, DDL; the pane's view state).
hooks/engine/ : the server side. It runs inside the Claude Code CLI and talks to Postgres through psql (or docker exec … psql ), one JSON document per request, built in SQL. The password travels in PGPASSWORD , never on a command line.
hooks/client/ : the client side. It draws the pane from one view value with the surface's own elements, so the same tree works on the terminal, desktop, VS Code and mobile.
hooks/register.tsx :

[truncated]

## Original Extract

Custom generic db view build in Claude Code using mod - agntflw/claude-dbview

GitHub - agntflw/claude-dbview: Custom generic db view build in Claude Code using mod · GitHub
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
agntflw
claude-dbview Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Custom generic db view build in Claude Code using mod
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
35 Commits 35 Commits Folders and files
.claude-plugin .claude-plugin dev dev docs/ images docs/ images hooks hooks tests tests types types .gitattributes .gitattributes .gitignore .gitignore LICENSE LICENSE README.md README.md tsconfig.json tsconfig.json View all files Repository files navigation
A PostgreSQL database view for Claude Code, in a pane beside the chat: every database and Postgres container as an IDE-style tree, with columns, constraints, rows and DDL. Read-only.
The pane on the desktop app, as it draws the sample database in dev/seed.sql . On the terminal it uses glyphs and no borders.
Install, once, for every project (user scope):
claude plugin install dbview --marketplace agntflw/claude-dbview --config psql_path=psql
Start a new Claude Code session (one already open does not see the plugin until you run /reload-plugins in it), then type /db . The pane opens on the databases it finds: your DATABASE_URL or PG* variables, and every running Postgres container.
To add another one: /db connect postgres://user:password@host:5432/database [name] .
Press a database, then a table. Switch between Columns, Rows (pick 50, 100, 500 or 1000) and DDL.
Paste this into Claude Code, in any folder:
Install the dbview plugin for Claude Code globally for me (user scope), from github.com/agntflw/claude-dbview.
1. Run: claude plugin install dbview --marketplace agntflw/claude-dbview --scope user --config psql_path=psql
If this Claude Code is too old for --marketplace, run instead:
claude plugin marketplace add agntflw/claude-dbview, then claude plugin install dbview@claude-dbview --scope user
2. Run: claude plugin list, and check dbview is installed and enabled.
3. Check what it needs: psql --version, and docker --version if I use Postgres in Docker.
Tell me what is missing and how to install it on this machine, but do not install it.
4. Tell me to start a new Claude Code session and type /db.
Add MySQL beside Postgres
dbview reads PostgreSQL only. To have Claude add MySQL (8.4 LTS, or pinned to 8.4.5) to your copy, paste this into Claude Code, in any folder:
Add MySQL support to my dbview plugin (github.com/agntflw/claude-dbview), beside PostgreSQL,
without changing how Postgres works.
Find the code first: use the claude-dbview folder that `claude plugin list` shows dbview is read from,
else one already on this machine, else clone the repository to ~/.claude/plugins-src/claude-dbview.
Work in that folder.
Set up a test database first:
- Run MySQL 8.4 in Docker: docker run -d --name mysql84 -e MYSQL_ROOT_PASSWORD=dev
-e MYSQL_DATABASE=shop -p 33306:3306 mysql:8.4 (the latest 8.4 LTS; write mysql:8.4.5 here to pin it).
- Wait until it accepts connections, then load a MySQL version of dev/seed.sql
(schemas become databases; serial becomes AUTO_INCREMENT; keep the same tables, keys and rows).
Then extend the mod, following its layout (types first, then engine, then pane):
1. types/index.d.ts: add an engine field ('postgres' | 'mysql') to DbSource, and keep every request shape.
2. hooks/engine/: add a MySQL driver next to psql.ts. Run the mysql client inside the container
(docker exec -i <container> mysql -N -B ...) or on the host for mysql:// URLs.
Pass the password in MYSQL_PWD, never on the command line.
Every session is read-only: SET SESSION TRANSACTION READ ONLY and max_execution_time=10000.
Build each answer as one JSON document with JSON_OBJECT and JSON_ARRAYAGG
from information_schema (catalog, columns, constraints, indexes).
Use SHOW CREATE TABLE / VIEW for the DDL.
3. hooks/engine/docker.ts: also discover mysql and mariadb containers
(user root with MYSQL_ROOT_PASSWORD, or MYSQL_USER and MYSQL_PASSWORD; database from MYSQL_DATABASE).
4. /db connect: accept mysql:// URLs as well as postgres://.
5. Pane: show the engine and version on each database row and in the switcher
(MySQL 8.4 · docker). MySQL has no schemas inside a database, so skip the schema level for it.
Sequences and materialized views do not exist there.
Prove it:
- Extend dev/engine-e2e.ts with MySQL checks against mysql84
(catalog, columns, a foreign
[truncated]
Each database is a root of the tree, with its schemas as packages inside, then Tables, Views, Materialized views, Sequences and Functions folders, then the objects and their columns. Open a table or view to see:
Columns : name, type, nullable, primary key or → schema.table(column) reference, defaults, then the constraints (keys, foreign keys, checks, not null), indexes, and the foreign keys elsewhere that point at it ( Referenced by ).
Rows : the first 50, 100, 500 or 1000 rows in key order, with the query shown.
DDL : the table as a migration would create it (its sequences, the table with defaults and named constraints, indexes, OWNED BY ), numbered, with Copy.
Foreign keys are links that look like plain text. Press a → schema.table(column) reference, a foreign key constraint, a ← schema.table(column) line under Referenced by, or a key column in the tree, and the table at the other end opens on the same tab, unfolded and scrolled into view in the tree. ← name at the top right goes back, one jump at a time.
Every Postgres container Docker is running shows up by itself, so the databases a migration test leaves behind (say mig_before and mig_after ) sit side by side in the tree to compare.
Everything is read-only: each request runs in a read-only session with a 10 second statement timeout.
At the prompt of a terminal session:
/plugin install dbview --marketplace agntflw/claude-dbview
Answer y to add the marketplace, then pick a scope (user is the first). From a shell:
claude plugin install dbview --marketplace agntflw/claude-dbview
In the desktop app, where /plugin does not run: click + next to the prompt box, then Plugins, Add marketplace with agntflw/claude-dbview , then Add plugin.
On a Claude Code older than v2.1.292, add the marketplace first, then install from it:
claude plugin marketplace add agntflw/claude-dbview
claude plugin install dbview@claude-dbview
To update: claude plugin update dbview@claude-dbview , or turn on auto-update for claude-dbview in /plugin › Marketplaces.
On a claude.ai Team or Enterprise plan, an Owner can also add dbview for everyone in Organization settings › Plugins & skills ; it then reaches each member's Claude Code as dbview@synced .
Start a new session, or run /reload-plugins in the open one. A plugin installed from a shell loads in sessions started after the install.
Check claude --version : mods (plugins with a hooks module, like this one) need a recent Claude Code. Run claude update if it is old.
Run claude plugin list : dbview should say enabled. In a session, /plugin shows any error loading it.
The "userConfig option not yet set" note after install is harmless: psql_path defaults to psql .
/db open the pane on every database it finds
/db connect postgres://u:p@host:5432/db [name] save a database under a name (default: its database name)
/db forget <name> forget a saved database
/db refresh list the databases again and reload the open ones
/db theme [system|light|dark] the pane's colours in the desktop app (default: system)
The pane lists, in order: the databases saved with /db connect , the one named by DATABASE_URL or the PG* variables (unless a saved one is the same), then each running Postgres container. Press a database to unfold it (its catalog loads the first time), or pick one in the switcher at the top. Press a schema or folder to fold it, press a table to open it (press again to fold its columns), and filter the tree by table or column name. On the terminal, with the pane focused: c , w , d switch to Columns, Rows, DDL; r refreshes; y copies the DDL; b goes back after a jump.
The pane follows its width: the tree and the open object sit side by side when it is wide, the object under the tree when it is narrow, and the tree takes about a third of a wide pane. In the Claude desktop app the pane scrolls as one page with the wheel or trackpad, and a table opens beside (or, when narrow, right under) the row you picked; once you stop scrolling, the open table comes back to the top of the window (scrolling down over the table itself leaves it, so its rows can be read). In a terminal each panel scrolls on its own: the wheel moves the part under the pointer, and over the bar under the open object it moves the columns sideways. The ▴ ▾ ◂ ▸ steps at the ends of the bars do the same with a click; the ◂ ▸ bar over a wide grid works in the desktop app too.
In the Claude desktop app (and VS Code) the pane follows the app's light or dark theme by itself: system , the default, draws its text in the app's own colour and its panels, pills and lines as faint see-through layers, so the same pane sits on either theme and changes with it. /db theme light or /db theme dark paints the panels in fixed warm off-whites or warm greys instead; pick the one that matches your app, since buttons always take the app's own text colour. /db theme system goes back. The choice is kept across sessions. The terminal keeps its own colours whatever the setting.
claude plugin marketplace update claude-dbview
claude plugin update dbview@claude-dbview
Then run /reload-plugins in an open session, or start a new one. Installed from a local folder instead: git pull in that folder, then /reload-plugins .
For saved and DATABASE_URL databases: psql (the PostgreSQL client) on the machine running Claude Code. The engine looks on the PATH and in the usual Homebrew, Postgres.app and Linux locations; set another path with the psql_path option in /config .
For containers: docker on that machine. The engine runs docker exec … psql inside the container as its POSTGRES_USER on its POSTGRES_DB , so it needs no published port, password or host psql.
types/index.d.ts : the contract between the two sides (sources, catalog, relation detail, preview, DDL; the pane's view state).
hooks/engine/ : the server side. It runs inside the Claude Code CLI and talks to Postgres through psql (or docker exec … psql ), one JSON document per request, built in SQL. The password travels in PGPASSWORD , never on a command line.
hooks/client/ : the client side. It draws the pane from one view value with the surface's own elements, so the same tree works on the terminal, desktop, VS Code and mobile.
hooks/register.tsx :

[truncated]
