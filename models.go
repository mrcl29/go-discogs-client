package discogs

// --- Common Models ---

// Image represents an image resource linked to a Release, Artist, or Label.
type Image struct {
	Type        string `json:"type"` // "primary" or "secondary"
	URI         string `json:"uri"`
	ResourceURL string `json:"resource_url"`
	URI150      string `json:"uri150"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
}

// Video represents a video resource (e.g., YouTube) linked to a Release or Master.
type Video struct {
	URI         string `json:"uri"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Embed       bool   `json:"embed"`
}

// Rating represents community or user rating statistics.
type Rating struct {
	Count   int     `json:"count"`
	Average float64 `json:"average"`
}

// Stats represents community ownership statistics (haves and wants).
type Stats struct {
	Community struct {
		InCollection int `json:"in_collection"`
		InWantlist   int `json:"in_wantlist"`
	} `json:"community"`
}

// --- Database Models ---

// Release represents a physical or digital music object in the Discogs database.
type Release struct {
	ID                int            `json:"id"`
	Title             string         `json:"title"`
	Year              int            `json:"year"`
	ResourceURL       string         `json:"resource_url"`
	URI               string         `json:"uri"`
	Artists           []ArtistSource `json:"artists"`
	Genres            []string       `json:"genres"`
	Styles            []string       `json:"styles"`
	Thumb             string         `json:"thumb"`
	Country           string         `json:"country"`
	Released          string         `json:"released"`
	ReleasedFormatted string         `json:"released_formatted"`
	Notes             string         `json:"notes"`
	MasterID          int            `json:"master_id"`
	MasterURL         string         `json:"master_url"`
	Tracklist         []Track        `json:"tracklist"`
	Images            []Image        `json:"images"`
	Videos            []Video        `json:"videos"`
	Labels            []LabelSource  `json:"labels"`
	Companies         []Company      `json:"companies"`
	ExtraArtists      []ArtistSource `json:"extraartists"`
	Community         CommunityStats `json:"community"`
	LowestPrice       float64        `json:"lowest_price"`
	NumForSale        int            `json:"num_for_sale"`
	DataQuality       string         `json:"data_quality"`
}

// ArtistSource represents an artist's participation in a Release.
type ArtistSource struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ANV         string `json:"anv"` // Artist Name Variation
	Join        string `json:"join"`
	Role        string `json:"role"`
	Tracks      string `json:"tracks"`
	ResourceURL string `json:"resource_url"`
}

// LabelSource represents a label's role in a Release.
type LabelSource struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Catno       string `json:"catno"`
	ResourceURL string `json:"resource_url"`
}

// Company represents a company (e.g., studio, manufacturer) involved in a Release.
type Company struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Catno          string `json:"catno"`
	EntityType     string `json:"entity_type"`
	EntityTypeName string `json:"entity_type_name"`
	ResourceURL    string `json:"resource_url"`
}

// Track represents a single audio track within a Release.
type Track struct {
	Position string `json:"position"`
	Type     string `json:"type_"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
}

// CommunityStats contains aggregate data from the Discogs community for a Release.
type CommunityStats struct {
	Status      string    `json:"status"`
	Rating      Rating    `json:"rating"`
	Have        int       `json:"have"`
	Want        int       `json:"want"`
	Submitter   UserRef   `json:"submitter"`
	Contributor []UserRef `json:"contributors"`
}

// UserRef is a lightweight reference to a Discogs user.
type UserRef struct {
	Username    string `json:"username"`
	ResourceURL string `json:"resource_url"`
}

// Master represents a master release, which groups together several physical/digital releases.
type Master struct {
	ID             int            `json:"id"`
	Title          string         `json:"title"`
	Year           int            `json:"year"`
	ResourceURL    string         `json:"resource_url"`
	URI            string         `json:"uri"`
	MainRelease    int            `json:"main_release"`
	MainReleaseURL string         `json:"main_release_url"`
	VersionsURL    string         `json:"versions_url"`
	Artists        []ArtistSource `json:"artists"`
	Genres         []string       `json:"genres"`
	Styles         []string       `json:"styles"`
	Videos         []Video        `json:"videos"`
	Images         []Image        `json:"images"`
	Tracklist      []Track        `json:"tracklist"`
	DataQuality    string         `json:"data_quality"`
}

// Artist represents a person or musical group in the Discogs database.
type Artist struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	ResourceURL    string   `json:"resource_url"`
	URI            string   `json:"uri"`
	Profile        string   `json:"profile"`
	ReleasesURL    string   `json:"releases_url"`
	NameVariations []string `json:"namevariations"`
	URLs           []string `json:"urls"`
	Members        []Member `json:"members"`
	Images         []Image  `json:"images"`
	DataQuality    string   `json:"data_quality"`
}

// Member represents a member of a musical group.
type Member struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Active      bool   `json:"active"`
	ResourceURL string `json:"resource_url"`
}

// Label represents a music label, record company, or music production entity.
type Label struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	ResourceURL string     `json:"resource_url"`
	URI         string     `json:"uri"`
	Profile     string     `json:"profile"`
	ReleasesURL string     `json:"releases_url"`
	ContactInfo string     `json:"contact_info"`
	Sublabels   []LabelRef `json:"sublabels"`
	URLs        []string   `json:"urls"`
	Images      []Image    `json:"images"`
	DataQuality string     `json:"data_quality"`
}

// LabelRef provides a lightweight reference to a Label.
type LabelRef struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
}

// SearchResult represents a single entry in a database search response.
type SearchResult struct {
	ID          int      `json:"id"`
	Type        string   `json:"type"` // "release", "master", "artist", or "label"
	Title       string   `json:"title"`
	Thumb       string   `json:"thumb"`
	ResourceURL string   `json:"resource_url"`
	URI         string   `json:"uri"`
	Year        string   `json:"year"`
	Format      []string `json:"format"`
	Label       []string `json:"label"`
	Genre       []string `json:"genre"`
	Style       []string `json:"style"`
	Country     string   `json:"country"`
	Barcode     []string `json:"barcode"`
	Catno       string   `json:"catno"`
	Community   struct {
		Want int `json:"want"`
		Have int `json:"have"`
	} `json:"community"`
}

// --- Marketplace Models ---

// Listing represents an item offered for sale in the Discogs Marketplace.
type Listing struct {
	ID              int        `json:"id"`
	ResourceURL     string     `json:"resource_url"`
	URI             string     `json:"uri"`
	Status          string     `json:"status"` // e.g., "For Sale", "Draft", "Sold"
	Price           Price      `json:"price"`
	Condition       string     `json:"condition"`
	SleeveCondition string     `json:"sleeve_condition"`
	Comments        string     `json:"comments"`
	AllowOffers     bool       `json:"allow_offers"`
	ShipsFrom       string     `json:"ships_from"`
	Posted          string     `json:"posted"`
	Seller          Seller     `json:"seller"`
	Release         ReleaseRef `json:"release"`
	Audio           bool       `json:"audio"`
}

// Price represents a monetary amount and its currency.
type Price struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}

// Seller represents a marketplace user who sells items.
type Seller struct {
	ID          int         `json:"id"`
	Username    string      `json:"username"`
	AvatarURL   string      `json:"avatar_url"`
	ResourceURL string      `json:"resource_url"`
	URL         string      `json:"url"`
	Stats       SellerStats `json:"stats"`
}

// SellerStats contains reputation and rating data for a marketplace seller.
type SellerStats struct {
	Rating string  `json:"rating"` // percentage as string
	Stars  float64 `json:"stars"`
	Total  int     `json:"total"`
}

// ReleaseRef is a lightweight summary of a Release used in Marketplace contexts.
type ReleaseRef struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Format      string `json:"format"`
	Catno       string `json:"catalog_number"`
	Year        int    `json:"year"`
	ResourceURL string `json:"resource_url"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
}

// Order represents a transaction between a buyer and a seller in the Discogs Marketplace.
type Order struct {
	ID                     string      `json:"id"`
	ResourceURL            string      `json:"resource_url"`
	MessagesURL            string      `json:"messages_url"`
	URI                    string      `json:"uri"`
	Status                 string      `json:"status"`
	NextStatus             []string    `json:"next_status"`
	Fee                    Price       `json:"fee"`
	Created                string      `json:"created"`
	LastActivity           string      `json:"last_activity"`
	Items                  []OrderItem `json:"items"`
	Shipping               Shipping    `json:"shipping"`
	ShippingAddress        string      `json:"shipping_address"`
	AdditionalInstructions string      `json:"additional_instructions"`
	Seller                 UserRef     `json:"seller"`
	Buyer                  UserRef     `json:"buyer"`
	Total                  Price       `json:"total"`
	Tracking               Tracking    `json:"tracking"`
}

// OrderItem represents a single item contained within an Order.
type OrderItem struct {
	ID              int        `json:"id"`
	Release         ReleaseRef `json:"release"`
	Price           Price      `json:"price"`
	MediaCondition  string     `json:"media_condition"`
	SleeveCondition string     `json:"sleeve_condition"`
}

// Shipping contains shipping cost and method details for an Order.
type Shipping struct {
	Currency string  `json:"currency"`
	Method   string  `json:"method"`
	Value    float64 `json:"value"`
}

// Tracking contains shipment tracking information.
type Tracking struct {
	Number  string `json:"number"`
	Carrier string `json:"carrier"`
	URL     string `json:"url"`
}

// --- User Models ---

// Profile contains public and private metadata for a Discogs user account.
type Profile struct {
	ID                   int     `json:"id"`
	Username             string  `json:"username"`
	Name                 string  `json:"name"`
	ResourceURL          string  `json:"resource_url"`
	URI                  string  `json:"uri"`
	AvatarURL            string  `json:"avatar_url"`
	BannerURL            string  `json:"banner_url"`
	HomePge              string  `json:"home_page"`
	Location             string  `json:"location"`
	Profile              string  `json:"profile"`
	Registered           string  `json:"registered"`
	Rank                 int     `json:"rank"`
	RatingAvg            float64 `json:"rating_avg"`
	NumCollection        int     `json:"num_collection"`
	NumWantlist          int     `json:"num_wantlist"`
	NumLists             int     `json:"num_lists"`
	NumPending           int     `json:"num_pending"`
	NumForSale           int     `json:"num_for_sale"`
	ReleasesContributed  int     `json:"releases_contributed"`
	ReleasesRated        int     `json:"releases_rated"`
	CurrAbbr             string  `json:"curr_abbr"`
	CollectionFoldersURL string  `json:"collection_folders_url"`
	CollectionFieldsURL  string  `json:"collection_fields_url"`
	WantlistURL          string  `json:"wantlist_url"`
	InventoryURL         string  `json:"inventory_url"`
}

// --- Collection & Wantlist Models ---

// Folder represents a named organization unit in a user's collection.
type Folder struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Count       int    `json:"count"`
	ResourceURL string `json:"resource_url"`
}

// CollectionItem represents a specific instance of a Release owned by a user.
type CollectionItem struct {
	InstanceID       int              `json:"instance_id"`
	Rating           int              `json:"rating"`
	FolderID         int              `json:"folder_id"`
	DateAdded        string           `json:"date_added"`
	ID               int              `json:"id"`
	BasicInformation BasicInformation `json:"basic_information"`
	Notes            []Note           `json:"notes"`
}

// BasicInformation groups common metadata shared across Collection, Wantlist, and List items.
type BasicInformation struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Year        int            `json:"year"`
	ResourceURL string         `json:"resource_url"`
	Thumb       string         `json:"thumb"`
	CoverImage  string         `json:"cover_image"`
	Artists     []ArtistSource `json:"artists"`
	Labels      []LabelSource  `json:"labels"`
	Formats     []Format       `json:"formats"`
	Genres      []string       `json:"genres"`
	Styles      []string       `json:"styles"`
}

// Format represents a media format (e.g., Vinyl, CD, etc.) and its descriptions.
type Format struct {
	Name         string   `json:"name"`
	Qty          string   `json:"qty"`
	Text         string   `json:"text"`
	Descriptions []string `json:"descriptions"`
}

// Note represents a custom field value attached to a Release instance.
type Note struct {
	FieldID int    `json:"field_id"`
	Value   string `json:"value"`
}

// Want represents a Release that a user has added to their Wantlist.
type Want struct {
	ID               int              `json:"id"`
	Rating           int              `json:"rating"`
	Notes            string           `json:"notes"`
	ResourceURL      string           `json:"resource_url"`
	BasicInformation BasicInformation `json:"basic_information"`
}

// List represents a custom-curated collection of Releases, Artists, or Labels.
type List struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Public      bool       `json:"public"`
	URI         string     `json:"uri"`
	ResourceURL string     `json:"resource_url"`
	DateAdded   string     `json:"date_added"`
	DateChanged string     `json:"date_changed"`
	Items       []ListItem `json:"items"`
}

// ListItem represents an individual entry in a user List.
type ListItem struct {
	ID           int    `json:"id"`
	Type         string `json:"type"` // "release", "artist", or "label"
	DisplayTitle string `json:"display_title"`
	Comment      string `json:"comment"`
	URI          string `json:"uri"`
	ResourceURL  string `json:"resource_url"`
	ImageURL     string `json:"image_url"`
}

// --- Response Wrappers ---

// ReleaseResponse is the root response for a single release query.
type ReleaseResponse Release

// MasterVersionsResponse represents a paginated list of master release versions.
type MasterVersionsResponse struct {
	Pagination Pagination `json:"pagination"`
	Versions   []Version  `json:"versions"`
}

// Version represents a specific physical or digital version of a Master Release.
type Version struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Format      string `json:"format"`
	Label       string `json:"label"`
	Catno       string `json:"catno"`
	Released    string `json:"released"`
	Status      string `json:"status"`
	ResourceURL string `json:"resource_url"`
	Thumb       string `json:"thumb"`
	Country     string `json:"country"`
}

// ArtistReleasesResponse represents a paginated list of an artist's releases.
type ArtistReleasesResponse struct {
	Pagination Pagination      `json:"pagination"`
	Releases   []ArtistRelease `json:"releases"`
}

// ArtistRelease is a summary of a release or master where an artist participated.
type ArtistRelease struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Type        string `json:"type"` // "release" or "master"
	Role        string `json:"role"`
	Year        int    `json:"year"`
	ResourceURL string `json:"resource_url"`
	Thumb       string `json:"thumb"`
	Artist      string `json:"artist"`
	MainRelease int    `json:"main_release"`
}

// LabelReleasesResponse represents a paginated list of a label's releases.
type LabelReleasesResponse struct {
	Pagination Pagination     `json:"pagination"`
	Releases   []LabelRelease `json:"releases"`
}

// LabelRelease is a summary of a release published by a label.
type LabelRelease struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Format      string `json:"format"`
	Catno       string `json:"catno"`
	Year        int    `json:"year"`
	ResourceURL string `json:"resource_url"`
	Status      string `json:"status"`
	Thumb       string `json:"thumb"`
}

// SearchResponse represents the full response from a database search.
type SearchResponse struct {
	Pagination Pagination     `json:"pagination"`
	Results    []SearchResult `json:"results"`
}

// InventoryResponse represents a paginated list of seller inventory listings.
type InventoryResponse struct {
	Pagination Pagination `json:"pagination"`
	Listings   []Listing  `json:"listings"`
}

// OrdersResponse represents a paginated list of marketplace orders.
type OrdersResponse struct {
	Pagination Pagination `json:"pagination"`
	Orders     []Order    `json:"orders"`
}

// CollectionFoldersResponse represents a list of collection folders.
type CollectionFoldersResponse struct {
	Folders []Folder `json:"folders"`
}

// CollectionItemsResponse represents a paginated list of items in a collection.
type CollectionItemsResponse struct {
	Pagination Pagination       `json:"pagination"`
	Releases   []CollectionItem `json:"releases"`
}

// WantlistResponse represents a paginated list of items in a user's wantlist.
type WantlistResponse struct {
	Pagination Pagination `json:"pagination"`
	Wants      []Want     `json:"wants"`
}

// UserListsResponse represents a paginated list of user-created lists.
type UserListsResponse struct {
	Pagination Pagination `json:"pagination"`
	Lists      []List     `json:"lists"`
}

// SubmissionsResponse represents a paginated list of database submissions and edits.
type SubmissionsResponse struct {
	Pagination  Pagination `json:"pagination"`
	Submissions struct {
		Artists  []Artist  `json:"artists"`
		Labels   []Label   `json:"labels"`
		Releases []Release `json:"releases"`
	} `json:"submissions"`
}

// ContributionsResponse represents a paginated list of releases contributed by a user.
type ContributionsResponse struct {
	Pagination    Pagination `json:"pagination"`
	Contributions []Release  `json:"contributions"`
}
