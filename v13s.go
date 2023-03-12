package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	"v13s/v13s/ent"
	"v13s/v13s/ent/illusttag"
	"v13s/v13s/ent/tag"

	"entgo.io/ent/dialect"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"modernc.org/sqlite"
)

// type definition start

type Options struct {
	SiteName        string `json:"siteName"`
	SiteDescription string `json:"siteDescription"`
	ServerPort      int    `json:"serverPort"`
}

type SqliteDriver struct {
	*sqlite.Driver
}

type IllustInfo struct {
	IllustID    string    `json:"illustID"`
	Hash        string    `json:"hash"`
	OriginalURL *string   `json:"originalURL"`
	HeightRatio float32   `json:"heightRatio"`
	Tags        []TagInfo `json:"tags"`
	CreatedAt   time.Time `json:"createdAt"`
}

type TagInfo struct {
	TagID        string `json:"tagID"`
	Type         string `json:"type"`
	OfficialName string `json:"officialName"`
}

// type definition end

// http request handler function definition start

func entryPointHandler(w http.ResponseWriter, r *http.Request, options Options) {
	template := template.Must(template.ParseFiles("frontend/index.html"))

	siteDescription := func() string {
		if options.SiteName != "" {
			var desc string = "Let’s organise and archive your precious masterpieces."
			if options.SiteDescription != "" {
				desc = options.SiteDescription
			}
			return fmt.Sprintf("%s - %s Powered by Vivliothikarios(v13s)", options.SiteName, desc)
		} else {
			return "Vivliothikarios(v13s) - Let’s organise and archive your precious masterpieces."
		}
	}()

	variables := map[string]string{
		"siteName":        options.SiteName,
		"siteDescription": siteDescription,
	}

	template.ExecuteTemplate(w, "index.html", variables)
}

func illustsRequestHandler(w http.ResponseWriter, r *http.Request, entC *ent.Client, ctx context.Context) {
	illustsData, _ := entC.Illust.Query().All(ctx)

	var tmpArr []IllustInfo
	for _, v := range illustsData {
		illustTags, _ := entC.IllustTag.
			Query().
			Where(illusttag.ParentIllustID(v.IllustID)).
			All(ctx)

		tags := func() []TagInfo {
			var tmpArr []TagInfo
			for _, v := range illustTags {
				rawTag, _ := entC.Tag.
					Query().
					Where(tag.TagID(v.PrentTagID)).
					First(ctx)
				tag := TagInfo{
					TagID:        rawTag.TagID,
					Type:         string(rawTag.Type),
					OfficialName: rawTag.OfficialName,
				}
				tmpArr = append(tmpArr, tag)
			}
			return tmpArr
		}()

		illust := IllustInfo{
			IllustID:    v.IllustID,
			Hash:        v.Hash,
			OriginalURL: &v.OriginalURL,
			HeightRatio: v.HeightRatio,
			Tags:        tags,
			CreatedAt:   v.CreatedAt,
		}
		tmpArr = append(tmpArr, illust)
	}
	illustJSON, _ := json.Marshal(tmpArr)
	w.WriteHeader(200)
	w.Write(illustJSON)
}

// http request handler function definition end

// utility function definition start

func (d SqliteDriver) Open(name string) (driver.Conn, error) {
	conn, err := d.Driver.Open(name)
	if err != nil {
		return conn, err
	}
	c := conn.(interface {
		Exec(stmt string, args []driver.Value) (driver.Result, error)
	})
	if _, err := c.Exec("PRAGMA foreign_keys = on;", nil); err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to enable enable foreign keys: %w", err)
	}
	return conn, nil
}

// utility function definition end

// major function definition start

func init() {
	sql.Register("sqlite3", SqliteDriver{Driver: &sqlite.Driver{}})
}

func main() {
	entC, _ := ent.Open(dialect.SQLite, "file:data/database/v13s.db?cache=shared")
	defer entC.Close()
	ctx := context.Background()

	if err := entC.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	//initialize(entC, ctx)

	_, err := os.Stat("v13s.config.json")
	if err != nil {
		options := Options{
			SiteName:        "",
			SiteDescription: "",
			ServerPort:      9090,
		}

		rawOptions, _ := json.MarshalIndent(options, "", "\t")

		f, _ := os.Create("v13s.config.json")
		f.Write(rawOptions)
	}

	rawOptions, _ := os.ReadFile("v13s.config.json")
	var options Options
	json.Unmarshal(rawOptions, &options)

	func() {
		errVal := func() string {
			if options.SiteName == "" {
				return "siteName"
			} else if options.ServerPort == 0 {
				return "serverPort"
			} else {
				return ""
			}
		}()
		if errVal != "" {
			fmt.Fprintf(os.Stderr, "Error: Option value \"%s\" (at v13s.config.json) is invalid.\n", errVal)
			os.Exit(2)
		}
	}()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/static/"))))
	http.HandleFunc("/asset/illust/.webp", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	http.Handle("/asset/illust/", http.StripPrefix("/asset/illust/", http.FileServer(http.Dir("data/illust/"))))
	http.HandleFunc("/api/illusts/", func(w http.ResponseWriter, r *http.Request) {
		illustsRequestHandler(w, r, entC, ctx)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		entryPointHandler(w, r, options)
	})
	http.ListenAndServe(fmt.Sprintf(":%d", options.ServerPort), nil)
}

// major function definition end

func initialize(entC *ent.Client, ctx context.Context) {
	illustID, _ := gonanoid.New()

	illust, _ := entC.Illust.
		Create().
		SetIllustID(illustID).
		SetHash("ce00ba358144f1f67445cad5ad5821c3").
		SetOriginalURL("pixiv.net/artworks/88359980").
		SetHeightRatio(1.2824).
		Save(ctx)

	var tagArr []*ent.Tag

	for _, v := range []map[string]string{
		{
			"type":         "WORK",
			"officialName": "ブルーアーカイブ",
		},
		{
			"type":         "CHARACTER",
			"officialName": "早瀬 ユウカ",
		},
	} {
		var tagType tag.Type = tag.Type(v["type"])
		tagID, _ := gonanoid.New()
		tag, _ := entC.Tag.
			Create().
			SetTagID(tagID).
			SetType(tagType).
			SetOfficialName(v["officialName"]).
			Save(ctx)
		tagArr = append(tagArr, tag)
	}

	for _, v := range tagArr {
		entC.IllustTag.
			Create().
			SetParentIllustID(illust.IllustID).
			SetPrentTagID(v.TagID).
			SetIllustTagID(illust.IllustID + v.TagID).
			Save(ctx)
	}

	fmt.Println(illustID)
}
