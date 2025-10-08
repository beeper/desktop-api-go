# Shared Response Types

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#Attachment">Attachment</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#BaseResponse">BaseResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#Message">Message</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#Reaction">Reaction</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#User">User</a>

# beeperdesktopapi

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#DownloadAssetResponse">DownloadAssetResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OpenResponse">OpenResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#SearchResponse">SearchResponse</a>

Methods:

- <code title="post /v1/download-asset">client.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#BeeperdesktopapiService.DownloadAsset">DownloadAsset</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#DownloadAssetParams">DownloadAssetParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#DownloadAssetResponse">DownloadAssetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/open">client.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#BeeperdesktopapiService.Open">Open</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OpenParams">OpenParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#OpenResponse">OpenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/search">client.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#BeeperdesktopapiService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#SearchParams">SearchParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#SearchResponse">SearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#Account">Account</a>

Methods:

- <code title="get /v1/accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Contacts

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ContactSearchResponse">ContactSearchResponse</a>

Methods:

- <code title="get /v1/accounts/{accountID}/contacts/search">client.Contacts.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ContactService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ContactSearchParams">ContactSearchParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ContactSearchResponse">ContactSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chats

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#Chat">Chat</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatNewResponse">ChatNewResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatListResponse">ChatListResponse</a>

Methods:

- <code title="post /v1/chats">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatNewParams">ChatNewParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatNewResponse">ChatNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/chats/{chatID}">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatGetParams">ChatGetParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#Chat">Chat</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/chats">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatListParams">ChatListParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/packages/pagination#CursorList">CursorList</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatListResponse">ChatListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/chats/{chatID}/archive">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatArchiveParams">ChatArchiveParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#BaseResponse">BaseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/chats/search">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatSearchParams">ChatSearchParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/packages/pagination#CursorSearch">CursorSearch</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#Chat">Chat</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Reminders

Methods:

- <code title="post /v1/chats/{chatID}/reminders">client.Chats.Reminders.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatReminderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatReminderNewParams">ChatReminderNewParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#BaseResponse">BaseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/chats/{chatID}/reminders">client.Chats.Reminders.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#ChatReminderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#BaseResponse">BaseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageSendResponse">MessageSendResponse</a>

Methods:

- <code title="get /v1/chats/{chatID}/messages">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageListParams">MessageListParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/packages/pagination#CursorList">CursorList</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#Message">Message</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messages/search">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageSearchParams">MessageSearchParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/packages/pagination#CursorSearch">CursorSearch</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/shared#Message">Message</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/chats/{chatID}/messages">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageSendParams">MessageSendParams</a>) (<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go#MessageSendResponse">MessageSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
