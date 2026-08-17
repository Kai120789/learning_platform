package grpc

type Client struct {
	UserClient     *UserClient
	AuthClient     *AuthClient
	GroupClient    *GroupClient
	LessonClient   *LessonClient
	ScheduleClient *ScheduleClient
	SubjectClient  *SubjectClient
	MediaClient    *MediaClient
	MaterialClient *MaterialClient
}

func NewClient(
	userGrpcUrl string,
	authGrpcUrl string,
	groupGrpcUrl string,
	lessonGrpcUrl string,
	scheduleGrpcUrl string,
	subjectGrpcUrl string,
	mediaGrpcUrl string,
	materialGrpcUrl string,
) (*Client, error) {
	userGrpcConnection, err := NewUserGrpcConnection(userGrpcUrl)
	if err != nil {
		return nil, err
	}

	authGrpcConnection, err := NewAuthGrpcConnection(authGrpcUrl)
	if err != nil {
		return nil, err
	}

	groupGrpcConnection, err := NewGroupGrpcConnection(groupGrpcUrl)
	if err != nil {
		return nil, err
	}

	lessonGrpcConnection, err := NewLessonGrpcConnection(lessonGrpcUrl)
	if err != nil {
		return nil, err
	}

	scheduleGrpcConnection, err := NewScheduleGrpcConnection(scheduleGrpcUrl)
	if err != nil {
		return nil, err
	}

	subjectGrpcConnection, err := NewSubjectGrpcConnection(subjectGrpcUrl)
	if err != nil {
		return nil, err
	}

	mediaGrpcConnection, err := NewMediaGrpcConnection(mediaGrpcUrl)
	if err != nil {
		return nil, err
	}

	materialGrpcConnection, err := NewMaterialGrpcConnection(materialGrpcUrl)
	if err != nil {
		return nil, err
	}

	return &Client{
		UserClient:     NewUserClient(userGrpcConnection),
		AuthClient:     NewAuthClient(authGrpcConnection),
		GroupClient:    NewGroupClient(groupGrpcConnection),
		LessonClient:   NewLessonClient(lessonGrpcConnection),
		ScheduleClient: NewScheduleClient(scheduleGrpcConnection),
		SubjectClient:  NewSubjectClient(subjectGrpcConnection),
		MediaClient:    NewMediaClient(mediaGrpcConnection),
		MaterialClient: NewMaterialClient(materialGrpcConnection),
	}, nil
}
