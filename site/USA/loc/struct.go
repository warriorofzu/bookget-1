package loc

import "time"

type ManifestsJson struct {
	ArticlesAndEssays interface{} `json:"articles_and_essays"`
	CiteThis          struct {
		Apa     string `json:"apa"`
		Chicago string `json:"chicago"`
		Mla     string `json:"mla"`
	} `json:"cite_this"`
	Item struct {
		Version          int64         `json:"_version_"`
		AccessRestricted bool          `json:"access_restricted"`
		Aka              []string      `json:"aka"`
		Campaigns        []interface{} `json:"campaigns"`
		ContributorNames []string      `json:"contributor_names"`
		Contributors     []struct {
			BuWeiLou                                   string `json:"bu wei lou,omitempty"`
			ChineseRareBookCollectionLibraryOfCongress string `json:"chinese rare book collection (library of congress),omitempty"`
			WeiJian                                    string `json:"wei, jian,omitempty"`
		} `json:"contributors"`
		ControlNumber    string   `json:"control_number"`
		CreatedPublished []string `json:"created_published"`
		Date             string   `json:"date"`
		Dates            []struct {
			Field1 string `json:"1712"`
		} `json:"dates"`
		Description      []string  `json:"description"`
		Digitized        bool      `json:"digitized"`
		DisplayOffsite   bool      `json:"display_offsite"`
		ExtractTimestamp time.Time `json:"extract_timestamp"`
		ExtractUrls      []string  `json:"extract_urls"`
		Format           []struct {
			Book string `json:"book"`
		} `json:"format"`
		Group       []string `json:"group"`
		Hassegments bool     `json:"hassegments"`
		Id          string   `json:"id"`
		ImageUrl    []string `json:"image_url"`
		Index       int      `json:"index"`
		Item        struct {
			Contributors     []string `json:"contributors"`
			CreatedPublished []string `json:"created_published"`
			Date             string   `json:"date"`
			Format           []string `json:"format"`
			Language         []string `json:"language"`
			Medium           []string `json:"medium"`
			Notes            []string `json:"notes"`
			OtherTitle       []string `json:"other_title"`
			Subjects         []string `json:"subjects"`
			Title            string   `json:"title"`
			TranslatedTitle  []string `json:"translated_title"`
		} `json:"item"`
		Language  []string `json:"language"`
		Languages []struct {
			Chinese string `json:"chinese"`
		} `json:"languages"`
		LibraryOfCongressControlNumber string   `json:"library_of_congress_control_number"`
		Medium                         []string `json:"medium"`
		MimeType                       []string `json:"mime_type"`
		Notes                          []string `json:"notes"`
		Number                         []string `json:"number"`
		NumberCarrierType              []string `json:"number_carrier_type"`
		NumberLccn                     []string `json:"number_lccn"`
		NumberOclc                     []string `json:"number_oclc"`
		NumberSourceModified           []string `json:"number_source_modified"`
		OnlineFormat                   []string `json:"online_format"`
		OriginalFormat                 []string `json:"original_format"`
		OtherFormats                   []struct {
			Label string `json:"label"`
			Link  string `json:"link"`
		} `json:"other_formats"`
		OtherTitle []string `json:"other_title"`
		Partof     []struct {
			Count int    `json:"count"`
			Title string `json:"title"`
			Url   string `json:"url"`
		} `json:"partof"`
		Resources []struct {
			Caption string `json:"caption"`
			Files   int    `json:"files"`
			Image   string `json:"image"`
			Url     string `json:"url"`
		} `json:"resources"`
		Rights          []string `json:"rights"`
		Score           float64  `json:"score"`
		ShelfId         string   `json:"shelf_id"`
		Site            []string `json:"site"`
		Subject         []string `json:"subject"`
		SubjectHeadings []string `json:"subject_headings"`
		Subjects        []struct {
			TianWen          string `json:"tian wen,omitempty"`
			TianWenSuanFaLei string `json:"tian wen suan fa lei,omitempty"`
			ZiBu             string `json:"zi bu,omitempty"`
		} `json:"subjects"`
		Timestamp       time.Time `json:"timestamp"`
		Title           string    `json:"title"`
		TranslatedTitle []string  `json:"translated_title"`
		Type            []string  `json:"type"`
		Url             string    `json:"url"`
	} `json:"item"`
	MoreLikeThis []struct {
		AccessRestricted bool          `json:"access_restricted"`
		Aka              []string      `json:"aka"`
		Campaigns        []interface{} `json:"campaigns"`
		Contributor      []string      `json:"contributor"`
		Date             string        `json:"date,omitempty"`
		Dates            []time.Time   `json:"dates,omitempty"`
		Description      []string      `json:"description"`
		Digitized        bool          `json:"digitized"`
		ExtractTimestamp time.Time     `json:"extract_timestamp"`
		Group            []string      `json:"group"`
		Hassegments      bool          `json:"hassegments"`
		Id               string        `json:"id"`
		ImageUrl         []string      `json:"image_url"`
		Index            int           `json:"index"`
		Language         []string      `json:"language"`
		MimeType         []string      `json:"mime_type"`
		Number           []string      `json:"number"`
		OnlineFormat     []string      `json:"online_format"`
		OriginalFormat   []string      `json:"original_format"`
		OtherTitle       []string      `json:"other_title"`
		Partof           []string      `json:"partof"`
		Score            float64       `json:"score"`
		ShelfId          string        `json:"shelf_id"`
		Site             []string      `json:"site"`
		Subject          []string      `json:"subject"`
		Timestamp        time.Time     `json:"timestamp"`
		Title            string        `json:"title"`
		Url              string        `json:"url"`
	} `json:"more_like_this"`
	Options struct {
		AccessGroup        []string      `json:"access_group"`
		AccessGroupRaw     string        `json:"access_group_raw"`
		All                interface{}   `json:"all"`
		ApiVersion         string        `json:"api_version"`
		AppContext         interface{}   `json:"app_context"`
		ApplicationVersion string        `json:"application_version"`
		Attribute          interface{}   `json:"attribute"`
		Attribute1         interface{}   `json:"attribute!"`
		AttributeMap       interface{}   `json:"attribute_map"`
		CacheTags          []string      `json:"cache_tags"`
		Callback           interface{}   `json:"callback"`
		Clip               interface{}   `json:"clip"`
		ClipImageWidth     interface{}   `json:"clip_image_width"`
		ClipRotation       interface{}   `json:"clip_rotation"`
		ContentFilter      interface{}   `json:"content_filter"`
		ContentReplacement string        `json:"content_replacement"`
		Count              interface{}   `json:"count"`
		Dates              interface{}   `json:"dates"`
		DefaultCount       int           `json:"default_count"`
		Delimiter          interface{}   `json:"delimiter"`
		DigitalId          interface{}   `json:"digital_id"`
		DisplayLevel       interface{}   `json:"display_level"`
		Distance           interface{}   `json:"distance"`
		DownloadOption     interface{}   `json:"downloadOption"`
		Duration           float64       `json:"duration"`
		Embed              []interface{} `json:"embed"`
		Embed1             []interface{} `json:"embed!"`
		ExcludeTerms       interface{}   `json:"excludeTerms"`
		FacetLimits        string        `json:"facetLimits"`
		FacetPrefix        interface{}   `json:"facetPrefix"`
		FacetCount         interface{}   `json:"facet_count"`
		FacetStyle         interface{}   `json:"facet_style"`
		Field              interface{}   `json:"field"`
		Format             string        `json:"format"`
		Host               string        `json:"host"`
		Ical               bool          `json:"ical"`
		Id                 string        `json:"id"`
		Iiif               bool          `json:"iiif"`
		Index              interface{}   `json:"index"`
		InputEncoding      string        `json:"inputEncoding"`
		IsPortal           interface{}   `json:"is_portal"`
		Item               interface{}   `json:"item"`
		Items              interface{}   `json:"items"`
		Keys               interface{}   `json:"keys"`
		Language           interface{}   `json:"language"`
		Latlong            interface{}   `json:"latlong"`
		Method             string        `json:"method"`
		NewSearch          interface{}   `json:"newSearch"`
		NewClipUrl         bool          `json:"new_clip_url"`
		Onsite             bool          `json:"onsite"`
		Operator           interface{}   `json:"operator"`
		OutputEncoding     string        `json:"outputEncoding"`
		PageHasCampaign    bool          `json:"page_has_campaign"`
		PathInfo           string        `json:"path_info"`
		Port               string        `json:"port"`
		Proxypath          interface{}   `json:"proxypath"`
		QueryString        string        `json:"query_string"`
		RedirectProxy      bool          `json:"redirect_proxy"`
		RedirectToItem     interface{}   `json:"redirect_to_item"`
		Referer            string        `json:"referer"`
		Region             string        `json:"region"`
		ReleaseId          int           `json:"release_id"`
		RequestParams      struct {
			CfChVerify  []string `json:"cf_ch_verify"`
			JschlAnswer []string `json:"jschl_answer"`
			JschlVc     []string `json:"jschl_vc"`
			Md          []string `json:"md"`
			Pass        []string `json:"pass"`
			R           []string `json:"r"`
		} `json:"request_params"`
		RequestUrl       string      `json:"request_url"`
		Resource         string      `json:"resource"`
		ResourceSequence interface{} `json:"resource_sequence"`
		Scheme           string      `json:"scheme"`
		SearchIn         interface{} `json:"searchIn"`
		SearchTerms      string      `json:"searchTerms"`
		Segments         interface{} `json:"segments"`
		SiteId           interface{} `json:"site_id"`
		SiteType         interface{} `json:"site_type"`
		SolrQuery        string      `json:"solrQuery"`
		SortBy           interface{} `json:"sortBy"`
		SortOrder        interface{} `json:"sortOrder"`
		StartPage        interface{} `json:"startPage"`
		Style            interface{} `json:"style"`
		Suggested        interface{} `json:"suggested"`
		Target           interface{} `json:"target"`
		Template         string      `json:"template"`
		Timestamp        float64     `json:"timestamp"`
		UnionFacets      string      `json:"unionFacets"`
		WebcastPermalink interface{} `json:"webcast_permalink"`
	} `json:"options"`
	RelatedItems []struct {
		AccessRestricted bool          `json:"access_restricted"`
		Aka              []string      `json:"aka"`
		Campaigns        []interface{} `json:"campaigns"`
		Contributor      []string      `json:"contributor"`
		Date             string        `json:"date,omitempty"`
		Dates            []time.Time   `json:"dates,omitempty"`
		Description      []string      `json:"description"`
		Digitized        bool          `json:"digitized"`
		ExtractTimestamp time.Time     `json:"extract_timestamp"`
		Group            []string      `json:"group"`
		Hassegments      bool          `json:"hassegments"`
		Id               string        `json:"id"`
		ImageUrl         []string      `json:"image_url"`
		Index            int           `json:"index"`
		Language         []string      `json:"language"`
		MimeType         []string      `json:"mime_type"`
		Number           []string      `json:"number"`
		OnlineFormat     []string      `json:"online_format"`
		OriginalFormat   []string      `json:"original_format"`
		OtherTitle       []string      `json:"other_title"`
		Partof           []string      `json:"partof"`
		Score            float64       `json:"score"`
		ShelfId          string        `json:"shelf_id"`
		Site             []string      `json:"site"`
		Subject          []string      `json:"subject"`
		Timestamp        time.Time     `json:"timestamp"`
		Title            string        `json:"title"`
		Url              string        `json:"url"`
	} `json:"related_items"`
	Resources []struct {
		Caption string        `json:"caption"`
		Files   [][]ImageFile `json:"files"`
		Image   string        `json:"image"`
		Url     string        `json:"url"`
	} `json:"resources"`
	Timestamp int64 `json:"timestamp"`
}
type ImageFile struct {
	Height   *int   `json:"height"`
	Levels   int    `json:"levels"`
	Mimetype string `json:"mimetype"`
	Url      string `json:"url"`
	Width    *int   `json:"width"`
	Info     string `json:"info,omitempty"`
	Size     int    `json:"size,omitempty"`
}
