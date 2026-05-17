# Shared Response Types

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared#APIError">APIError</a>
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
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AccountBridge">AccountBridge</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AccountGetResponse">AccountGetResponse</a>

Methods:

- <code title="get /v1/accounts/{accountID}">client.Accounts.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AccountGetResponse">AccountGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Contacts

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AccountContactSearchResponse">AccountContactSearchResponse</a>

Methods:

- <code title="get /v1/accounts/{accountID}/contacts/list">client.Accounts.Contacts.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AccountContactService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AccountContactListParams">AccountContactListParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/packages/pagination#CursorSearch">CursorSearch</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5/shared#User">User</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/accounts/{accountID}/contacts">client.Accounts.Contacts.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AccountContactService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AccountContactSearchParams">AccountContactSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#AccountContactSearchResponse">AccountContactSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Bridges

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#Bridge">Bridge</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#CookieField">CookieField</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#DisappearingTimerCapability">DisappearingTimerCapability</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#GroupFieldCapability">GroupFieldCapability</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#GroupTypeCapabilities">GroupTypeCapabilities</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#LoginFlow">LoginFlow</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#LoginInputField">LoginInputField</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#LoginSession">LoginSession</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ProvisioningCapabilities">ProvisioningCapabilities</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ResolveIdentifierCapabilities">ResolveIdentifierCapabilities</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeGetResponse">BridgeGetResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeListResponse">BridgeListResponse</a>

Methods:

- <code title="get /v1/bridges/{bridgeID}">client.Bridges.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bridgeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeGetResponse">BridgeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/bridges">client.Bridges.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeListResponse">BridgeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/bridges/{bridgeID}/capabilities">client.Bridges.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeService.GetCapabilities">GetCapabilities</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bridgeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#ProvisioningCapabilities">ProvisioningCapabilities</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## LoginFlows

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeLoginFlowListResponse">BridgeLoginFlowListResponse</a>

Methods:

- <code title="get /v1/bridges/{bridgeID}/login-flows">client.Bridges.LoginFlows.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeLoginFlowService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bridgeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeLoginFlowListResponse">BridgeLoginFlowListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Connections

## LoginSessions

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeLoginSessionCancelResponse">BridgeLoginSessionCancelResponse</a>

Methods:

- <code title="post /v1/bridges/{bridgeID}/login-sessions">client.Bridges.LoginSessions.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeLoginSessionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bridgeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeLoginSessionNewParams">BridgeLoginSessionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#LoginSession">LoginSession</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/bridges/{bridgeID}/login-sessions/{loginSessionID}">client.Bridges.LoginSessions.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeLoginSessionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, loginSessionID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeLoginSessionGetParams">BridgeLoginSessionGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#LoginSession">LoginSession</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/bridges/{bridgeID}/login-sessions/{loginSessionID}">client.Bridges.LoginSessions.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeLoginSessionService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, loginSessionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeLoginSessionCancelParams">BridgeLoginSessionCancelParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeLoginSessionCancelResponse">BridgeLoginSessionCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Steps

Methods:

- <code title="post /v1/bridges/{bridgeID}/login-sessions/{loginSessionID}/steps/{stepID}">client.Bridges.LoginSessions.Steps.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeLoginSessionStepService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stepID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#BridgeLoginSessionStepSubmitParams">BridgeLoginSessionStepSubmitParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v5#LoginSession">LoginSession</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

## Login

### Verification

#### RecoveryKey

##### Reset

## Verifications

### Qr

### SAS
