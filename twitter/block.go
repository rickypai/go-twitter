package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

// BlockService provides methods for blocking specific user.
type BlockService struct {
	sling *sling.Sling
}

// newBlockService returns a new BlockService.
func newBlockService(sling *sling.Sling) *BlockService {
	return &BlockService{
		sling: sling.Path("blocks/"),
	}
}

// BlockCreateParams are the parameters for BlockService.Create.
type BlockCreateParams struct {
	ScreenName      string `url:"screen_name,omitempty,comma"`
	UserID          int64  `url:"user_id,omitempty,comma"`
	IncludeEntities *bool  `url:"include_entities,omitempty"` // whether 'status' should include entities
	SkipStatus      *bool  `url:"skip_status,omitempty"`
}

// Create a block for specific user, return the user blocked as Entity.
// https://developer.twitter.com/en/docs/accounts-and-users/mute-block-report-users/api-reference/post-blocks-create
func (s *BlockService) Create(params *BlockCreateParams) (User, *http.Response, error) {
	users := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("create.json").QueryStruct(params).Receive(users, apiError)
	return *users, resp, relevantError(err, *apiError)
}

// BlockDestroyParams are the parameters for BlockService.Destroy.
type BlockDestroyParams struct {
	ScreenName      string `url:"screen_name,omitempty,comma"`
	UserID          int64  `url:"user_id,omitempty,comma"`
	IncludeEntities *bool  `url:"include_entities,omitempty"` // whether 'status' should include entities
	SkipStatus      *bool  `url:"skip_status,omitempty"`
}

// Destroy the block for specific user, return the user unblocked as Entity.
// https://developer.twitter.com/en/docs/accounts-and-users/mute-block-report-users/api-reference/post-blocks-destroy
func (s *BlockService) Destroy(params *BlockDestroyParams) (User, *http.Response, error) {
	users := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("destroy.json").QueryStruct(params).Receive(users, apiError)
	return *users, resp, relevantError(err, *apiError)
}

// BlockIDParams are the parameters for BlockService.Ids
type BlockIDParams struct {
	Cursor int64 `url:"cursor,omitempty"`
	Count  int   `url:"count,omitempty"`
}

// IDs returns a cursored collection of user ids that the specified user is following.
// https://dev.twitter.com/rest/reference/get/blocks/ids
func (s *BlockService) IDs(params *BlockIDParams) (*CursoredIDs, *http.Response, error) {
	ids := new(CursoredIDs)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("ids.json").QueryStruct(params).Receive(ids, apiError)
	return ids, resp, relevantError(err, *apiError)
}

// BlockListParams are the parameters for BlockService.List
type BlockListParams struct {
	Cursor              int64 `url:"cursor,omitempty"`
	Count               int   `url:"count,omitempty"`
	SkipStatus          *bool `url:"skip_status,omitempty"`
	IncludeUserEntities *bool `url:"include_user_entities,omitempty"`
}

// List returns a cursored collection of Users that the specified user is following.
// https://dev.twitter.com/rest/reference/get/blocks/list
func (s *BlockService) List(params *BlockListParams) (*CursoredUsers, *http.Response, error) {
	friends := new(CursoredUsers)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("list.json").QueryStruct(params).Receive(friends, apiError)
	return friends, resp, relevantError(err, *apiError)
}
