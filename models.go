package discogs

// --- Common Models ---

// Image represents an image resource linked to a Release, Artist, or Label.
type Image struct {
	// Type is the image category: "primary" or "secondary".
	Type string `json:"type"`
	// URI is the full path to the image resource.
	URI string `json:"uri"`
	// ResourceURL is the API endpoint to retrieve image metadata.
	ResourceURL string `json:"resource_url"`
	// URI150 is the path to a 150x150 thumbnail version of the image.
	URI150 string `json:"uri150"`
	// Width is the image width in pixels.
	Width int `json:"width"`
	// Height is the image height in pixels.
	Height int `json:"height"`
}

// Video represents a video resource (e.g., YouTube) linked to a Release or Master.
type Video struct {
	// URI is the location of the video.
	URI string `json:"uri"`
	// Title is the name of the video.
	Title string `json:"title"`
	// Description is a short summary of the video content.
	Description string `json:"description"`
	// Duration is the length of the video in seconds.
	Duration int `json:"duration"`
	// Embed indicates if the video can be embedded in other pages.
	Embed bool `json:"embed"`
}

// Rating represents community or user rating statistics.
type Rating struct {
	// Count is the total number of ratings submitted.
	Count int `json:"count"`
	// Average is the mean rating value (typically 1.0 to 5.0).
	Average float64 `json:"average"`
}

// Stats represents community ownership statistics (haves and wants).
type Stats struct {
	Community struct {
		// InCollection is the number of users who have this release in their collection.
		InCollection int `json:"in_collection"`
		// InWantlist is the number of users who have this release in their wantlist.
		InWantlist int `json:"in_wantlist"`
	} `json:"community"`
}

// --- Database Models ---

// Release represents a physical or digital music object in the Discogs database.
type Release struct {
	// ID is the unique identifier for the release.
	ID int `json:"id"`
	// Title is the name of the release.
	Title string `json:"title"`
	// Year is the release year.
	Year int `json:"year"`
	// ResourceURL is the API endpoint for this release.
	ResourceURL string `json:"resource_url"`
	// URI is the web URL for this release.
	URI string `json:"uri"`
	// Artists is a list of primary artists associated with the release.
	Artists []ArtistSource `json:"artists"`
	// Genres is a list of high-level musical genres (e.g., "Rock", "Electronic").
	Genres []string `json:"genres"`
	// Styles is a list of specific musical styles (e.g., "Grunge", "Techno").
	Styles []string `json:"styles"`
	// Thumb is a URL to a small thumbnail image of the release.
	Thumb string `json:"thumb"`
	// Country is the country of origin for the release.
	Country string `json:"country"`
	// Released is the specific release date string (often YYYY-MM-DD).
	Released string `json:"released"`
	// ReleasedFormatted is a human-readable version of the release date.
	ReleasedFormatted string `json:"released_formatted"`
	// Notes contains additional descriptive text about the release.
	Notes string `json:"notes"`
	// MasterID is the ID of the master release this version belongs to.
	MasterID int `json:"master_id"`
	// MasterURL is the API endpoint for the master release.
	MasterURL string `json:"master_url"`
	// Tracklist is the ordered list of audio tracks.
	Tracklist []Track `json:"tracklist"`
	// Images is a list of images associated with the release.
	Images []Image `json:"images"`
	// Videos is a list of videos associated with the release.
	Videos []Video `json:"videos"`
	// Labels is a list of labels that published the release.
	Labels []LabelSource `json:"labels"`
	// Companies is a list of companies involved in the production or distribution.
	Companies []Company `json:"companies"`
	// ExtraArtists is a list of non-primary artists (e.g., producers, engineers).
	ExtraArtists []ArtistSource `json:"extraartists"`
	// Community contains community statistics and metadata.
	Community CommunityStats `json:"community"`
	// LowestPrice is the current minimum price in the marketplace.
	LowestPrice float64 `json:"lowest_price"`
	// NumForSale is the number of copies currently available for purchase.
	NumForSale int `json:"num_for_sale"`
	// DataQuality is the Discogs data quality rating for this entry.
	DataQuality string `json:"data_quality"`
}

// ArtistSource represents an artist's participation in a Release.
type ArtistSource struct {
	// ID is the unique identifier for the artist.
	ID int `json:"id"`
	// Name is the artist's name.
	Name string `json:"name"`
	// ANV is the Artist Name Variation used on this specific release.
	ANV string `json:"anv"`
	// Join is a string used to join this artist with the next one (e.g., "feat.").
	Join string `json:"join"`
	// Role is the artist's specific contribution (e.g., "Producer").
	Role string `json:"role"`
	// Tracks is a string describing which tracks the artist participated in.
	Tracks string `json:"tracks"`
	// ResourceURL is the API endpoint for the artist.
	ResourceURL string `json:"resource_url"`
}

// LabelSource represents a label's role in a Release.
type LabelSource struct {
	// ID is the unique identifier for the label.
	ID int `json:"id"`
	// Name is the label's name.
	Name string `json:"name"`
	// Catno is the catalog number assigned by the label.
	Catno string `json:"catno"`
	// ResourceURL is the API endpoint for the label.
	ResourceURL string `json:"resource_url"`
}

// Company represents a company (e.g., studio, manufacturer) involved in a Release.
type Company struct {
	// ID is the unique identifier for the company.
	ID int `json:"id"`
	// Name is the company's name.
	Name string `json:"name"`
	// Catno is the catalog number assigned by the company.
	Catno string `json:"catno"`
	// EntityType is the numeric code for the company type.
	EntityType string `json:"entity_type"`
	// EntityTypeName is the descriptive name for the company type (e.g., "Studio").
	EntityTypeName string `json:"entity_type_name"`
	// ResourceURL is the API endpoint for the company.
	ResourceURL string `json:"resource_url"`
}

// Track represents a single audio track within a Release.
type Track struct {
	// Position is the track's position in the tracklist (e.g., "A1", "1").
	Position string `json:"position"`
	// Type is the track type (e.g., "track", "heading").
	Type string `json:"type_"`
	// Title is the track name.
	Title string `json:"title"`
	// Duration is the track length as a string (e.g., "3:45").
	Duration string `json:"duration"`
}

// CommunityStats contains aggregate data from the Discogs community for a Release.
type CommunityStats struct {
	// Status is the moderation status of the release entry.
	Status string `json:"status"`
	// Rating contains community rating statistics.
	Rating Rating `json:"rating"`
	// Have is the number of users who have this release in their collection.
	Have int `json:"have"`
	// Want is the number of users who have this release in their wantlist.
	Want int `json:"want"`
	// Submitter is a reference to the user who submitted the release.
	Submitter UserRef `json:"submitter"`
	// Contributor is a list of users who have edited the release entry.
	Contributor []UserRef `json:"contributors"`
}

// UserRef is a lightweight reference to a Discogs user.
type UserRef struct {
	// Username is the user's unique name.
	Username string `json:"username"`
	// ResourceURL is the API endpoint for the user's profile.
	ResourceURL string `json:"resource_url"`
}

// Master represents a master release, which groups together several physical/digital releases.
type Master struct {
	// ID is the unique identifier for the master release.
	ID int `json:"id"`
	// Title is the name of the master release.
	Title string `json:"title"`
	// Year is the earliest release year associated with this master.
	Year int `json:"year"`
	// ResourceURL is the API endpoint for this master release.
	ResourceURL string `json:"resource_url"`
	// URI is the web URL for this master release.
	URI string `json:"uri"`
	// MainRelease is the ID of the primary version of this master.
	MainRelease int `json:"main_release"`
	// MainReleaseURL is the API endpoint for the main release.
	MainReleaseURL string `json:"main_release_url"`
	// VersionsURL is the API endpoint to retrieve all versions of this master.
	VersionsURL string `json:"versions_url"`
	// Artists is a list of primary artists associated with the master release.
	Artists []ArtistSource `json:"artists"`
	// Genres is a list of high-level musical genres.
	Genres []string `json:"genres"`
	// Styles is a list of specific musical styles.
	Styles []string `json:"styles"`
	// Videos is a list of videos associated with the master release.
	Videos []Video `json:"videos"`
	// Images is a list of images associated with the master release.
	Images []Image `json:"images"`
	// Tracklist is the ordered list of audio tracks.
	Tracklist []Track `json:"tracklist"`
	// DataQuality is the Discogs data quality rating for this entry.
	DataQuality string `json:"data_quality"`
}

// Artist represents a person or musical group in the Discogs database.
type Artist struct {
	// ID is the unique identifier for the artist.
	ID int `json:"id"`
	// Name is the artist's name.
	Name string `json:"name"`
	// ResourceURL is the API endpoint for this artist.
	ResourceURL string `json:"resource_url"`
	// URI is the web URL for this artist.
	URI string `json:"uri"`
	// Profile is a descriptive biography of the artist.
	Profile string `json:"profile"`
	// ReleasesURL is the API endpoint to retrieve all releases by this artist.
	ReleasesURL string `json:"releases_url"`
	// NameVariations is a list of alternative names used by the artist.
	NameVariations []string `json:"namevariations"`
	// URLs is a list of official or related websites for the artist.
	URLs []string `json:"urls"`
	// Members is a list of group members (if the artist is a group).
	Members []Member `json:"members"`
	// Images is a list of images associated with the artist.
	Images []Image `json:"images"`
	// DataQuality is the Discogs data quality rating for this entry.
	DataQuality string `json:"data_quality"`
}

// Member represents a member of a musical group.
type Member struct {
	// ID is the unique identifier for the member (who is also an artist).
	ID int `json:"id"`
	// Name is the member's name.
	Name string `json:"name"`
	// Active indicates if the member is currently part of the group.
	Active bool `json:"active"`
	// ResourceURL is the API endpoint for the member's artist profile.
	ResourceURL string `json:"resource_url"`
}

// Label represents a music label, record company, or music production entity.
type Label struct {
	// ID is the unique identifier for the label.
	ID int `json:"id"`
	// Name is the label's name.
	Name string `json:"name"`
	// ResourceURL is the API endpoint for this label.
	ResourceURL string `json:"resource_url"`
	// URI is the web URL for this label.
	URI string `json:"uri"`
	// Profile is a descriptive history or biography of the label.
	Profile string `json:"profile"`
	// ReleasesURL is the API endpoint to retrieve all releases published by this label.
	ReleasesURL string `json:"releases_url"`
	// ContactInfo is the mailing or email address for the label.
	ContactInfo string `json:"contact_info"`
	// Sublabels is a list of labels owned or managed by this label.
	Sublabels []LabelRef `json:"sublabels"`
	// URLs is a list of official or related websites for the label.
	URLs []string `json:"urls"`
	// Images is a list of images associated with the label.
	Images []Image `json:"images"`
	// DataQuality is the Discogs data quality rating for this entry.
	DataQuality string `json:"data_quality"`
}

// LabelRef provides a lightweight reference to a Label.
type LabelRef struct {
	// ID is the unique identifier for the label.
	ID int `json:"id"`
	// Name is the label's name.
	Name string `json:"name"`
	// ResourceURL is the API endpoint for the label.
	ResourceURL string `json:"resource_url"`
}

// SearchResult represents a single entry in a database search response.
type SearchResult struct {
	// ID is the unique identifier for the result.
	ID int `json:"id"`
	// Type is the resource type: "release", "master", "artist", or "label".
	Type string `json:"type"`
	// Title is the name of the result.
	Title string `json:"title"`
	// Thumb is a URL to a small thumbnail image.
	Thumb string `json:"thumb"`
	// ResourceURL is the API endpoint for the result.
	ResourceURL string `json:"resource_url"`
	// URI is the web URL for the result.
	URI string `json:"uri"`
	// Year is the release year (as a string).
	Year string `json:"year"`
	// Format is a list of formats associated with the result.
	Format []string `json:"format"`
	// Label is a list of labels associated with the result.
	Label []string `json:"label"`
	// Genre is a list of genres.
	Genre []string `json:"genre"`
	// Style is a list of styles.
	Style []string `json:"style"`
	// Country is the country of origin.
	Country string `json:"country"`
	// Barcode is a list of barcodes found on the release.
	Barcode []string `json:"barcode"`
	// Catno is the catalog number.
	Catno string `json:"catno"`
	// Community contains community statistics.
	Community struct {
		// Want is the number of users who want this.
		Want int `json:"want"`
		// Have is the number of users who have this.
		Have int `json:"have"`
	} `json:"community"`
}

// --- Marketplace Models ---

// Listing represents an item offered for sale in the Discogs Marketplace.
type Listing struct {
	// ID is the unique identifier for the listing.
	ID int `json:"id"`
	// ResourceURL is the API endpoint for this listing.
	ResourceURL string `json:"resource_url"`
	// URI is the web URL for this listing.
	URI string `json:"uri"`
	// Status is the listing state: e.g., "For Sale", "Draft", "Sold".
	Status string `json:"status"`
	// Price contains the item's cost and currency.
	Price Price `json:"price"`
	// Condition is the media condition (e.g., "Near Mint (NM or M-)").
	Condition string `json:"condition"`
	// SleeveCondition is the sleeve condition (e.g., "Very Good Plus (VG+)").
	SleeveCondition string `json:"sleeve_condition"`
	// Comments is a descriptive note from the seller.
	Comments string `json:"comments"`
	// AllowOffers indicates if the seller accepts price negotiations.
	AllowOffers bool `json:"allow_offers"`
	// ShipsFrom is the location the item is sent from.
	ShipsFrom string `json:"ships_from"`
	// Posted is the date the listing was created.
	Posted string `json:"posted"`
	// Seller is a reference to the user selling the item.
	Seller Seller `json:"seller"`
	// Release is a reference to the release being sold.
	Release ReleaseRef `json:"release"`
	// Audio indicates if the listing includes a sample audio clip.
	Audio bool `json:"audio"`
}

// Price represents a monetary amount and its currency.
type Price struct {
	// Currency is the 3-letter currency code (e.g., "USD", "EUR").
	Currency string `json:"currency"`
	// Value is the numeric price amount.
	Value float64 `json:"value"`
}

// Seller represents a marketplace user who sells items.
type Seller struct {
	// ID is the unique identifier for the seller.
	ID int `json:"id"`
	// Username is the seller's unique name.
	Username string `json:"username"`
	// AvatarURL is a URL to the seller's profile image.
	AvatarURL string `json:"avatar_url"`
	// ResourceURL is the API endpoint for the seller's profile.
	ResourceURL string `json:"resource_url"`
	// URL is the web URL for the seller's store.
	URL string `json:"url"`
	// Stats contains seller reputation and rating data.
	Stats SellerStats `json:"stats"`
}

// SellerStats contains reputation and rating data for a marketplace seller.
type SellerStats struct {
	// Rating is the percentage of positive feedback (as a string).
	Rating string `json:"rating"`
	// Stars is the average star rating.
	Stars float64 `json:"stars"`
	// Total is the total number of feedback ratings received.
	Total int `json:"total"`
}

// ReleaseRef is a lightweight summary of a Release used in Marketplace contexts.
type ReleaseRef struct {
	// ID is the unique identifier for the release.
	ID int `json:"id"`
	// Title is the name of the release.
	Title string `json:"title"`
	// Artist is the primary artist's name.
	Artist string `json:"artist"`
	// Format is a descriptive string of the media format.
	Format string `json:"format"`
	// Catno is the catalog number.
	Catno string `json:"catalog_number"`
	// Year is the release year.
	Year int `json:"year"`
	// ResourceURL is the API endpoint for the release.
	ResourceURL string `json:"resource_url"`
	// Thumbnail is a URL to a small release image.
	Thumbnail string `json:"thumbnail"`
	// Description is a descriptive summary of the release.
	Description string `json:"description"`
}

// Order represents a transaction between a buyer and a seller in the Discogs Marketplace.
type Order struct {
	// ID is the unique identifier for the order.
	ID string `json:"id"`
	// ResourceURL is the API endpoint for this order.
	ResourceURL string `json:"resource_url"`
	// MessagesURL is the API endpoint to retrieve the order communication log.
	MessagesURL string `json:"messages_url"`
	// URI is the web URL for this order.
	URI string `json:"uri"`
	// Status is the order state: e.g., "New Order", "Shipped", "Cancelled".
	Status string `json:"status"`
	// NextStatus is a list of valid status transitions for this order.
	NextStatus []string `json:"next_status"`
	// Fee is the Discogs selling fee for this order.
	Fee Price `json:"fee"`
	// Created is the date the order was placed.
	Created string `json:"created"`
	// LastActivity is the date of the most recent event in the order.
	LastActivity string `json:"last_activity"`
	// Items is the list of items purchased in this order.
	Items []OrderItem `json:"items"`
	// Shipping contains shipping cost and method details.
	Shipping Shipping `json:"shipping"`
	// ShippingAddress is the buyer's delivery address.
	ShippingAddress string `json:"shipping_address"`
	// AdditionalInstructions is an optional note from the buyer.
	AdditionalInstructions string `json:"additional_instructions"`
	// Seller is a reference to the user selling the items.
	Seller UserRef `json:"seller"`
	// Buyer is a reference to the user purchasing the items.
	Buyer UserRef `json:"buyer"`
	// Total is the final order amount (including shipping and fees).
	Total Price `json:"total"`
	// Tracking contains shipment tracking information.
	Tracking Tracking `json:"tracking"`
}

// OrderItem represents a single item contained within an Order.
type OrderItem struct {
	// ID is the unique identifier for the item instance.
	ID int `json:"id"`
	// Release is a reference to the release being purchased.
	Release ReleaseRef `json:"release"`
	// Price is the cost of this specific item.
	Price Price `json:"price"`
	// MediaCondition is the condition of the physical media.
	MediaCondition string `json:"media_condition"`
	// SleeveCondition is the condition of the sleeve.
	SleeveCondition string `json:"sleeve_condition"`
}

// Shipping contains shipping cost and method details for an Order.
type Shipping struct {
	// Currency is the 3-letter currency code for shipping costs.
	Currency string `json:"currency"`
	// Method is the name of the shipping service used.
	Method string `json:"method"`
	// Value is the numeric shipping cost.
	Value float64 `json:"value"`
}

// Tracking contains shipment tracking information.
type Tracking struct {
	// Number is the unique shipment tracking ID.
	Number string `json:"number"`
	// Carrier is the name of the shipping company (e.g., "USPS", "FedEx").
	Carrier string `json:"carrier"`
	// URL is a link to the carrier's tracking page.
	URL string `json:"url"`
}

// --- User Models ---

// Profile contains public and private metadata for a Discogs user account.
type Profile struct {
	// ID is the unique identifier for the user.
	ID int `json:"id"`
	// Username is the user's unique name.
	Username string `json:"username"`
	// Name is the user's real name or display name.
	Name string `json:"name"`
	// ResourceURL is the API endpoint for this profile.
	ResourceURL string `json:"resource_url"`
	// URI is the web URL for this profile.
	URI string `json:"uri"`
	// AvatarURL is a URL to the user's avatar image.
	AvatarURL string `json:"avatar_url"`
	// BannerURL is a URL to the user's profile banner image.
	BannerURL string `json:"banner_url"`
	// HomePge is the user's personal website URL.
	HomePge string `json:"home_page"`
	// Location is the user's geographical location.
	Location string `json:"location"`
	// Profile is a descriptive biography written by the user.
	Profile string `json:"profile"`
	// Registered is the date the user joined Discogs.
	Registered string `json:"registered"`
	// Rank is the user's contribution rank in the community.
	Rank int `json:"rank"`
	// RatingAvg is the average rating received as a seller.
	RatingAvg float64 `json:"rating_avg"`
	// NumCollection is the total number of items in the user's collection.
	NumCollection int `json:"num_collection"`
	// NumWantlist is the total number of items in the user's wantlist.
	NumWantlist int `json:"num_wantlist"`
	// NumLists is the total number of lists created by the user.
	NumLists int `json:"num_lists"`
	// NumPending is the number of pending submissions by the user.
	NumPending int `json:"num_pending"`
	// NumForSale is the number of items the user currently has for sale.
	NumForSale int `json:"num_for_sale"`
	// ReleasesContributed is the number of releases added by the user.
	ReleasesContributed int `json:"releases_contributed"`
	// ReleasesRated is the number of releases rated by the user.
	ReleasesRated int `json:"releases_rated"`
	// CurrAbbr is the user's preferred currency abbreviation.
	CurrAbbr string `json:"curr_abbr"`
	// CollectionFoldersURL is the API endpoint to retrieve the user's collection folders.
	CollectionFoldersURL string `json:"collection_folders_url"`
	// CollectionFieldsURL is the API endpoint to retrieve the user's custom collection fields.
	CollectionFieldsURL string `json:"collection_fields_url"`
	// WantlistURL is the API endpoint to retrieve the user's wantlist.
	WantlistURL string `json:"wantlist_url"`
	// InventoryURL is the API endpoint to retrieve the user's marketplace inventory.
	InventoryURL string `json:"inventory_url"`
}

// --- Collection & Wantlist Models ---

// Folder represents a named organization unit in a user's collection.
type Folder struct {
	// ID is the unique identifier for the folder.
	ID int `json:"id"`
	// Name is the folder's name.
	Name string `json:"name"`
	// Count is the number of releases in this folder.
	Count int `json:"count"`
	// ResourceURL is the API endpoint for this folder.
	ResourceURL string `json:"resource_url"`
}

// CollectionItem represents a specific instance of a Release owned by a user.
type CollectionItem struct {
	// InstanceID is the unique identifier for this specific copy of a release.
	InstanceID int `json:"instance_id"`
	// Rating is the user's personal rating (typically 1-5).
	Rating int `json:"rating"`
	// FolderID is the ID of the folder containing this instance.
	FolderID int `json:"folder_id"`
	// DateAdded is the date the item was added to the collection.
	DateAdded string `json:"date_added"`
	// ID is the release ID.
	ID int `json:"id"`
	// BasicInformation groups common release metadata.
	BasicInformation BasicInformation `json:"basic_information"`
	// Notes is a list of user-defined custom field values.
	Notes []Note `json:"notes"`
}

// BasicInformation groups common metadata shared across Collection, Wantlist, and List items.
type BasicInformation struct {
	// ID is the unique identifier for the release.
	ID int `json:"id"`
	// Title is the name of the release.
	Title string `json:"title"`
	// Year is the release year.
	Year int `json:"year"`
	// ResourceURL is the API endpoint for the release.
	ResourceURL string `json:"resource_url"`
	// Thumb is a URL to a small thumbnail image.
	Thumb string `json:"thumb"`
	// CoverImage is a URL to a larger cover image.
	CoverImage string `json:"cover_image"`
	// Artists is a list of primary artists.
	Artists []ArtistSource `json:"artists"`
	// Labels is a list of publishing labels.
	Labels []LabelSource `json:"labels"`
	// Formats is a list of physical/digital media formats.
	Formats []Format `json:"formats"`
	// Genres is a list of musical genres.
	Genres []string `json:"genres"`
	// Styles is a list of musical styles.
	Styles []string `json:"styles"`
}

// Format represents a media format (e.g., Vinyl, CD, etc.) and its descriptions.
type Format struct {
	// Name is the high-level format name.
	Name string `json:"name"`
	// Qty is the number of items in the set (e.g., "2" for a 2xLP).
	Qty string `json:"qty"`
	// Text is additional descriptive text about the format.
	Text string `json:"text"`
	// Descriptions is a list of specific format attributes (e.g., "LP", "Album").
	Descriptions []string `json:"descriptions"`
}

// Note represents a custom field value attached to a Release instance.
type Note struct {
	// FieldID is the unique identifier for the custom field.
	FieldID int `json:"field_id"`
	// Value is the text content of the note.
	Value string `json:"value"`
}

// Want represents a Release that a user has added to their Wantlist.
type Want struct {
	// ID is the unique identifier for the release.
	ID int `json:"id"`
	// Rating is the user's rating for the wanted item.
	Rating int `json:"rating"`
	// Notes is a personal note about why the item is wanted.
	Notes string `json:"notes"`
	// ResourceURL is the API endpoint for this wantlist entry.
	ResourceURL string `json:"resource_url"`
	// BasicInformation groups common release metadata.
	BasicInformation BasicInformation `json:"basic_information"`
}

// List represents a custom-curated collection of Releases, Artists, or Labels.
type List struct {
	// ID is the unique identifier for the list.
	ID int `json:"id"`
	// Name is the list's title.
	Name string `json:"name"`
	// Description is a short summary of the list's purpose or theme.
	Description string `json:"description"`
	// Public indicates if the list is visible to other users.
	Public bool `json:"public"`
	// URI is the web URL for this list.
	URI string `json:"uri"`
	// ResourceURL is the API endpoint for this list.
	ResourceURL string `json:"resource_url"`
	// DateAdded is the date the list was created.
	DateAdded string `json:"date_added"`
	// DateChanged is the date the list was last modified.
	DateChanged string `json:"date_changed"`
	// Items is the list of entries contained within the list.
	Items []ListItem `json:"items"`
}

// ListItem represents an individual entry in a user List.
type ListItem struct {
	// ID is the identifier for the specific entry (Release, Artist, or Label ID).
	ID int `json:"id"`
	// Type is the entry type: "release", "artist", or "label".
	Type string `json:"type"`
	// DisplayTitle is the formatted name of the entry.
	DisplayTitle string `json:"display_title"`
	// Comment is an optional personal note about this entry.
	Comment string `json:"comment"`
	// URI is the web URL for the entry.
	URI string `json:"uri"`
	// ResourceURL is the API endpoint for the entry.
	ResourceURL string `json:"resource_url"`
	// ImageURL is a URL to the entry's image.
	ImageURL string `json:"image_url"`
}

// --- Response Wrappers ---

// ReleaseResponse is the root response for a single release query.
type ReleaseResponse Release

// MasterVersionsResponse represents a paginated list of master release versions.
type MasterVersionsResponse struct {
	// Pagination contains metadata about the paginated results.
	Pagination Pagination `json:"pagination"`
	// Versions is the list of release versions belonging to the master.
	Versions []Version `json:"versions"`
}

// Version represents a specific physical or digital version of a Master Release.
type Version struct {
	// ID is the unique identifier for this version (Release ID).
	ID int `json:"id"`
	// Title is the release title.
	Title string `json:"title"`
	// Format is a descriptive string of the media format.
	Format string `json:"format"`
	// Label is the name of the primary label.
	Label string `json:"label"`
	// Catno is the catalog number.
	Catno string `json:"catno"`
	// Released is the release date string.
	Released string `json:"released"`
	// Status is the release status (e.g., "Accepted").
	Status string `json:"status"`
	// ResourceURL is the API endpoint for this version.
	ResourceURL string `json:"resource_url"`
	// Thumb is a URL to a small thumbnail image.
	Thumb string `json:"thumb"`
	// Country is the country of origin.
	Country string `json:"country"`
}

// ArtistReleasesResponse represents a paginated list of an artist's releases.
type ArtistReleasesResponse struct {
	// Pagination contains metadata about the paginated results.
	Pagination Pagination `json:"pagination"`
	// Releases is the list of releases where the artist participated.
	Releases []ArtistRelease `json:"releases"`
}

// ArtistRelease is a summary of a release or master where an artist participated.
type ArtistRelease struct {
	// ID is the identifier for the release or master.
	ID int `json:"id"`
	// Title is the name of the release or master.
	Title string `json:"title"`
	// Type is the resource type: "release" or "master".
	Type string `json:"type"`
	// Role is the artist's specific contribution (e.g., "Main", "Producer").
	Role string `json:"role"`
	// Year is the release year.
	Year int `json:"year"`
	// ResourceURL is the API endpoint for the entry.
	ResourceURL string `json:"resource_url"`
	// Thumb is a URL to a small thumbnail image.
	Thumb string `json:"thumb"`
	// Artist is the name of the artist.
	Artist string `json:"artist"`
	// MainRelease is the ID of the primary version (if Type is "master").
	MainRelease int `json:"main_release"`
}

// LabelReleasesResponse represents a paginated list of a label's releases.
type LabelReleasesResponse struct {
	// Pagination contains metadata about the paginated results.
	Pagination Pagination `json:"pagination"`
	// Releases is the list of releases published by the label.
	Releases []LabelRelease `json:"releases"`
}

// LabelRelease is a summary of a release published by a label.
type LabelRelease struct {
	// ID is the identifier for the release.
	ID int `json:"id"`
	// Title is the release title.
	Title string `json:"title"`
	// Artist is the primary artist's name.
	Artist string `json:"artist"`
	// Format is a descriptive string of the media format.
	Format string `json:"format"`
	// Catno is the catalog number.
	Catno string `json:"catno"`
	// Year is the release year.
	Year int `json:"year"`
	// ResourceURL is the API endpoint for the release.
	ResourceURL string `json:"resource_url"`
	// Status is the release status.
	Status string `json:"status"`
	// Thumb is a URL to a small thumbnail image.
	Thumb string `json:"thumb"`
}

// SearchResponse represents the full response from a database search.
type SearchResponse struct {
	// Pagination contains metadata about the paginated results.
	Pagination Pagination `json:"pagination"`
	// Results is the list of search matches.
	Results []SearchResult `json:"results"`
}

// InventoryResponse represents a paginated list of seller inventory listings.
type InventoryResponse struct {
	// Pagination contains metadata about the paginated results.
	Pagination Pagination `json:"pagination"`
	// Listings is the list of marketplace listings.
	Listings []Listing `json:"listings"`
}

// OrdersResponse represents a paginated list of marketplace orders.
type OrdersResponse struct {
	// Pagination contains metadata about the paginated results.
	Pagination Pagination `json:"pagination"`
	// Orders is the list of marketplace orders.
	Orders []Order `json:"orders"`
}

// CollectionFoldersResponse represents a list of collection folders.
type CollectionFoldersResponse struct {
	// Folders is the list of user collection folders.
	Folders []Folder `json:"folders"`
}

// CollectionItemsResponse represents a paginated list of items in a collection.
type CollectionItemsResponse struct {
	// Pagination contains metadata about the paginated results.
	Pagination Pagination `json:"pagination"`
	// Releases is the list of collection item instances.
	Releases []CollectionItem `json:"releases"`
}

// WantlistResponse represents a paginated list of items in a user's wantlist.
type WantlistResponse struct {
	// Pagination contains metadata about the paginated results.
	Pagination Pagination `json:"pagination"`
	// Wants is the list of wanted release entries.
	Wants []Want `json:"wants"`
}

// UserListsResponse represents a paginated list of user-created lists.
type UserListsResponse struct {
	// Pagination contains metadata about the paginated results.
	Pagination Pagination `json:"pagination"`
	// Lists is the list of user curated lists.
	Lists []List `json:"lists"`
}

// SubmissionsResponse represents a paginated list of database submissions and edits.
type SubmissionsResponse struct {
	// Pagination contains metadata about the paginated results.
	Pagination Pagination `json:"pagination"`
	// Submissions groups the different types of submissions.
	Submissions struct {
		// Artists is a list of submitted artists.
		Artists []Artist `json:"artists"`
		// Labels is a list of submitted labels.
		Labels []Label `json:"labels"`
		// Releases is a list of submitted releases.
		Releases []Release `json:"releases"`
	} `json:"submissions"`
}

// ContributionsResponse represents a paginated list of releases contributed by a user.
type ContributionsResponse struct {
	// Pagination contains metadata about the paginated results.
	Pagination Pagination `json:"pagination"`
	// Contributions is the list of contributed release metadata.
	Contributions []Release `json:"contributions"`
}
