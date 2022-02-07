package lib

import "encoding/json"

const CurrentMetadataVersion int = 0

var MetadataID = []byte("_metadata")

// AppMetadata stores metadata about the
type AppMetadata struct {
	Version int
}

func NewMetadata() *AppMetadata {
	return &AppMetadata{
		Version: CurrentMetadataVersion,
	}
}

func (m *AppMetadata) MarshalBytes() (id []byte, body []byte, err error) {
	body, err = json.Marshal(m)
	if err != nil {
		return nil, nil, err
	}
	return MetadataID, body, nil
}
