package applemusic

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestCatalogService_GetAlbum(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/albums/310730204", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(albumsJSON)
	})

	got, _, err := client.Catalog.GetAlbum(context.Background(), "us", "310730204", nil)
	if err != nil {
		t.Errorf("Catalog.GetAlbum returned error: %v", err)
	}
	if want := albums; !reflect.DeepEqual(got, want) {
		t.Errorf("Catalog.GetAlbumsByIds = %+v, want %+v", got, want)
	}
}

func TestCatalogService_GetAlbumsByIds(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/albums", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"ids": "310730204,190758914",
		})

		w.WriteHeader(http.StatusOK)
	})

	_, _, err := client.Catalog.GetAlbumsByIds(context.Background(), "us", []string{"310730204", "190758914"}, nil)
	if err != nil {
		t.Errorf("Catalog.GetAlbumsByIds returned error: %v", err)
	}
}

var albumsJSON = []byte(`{
  "data": [
    {
      "attributes": {
        "artistName": "Bruce Springsteen",
        "artwork": {
          "bgColor": "ffffff",
          "height": 1500,
          "textColor1": "0c0b09",
          "textColor2": "2a2724",
          "textColor3": "3d3c3a",
          "textColor4": "555250",
          "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpg",
          "width": 1500
        },
        "copyright": "\u2117 1975 Bruce Springsteen",
        "editorialNotes": {
          "short": "Springsteen's third album was the one that broke it all open for him.",
          "standard": "Springsteen's third album was the one that broke it all open for him, planting his tales of Jersey girls, cars, and nights spent sleeping on the beach firmly in the Top Five. He shot for an unholy hybrid of Orbison, Dylan and Spector \u2014 and actually reached it. \"Come take my hand,\" he invited in the opening lines. \"We're ridin' out tonight to case the Promised Land.\" Soon after this album, he'd discover the limits of such dreams, but here, it's a wide-open road: Even the tales of petty crime (\"Meeting Across the River\") and teen-gang violence (\"Jungleland\") are invested with all the wit and charm you can handle. Bruce's catalog is filled with one-of-a-kind albums from <i>The Wild, The Innocent and the E Street Shuffle</i> to <i>The Ghost of Tom Joad</i>. Forty years on, <i>Born to Run</i> still sits near the very top of that stack."
        },
        "genreNames": [
          "Rock",
          "Music",
          "Arena Rock",
          "Rock & Roll",
          "Pop",
          "Pop/Rock"
        ],
        "isComplete": true,
        "isSingle": false,
        "name": "Born to Run",
        "playParams": {
          "id": "310730204",
          "kind": "album"
        },
        "releaseDate": "1975-08-25",
        "trackCount": 8,
        "url": "https://itunes.apple.com/us/album/born-to-run/id310730204"
      },
      "href": "/v1/catalog/us/albums/310730204",
      "id": "310730204",
      "relationships": {
        "artists": {
          "data": [
            {
              "href": "/v1/catalog/us/artists/178834",
              "id": "178834",
              "type": "artists"
            }
          ],
          "href": "/v1/catalog/us/albums/310730204/artists"
        },
        "tracks": {
          "data": [
            {
              "attributes": {
                "artistName": "Bruce Springsteen",
                "artwork": {
                  "bgColor": "ffffff",
                  "height": 1500,
                  "textColor1": "0c0b09",
                  "textColor2": "2a2724",
                  "textColor3": "3d3c3a",
                  "textColor4": "555250",
                  "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpg",
                  "width": 1500
                },
                "composerName": "Bruce Springsteen",
                "discNumber": 1,
                "durationInMillis": 289186,
                "genreNames": [
                  "Rock",
                  "Music",
                  "Arena Rock",
                  "Rock & Roll",
                  "Pop",
                  "Pop/Rock"
                ],
                "name": "Thunder Road",
                "playParams": {
                  "id": "310730206",
                  "kind": "song"
                },
                "releaseDate": "1975-08-25",
                "trackNumber": 1,
                "url": "https://itunes.apple.com/us/album/thunder-road/id310730204?i=310730206"
              },
              "href": "/v1/catalog/us/songs/310730206",
              "id": "310730206",
              "type": "songs"
            },
            {
              "attributes": {
                "artistName": "Bruce Springsteen",
                "artwork": {
                  "bgColor": "ffffff",
                  "height": 1500,
                  "textColor1": "0c0b09",
                  "textColor2": "2a2724",
                  "textColor3": "3d3c3a",
                  "textColor4": "555250",
                  "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpg",
                  "width": 1500
                },
                "composerName": "Bruce Springsteen",
                "discNumber": 1,
                "durationInMillis": 191315,
                "genreNames": [
                  "Rock",
                  "Music",
                  "Arena Rock",
                  "Rock & Roll",
                  "Pop",
                  "Pop/Rock"
                ],
                "name": "Tenth Avenue Freeze-Out",
                "playParams": {
                  "id": "310730208",
                  "kind": "song"
                },
                "releaseDate": "1975-08-25",
                "trackNumber": 2,
                "url": "https://itunes.apple.com/us/album/tenth-avenue-freeze-out/id310730204?i=310730208"
              },
              "href": "/v1/catalog/us/songs/310730208",
              "id": "310730208",
              "type": "songs"
            },
            {
              "attributes": {
                "artistName": "Bruce Springsteen",
                "artwork": {
                  "bgColor": "ffffff",
                  "height": 1500,
                  "textColor1": "0c0b09",
                  "textColor2": "2a2724",
                  "textColor3": "3d3c3a",
                  "textColor4": "555250",
                  "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpg",
                  "width": 1500
                },
                "composerName": "Bruce Springsteen",
                "discNumber": 1,
                "durationInMillis": 181341,
                "genreNames": [
                  "Rock",
                  "Music",
                  "Arena Rock",
                  "Rock & Roll",
                  "Pop",
                  "Pop/Rock"
                ],
                "name": "Night",
                "playParams": {
                  "id": "310730209",
                  "kind": "song"
                },
                "releaseDate": "1975-08-25",
                "trackNumber": 3,
                "url": "https://itunes.apple.com/us/album/night/id310730204?i=310730209"
              },
              "href": "/v1/catalog/us/songs/310730209",
              "id": "310730209",
              "type": "songs"
            },
            {
              "attributes": {
                "artistName": "Bruce Springsteen",
                "artwork": {
                  "bgColor": "ffffff",
                  "height": 1500,
                  "textColor1": "0c0b09",
                  "textColor2": "2a2724",
                  "textColor3": "3d3c3a",
                  "textColor4": "555250",
                  "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpg",
                  "width": 1500
                },
                "composerName": "Bruce Springsteen",
                "discNumber": 1,
                "durationInMillis": 391764,
                "genreNames": [
                  "Rock",
                  "Music",
                  "Arena Rock",
                  "Rock & Roll",
                  "Pop",
                  "Pop/Rock"
                ],
                "name": "Backstreets",
                "playParams": {
                  "id": "310730213",
                  "kind": "song"
                },
                "releaseDate": "1975-08-25",
                "trackNumber": 4,
                "url": "https://itunes.apple.com/us/album/backstreets/id310730204?i=310730213"
              },
              "href": "/v1/catalog/us/songs/310730213",
              "id": "310730213",
              "type": "songs"
            },
            {
              "attributes": {
                "artistName": "Bruce Springsteen",
                "artwork": {
                  "bgColor": "ffffff",
                  "height": 1500,
                  "textColor1": "0c0b09",
                  "textColor2": "2a2724",
                  "textColor3": "3d3c3a",
                  "textColor4": "555250",
                  "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpg",
                  "width": 1500
                },
                "composerName": "Bruce Springsteen",
                "discNumber": 1,
                "durationInMillis": 269933,
                "genreNames": [
                  "Rock",
                  "Music",
                  "Arena Rock",
                  "Rock & Roll",
                  "Pop",
                  "Pop/Rock"
                ],
                "name": "Born to Run",
                "playParams": {
                  "id": "310730214",
                  "kind": "song"
                },
                "releaseDate": "1975-08-25",
                "trackNumber": 5,
                "url": "https://itunes.apple.com/us/album/born-to-run/id310730204?i=310730214"
              },
              "href": "/v1/catalog/us/songs/310730214",
              "id": "310730214",
              "type": "songs"
            },
            {
              "attributes": {
                "artistName": "Bruce Springsteen",
                "artwork": {
                  "bgColor": "ffffff",
                  "height": 1500,
                  "textColor1": "0c0b09",
                  "textColor2": "2a2724",
                  "textColor3": "3d3c3a",
                  "textColor4": "555250",
                  "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpg",
                  "width": 1500
                },
                "composerName": "Bruce Springsteen",
                "discNumber": 1,
                "durationInMillis": 270261,
                "genreNames": [
                  "Rock",
                  "Music",
                  "Arena Rock",
                  "Rock & Roll",
                  "Pop",
                  "Pop/Rock"
                ],
                "name": "She's the One",
                "playParams": {
                  "id": "310730220",
                  "kind": "song"
                },
                "releaseDate": "1975-08-25",
                "trackNumber": 6,
                "url": "https://itunes.apple.com/us/album/shes-the-one/id310730204?i=310730220"
              },
              "href": "/v1/catalog/us/songs/310730220",
              "id": "310730220",
              "type": "songs"
            },
            {
              "attributes": {
                "artistName": "Bruce Springsteen",
                "artwork": {
                  "bgColor": "ffffff",
                  "height": 1500,
                  "textColor1": "0c0b09",
                  "textColor2": "2a2724",
                  "textColor3": "3d3c3a",
                  "textColor4": "555250",
                  "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpg",
                  "width": 1500
                },
                "composerName": "Bruce Springsteen",
                "discNumber": 1,
                "durationInMillis": 198721,
                "genreNames": [
                  "Rock",
                  "Music",
                  "Arena Rock",
                  "Rock & Roll",
                  "Pop",
                  "Pop/Rock"
                ],
                "name": "Meeting Across the River",
                "playParams": {
                  "id": "310730222",
                  "kind": "song"
                },
                "releaseDate": "1975-08-25",
                "trackNumber": 7,
                "url": "https://itunes.apple.com/us/album/meeting-across-the-river/id310730204?i=310730222"
              },
              "href": "/v1/catalog/us/songs/310730222",
              "id": "310730222",
              "type": "songs"
            },
            {
              "attributes": {
                "artistName": "Bruce Springsteen",
                "artwork": {
                  "bgColor": "ffffff",
                  "height": 1500,
                  "textColor1": "0c0b09",
                  "textColor2": "2a2724",
                  "textColor3": "3d3c3a",
                  "textColor4": "555250",
                  "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpg",
                  "width": 1500
                },
                "composerName": "Bruce Springsteen",
                "discNumber": 1,
                "durationInMillis": 573496,
                "genreNames": [
                  "Rock",
                  "Music",
                  "Arena Rock",
                  "Rock & Roll",
                  "Pop",
                  "Pop/Rock"
                ],
                "name": "Jungleland",
                "playParams": {
                  "id": "310730227",
                  "kind": "song"
                },
                "releaseDate": "1975-08-25",
                "trackNumber": 8,
                "url": "https://itunes.apple.com/us/album/jungleland/id310730204?i=310730227"
              },
              "href": "/v1/catalog/us/songs/310730227",
              "id": "310730227",
              "type": "songs"
            }
          ],
          "href": "/v1/catalog/us/albums/310730204/tracks"
        }
      },
      "type": "albums"
    }
  ]
}`)

var albums = &Albums{
	Data: []Album{
		{
			Attributes: AlbumAttributes{
				ArtistName: "Bruce Springsteen",
				Artwork: Artwork{
					BgColor:    "ffffff",
					Height:     1500,
					TextColor1: "0c0b09",
					TextColor2: "2a2724",
					TextColor3: "3d3c3a",
					TextColor4: "555250",
					URL:        "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpg",
					Width:      1500,
				},
				Copyright: "\u2117 1975 Bruce Springsteen",
				EditorialNotes: EditorialNotes{
					Short:    "Springsteen's third album was the one that broke it all open for him.",
					Standard: "Springsteen's third album was the one that broke it all open for him, planting his tales of Jersey girls, cars, and nights spent sleeping on the beach firmly in the Top Five. He shot for an unholy hybrid of Orbison, Dylan and Spector \u2014 and actually reached it. \"Come take my hand,\" he invited in the opening lines. \"We're ridin' out tonight to case the Promised Land.\" Soon after this album, he'd discover the limits of such dreams, but here, it's a wide-open road: Even the tales of petty crime (\"Meeting Across the River\") and teen-gang violence (\"Jungleland\") are invested with all the wit and charm you can handle. Bruce's catalog is filled with one-of-a-kind albums from <i>The Wild, The Innocent and the E Street Shuffle</i> to <i>The Ghost of Tom Joad</i>. Forty years on, <i>Born to Run</i> still sits near the very top of that stack.",
				},
				GenreNames: []string{
					"Rock",
					"Music",
					"Arena Rock",
					"Rock & Roll",
					"Pop",
					"Pop/Rock",
				},
				IsComplete: true,
				IsSingle:   false,
				Name:       "Born to Run",
				PlayParams: PlayParameters{
					Id:   "310730204",
					Kind: "album",
				},
				ReleaseDate: "1975-08-25",
				TrackCount:  8,
				URL:         "https://itunes.apple.com/us/album/born-to-run/id310730204",
			},
			Href: "/v1/catalog/us/albums/310730204",
			Id:   "310730204",
			Relationships: AlbumRelationships{
				Artists: Artists{
					Data: []Artist{
						{
							Href: "/v1/catalog/us/artists/178834",
							Id:   "178834",
							Type: "artists",
						},
					},
					Href: "/v1/catalog/us/albums/310730204/artists",
				},
				Tracks: json.RawMessage(`{
          "data": [
            {
              "attributes": {
                "artistName": "Bruce Springsteen",
                "artwork": {
                  "bgColor": "ffffff",
                  "height": 1500,
                  "textColor1": "0c0b09",
                  "textColor2": "2a2724",
                  "textColor3": "3d3c3a",
                  "textColor4": "555250",
                  "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpg",
                  "width": 1500
                },
                "composerName": "Bruce Springsteen",
                "discNumber": 1,
                "durationInMillis": 289186,
                "genreNames": [
                  "Rock",
                  "Music",
                  "Arena Rock",
                  "Rock & Roll",
                  "Pop",
                  "Pop/Rock"
                ],
                "name": "Thunder Road",
                "playParams": {
                  "id": "310730206",
                  "kind": "song"
                },
                "releaseDate": "1975-08-25",
                "trackNumber": 1,
                "url": "https://itunes.apple.com/us/album/thunder-road/id310730204?i=310730206"
              },
              "href": "/v1/catalog/us/songs/310730206",
              "id": "310730206",
              "type": "songs"
            },
            {
              "attributes": {
                "artistName": "Bruce Springsteen",
                "artwork": {
                  "bgColor": "ffffff",
                  "height": 1500,
                  "textColor1": "0c0b09",
                  "textColor2": "2a2724",
                  "textColor3": "3d3c3a",
                  "textColor4": "555250",
                  "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpg",
                  "width": 1500
                },
                "composerName": "Bruce Springsteen",
                "discNumber": 1,
                "durationInMillis": 191315,
                "genreNames": [
                  "Rock",
                  "Music",
                  "Arena Rock",
                  "Rock & Roll",
                  "Pop",
                  "Pop/Rock"
                ],
                "name": "Tenth Avenue Freeze-Out",
                "playParams": {
                  "id": "310730208",
                  "kind": "song"
                },
                "releaseDate": "1975-08-25",
                "trackNumber": 2,
                "url": "https://itunes.apple.com/us/album/tenth-avenue-freeze-out/id310730204?i=310730208"
              },
              "href": "/v1/catalog/us/songs/310730208",
              "id": "310730208",
              "type": "songs"
            },
            {
              "attributes": {
                "artistName": "Bruce Springsteen",
                "artwork": {
                  "bgColor": "ffffff",
                  "height": 1500,
                  "textColor1": "0c0b09",
                  "textColor2": "2a2724",
                  "textColor3": "3d3c3a",
                  "textColor4": "555250",
                  "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpg",
                  "width": 1500
                },
                "composerName": "Bruce Springsteen",
                "discNumber": 1,
                "durationInMillis": 181341,
                "genreNames": [
                  "Rock",
                  "Music",
                  "Arena Rock",
                  "Rock & Roll",
                  "Pop",
                  "Pop/Rock"
                ],
                "name": "Night",
                "playParams": {
                  "id": "310730209",
                  "kind": "song"
                },
                "releaseDate": "1975-08-25",
                "trackNumber": 3,
                "url": "https://itunes.apple.com/us/album/night/id310730204?i=310730209"
              },
              "href": "/v1/catalog/us/songs/310730209",
              "id": "310730209",
              "type": "songs"
            },
            {
              "attributes": {
                "artistName": "Bruce Springsteen",
                "artwork": {
                  "bgColor": "ffffff",
                  "height": 1500,
                  "textColor1": "0c0b09",
                  "textColor2": "2a2724",
                  "textColor3": "3d3c3a",
                  "textColor4": "555250",
                  "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpg",
                  "width": 1500
                },
                "composerName": "Bruce Springsteen",
                "discNumber": 1,
                "durationInMillis": 391764,
                "genreNames": [
                  "Rock",
                  "Music",
                  "Arena Rock",
                  "Rock & Roll",
                  "Pop",
                  "Pop/Rock"
                ],
                "name": "Backstreets",
                "playParams": {
                  "id": "310730213",
                  "kind": "song"
                },
                "releaseDate": "1975-08-25",
                "trackNumber": 4,
                "url": "https://itunes.apple.com/us/album/backstreets/id310730204?i=310730213"
              },
              "href": "/v1/catalog/us/songs/310730213",
              "id": "310730213",
              "type": "songs"
            },
            {
              "attributes": {
                "artistName": "Bruce Springsteen",
                "artwork": {
                  "bgColor": "ffffff",
                  "height": 1500,
                  "textColor1": "0c0b09",
                  "textColor2": "2a2724",
                  "textColor3": "3d3c3a",
                  "textColor4": "555250",
                  "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpg",
                  "width": 1500
                },
                "composerName": "Bruce Springsteen",
                "discNumber": 1,
                "durationInMillis": 269933,
                "genreNames": [
                  "Rock",
                  "Music",
                  "Arena Rock",
                  "Rock & Roll",
                  "Pop",
                  "Pop/Rock"
                ],
                "name": "Born to Run",
                "playParams": {
                  "id": "310730214",
                  "kind": "song"
                },
                "releaseDate": "1975-08-25",
                "trackNumber": 5,
                "url": "https://itunes.apple.com/us/album/born-to-run/id310730204?i=310730214"
              },
              "href": "/v1/catalog/us/songs/310730214",
              "id": "310730214",
              "type": "songs"
            },
            {
              "attributes": {
                "artistName": "Bruce Springsteen",
                "artwork": {
                  "bgColor": "ffffff",
                  "height": 1500,
                  "textColor1": "0c0b09",
                  "textColor2": "2a2724",
                  "textColor3": "3d3c3a",
                  "textColor4": "555250",
                  "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpg",
                  "width": 1500
                },
                "composerName": "Bruce Springsteen",
                "discNumber": 1,
                "durationInMillis": 270261,
                "genreNames": [
                  "Rock",
                  "Music",
                  "Arena Rock",
                  "Rock & Roll",
                  "Pop",
                  "Pop/Rock"
                ],
                "name": "She's the One",
                "playParams": {
                  "id": "310730220",
                  "kind": "song"
                },
                "releaseDate": "1975-08-25",
                "trackNumber": 6,
                "url": "https://itunes.apple.com/us/album/shes-the-one/id310730204?i=310730220"
              },
              "href": "/v1/catalog/us/songs/310730220",
              "id": "310730220",
              "type": "songs"
            },
            {
              "attributes": {
                "artistName": "Bruce Springsteen",
                "artwork": {
                  "bgColor": "ffffff",
                  "height": 1500,
                  "textColor1": "0c0b09",
                  "textColor2": "2a2724",
                  "textColor3": "3d3c3a",
                  "textColor4": "555250",
                  "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpg",
                  "width": 1500
                },
                "composerName": "Bruce Springsteen",
                "discNumber": 1,
                "durationInMillis": 198721,
                "genreNames": [
                  "Rock",
                  "Music",
                  "Arena Rock",
                  "Rock & Roll",
                  "Pop",
                  "Pop/Rock"
                ],
                "name": "Meeting Across the River",
                "playParams": {
                  "id": "310730222",
                  "kind": "song"
                },
                "releaseDate": "1975-08-25",
                "trackNumber": 7,
                "url": "https://itunes.apple.com/us/album/meeting-across-the-river/id310730204?i=310730222"
              },
              "href": "/v1/catalog/us/songs/310730222",
              "id": "310730222",
              "type": "songs"
            },
            {
              "attributes": {
                "artistName": "Bruce Springsteen",
                "artwork": {
                  "bgColor": "ffffff",
                  "height": 1500,
                  "textColor1": "0c0b09",
                  "textColor2": "2a2724",
                  "textColor3": "3d3c3a",
                  "textColor4": "555250",
                  "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpg",
                  "width": 1500
                },
                "composerName": "Bruce Springsteen",
                "discNumber": 1,
                "durationInMillis": 573496,
                "genreNames": [
                  "Rock",
                  "Music",
                  "Arena Rock",
                  "Rock & Roll",
                  "Pop",
                  "Pop/Rock"
                ],
                "name": "Jungleland",
                "playParams": {
                  "id": "310730227",
                  "kind": "song"
                },
                "releaseDate": "1975-08-25",
                "trackNumber": 8,
                "url": "https://itunes.apple.com/us/album/jungleland/id310730204?i=310730227"
              },
              "href": "/v1/catalog/us/songs/310730227",
              "id": "310730227",
              "type": "songs"
            }
          ],
          "href": "/v1/catalog/us/albums/310730204/tracks"
        }`),
			},
			Type: "albums",
		},
	},
}
