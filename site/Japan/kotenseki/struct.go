package kotenseki

type Canvases struct {
	ImgUrls  []string
	IiifUrls []string
	Size     int
}

type Manifest struct {
	Context  string `json:"@context"`
	Id       string `json:"@id"`
	Type     string `json:"@type"`
	Metadata []struct {
		Label string `json:"label"`
		Value string `json:"value"`
	} `json:"metadata"`
	Label            string `json:"label"`
	Attribution      string `json:"attribution"`
	License          string `json:"license"`
	ViewingDirection string `json:"viewingDirection"`
	ViewingHint      string `json:"viewingHint"`
	Sequences        []struct {
		Id       string `json:"@id"`
		Type     string `json:"@type"`
		Canvases []struct {
			Id     string `json:"@id"`
			Type   string `json:"@type"`
			Label  string `json:"label"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
			Images []struct {
				Id         string `json:"@id"`
				Type       string `json:"@type"`
				Motivation string `json:"motivation"`
				Resource   struct {
					Id           string `json:"@id"`
					Type         string `json:"@type"`
					Format       string `json:"format"`
					Height       int    `json:"height"`
					Width        int    `json:"width"`
					DcIdentifier string `json:"dc:identifier"`
					Service      struct {
						Context string `json:"@context"`
						Id      string `json:"@id"`
						Profile string `json:"profile"`
					} `json:"service"`
				} `json:"resource"`
				On string `json:"on"`
			} `json:"images"`
			OtherContent []struct {
				Context string `json:"@context"`
				Id      string `json:"@id"`
				Type    string `json:"@type"`
			} `json:"otherContent,omitempty"`
		} `json:"canvases"`
	} `json:"sequences"`
}

type WsSearch2 struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total    int     `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			Id     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				SBid           string `json:"s.bid"`
				DBid           string `json:"d.bid"`
				SHshomeih      string `json:"s.hshomeih"`
				DHshomeih      string `json:"d.hshomeih"`
				EnSHshomeih    string `json:"en.s.hshomeih"`
				EnDHshomeih    string `json:"en.d.hshomeih"`
				SHshomeiy      string `json:"s.hshomeiy"`
				DHshomeir      string `json:"d.hshomeir"`
				DHshomeiy      string `json:"d.hshomeiy"`
				SHshomeir      string `json:"s.hshomeir"`
				EnSHshomeiy    string `json:"en.s.hshomeiy"`
				EnDHshomeiy    string `json:"en.d.hshomeiy"`
				SKansha        string `json:"s.kansha"`
				DKansha        string `json:"d.kansha"`
				OKansha        string `json:"o.kansha"`
				EnSKansha      string `json:"en.s.kansha"`
				EnDKansha      string `json:"en.d.kansha"`
				SKeitai        string `json:"s.keitai"`
				DKeitai        string `json:"d.keitai"`
				SSatsusu       string `json:"s.satsusu"`
				DSatsusu       string `json:"d.satsusu"`
				SBchuki        string `json:"s.bchuki"`
				SCallno        string `json:"s.callno"`
				DCallno        string `json:"d.callno"`
				SBshubetsu     string `json:"s.bshubetsu"`
				DBshubetsu     string `json:"d.bshubetsu"`
				OBshubetsu     string `json:"o.bshubetsu"`
				SGdate         string `json:"s.gdate"`
				SWid           string `json:"s.wid"`
				SCid           string `json:"s.cid"`
				SKzsj          string `json:"s.kzsj"`
				DMshomeihoka   string `json:"d.mshomeihoka"`
				EnDMshomeihoka string `json:"en.d.mshomeihoka"`
				DCallnoh       string `json:"d.callnoh"`
				DBshubetsuh    string `json:"d.bshubetsuh"`
				EnDBshubetsuh  string `json:"en.d.bshubetsuh"`
				DHshomei       string `json:"d.hshomei"`
				EnDHshomei     string `json:"en.d.hshomei"`
				TKshomei       []struct {
					SNum       string `json:"s.num"`
					SKshomeih  string `json:"s.kshomeih"`
					DKshomeih  string `json:"d.kshomeih"`
					SKshomeiy  string `json:"s.kshomeiy"`
					DKshomeir  string `json:"d.kshomeir"`
					DKshomeiy  string `json:"d.kshomeiy"`
					SKshomeir  string `json:"s.kshomeir"`
					DKshomei   string `json:"d.kshomei"`
					EnDKshomei string `json:"en.d.kshomei"`
				} `json:"t.kshomei"`
				DShomeih       []string `json:"d.shomeih"`
				EnDShomeih     []string `json:"en.d.shomeih"`
				DShomeiy       []string `json:"d.shomeiy"`
				EnDShomeiy     []string `json:"en.d.shomeiy"`
				OMshomeihokay  string   `json:"o.mshomeihokay"`
				SDigcallno     string   `json:"s.digcallno"`
				DDigcallno     string   `json:"d.digcallno"`
				SDserviceclass string   `json:"s.dserviceclass"`
				DDserviceclass string   `json:"d.dserviceclass"`
				DFilmkeitai    string   `json:"d.filmkeitai"`
				DServiceclass  string   `json:"d.serviceclass"`
				STshomeih      string   `json:"s.tshomeih"`
				DTshomeih      string   `json:"d.tshomeih"`
				EnSTshomeih    string   `json:"en.s.tshomeih"`
				EnDTshomeih    string   `json:"en.d.tshomeih"`
				STshomeiy      string   `json:"s.tshomeiy"`
				DTshomeir      string   `json:"d.tshomeir"`
				DTshomeiy      string   `json:"d.tshomeiy"`
				STshomeir      string   `json:"s.tshomeir"`
				EnSTshomeiy    string   `json:"en.s.tshomeiy"`
				EnDTshomeiy    string   `json:"en.d.tshomeiy"`
				SKokusho       string   `json:"s.kokusho"`
				DKokusho       string   `json:"d.kokusho"`
				OKokusho       string   `json:"o.kokusho"`
				SDomeishono    string   `json:"s.domeishono"`
				SKansatsu      string   `json:"s.kansatsu"`
				SSeiritsu      string   `json:"s.seiritsu"`
				SWchuki        string   `json:"s.wchuki"`
				DWchuki        string   `json:"d.wchuki"`
				DWshubetsu     string   `json:"d.wshubetsu"`
				TBshomei       []struct {
					SBshomeino string `json:"s.bshomeino"`
					SBshomeih  string `json:"s.bshomeih"`
					DBshomeih  string `json:"d.bshomeih"`
					SBshomeiy  string `json:"s.bshomeiy"`
					DBshomeir  string `json:"d.bshomeir"`
					DBshomeiy  string `json:"d.bshomeiy"`
					SBshomeir  string `json:"s.bshomeir"`
					DBshomei   string `json:"d.bshomei"`
					EnDBshomei string `json:"en.d.bshomei"`
				} `json:"t.bshomei"`
				TWorkAuth []struct {
					SAid      string `json:"s.aid"`
					STchoshah string `json:"s.tchoshah"`
					DTchoshah string `json:"d.tchoshah"`
					STchoshay string `json:"s.tchoshay"`
					DTchoshar string `json:"d.tchoshar"`
					DTchoshay string `json:"d.tchoshay"`
					STchoshar string `json:"s.tchoshar"`
					TAlias    []struct {
						SBeshono string `json:"s.beshono"`
						SBeshoh  string `json:"s.beshoh"`
						DBeshoh  string `json:"d.beshoh"`
						SBeshoy  string `json:"s.beshoy"`
						DBeshor  string `json:"d.beshor"`
						DBeshoy  string `json:"d.beshoy"`
						SBeshor  string `json:"s.beshor"`
					} `json:"t.Alias"`
				} `json:"t.Work.auth"`
				DChosha        []string      `json:"d.chosha"`
				EnDChosha      []string      `json:"en.d.chosha"`
				DMchosha       string        `json:"d.mchosha"`
				EnDMchosha     string        `json:"en.d.mchosha"`
				SSid           string        `json:"s.sid"`
				DCryakushohh   string        `json:"d.cryakushohh"`
				EnDCryakushohh string        `json:"en.d.cryakushohh"`
				DCollect       string        `json:"d.Collect"`
				EnDCollect     string        `json:"en.d.Collect"`
				TMokuji        []interface{} `json:"t.mokuji"`
				TIiiflink      []struct {
					DIiiflabel     string `json:"d.iiiflabel"`
					DIIIFMani      string `json:"d.IIIFMani"`
					SSearchFlgIIIF string `json:"s.searchFlgIIIF"`
				} `json:"t.iiiflink"`
				SVolume  string `json:"s.volume"`
				TSuggest []struct {
					SSuggest string `json:"s.suggest"`
				} `json:"t.suggest"`
				ORegsort string `json:"o.regsort"`
				SPubflg  string `json:"s.pubflg"`
				Thumb    string `json:"thumb"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type WsSearch struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total    int     `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			Id     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				SBid           string   `json:"s.bid"`
				DBid           string   `json:"d.bid"`
				SHshomeih      string   `json:"s.hshomeih"`
				DHshomeih      string   `json:"d.hshomeih"`
				EnSHshomeih    string   `json:"en.s.hshomeih"`
				EnDHshomeih    string   `json:"en.d.hshomeih"`
				SHshomeiy      string   `json:"s.hshomeiy"`
				DHshomeir      string   `json:"d.hshomeir"`
				DHshomeiy      string   `json:"d.hshomeiy"`
				SHshomeir      string   `json:"s.hshomeir"`
				EnSHshomeiy    string   `json:"en.s.hshomeiy"`
				EnDHshomeiy    string   `json:"en.d.hshomeiy"`
				SKansu         string   `json:"s.kansu"`
				DKansu         string   `json:"d.kansu"`
				SKansha        string   `json:"s.kansha"`
				DKansha        string   `json:"d.kansha"`
				OKansha        string   `json:"o.kansha"`
				EnSKansha      string   `json:"en.s.kansha"`
				EnDKansha      string   `json:"en.d.kansha"`
				SSatsusu       string   `json:"s.satsusu"`
				DSatsusu       string   `json:"d.satsusu"`
				SBchuki        string   `json:"s.bchuki"`
				SCallno        string   `json:"s.callno"`
				DCallno        string   `json:"d.callno"`
				SBshubetsu     string   `json:"s.bshubetsu"`
				DBshubetsu     string   `json:"d.bshubetsu"`
				OBshubetsu     string   `json:"o.bshubetsu"`
				SGdate         string   `json:"s.gdate"`
				SWid           string   `json:"s.wid"`
				SCid           string   `json:"s.cid"`
				SPubflg        string   `json:"s.pubflg"`
				SKzsj          string   `json:"s.kzsj"`
				DMshomeihoka   string   `json:"d.mshomeihoka"`
				EnDMshomeihoka string   `json:"en.d.mshomeihoka"`
				DCallnoh       string   `json:"d.callnoh"`
				DBshubetsuh    string   `json:"d.bshubetsuh"`
				EnDBshubetsuh  string   `json:"en.d.bshubetsuh"`
				DHshomei       string   `json:"d.hshomei"`
				EnDHshomei     string   `json:"en.d.hshomei"`
				DKanshah       []string `json:"d.kanshah"`
				DDoi           string   `json:"d.doi"`
				EnDKanshah     []string `json:"en.d.kanshah"`
				DShomeih       []string `json:"d.shomeih"`
				EnDShomeih     []string `json:"en.d.shomeih"`
				DShomeiy       []string `json:"d.shomeiy"`
				EnDShomeiy     []string `json:"en.d.shomeiy"`
				OMshomeihokay  string   `json:"o.mshomeihokay"`
				TKchosha       []struct {
					SNum      string `json:"s.num"`
					SKchoshah string `json:"s.kchoshah"`
					DKchoshah string `json:"d.kchoshah"`
					DChosha   string `json:"d.chosha"`
					EnDChosha string `json:"en.d.chosha"`
				} `json:"t.kchosha"`
				TShuppan []struct {
					SNum    string `json:"s.num"`
					SKannen string `json:"s.kannen"`
					DKannen string `json:"d.kannen"`
				} `json:"t.shuppan"`
				DKanshahoka   string `json:"d.kanshahoka"`
				EnDKanshahoka string `json:"en.d.kanshahoka"`
				SDigcallno    string `json:"s.digcallno"`
				DDigcallno    string `json:"d.digcallno"`
				DFilmkeitai   string `json:"d.filmkeitai"`
				STshomeih     string `json:"s.tshomeih"`
				DTshomeih     string `json:"d.tshomeih"`
				EnSTshomeih   string `json:"en.s.tshomeih"`
				EnDTshomeih   string `json:"en.d.tshomeih"`
				STshomeiy     string `json:"s.tshomeiy"`
				DTshomeir     string `json:"d.tshomeir"`
				DTshomeiy     string `json:"d.tshomeiy"`
				STshomeir     string `json:"s.tshomeir"`
				EnSTshomeiy   string `json:"en.s.tshomeiy"`
				EnDTshomeiy   string `json:"en.d.tshomeiy"`
				SKokusho      string `json:"s.kokusho"`
				DKokusho      string `json:"d.kokusho"`
				OKokusho      string `json:"o.kokusho"`
				SDomeishono   string `json:"s.domeishono"`
				SKansatsu     string `json:"s.kansatsu"`
				SWkeyword     string `json:"s.wkeyword"`
				DWkeyword     string `json:"d.wkeyword"`
				DWshubetsu    string `json:"d.wshubetsu"`
				SKshozai      string `json:"s.kshozai"`
				TWorkAuth     []struct {
					SAid      string `json:"s.aid"`
					STchoshah string `json:"s.tchoshah"`
					DTchoshah string `json:"d.tchoshah"`
					STchoshay string `json:"s.tchoshay"`
					DTchoshar string `json:"d.tchoshar"`
					DTchoshay string `json:"d.tchoshay"`
					STchoshar string `json:"s.tchoshar"`
				} `json:"t.Work.auth"`
				DChosha        []string `json:"d.chosha"`
				EnDChosha      []string `json:"en.d.chosha"`
				DMchosha       string   `json:"d.mchosha"`
				EnDMchosha     string   `json:"en.d.mchosha"`
				SSid           string   `json:"s.sid"`
				DCryakushohh   string   `json:"d.cryakushohh"`
				EnDCryakushohh string   `json:"en.d.cryakushohh"`
				DCollect       string   `json:"d.Collect"`
				EnDCollect     string   `json:"en.d.Collect"`
				TBibgazo       []struct {
					SIid       string `json:"s.iid"`
					DSatustart string `json:"d.satustart"`
					DKomastart string `json:"d.komastart"`
					SFlgdl     string `json:"s.flgdl"`
				} `json:"t.bibgazo"`
				TMokuji  []interface{} `json:"t.mokuji"`
				SIid     string        `json:"s.iid"`
				SVolume  string        `json:"s.volume"`
				SLicense []string      `json:"s.license"`
				TSuggest []struct {
					SSuggest string `json:"s.suggest"`
				} `json:"t.suggest"`
				Thumb    string `json:"thumb"`
				ORegsort string `json:"o.regsort"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
