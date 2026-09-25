# Shared Response Types

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared#APIError">APIError</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared#Attachment">Attachment</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared#AttachmentCapabilities">AttachmentCapabilities</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared#ChatCapabilities">ChatCapabilities</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared#ChatDraft">ChatDraft</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared#ChatStateCapabilities">ChatStateCapabilities</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared#DraftAttachment">DraftAttachment</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared#LinkPreview">LinkPreview</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared#Message">Message</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared#Reaction">Reaction</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared#SendStatus">SendStatus</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared#User">User</a>

# beeperdesktopapi

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#FocusResponse">FocusResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#SearchResponse">SearchResponse</a>

Methods:

- <code title="post /v1/focus">client.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BeeperdesktopapiService.Focus">Focus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#FocusParams">FocusParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#FocusResponse">FocusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/search">client.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BeeperdesktopapiService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#SearchParams">SearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#SearchResponse">SearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#Account">Account</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AccountBridge">AccountBridge</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AccountGetResponse">AccountGetResponse</a>

Methods:

- <code title="get /v1/accounts/{accountID}">client.Accounts.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AccountGetResponse">AccountGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Contacts

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AccountContactSearchResponse">AccountContactSearchResponse</a>

Methods:

- <code title="get /v1/accounts/{accountID}/contacts/list">client.Accounts.Contacts.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AccountContactService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AccountContactListParams">AccountContactListParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/packages/pagination#CursorSearch">CursorSearch</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared#User">User</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/accounts/{accountID}/contacts">client.Accounts.Contacts.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AccountContactService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AccountContactSearchParams">AccountContactSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AccountContactSearchResponse">AccountContactSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Bridges

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#Bridge">Bridge</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLogin">BridgeLogin</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#CookieField">CookieField</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#DisappearingTimerCapability">DisappearingTimerCapability</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#GroupFieldCapability">GroupFieldCapability</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#GroupTypeCapabilities">GroupTypeCapabilities</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#LoginFlow">LoginFlow</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#LoginInputField">LoginInputField</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#LoginSession">LoginSession</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ProvisioningCapabilities">ProvisioningCapabilities</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ResolveIdentifierCapabilities">ResolveIdentifierCapabilities</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeGetResponse">BridgeGetResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeListResponse">BridgeListResponse</a>

Methods:

- <code title="get /v1/bridges/{bridgeID}">client.Bridges.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bridgeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeGetResponse">BridgeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/bridges">client.Bridges.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeListResponse">BridgeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/bridges/{bridgeID}/capabilities">client.Bridges.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeService.GetCapabilities">GetCapabilities</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bridgeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ProvisioningCapabilities">ProvisioningCapabilities</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## LoginFlows

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginFlowListResponse">BridgeLoginFlowListResponse</a>

Methods:

- <code title="get /v1/bridges/{bridgeID}/login-flows">client.Bridges.LoginFlows.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginFlowService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bridgeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginFlowListResponse">BridgeLoginFlowListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Logins

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginListResponse">BridgeLoginListResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginRemoveResponse">BridgeLoginRemoveResponse</a>

Methods:

- <code title="get /v1/bridges/{bridgeID}/logins/{loginID}">client.Bridges.Logins.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, loginID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginGetParams">BridgeLoginGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLogin">BridgeLogin</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/bridges/{bridgeID}/logins">client.Bridges.Logins.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bridgeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginListResponse">BridgeLoginListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/bridges/{bridgeID}/logins/{loginID}/remove">client.Bridges.Logins.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, loginID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginRemoveParams">BridgeLoginRemoveParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginRemoveResponse">BridgeLoginRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## LoginSessions

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginSessionCancelResponse">BridgeLoginSessionCancelResponse</a>

Methods:

- <code title="post /v1/bridges/{bridgeID}/login-sessions">client.Bridges.LoginSessions.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginSessionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, bridgeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginSessionNewParams">BridgeLoginSessionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#LoginSession">LoginSession</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/bridges/{bridgeID}/login-sessions/{loginSessionID}">client.Bridges.LoginSessions.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginSessionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, loginSessionID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginSessionGetParams">BridgeLoginSessionGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#LoginSession">LoginSession</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/bridges/{bridgeID}/login-sessions/{loginSessionID}">client.Bridges.LoginSessions.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginSessionService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, loginSessionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginSessionCancelParams">BridgeLoginSessionCancelParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginSessionCancelResponse">BridgeLoginSessionCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Steps

Methods:

- <code title="post /v1/bridges/{bridgeID}/login-sessions/{loginSessionID}/steps/{stepID}">client.Bridges.LoginSessions.Steps.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginSessionStepService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stepID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#BridgeLoginSessionStepSubmitParams">BridgeLoginSessionStepSubmitParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#LoginSession">LoginSession</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chats

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#Chat">Chat</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatNewResponse">ChatNewResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatListResponse">ChatListResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatStartResponse">ChatStartResponse</a>

Methods:

- <code title="post /v1/chats">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatNewParams">ChatNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatNewResponse">ChatNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/chats/{chatID}">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatGetParams">ChatGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#Chat">Chat</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/chats/{chatID}">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatUpdateParams">ChatUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#Chat">Chat</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/chats">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatListParams">ChatListParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/packages/pagination#CursorNoLimit">CursorNoLimit</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatListResponse">ChatListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/chats/{chatID}/archive">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatArchiveParams">ChatArchiveParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/chats/{chatID}/read">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatService.MarkRead">MarkRead</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatMarkReadParams">ChatMarkReadParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#Chat">Chat</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/chats/{chatID}/unread">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatService.MarkUnread">MarkUnread</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatMarkUnreadParams">ChatMarkUnreadParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#Chat">Chat</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/chats/{chatID}/notify-anyway">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatService.NotifyAnyway">NotifyAnyway</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatNotifyAnywayParams">ChatNotifyAnywayParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#Chat">Chat</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/chats/search">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatSearchParams">ChatSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/packages/pagination#CursorSearch">CursorSearch</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#Chat">Chat</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/chats/start">client.Chats.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatStartParams">ChatStartParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatStartResponse">ChatStartResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Reminders

Methods:

- <code title="post /v1/chats/{chatID}/reminders">client.Chats.Reminders.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatReminderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatReminderNewParams">ChatReminderNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /v1/chats/{chatID}/reminders">client.Chats.Reminders.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatReminderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Messages

### Reactions

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatMessageReactionDeleteResponse">ChatMessageReactionDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatMessageReactionAddResponse">ChatMessageReactionAddResponse</a>

Methods:

- <code title="delete /v1/chats/{chatID}/messages/{messageID}/reactions/{reactionKey}">client.Chats.Messages.Reactions.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatMessageReactionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, reactionKey <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatMessageReactionDeleteParams">ChatMessageReactionDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatMessageReactionDeleteResponse">ChatMessageReactionDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/chats/{chatID}/messages/{messageID}/reactions">client.Chats.Messages.Reactions.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatMessageReactionService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatMessageReactionAddParams">ChatMessageReactionAddParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#ChatMessageReactionAddResponse">ChatMessageReactionAddResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Labels

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#Label">Label</a>

Methods:

- <code title="get /v1/labels">client.Labels.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#LabelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#Label">Label</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#MessageUpdateResponse">MessageUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#MessageSendResponse">MessageSendResponse</a>

Methods:

- <code title="get /v1/chats/{chatID}/messages/{messageID}">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#MessageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#MessageGetParams">MessageGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared#Message">Message</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/chats/{chatID}/messages/{messageID}">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#MessageService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#MessageUpdateParams">MessageUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#MessageUpdateResponse">MessageUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/chats/{chatID}/messages">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#MessageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#MessageListParams">MessageListParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/packages/pagination#CursorNoLimit">CursorNoLimit</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared#Message">Message</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/chats/{chatID}/messages/{messageID}">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#MessageService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#MessageDeleteParams">MessageDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/messages/search">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#MessageService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#MessageSearchParams">MessageSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/packages/pagination#CursorSearch">CursorSearch</a>[<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared">shared</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6/shared#Message">Message</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/chats/{chatID}/messages">client.Messages.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#MessageService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, chatID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#MessageSendParams">MessageSendParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#MessageSendResponse">MessageSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AssetDownloadResponse">AssetDownloadResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AssetUploadResponse">AssetUploadResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AssetUploadBase64Response">AssetUploadBase64Response</a>

Methods:

- <code title="post /v1/assets/download">client.Assets.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AssetService.Download">Download</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AssetDownloadParams">AssetDownloadParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AssetDownloadResponse">AssetDownloadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/assets/serve">client.Assets.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AssetService.Serve">Serve</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AssetServeParams">AssetServeParams</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/assets/upload">client.Assets.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AssetService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AssetUploadParams">AssetUploadParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AssetUploadResponse">AssetUploadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/assets/upload/base64">client.Assets.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AssetService.UploadBase64">UploadBase64</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AssetUploadBase64Params">AssetUploadBase64Params</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AssetUploadBase64Response">AssetUploadBase64Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Info

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#InfoGetResponse">InfoGetResponse</a>

Methods:

- <code title="get /v1/info">client.Info.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#InfoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#InfoGetResponse">InfoGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# App

## Setup

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupGetResponse">AppSetupGetResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupRegisterResponse">AppSetupRegisterResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupResponseResponseUnion">AppSetupResponseResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupStartResponse">AppSetupStartResponse</a>

Methods:

- <code title="get /v1/app/setup">client.App.Setup.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupGetResponse">AppSetupGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/app/setup/email">client.App.Setup.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupService.Email">Email</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupEmailParams">AppSetupEmailParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/app/setup/register">client.App.Setup.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupService.Register">Register</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupRegisterParams">AppSetupRegisterParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupRegisterResponse">AppSetupRegisterResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/app/setup/response">client.App.Setup.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupService.Response">Response</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupResponseParams">AppSetupResponseParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupResponseResponseUnion">AppSetupResponseResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/app/setup/start">client.App.Setup.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupStartResponse">AppSetupStartResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### RecoveryKey

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupRecoveryKeyVerifyResponse">AppSetupRecoveryKeyVerifyResponse</a>

Methods:

- <code title="post /v1/app/setup/verification/recovery-key">client.App.Setup.RecoveryKey.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupRecoveryKeyService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupRecoveryKeyVerifyParams">AppSetupRecoveryKeyVerifyParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupRecoveryKeyVerifyResponse">AppSetupRecoveryKeyVerifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Reset

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupRecoveryKeyResetNewResponse">AppSetupRecoveryKeyResetNewResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupRecoveryKeyResetConfirmResponse">AppSetupRecoveryKeyResetConfirmResponse</a>

Methods:

- <code title="post /v1/app/setup/verification/recovery-key/reset">client.App.Setup.RecoveryKey.Reset.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupRecoveryKeyResetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupRecoveryKeyResetNewParams">AppSetupRecoveryKeyResetNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupRecoveryKeyResetNewResponse">AppSetupRecoveryKeyResetNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/app/setup/verification/recovery-key/reset/confirm">client.App.Setup.RecoveryKey.Reset.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupRecoveryKeyResetService.Confirm">Confirm</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupRecoveryKeyResetConfirmParams">AppSetupRecoveryKeyResetConfirmParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupRecoveryKeyResetConfirmResponse">AppSetupRecoveryKeyResetConfirmResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Verifications

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationNewResponse">AppSetupVerificationNewResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationGetResponse">AppSetupVerificationGetResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationListResponse">AppSetupVerificationListResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationAcceptResponse">AppSetupVerificationAcceptResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationCancelResponse">AppSetupVerificationCancelResponse</a>

Methods:

- <code title="post /v1/app/setup/verifications">client.App.Setup.Verifications.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationNewParams">AppSetupVerificationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationNewResponse">AppSetupVerificationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/app/setup/verifications/{verificationID}">client.App.Setup.Verifications.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, verificationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationGetResponse">AppSetupVerificationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/app/setup/verifications">client.App.Setup.Verifications.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationListResponse">AppSetupVerificationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/app/setup/verifications/{verificationID}/accept">client.App.Setup.Verifications.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationService.Accept">Accept</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, verificationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationAcceptResponse">AppSetupVerificationAcceptResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/app/setup/verifications/{verificationID}/cancel">client.App.Setup.Verifications.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, verificationID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationCancelParams">AppSetupVerificationCancelParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationCancelResponse">AppSetupVerificationCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### QR

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationQRConfirmScannedResponse">AppSetupVerificationQRConfirmScannedResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationQRScanResponse">AppSetupVerificationQRScanResponse</a>

Methods:

- <code title="post /v1/app/setup/verifications/{verificationID}/qr/confirm-scanned">client.App.Setup.Verifications.QR.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationQRService.ConfirmScanned">ConfirmScanned</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, verificationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationQRConfirmScannedResponse">AppSetupVerificationQRConfirmScannedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/app/setup/verifications/qr/scan">client.App.Setup.Verifications.QR.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationQRService.Scan">Scan</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationQRScanParams">AppSetupVerificationQRScanParams</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationQRScanResponse">AppSetupVerificationQRScanResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### SAS

Response Types:

- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationSASConfirmResponse">AppSetupVerificationSASConfirmResponse</a>
- <a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationSASStartResponse">AppSetupVerificationSASStartResponse</a>

Methods:

- <code title="post /v1/app/setup/verifications/{verificationID}/sas/confirm">client.App.Setup.Verifications.SAS.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationSASService.Confirm">Confirm</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, verificationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationSASConfirmResponse">AppSetupVerificationSASConfirmResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/app/setup/verifications/{verificationID}/sas/start">client.App.Setup.Verifications.SAS.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationSASService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, verificationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6">beeperdesktopapi</a>.<a href="https://pkg.go.dev/github.com/beeper/desktop-api-go/v6#AppSetupVerificationSASStartResponse">AppSetupVerificationSASStartResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
