package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

// MuteService provides methods for muting specific user.
type MuteService struct {
	sling *sling.Sling
}

// newMuteService returns a new MuteService.
func newMuteService(sling *sling.Sling) *MuteService {
	return &MuteService{
		sling: sling.Path("mutes/users/"),
	}
}

// MuteCreateParams are the parameters for MuteService.Create.
type MuteCreateParams struct {
	ScreenName      string `url:"screen_name,omitempty,comma"`
	UserID          int64  `url:"user_id,omitempty,comma"`
	IncludeEntities *bool  `url:"include_entities,omitempty"` // whether 'status' should include entities
	SkipStatus      *bool  `url:"skip_status,omitempty"`
}

// Create a mute for specific user, return the user muteed as Entity.
// https://developer.twitter.com/en/docs/accounts-and-users/mute-mute-report-users/api-reference/post-mutes-create
func (s *MuteService) Create(params *MuteCreateParams) (User, *http.Response, error) {
	users := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("create.json").QueryStruct(params).Receive(users, apiError)
	return *users, resp, relevantError(err, *apiError)
}

// MuteDestroyParams are the parameters for MuteService.Destroy.
type MuteDestroyParams struct {
	ScreenName      string `url:"screen_name,omitempty,comma"`
	UserID          int64  `url:"user_id,omitempty,comma"`
	IncludeEntities *bool  `url:"include_entities,omitempty"` // whether 'status' should include entities
	SkipStatus      *bool  `url:"skip_status,omitempty"`
}

// Destroy the mute for specific user, return the user unmuteed as Entity.
// https://developer.twitter.com/en/docs/accounts-and-users/mute-mute-report-users/api-reference/post-mutes-destroy
func (s *MuteService) Destroy(params *MuteDestroyParams) (User, *http.Response, error) {
	users := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("destroy.json").QueryStruct(params).Receive(users, apiError)
	return *users, resp, relevantError(err, *apiError)
}

// MuteIDParams are the parameters for MuteService.Ids
type MuteIDParams struct {
	Cursor int64 `url:"cursor,omitempty"`
	Count  int   `url:"count,omitempty"`
}

// IDs returns a cursored collection of user ids that the specified user is following.
// https://dev.twitter.com/rest/reference/get/mutes/ids
func (s *MuteService) IDs(params *MuteIDParams) (*CursoredIDs, *http.Response, error) {
	ids := new(CursoredIDs)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("ids.json").QueryStruct(params).Receive(ids, apiError)
	return ids, resp, relevantError(err, *apiError)
}

// MuteListParams are the parameters for MuteService.List
type MuteListParams struct {
	Cursor              int64 `url:"cursor,omitempty"`
	Count               int   `url:"count,omitempty"`
	SkipStatus          *bool `url:"skip_status,omitempty"`
	IncludeUserEntities *bool `url:"include_user_entities,omitempty"`
}

// List returns a cursored collection of Users that the specified user is following.
// https://dev.twitter.com/rest/reference/get/mutes/list
func (s *MuteService) List(params *MuteListParams) (*CursoredUsers, *http.Response, error) {
	friends := new(CursoredUsers)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("list.json").QueryStruct(params).Receive(friends, apiError)
	return friends, resp, relevantError(err, *apiError)
}
