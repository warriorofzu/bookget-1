package familysearch

type Downloader struct {
	Index    int
	Url      string
	Domain   string
	SavePath string
	BookId   string
	UrlType  uint8 //URL类型识别，同一个站有多种URL
}

//{"type":"image-data","args":{"imageURL":"https://www.familysearch.org/ark:/61903/3:1:3QSQ-G97V-BRNQ?cc=1787988&wc=3XLV-ZVP:1022211401,1021934502,1021934702,1021936102,1022506101","state":{"wc":"3XPX-Y4M:1022211401,1021934502,1021934702,1021936102,1022505901","cc":"1787988","imageOrFilmUrl":"/ark:/61903/3:1:3QS7-L97V-7XZ3","collectionContext":"1787988","viewMode":"i","selectedImageIndex":-1,"waypointContext":"/service/cds/recapi/waypoints/3XPX-Y4M:1022211401,1021934502,1021934702,1021936102,1022505901"},"locale":"zh"}}
type ImageData struct {
	Type string `json:"type"`
	Args struct {
		ImageURL string `json:"imageURL"`
		State    struct {
			Wc                 string `json:"wc"`
			Cc                 string `json:"cc"`
			ImageOrFilmUrl     string `json:"imageOrFilmUrl"`
			CollectionContext  string `json:"collectionContext"`
			ViewMode           string `json:"viewMode"`
			SelectedImageIndex int    `json:"selectedImageIndex"`
			WaypointContext    string `json:"waypointContext"`
		} `json:"state"`
		Locale string `json:"locale"`
	} `json:"args"`
}

//{"type":"film-data","args":{"dgsNum":"005547019","state":{"imageOrFilmUrl":"","viewMode":"i","selectedImageIndex":-1},"locale":"zh"}}
type FilmData struct {
	Type string `json:"type"`
	Args struct {
		WaypointURL string `json:"waypointURL"`
		DgsNum      string `json:"dgsNum"`
		State       struct {
			ImageOrFilmUrl     string `json:"imageOrFilmUrl"`
			ViewMode           string `json:"viewMode"`
			SelectedImageIndex int    `json:"selectedImageIndex"`
		} `json:"state"`
		Locale string `json:"locale"`
	} `json:"args"`
}

type ResultImageData struct {
	ImageURL    string        `json:"imageURL"`
	ArkId       string        `json:"arkId"`
	DgsNum      string        `json:"dgsNum"`
	Collections []interface{} `json:"collections"`
}

type ResultFilmData struct {
	DgsNum         string      `json:"dgsNum"`
	Images         []string    `json:"images"`
	Type           string      `json:"type"`
	WaypointCrumbs interface{} `json:"waypointCrumbs"`
	WaypointURL    interface{} `json:"waypointURL"`
	Templates      struct {
		DasTemplate string `json:"dasTemplate"`
		DzTemplate  string `json:"dzTemplate"`
	} `json:"templates"`
}

type ResultError struct {
	Error struct {
		Message     string   `json:"message"`
		FailedRoles []string `json:"failedRoles"`
		StatusCode  int      `json:"statusCode"`
	} `json:"error"`
}

//家谱图像 https://www.familysearch.org/records/images/api/imageDetails/groups/M94F-78D?properties&changeLog&coverageIndex=null
type Canvases struct {
	ImageUrls []string
	IiifUrls  []string
	Size      int
}
type ImageGroups struct {
	Groups []struct {
		Id               string   `json:"id"`
		Created          int64    `json:"created"`
		Modified         int64    `json:"modified"`
		ExternalId       string   `json:"externalId"`
		PartitionKey     string   `json:"partitionKey"`
		ParentIds        []string `json:"parentIds"`
		GroupName        string   `json:"groupName"`
		VolumeSetTitle   string   `json:"volumeSetTitle"`
		ChildCount       int      `json:"childCount"`
		ImageApids       []string `json:"imageApids"`
		ImageUrls        []string `json:"imageUrls"`
		ParentGroupId    string   `json:"parentGroupId"`
		ParentImageCount int      `json:"parentImageCount"`
	} `json:"groups"`
	VolumeSet struct {
		Id                       string `json:"id"`
		Created                  int64  `json:"created"`
		Modified                 int64  `json:"modified"`
		GroupName                string `json:"groupName"`
		Title                    string `json:"title"`
		VolumeSetTitle           string `json:"volumeSetTitle"`
		VolumeSetFirstAncestor   string `json:"volumeSetFirstAncestor"`
		VolumeSetMigrantAncestor string `json:"volumeSetMigrantAncestor"`
		VolumeSetTotalVolumes    string `json:"volumeSetTotalVolumes"`
		VolumeSetLostVolumes     string `json:"volumeSetLostVolumes"`
		ChildCount               int    `json:"childCount"`
		IndexedChildCount        int    `json:"indexedChildCount"`
		ModifiedDateTime         string `json:"modifiedDateTime"`
		CreatedDateTime          string `json:"createdDateTime"`
	} `json:"volumeSet"`
}
