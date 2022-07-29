package gohttp

import "sync"

var (
	mockupServer = mockServer{
		mocks: make(map[string]*Mock),
	}
)

type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex

	mocks map[string]*Mock
}

func StartMockServer() {
	//this makes sure that the server is concurrent safe.
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enabled = true
}

func StopMockServer() {
	//this makes sure that the server is concurrent safe.
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enabled = false
}

func AddMock(mock Mock) {
	//this makes sure that the server is concurrent safe.
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	key := mockupServer.getMockKey(mock.Method, mock.Url, mock.RequestBody) + mock.Url + mock.RequestBody
	mockupServer.mocks[key] = &mock
}

func (m *mockServer) getMockKey(method, url, body string) string {
	return method + url + body
}

func (m *mockServer) getMock(method, url, body string) *Mock {
	return m.mocks[m.getMockKey(method, url, body)]
}
