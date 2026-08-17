package mediaDto

type UploadFile struct {
	FileName    string
	ContentType string
	Size        int64
	Data        []byte
}
