package pkg

type MockSpotifyAPI struct{}

func (m *MockSpotifyAPI) Play() error {
	return nil
}

func (m *MockSpotifyAPI) Pause() error {
	return nil
}
