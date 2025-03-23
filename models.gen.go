// Package hoarder provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package hoarder

import (
	"encoding/json"

	"github.com/oapi-codegen/runtime"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for BookmarkAssetsAssetType.
const (
	BookmarkAssetsAssetTypeAssetScreenshot   BookmarkAssetsAssetType = "assetScreenshot"
	BookmarkAssetsAssetTypeBannerImage       BookmarkAssetsAssetType = "bannerImage"
	BookmarkAssetsAssetTypeBookmarkAsset     BookmarkAssetsAssetType = "bookmarkAsset"
	BookmarkAssetsAssetTypeFullPageArchive   BookmarkAssetsAssetType = "fullPageArchive"
	BookmarkAssetsAssetTypePrecrawledArchive BookmarkAssetsAssetType = "precrawledArchive"
	BookmarkAssetsAssetTypeScreenshot        BookmarkAssetsAssetType = "screenshot"
	BookmarkAssetsAssetTypeUnknown           BookmarkAssetsAssetType = "unknown"
	BookmarkAssetsAssetTypeVideo             BookmarkAssetsAssetType = "video"
)

// Defines values for BookmarkContent0Type.
const (
	BookmarkContent0TypeLink BookmarkContent0Type = "link"
)

// Defines values for BookmarkContent1Type.
const (
	BookmarkContent1TypeText BookmarkContent1Type = "text"
)

// Defines values for BookmarkContent2AssetType.
const (
	BookmarkContent2AssetTypeImage BookmarkContent2AssetType = "image"
	BookmarkContent2AssetTypePdf   BookmarkContent2AssetType = "pdf"
)

// Defines values for BookmarkContent2Type.
const (
	BookmarkContent2TypeAsset BookmarkContent2Type = "asset"
)

// Defines values for BookmarkContent3Type.
const (
	BookmarkContent3TypeUnknown BookmarkContent3Type = "unknown"
)

// Defines values for BookmarkTaggingStatus.
const (
	Failure BookmarkTaggingStatus = "failure"
	Pending BookmarkTaggingStatus = "pending"
	Success BookmarkTaggingStatus = "success"
)

// Defines values for BookmarkTagsAttachedBy.
const (
	Ai    BookmarkTagsAttachedBy = "ai"
	Human BookmarkTagsAttachedBy = "human"
)

// Defines values for HighlightColor.
const (
	HighlightColorBlue   HighlightColor = "blue"
	HighlightColorGreen  HighlightColor = "green"
	HighlightColorRed    HighlightColor = "red"
	HighlightColorYellow HighlightColor = "yellow"
)

// Defines values for ListType.
const (
	ListTypeManual ListType = "manual"
	ListTypeSmart  ListType = "smart"
)

// Defines values for PostBookmarksJSONBody0Type.
const (
	PostBookmarksJSONBody0TypeLink PostBookmarksJSONBody0Type = "link"
)

// Defines values for PostBookmarksJSONBody1Type.
const (
	PostBookmarksJSONBody1TypeText PostBookmarksJSONBody1Type = "text"
)

// Defines values for PostBookmarksJSONBody2AssetType.
const (
	PostBookmarksJSONBody2AssetTypeImage PostBookmarksJSONBody2AssetType = "image"
	PostBookmarksJSONBody2AssetTypePdf   PostBookmarksJSONBody2AssetType = "pdf"
)

// Defines values for PostBookmarksJSONBody2Type.
const (
	PostBookmarksJSONBody2TypeAsset PostBookmarksJSONBody2Type = "asset"
)

// Defines values for PostBookmarksBookmarkIdAssetsJSONBodyAssetType.
const (
	AssetScreenshot   PostBookmarksBookmarkIdAssetsJSONBodyAssetType = "assetScreenshot"
	BannerImage       PostBookmarksBookmarkIdAssetsJSONBodyAssetType = "bannerImage"
	BookmarkAsset     PostBookmarksBookmarkIdAssetsJSONBodyAssetType = "bookmarkAsset"
	FullPageArchive   PostBookmarksBookmarkIdAssetsJSONBodyAssetType = "fullPageArchive"
	PrecrawledArchive PostBookmarksBookmarkIdAssetsJSONBodyAssetType = "precrawledArchive"
	Screenshot        PostBookmarksBookmarkIdAssetsJSONBodyAssetType = "screenshot"
	Unknown           PostBookmarksBookmarkIdAssetsJSONBodyAssetType = "unknown"
	Video             PostBookmarksBookmarkIdAssetsJSONBodyAssetType = "video"
)

// Defines values for PostHighlightsJSONBodyColor.
const (
	PostHighlightsJSONBodyColorBlue   PostHighlightsJSONBodyColor = "blue"
	PostHighlightsJSONBodyColorGreen  PostHighlightsJSONBodyColor = "green"
	PostHighlightsJSONBodyColorRed    PostHighlightsJSONBodyColor = "red"
	PostHighlightsJSONBodyColorYellow PostHighlightsJSONBodyColor = "yellow"
)

// Defines values for PatchHighlightsHighlightIdJSONBodyColor.
const (
	Blue   PatchHighlightsHighlightIdJSONBodyColor = "blue"
	Green  PatchHighlightsHighlightIdJSONBodyColor = "green"
	Red    PatchHighlightsHighlightIdJSONBodyColor = "red"
	Yellow PatchHighlightsHighlightIdJSONBodyColor = "yellow"
)

// Defines values for PostListsJSONBodyType.
const (
	PostListsJSONBodyTypeManual PostListsJSONBodyType = "manual"
	PostListsJSONBodyTypeSmart  PostListsJSONBodyType = "smart"
)

// AssetId defines model for AssetId.
type AssetId = string

// Bookmark defines model for Bookmark.
type Bookmark struct {
	Archived bool `json:"archived"`
	Assets   []struct {
		AssetType BookmarkAssetsAssetType `json:"assetType"`
		Id        string                  `json:"id"`
	} `json:"assets"`
	Content       Bookmark_Content       `json:"content"`
	CreatedAt     string                 `json:"createdAt"`
	Favourited    bool                   `json:"favourited"`
	Id            string                 `json:"id"`
	ModifiedAt    *string                `json:"modifiedAt"`
	Note          *string                `json:"note"`
	Summary       *string                `json:"summary"`
	TaggingStatus *BookmarkTaggingStatus `json:"taggingStatus"`
	Tags          []struct {
		AttachedBy BookmarkTagsAttachedBy `json:"attachedBy"`
		Id         string                 `json:"id"`
		Name       string                 `json:"name"`
	} `json:"tags"`
	Title *string `json:"title"`
}

// BookmarkAssetsAssetType defines model for Bookmark.Assets.AssetType.
type BookmarkAssetsAssetType string

// BookmarkContent0 defines model for .
type BookmarkContent0 struct {
	CrawledAt                *string              `json:"crawledAt"`
	Description              *string              `json:"description"`
	Favicon                  *string              `json:"favicon"`
	FullPageArchiveAssetId   *string              `json:"fullPageArchiveAssetId"`
	HtmlContent              *string              `json:"htmlContent"`
	ImageAssetId             *string              `json:"imageAssetId"`
	ImageUrl                 *string              `json:"imageUrl"`
	PrecrawledArchiveAssetId *string              `json:"precrawledArchiveAssetId"`
	ScreenshotAssetId        *string              `json:"screenshotAssetId"`
	Title                    *string              `json:"title"`
	Type                     BookmarkContent0Type `json:"type"`
	Url                      string               `json:"url"`
	VideoAssetId             *string              `json:"videoAssetId"`
}

// BookmarkContent0Type defines model for Bookmark.Content.0.Type.
type BookmarkContent0Type string

// BookmarkContent1 defines model for .
type BookmarkContent1 struct {
	SourceUrl *string              `json:"sourceUrl"`
	Text      string               `json:"text"`
	Type      BookmarkContent1Type `json:"type"`
}

// BookmarkContent1Type defines model for Bookmark.Content.1.Type.
type BookmarkContent1Type string

// BookmarkContent2 defines model for .
type BookmarkContent2 struct {
	AssetId   string                    `json:"assetId"`
	AssetType BookmarkContent2AssetType `json:"assetType"`
	FileName  *string                   `json:"fileName"`
	Size      *float32                  `json:"size"`
	SourceUrl *string                   `json:"sourceUrl"`
	Type      BookmarkContent2Type      `json:"type"`
}

// BookmarkContent2AssetType defines model for Bookmark.Content.2.AssetType.
type BookmarkContent2AssetType string

// BookmarkContent2Type defines model for Bookmark.Content.2.Type.
type BookmarkContent2Type string

// BookmarkContent3 defines model for .
type BookmarkContent3 struct {
	Type BookmarkContent3Type `json:"type"`
}

// BookmarkContent3Type defines model for Bookmark.Content.3.Type.
type BookmarkContent3Type string

// Bookmark_Content defines model for Bookmark.Content.
type Bookmark_Content struct {
	union json.RawMessage
}

// BookmarkTaggingStatus defines model for Bookmark.TaggingStatus.
type BookmarkTaggingStatus string

// BookmarkTagsAttachedBy defines model for Bookmark.Tags.AttachedBy.
type BookmarkTagsAttachedBy string

// BookmarkId defines model for BookmarkId.
type BookmarkId = string

// Cursor defines model for Cursor.
type Cursor = string

// Highlight defines model for Highlight.
type Highlight struct {
	BookmarkId  string          `json:"bookmarkId"`
	Color       *HighlightColor `json:"color,omitempty"`
	CreatedAt   string          `json:"createdAt"`
	EndOffset   float32         `json:"endOffset"`
	Id          string          `json:"id"`
	Note        *string         `json:"note"`
	StartOffset float32         `json:"startOffset"`
	Text        *string         `json:"text"`
	UserId      string          `json:"userId"`
}

// HighlightColor defines model for Highlight.Color.
type HighlightColor string

// HighlightId defines model for HighlightId.
type HighlightId = string

// List defines model for List.
type List struct {
	Icon     string    `json:"icon"`
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	ParentId *string   `json:"parentId"`
	Query    *string   `json:"query"`
	Type     *ListType `json:"type,omitempty"`
}

// ListType defines model for List.Type.
type ListType string

// ListId defines model for ListId.
type ListId = string

// PaginatedBookmarks defines model for PaginatedBookmarks.
type PaginatedBookmarks struct {
	Bookmarks  []Bookmark `json:"bookmarks"`
	NextCursor *string    `json:"nextCursor"`
}

// PaginatedHighlights defines model for PaginatedHighlights.
type PaginatedHighlights struct {
	Highlights []Highlight `json:"highlights"`
	NextCursor *string     `json:"nextCursor"`
}

// Tag defines model for Tag.
type Tag struct {
	Id                         string  `json:"id"`
	Name                       string  `json:"name"`
	NumBookmarks               float32 `json:"numBookmarks"`
	NumBookmarksByAttachedType struct {
		Ai    *float32 `json:"ai,omitempty"`
		Human *float32 `json:"human,omitempty"`
	} `json:"numBookmarksByAttachedType"`
}

// TagId defines model for TagId.
type TagId = string

// GetBookmarksParams defines parameters for GetBookmarks.
type GetBookmarksParams struct {
	Archived   *bool    `form:"archived,omitempty" json:"archived,omitempty"`
	Favourited *bool    `form:"favourited,omitempty" json:"favourited,omitempty"`
	Limit      *float32 `form:"limit,omitempty" json:"limit,omitempty"`
	Cursor     *Cursor  `form:"cursor,omitempty" json:"cursor,omitempty"`
}

// PostBookmarksJSONBody defines parameters for PostBookmarks.
type PostBookmarksJSONBody struct {
	Archived   *bool   `json:"archived,omitempty"`
	CreatedAt  *string `json:"createdAt"`
	Favourited *bool   `json:"favourited,omitempty"`
	Note       *string `json:"note,omitempty"`
	Summary    *string `json:"summary,omitempty"`
	Title      *string `json:"title"`
	union      json.RawMessage
}

// PostBookmarksJSONBody0 defines parameters for PostBookmarks.
type PostBookmarksJSONBody0 struct {
	PrecrawledArchiveId *string                    `json:"precrawledArchiveId,omitempty"`
	Type                PostBookmarksJSONBody0Type `json:"type"`
	Url                 string                     `json:"url"`
}

// PostBookmarksJSONBody0Type defines parameters for PostBookmarks.
type PostBookmarksJSONBody0Type string

// PostBookmarksJSONBody1 defines parameters for PostBookmarks.
type PostBookmarksJSONBody1 struct {
	SourceUrl *string                    `json:"sourceUrl,omitempty"`
	Text      string                     `json:"text"`
	Type      PostBookmarksJSONBody1Type `json:"type"`
}

// PostBookmarksJSONBody1Type defines parameters for PostBookmarks.
type PostBookmarksJSONBody1Type string

// PostBookmarksJSONBody2 defines parameters for PostBookmarks.
type PostBookmarksJSONBody2 struct {
	AssetId   string                          `json:"assetId"`
	AssetType PostBookmarksJSONBody2AssetType `json:"assetType"`
	FileName  *string                         `json:"fileName,omitempty"`
	SourceUrl *string                         `json:"sourceUrl,omitempty"`
	Type      PostBookmarksJSONBody2Type      `json:"type"`
}

// PostBookmarksJSONBody2AssetType defines parameters for PostBookmarks.
type PostBookmarksJSONBody2AssetType string

// PostBookmarksJSONBody2Type defines parameters for PostBookmarks.
type PostBookmarksJSONBody2Type string

// GetBookmarksSearchParams defines parameters for GetBookmarksSearch.
type GetBookmarksSearchParams struct {
	Q      string   `form:"q" json:"q"`
	Limit  *float32 `form:"limit,omitempty" json:"limit,omitempty"`
	Cursor *Cursor  `form:"cursor,omitempty" json:"cursor,omitempty"`
}

// PatchBookmarksBookmarkIdJSONBody defines parameters for PatchBookmarksBookmarkId.
type PatchBookmarksBookmarkIdJSONBody struct {
	Archived   *bool   `json:"archived,omitempty"`
	CreatedAt  *string `json:"createdAt"`
	Favourited *bool   `json:"favourited,omitempty"`
	Note       *string `json:"note,omitempty"`
	Summary    *string `json:"summary"`
	Title      *string `json:"title"`
}

// PostBookmarksBookmarkIdAssetsJSONBody defines parameters for PostBookmarksBookmarkIdAssets.
type PostBookmarksBookmarkIdAssetsJSONBody struct {
	AssetType PostBookmarksBookmarkIdAssetsJSONBodyAssetType `json:"assetType"`
	Id        string                                         `json:"id"`
}

// PostBookmarksBookmarkIdAssetsJSONBodyAssetType defines parameters for PostBookmarksBookmarkIdAssets.
type PostBookmarksBookmarkIdAssetsJSONBodyAssetType string

// PutBookmarksBookmarkIdAssetsAssetIdJSONBody defines parameters for PutBookmarksBookmarkIdAssetsAssetId.
type PutBookmarksBookmarkIdAssetsAssetIdJSONBody struct {
	AssetId string `json:"assetId"`
}

// DeleteBookmarksBookmarkIdTagsJSONBody defines parameters for DeleteBookmarksBookmarkIdTags.
type DeleteBookmarksBookmarkIdTagsJSONBody struct {
	Tags []struct {
		TagId   *string `json:"tagId,omitempty"`
		TagName *string `json:"tagName,omitempty"`
	} `json:"tags"`
}

// PostBookmarksBookmarkIdTagsJSONBody defines parameters for PostBookmarksBookmarkIdTags.
type PostBookmarksBookmarkIdTagsJSONBody struct {
	Tags []struct {
		TagId   *string `json:"tagId,omitempty"`
		TagName *string `json:"tagName,omitempty"`
	} `json:"tags"`
}

// GetHighlightsParams defines parameters for GetHighlights.
type GetHighlightsParams struct {
	Limit  *float32 `form:"limit,omitempty" json:"limit,omitempty"`
	Cursor *Cursor  `form:"cursor,omitempty" json:"cursor,omitempty"`
}

// PostHighlightsJSONBody defines parameters for PostHighlights.
type PostHighlightsJSONBody struct {
	BookmarkId  string                       `json:"bookmarkId"`
	Color       *PostHighlightsJSONBodyColor `json:"color,omitempty"`
	EndOffset   float32                      `json:"endOffset"`
	Note        *string                      `json:"note"`
	StartOffset float32                      `json:"startOffset"`
	Text        *string                      `json:"text"`
}

// PostHighlightsJSONBodyColor defines parameters for PostHighlights.
type PostHighlightsJSONBodyColor string

// PatchHighlightsHighlightIdJSONBody defines parameters for PatchHighlightsHighlightId.
type PatchHighlightsHighlightIdJSONBody struct {
	Color *PatchHighlightsHighlightIdJSONBodyColor `json:"color,omitempty"`
}

// PatchHighlightsHighlightIdJSONBodyColor defines parameters for PatchHighlightsHighlightId.
type PatchHighlightsHighlightIdJSONBodyColor string

// PostListsJSONBody defines parameters for PostLists.
type PostListsJSONBody struct {
	Icon     string                 `json:"icon"`
	Name     string                 `json:"name"`
	ParentId *string                `json:"parentId"`
	Query    *string                `json:"query,omitempty"`
	Type     *PostListsJSONBodyType `json:"type,omitempty"`
}

// PostListsJSONBodyType defines parameters for PostLists.
type PostListsJSONBodyType string

// PatchListsListIdJSONBody defines parameters for PatchListsListId.
type PatchListsListIdJSONBody struct {
	Icon     *string `json:"icon,omitempty"`
	Name     *string `json:"name,omitempty"`
	ParentId *string `json:"parentId"`
	Query    *string `json:"query,omitempty"`
}

// GetListsListIdBookmarksParams defines parameters for GetListsListIdBookmarks.
type GetListsListIdBookmarksParams struct {
	Limit  *float32 `form:"limit,omitempty" json:"limit,omitempty"`
	Cursor *Cursor  `form:"cursor,omitempty" json:"cursor,omitempty"`
}

// PatchTagsTagIdJSONBody defines parameters for PatchTagsTagId.
type PatchTagsTagIdJSONBody struct {
	Name *string `json:"name,omitempty"`
}

// GetTagsTagIdBookmarksParams defines parameters for GetTagsTagIdBookmarks.
type GetTagsTagIdBookmarksParams struct {
	Limit  *float32 `form:"limit,omitempty" json:"limit,omitempty"`
	Cursor *Cursor  `form:"cursor,omitempty" json:"cursor,omitempty"`
}

// PostBookmarksJSONRequestBody defines body for PostBookmarks for application/json ContentType.
type PostBookmarksJSONRequestBody PostBookmarksJSONBody

// PatchBookmarksBookmarkIdJSONRequestBody defines body for PatchBookmarksBookmarkId for application/json ContentType.
type PatchBookmarksBookmarkIdJSONRequestBody PatchBookmarksBookmarkIdJSONBody

// PostBookmarksBookmarkIdAssetsJSONRequestBody defines body for PostBookmarksBookmarkIdAssets for application/json ContentType.
type PostBookmarksBookmarkIdAssetsJSONRequestBody PostBookmarksBookmarkIdAssetsJSONBody

// PutBookmarksBookmarkIdAssetsAssetIdJSONRequestBody defines body for PutBookmarksBookmarkIdAssetsAssetId for application/json ContentType.
type PutBookmarksBookmarkIdAssetsAssetIdJSONRequestBody PutBookmarksBookmarkIdAssetsAssetIdJSONBody

// DeleteBookmarksBookmarkIdTagsJSONRequestBody defines body for DeleteBookmarksBookmarkIdTags for application/json ContentType.
type DeleteBookmarksBookmarkIdTagsJSONRequestBody DeleteBookmarksBookmarkIdTagsJSONBody

// PostBookmarksBookmarkIdTagsJSONRequestBody defines body for PostBookmarksBookmarkIdTags for application/json ContentType.
type PostBookmarksBookmarkIdTagsJSONRequestBody PostBookmarksBookmarkIdTagsJSONBody

// PostHighlightsJSONRequestBody defines body for PostHighlights for application/json ContentType.
type PostHighlightsJSONRequestBody PostHighlightsJSONBody

// PatchHighlightsHighlightIdJSONRequestBody defines body for PatchHighlightsHighlightId for application/json ContentType.
type PatchHighlightsHighlightIdJSONRequestBody PatchHighlightsHighlightIdJSONBody

// PostListsJSONRequestBody defines body for PostLists for application/json ContentType.
type PostListsJSONRequestBody PostListsJSONBody

// PatchListsListIdJSONRequestBody defines body for PatchListsListId for application/json ContentType.
type PatchListsListIdJSONRequestBody PatchListsListIdJSONBody

// PatchTagsTagIdJSONRequestBody defines body for PatchTagsTagId for application/json ContentType.
type PatchTagsTagIdJSONRequestBody PatchTagsTagIdJSONBody

// AsBookmarkContent0 returns the union data inside the Bookmark_Content as a BookmarkContent0
func (t Bookmark_Content) AsBookmarkContent0() (BookmarkContent0, error) {
	var body BookmarkContent0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBookmarkContent0 overwrites any union data inside the Bookmark_Content as the provided BookmarkContent0
func (t *Bookmark_Content) FromBookmarkContent0(v BookmarkContent0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBookmarkContent0 performs a merge with any union data inside the Bookmark_Content, using the provided BookmarkContent0
func (t *Bookmark_Content) MergeBookmarkContent0(v BookmarkContent0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBookmarkContent1 returns the union data inside the Bookmark_Content as a BookmarkContent1
func (t Bookmark_Content) AsBookmarkContent1() (BookmarkContent1, error) {
	var body BookmarkContent1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBookmarkContent1 overwrites any union data inside the Bookmark_Content as the provided BookmarkContent1
func (t *Bookmark_Content) FromBookmarkContent1(v BookmarkContent1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBookmarkContent1 performs a merge with any union data inside the Bookmark_Content, using the provided BookmarkContent1
func (t *Bookmark_Content) MergeBookmarkContent1(v BookmarkContent1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBookmarkContent2 returns the union data inside the Bookmark_Content as a BookmarkContent2
func (t Bookmark_Content) AsBookmarkContent2() (BookmarkContent2, error) {
	var body BookmarkContent2
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBookmarkContent2 overwrites any union data inside the Bookmark_Content as the provided BookmarkContent2
func (t *Bookmark_Content) FromBookmarkContent2(v BookmarkContent2) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBookmarkContent2 performs a merge with any union data inside the Bookmark_Content, using the provided BookmarkContent2
func (t *Bookmark_Content) MergeBookmarkContent2(v BookmarkContent2) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBookmarkContent3 returns the union data inside the Bookmark_Content as a BookmarkContent3
func (t Bookmark_Content) AsBookmarkContent3() (BookmarkContent3, error) {
	var body BookmarkContent3
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBookmarkContent3 overwrites any union data inside the Bookmark_Content as the provided BookmarkContent3
func (t *Bookmark_Content) FromBookmarkContent3(v BookmarkContent3) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBookmarkContent3 performs a merge with any union data inside the Bookmark_Content, using the provided BookmarkContent3
func (t *Bookmark_Content) MergeBookmarkContent3(v BookmarkContent3) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t Bookmark_Content) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Bookmark_Content) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}
