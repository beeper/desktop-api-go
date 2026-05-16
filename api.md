# Shared Response Types

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared#Attachment">Attachment</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared#Message">Message</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared#Reaction">Reaction</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared#User">User</a>

# beeperdesktopapi

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#FocusResponse">FocusResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#SearchResponse">SearchResponse</a>

Methods:

- <code title="post /v1/focus">client.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BeeperdesktopapiService.Focus">Focus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#FocusParams">FocusParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#FocusResponse">FocusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/search">client.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BeeperdesktopapiService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#SearchParams">SearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#SearchResponse">SearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#Account">Account</a>

Methods:

- <code title="get /v1/accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Contacts

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AccountContactSearchResponse">AccountContactSearchResponse</a>

Methods:

- <code title="get /v1/accounts/{accountID}/contacts/list">client.Accounts.Contacts.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AccountContactService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AccountContactListParams">AccountContactListParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/packages/pagination#CursorSearch">CursorSearch</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared#User">User</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/accounts/{accountID}/contacts">client.Accounts.Contacts.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AccountContactService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AccountContactSearchParams">AccountContactSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AccountContactSearchResponse">AccountContactSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Bridges

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeAvailability">BridgeAvailability</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeListResponse">BridgeListResponse</a>

Methods:

- <code title="get /v1/bridges">client.Bridges.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeListResponse">BridgeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chats

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#Chat">Chat</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatNewResponse">ChatNewResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatListResponse">ChatListResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatStartResponse">ChatStartResponse</a>

Methods:

- <code title="post /v1/chats">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatNewParams">ChatNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatNewResponse">ChatNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/chats/{chatID}">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatGetParams">ChatGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#Chat">Chat</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/chats/{chatID}">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatUpdateParams">ChatUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#Chat">Chat</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/chats">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatListParams">ChatListParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/packages/pagination#CursorNoLimit">CursorNoLimit</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatListResponse">ChatListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/chats/{chatID}/archive">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatArchiveParams">ChatArchiveParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/chats/{chatID}/read">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatService.MarkRead">MarkRead</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatMarkReadParams">ChatMarkReadParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#Chat">Chat</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/chats/{chatID}/unread">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatService.MarkUnread">MarkUnread</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatMarkUnreadParams">ChatMarkUnreadParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#Chat">Chat</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/chats/{chatID}/notify-anyway">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatService.NotifyAnyway">NotifyAnyway</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatNotifyAnywayParams">ChatNotifyAnywayParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#Chat">Chat</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/chats/search">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatSearchParams">ChatSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/packages/pagination#CursorSearch">CursorSearch</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#Chat">Chat</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/chats/start">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatStartParams">ChatStartParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatStartResponse">ChatStartResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Reminders

Methods:

- <code title="post /v1/chats/{chatID}/reminders">client.Chats.Reminders.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatReminderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatReminderNewParams">ChatReminderNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /v1/chats/{chatID}/reminders">client.Chats.Reminders.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatReminderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Messages

### Reactions

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatMessageReactionDeleteResponse">ChatMessageReactionDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatMessageReactionAddResponse">ChatMessageReactionAddResponse</a>

Methods:

- <code title="delete /v1/chats/{chatID}/messages/{messageID}/reactions/{reactionKey}">client.Chats.Messages.Reactions.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatMessageReactionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, reactionKey <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatMessageReactionDeleteParams">ChatMessageReactionDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatMessageReactionDeleteResponse">ChatMessageReactionDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/chats/{chatID}/messages/{messageID}/reactions">client.Chats.Messages.Reactions.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatMessageReactionService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatMessageReactionAddParams">ChatMessageReactionAddParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ChatMessageReactionAddResponse">ChatMessageReactionAddResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MessageUpdateResponse">MessageUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MessageSendResponse">MessageSendResponse</a>

Methods:

- <code title="get /v1/chats/{chatID}/messages/{messageID}">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MessageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MessageGetParams">MessageGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared#Message">Message</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/chats/{chatID}/messages/{messageID}">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MessageService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MessageUpdateParams">MessageUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MessageUpdateResponse">MessageUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/chats/{chatID}/messages">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MessageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MessageListParams">MessageListParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/packages/pagination#CursorNoLimit">CursorNoLimit</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared#Message">Message</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/chats/{chatID}/messages/{messageID}">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MessageService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MessageDeleteParams">MessageDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/messages/search">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MessageService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MessageSearchParams">MessageSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/packages/pagination#CursorSearch">CursorSearch</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared#Message">Message</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/chats/{chatID}/messages">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MessageService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MessageSendParams">MessageSendParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MessageSendResponse">MessageSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AssetDownloadResponse">AssetDownloadResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AssetUploadResponse">AssetUploadResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AssetUploadBase64Response">AssetUploadBase64Response</a>

Methods:

- <code title="post /v1/assets/download">client.Assets.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AssetService.Download">Download</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AssetDownloadParams">AssetDownloadParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AssetDownloadResponse">AssetDownloadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/assets/serve">client.Assets.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AssetService.Serve">Serve</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AssetServeParams">AssetServeParams</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/assets/upload">client.Assets.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AssetService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AssetUploadParams">AssetUploadParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AssetUploadResponse">AssetUploadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/assets/upload/base64">client.Assets.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AssetService.UploadBase64">UploadBase64</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AssetUploadBase64Params">AssetUploadBase64Params</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AssetUploadBase64Response">AssetUploadBase64Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Info

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#InfoGetResponse">InfoGetResponse</a>

Methods:

- <code title="get /v1/info">client.Info.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#InfoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#InfoGetResponse">InfoGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# App

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppStatusResponse">AppStatusResponse</a>

Methods:

- <code title="get /v1/app/status">client.App.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppService.Status">Status</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppStatusResponse">AppStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Login

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppLoginEmailResponse">AppLoginEmailResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppLoginRegisterResponse">AppLoginRegisterResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppLoginResponseResponseUnion">AppLoginResponseResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppLoginStartResponse">AppLoginStartResponse</a>

Methods:

- <code title="post /v1/app/login/email">client.App.Login.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppLoginService.Email">Email</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppLoginEmailParams">AppLoginEmailParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppLoginEmailResponse">AppLoginEmailResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/app/login/register">client.App.Login.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppLoginService.Register">Register</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppLoginRegisterParams">AppLoginRegisterParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppLoginRegisterResponse">AppLoginRegisterResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/app/login/response">client.App.Login.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppLoginService.Response">Response</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppLoginResponseParams">AppLoginResponseParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppLoginResponseResponseUnion">AppLoginResponseResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/app/login/start">client.App.Login.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppLoginService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppLoginStartResponse">AppLoginStartResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## E2ee

### RecoveryCode

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeRecoveryCodeMarkBackedUpResponse">AppE2eeRecoveryCodeMarkBackedUpResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeRecoveryCodeVerifyResponse">AppE2eeRecoveryCodeVerifyResponse</a>

Methods:

- <code title="post /v1/app/e2ee/recovery-code/mark-backed-up">client.App.E2ee.RecoveryCode.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeRecoveryCodeService.MarkBackedUp">MarkBackedUp</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeRecoveryCodeMarkBackedUpResponse">AppE2eeRecoveryCodeMarkBackedUpResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/app/e2ee/recovery-code/verify">client.App.E2ee.RecoveryCode.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeRecoveryCodeService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeRecoveryCodeVerifyParams">AppE2eeRecoveryCodeVerifyParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeRecoveryCodeVerifyResponse">AppE2eeRecoveryCodeVerifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Reset

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeRecoveryCodeResetNewResponse">AppE2eeRecoveryCodeResetNewResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeRecoveryCodeResetConfirmResponse">AppE2eeRecoveryCodeResetConfirmResponse</a>

Methods:

- <code title="post /v1/app/e2ee/recovery-code/reset">client.App.E2ee.RecoveryCode.Reset.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeRecoveryCodeResetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeRecoveryCodeResetNewParams">AppE2eeRecoveryCodeResetNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeRecoveryCodeResetNewResponse">AppE2eeRecoveryCodeResetNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/app/e2ee/recovery-code/reset/confirm">client.App.E2ee.RecoveryCode.Reset.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeRecoveryCodeResetService.Confirm">Confirm</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeRecoveryCodeResetConfirmParams">AppE2eeRecoveryCodeResetConfirmParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeRecoveryCodeResetConfirmResponse">AppE2eeRecoveryCodeResetConfirmResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Verification

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationNewResponse">AppE2eeVerificationNewResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationAcceptResponse">AppE2eeVerificationAcceptResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationCancelResponse">AppE2eeVerificationCancelResponse</a>

Methods:

- <code title="post /v1/app/e2ee/verification">client.App.E2ee.Verification.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationNewParams">AppE2eeVerificationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationNewResponse">AppE2eeVerificationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/app/e2ee/verification/{verificationID}/accept">client.App.E2ee.Verification.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationService.Accept">Accept</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, verificationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationAcceptResponse">AppE2eeVerificationAcceptResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/app/e2ee/verification/{verificationID}/cancel">client.App.E2ee.Verification.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, verificationID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationCancelParams">AppE2eeVerificationCancelParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationCancelResponse">AppE2eeVerificationCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Qr

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationQrConfirmScannedResponse">AppE2eeVerificationQrConfirmScannedResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationQrScanResponse">AppE2eeVerificationQrScanResponse</a>

Methods:

- <code title="post /v1/app/e2ee/verification/{verificationID}/qr/confirm-scanned">client.App.E2ee.Verification.Qr.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationQrService.ConfirmScanned">ConfirmScanned</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, verificationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationQrConfirmScannedResponse">AppE2eeVerificationQrConfirmScannedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/app/e2ee/verification/qr/scan">client.App.E2ee.Verification.Qr.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationQrService.Scan">Scan</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationQrScanParams">AppE2eeVerificationQrScanParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationQrScanResponse">AppE2eeVerificationQrScanResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Sas

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationSaConfirmResponse">AppE2eeVerificationSaConfirmResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationSaStartResponse">AppE2eeVerificationSaStartResponse</a>

Methods:

- <code title="post /v1/app/e2ee/verification/{verificationID}/sas/confirm">client.App.E2ee.Verification.Sas.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationSaService.Confirm">Confirm</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, verificationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationSaConfirmResponse">AppE2eeVerificationSaConfirmResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/app/e2ee/verification/{verificationID}/sas/start">client.App.E2ee.Verification.Sas.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationSaService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, verificationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AppE2eeVerificationSaStartResponse">AppE2eeVerificationSaStartResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Matrix

## Users

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixUserGetProfileResponse">MatrixUserGetProfileResponse</a>

Methods:

- <code title="get /_matrix/client/v3/profile/{userId}">client.Matrix.Users.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixUserService.GetProfile">GetProfile</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixUserGetProfileResponse">MatrixUserGetProfileResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### AccountData

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixUserAccountDataGetResponse">MatrixUserAccountDataGetResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixUserAccountDataUpdateResponse">MatrixUserAccountDataUpdateResponse</a>

Methods:

- <code title="get /_matrix/client/v3/user/{userId}/account_data/{type}">client.Matrix.Users.AccountData.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixUserAccountDataService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, type\_ <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixUserAccountDataGetParams">MatrixUserAccountDataGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixUserAccountDataGetResponse">MatrixUserAccountDataGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /_matrix/client/v3/user/{userId}/account_data/{type}">client.Matrix.Users.AccountData.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixUserAccountDataService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, type\_ <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixUserAccountDataUpdateParams">MatrixUserAccountDataUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixUserAccountDataUpdateResponse">MatrixUserAccountDataUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Rooms

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomNewResponse">MatrixRoomNewResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomJoinResponse">MatrixRoomJoinResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomLeaveResponse">MatrixRoomLeaveResponse</a>

Methods:

- <code title="post /_matrix/client/v3/createRoom">client.Matrix.Rooms.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomNewParams">MatrixRoomNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomNewResponse">MatrixRoomNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /_matrix/client/v3/join/{roomIdOrAlias}">client.Matrix.Rooms.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomService.Join">Join</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, roomIDOrAlias <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomJoinParams">MatrixRoomJoinParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomJoinResponse">MatrixRoomJoinResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /_matrix/client/v3/rooms/{roomId}/leave">client.Matrix.Rooms.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomService.Leave">Leave</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, roomID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomLeaveParams">MatrixRoomLeaveParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomLeaveResponse">MatrixRoomLeaveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### AccountData

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomAccountDataGetResponse">MatrixRoomAccountDataGetResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomAccountDataUpdateResponse">MatrixRoomAccountDataUpdateResponse</a>

Methods:

- <code title="get /_matrix/client/v3/user/{userId}/rooms/{roomId}/account_data/{type}">client.Matrix.Rooms.AccountData.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomAccountDataService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, type\_ <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomAccountDataGetParams">MatrixRoomAccountDataGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomAccountDataGetResponse">MatrixRoomAccountDataGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /_matrix/client/v3/user/{userId}/rooms/{roomId}/account_data/{type}">client.Matrix.Rooms.AccountData.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomAccountDataService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, type\_ <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomAccountDataUpdateParams">MatrixRoomAccountDataUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomAccountDataUpdateResponse">MatrixRoomAccountDataUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### State

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomStateGetResponse">MatrixRoomStateGetResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomStateListResponse">MatrixRoomStateListResponse</a>

Methods:

- <code title="get /_matrix/client/v3/rooms/{roomId}/state/{eventType}/{stateKey}">client.Matrix.Rooms.State.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomStateService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stateKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomStateGetParams">MatrixRoomStateGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomStateGetResponse">MatrixRoomStateGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /_matrix/client/v3/rooms/{roomId}/state">client.Matrix.Rooms.State.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomStateService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, roomID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*[]<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomStateListResponse">MatrixRoomStateListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Events

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomEventGetResponse">MatrixRoomEventGetResponse</a>

Methods:

- <code title="get /_matrix/client/v3/rooms/{roomId}/event/{eventId}">client.Matrix.Rooms.Events.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomEventService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, eventID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomEventGetParams">MatrixRoomEventGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixRoomEventGetResponse">MatrixRoomEventGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Bridges

### Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthListFlowsResponse">MatrixBridgeAuthListFlowsResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthListLoginsResponse">MatrixBridgeAuthListLoginsResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthLogoutResponse">MatrixBridgeAuthLogoutResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthStartLoginResponseUnion">MatrixBridgeAuthStartLoginResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthSubmitCookiesResponseUnion">MatrixBridgeAuthSubmitCookiesResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthSubmitUserInputResponseUnion">MatrixBridgeAuthSubmitUserInputResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthWaitForStepResponseUnion">MatrixBridgeAuthWaitForStepResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthWhoamiResponse">MatrixBridgeAuthWhoamiResponse</a>

Methods:

- <code title="get /_matrix/client/unstable/com.beeper.bridge/{bridgeID}/_matrix/provision/v3/login/flows">client.Matrix.Bridges.Auth.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthService.ListFlows">ListFlows</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bridgeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthListFlowsResponse">MatrixBridgeAuthListFlowsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /_matrix/client/unstable/com.beeper.bridge/{bridgeID}/_matrix/provision/v3/logins">client.Matrix.Bridges.Auth.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthService.ListLogins">ListLogins</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bridgeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthListLoginsResponse">MatrixBridgeAuthListLoginsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /_matrix/client/unstable/com.beeper.bridge/{bridgeID}/_matrix/provision/v3/logout/{loginID}">client.Matrix.Bridges.Auth.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthService.Logout">Logout</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, loginID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthLogoutParams">MatrixBridgeAuthLogoutParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthLogoutResponse">MatrixBridgeAuthLogoutResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /_matrix/client/unstable/com.beeper.bridge/{bridgeID}/_matrix/provision/v3/login/start/{flowID}">client.Matrix.Bridges.Auth.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthService.StartLogin">StartLogin</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, flowID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthStartLoginParams">MatrixBridgeAuthStartLoginParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthStartLoginResponseUnion">MatrixBridgeAuthStartLoginResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /_matrix/client/unstable/com.beeper.bridge/{bridgeID}/_matrix/provision/v3/login/step/{loginProcessID}/{stepID}/cookies">client.Matrix.Bridges.Auth.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthService.SubmitCookies">SubmitCookies</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stepID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthSubmitCookiesParams">MatrixBridgeAuthSubmitCookiesParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthSubmitCookiesResponseUnion">MatrixBridgeAuthSubmitCookiesResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /_matrix/client/unstable/com.beeper.bridge/{bridgeID}/_matrix/provision/v3/login/step/{loginProcessID}/{stepID}/user_input">client.Matrix.Bridges.Auth.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthService.SubmitUserInput">SubmitUserInput</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stepID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthSubmitUserInputParams">MatrixBridgeAuthSubmitUserInputParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthSubmitUserInputResponseUnion">MatrixBridgeAuthSubmitUserInputResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /_matrix/client/unstable/com.beeper.bridge/{bridgeID}/_matrix/provision/v3/login/step/{loginProcessID}/{stepID}/display_and_wait">client.Matrix.Bridges.Auth.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthService.WaitForStep">WaitForStep</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stepID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthWaitForStepParams">MatrixBridgeAuthWaitForStepParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthWaitForStepResponseUnion">MatrixBridgeAuthWaitForStepResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /_matrix/client/unstable/com.beeper.bridge/{bridgeID}/_matrix/provision/v3/whoami">client.Matrix.Bridges.Auth.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthService.Whoami">Whoami</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bridgeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeAuthWhoamiResponse">MatrixBridgeAuthWhoamiResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Contacts

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeContactListResponse">MatrixBridgeContactListResponse</a>

Methods:

- <code title="get /_matrix/client/unstable/com.beeper.bridge/{bridgeID}/_matrix/provision/v3/contacts">client.Matrix.Bridges.Contacts.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeContactService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bridgeID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeContactListParams">MatrixBridgeContactListParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeContactListResponse">MatrixBridgeContactListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Users

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeUserResolveResponse">MatrixBridgeUserResolveResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeUserSearchResponse">MatrixBridgeUserSearchResponse</a>

Methods:

- <code title="get /_matrix/client/unstable/com.beeper.bridge/{bridgeID}/_matrix/provision/v3/resolve_identifier/{identifier}">client.Matrix.Bridges.Users.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeUserService.Resolve">Resolve</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeUserResolveParams">MatrixBridgeUserResolveParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeUserResolveResponse">MatrixBridgeUserResolveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /_matrix/client/unstable/com.beeper.bridge/{bridgeID}/_matrix/provision/v3/search_users">client.Matrix.Bridges.Users.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeUserService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bridgeID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeUserSearchParams">MatrixBridgeUserSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeUserSearchResponse">MatrixBridgeUserSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Rooms

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeRoomNewDmResponse">MatrixBridgeRoomNewDmResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeRoomNewGroupResponse">MatrixBridgeRoomNewGroupResponse</a>

Methods:

- <code title="post /_matrix/client/unstable/com.beeper.bridge/{bridgeID}/_matrix/provision/v3/create_dm/{identifier}">client.Matrix.Bridges.Rooms.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeRoomService.NewDm">NewDm</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeRoomNewDmParams">MatrixBridgeRoomNewDmParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeRoomNewDmResponse">MatrixBridgeRoomNewDmResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /_matrix/client/unstable/com.beeper.bridge/{bridgeID}/_matrix/provision/v3/create_group/{groupType}">client.Matrix.Bridges.Rooms.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeRoomService.NewGroup">NewGroup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, groupType <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeRoomNewGroupParams">MatrixBridgeRoomNewGroupParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeRoomNewGroupResponse">MatrixBridgeRoomNewGroupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Capabilities

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeCapabilityGetResponse">MatrixBridgeCapabilityGetResponse</a>

Methods:

- <code title="get /_matrix/client/unstable/com.beeper.bridge/{bridgeID}/_matrix/provision/v3/capabilities">client.Matrix.Bridges.Capabilities.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeCapabilityService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bridgeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#MatrixBridgeCapabilityGetResponse">MatrixBridgeCapabilityGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
