package mediaDto

type FileDataType struct {
	FileMetadata FileMetadata `json:"file_metadata"`
	Data         []byte       `json:"data"`
}
